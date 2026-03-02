package api

import (
	"fmt"
	"net/http"
	transaction_cache "pointofsale/internal/cache/api/transaction"
	"pointofsale/internal/domain/requests"
	response_api "pointofsale/internal/mapper"
	"pointofsale/internal/pb"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/logger"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transactionHandleApi struct {
	client     pb.TransactionServiceClient
	logger     logger.LoggerInterface
	mapping    response_api.TransactionResponseMapper
	apiHandler errors.ApiHandler
	cache      transaction_cache.TransactionMencache
}

func NewHandlerTransaction(
	router *echo.Echo,
	client pb.TransactionServiceClient,
	logger logger.LoggerInterface,
	mapping response_api.TransactionResponseMapper,
	apiHandler errors.ApiHandler,
	cache transaction_cache.TransactionMencache,
) *transactionHandleApi {
	transactionHandle := &transactionHandleApi{
		client:     client,
		logger:     logger,
		mapping:    mapping,
		apiHandler: apiHandler,
		cache:      cache,
	}

	routerTransaction := router.Group("/api/transaction")

	routerTransaction.GET("", transactionHandle.FindAllTransaction)
	routerTransaction.GET("/:id", transactionHandle.FindById)
	routerTransaction.GET("/merchant/:merchant_id", transactionHandle.FindByMerchant)
	routerTransaction.GET("/active", transactionHandle.FindByActive)
	routerTransaction.GET("/trashed", transactionHandle.FindByTrashed)

	routerTransaction.GET("/monthly-success", transactionHandle.FindMonthStatusSuccess)
	routerTransaction.GET("/yearly-success", transactionHandle.FindYearStatusSuccess)
	routerTransaction.GET("/monthly-failed", transactionHandle.FindMonthStatusFailed)
	routerTransaction.GET("/yearly-failed", transactionHandle.FindYearStatusFailed)

	routerTransaction.GET("/merchant/monthly-success", transactionHandle.FindMonthStatusSuccessByMerchant)
	routerTransaction.GET("/merchant/yearly-success", transactionHandle.FindYearStatusSuccessByMerchant)
	routerTransaction.GET("/merchant/monthly-failed", transactionHandle.FindMonthStatusFailedByMerchant)
	routerTransaction.GET("/merchant/yearly-failed", transactionHandle.FindYearStatusFailedByMerchant)

	routerTransaction.GET("/monthly-method-success", transactionHandle.FindMonthMethodSuccess)
	routerTransaction.GET("/yearly-method-success", transactionHandle.FindYearMethodSuccess)

	routerTransaction.GET("/merchant/monthly-method-success/:merchant_id", transactionHandle.FindMonthMethodByMerchantSuccess)
	routerTransaction.GET("/merchant/yearly-method-success/:merchant_id", transactionHandle.FindYearMethodByMerchantSuccess)

	routerTransaction.GET("/monthly-method-failed", transactionHandle.FindMonthMethodFailed)
	routerTransaction.GET("/yearly-method-failed", transactionHandle.FindYearMethodFailed)

	routerTransaction.GET("/merchant/monthly-method-failed/:merchant_id", transactionHandle.FindMonthMethodByMerchantFailed)
	routerTransaction.GET("/merchant/yearly-method-failed/:merchant_id", transactionHandle.FindYearMethodByMerchantFailed)

	routerTransaction.POST("/create", apiHandler.Handle("create", transactionHandle.Create))
	routerTransaction.POST("/update/:id", apiHandler.Handle("update", transactionHandle.Update))

	routerTransaction.POST("/trashed/:id", apiHandler.Handle("trashed", transactionHandle.TrashedTransaction))
	routerTransaction.POST("/restore/:id", apiHandler.Handle("restore", transactionHandle.RestoreTransaction))
	routerTransaction.DELETE("/permanent/:id", apiHandler.Handle("delete", transactionHandle.DeleteTransactionPermanent))

	routerTransaction.POST("/restore/all", apiHandler.Handle("restore-all", transactionHandle.RestoreAllTransaction))
	routerTransaction.POST("/permanent/all", apiHandler.Handle("delete-all", transactionHandle.DeleteAllTransactionPermanent))

	return transactionHandle
}

// @Security Bearer
// @Summary Find all transactions
// @Tags Transaction
// @Description Retrieve a list of all transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransaction "List of transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction [get]
func (h *transactionHandleApi) FindAllTransaction(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllTransaction{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedTransactionsCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to fetch transactions", zap.Error(err))
		return h.handleGrpcError(err, "FindAllTransaction")
	}

	so := h.mapping.ToApiResponsePaginationTransaction(res)

	h.cache.SetCachedTransactionsCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find all transactions by merchant
// @Tags Transaction
// @Description Retrieve a list of all transactions filtered by merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransaction "List of transactions"
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction [get]
func (h *transactionHandleApi) FindByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))

	if err != nil || merchantID <= 0 {
		h.logger.Debug("Invalid merchant ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid merchant ID")
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllTransactionByMerchant{
		MerchantID: merchantID,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
	}

	if cached, found := h.cache.GetCachedTransactionByMerchant(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllTransactionMerchantRequest{
		MerchantId: int32(merchantID),
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Search:     search,
	}

	res, err := h.client.FindByMerchant(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to fetch transaction", zap.Error(err))
		return h.handleGrpcError(err, "FindByMerchant")
	}

	so := h.mapping.ToApiResponsePaginationTransaction(res)

	h.cache.SetCachedTransactionByMerchant(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find transaction by ID
// @Tags Transaction
// @Description Retrieve a transaction by ID
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransaction "Transaction data"
// @Failure 400 {object} response.ErrorResponse "Invalid transaction ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/{id} [get]
func (h *transactionHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid transaction ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid transaction ID")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedTransactionCache(ctx, id); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindByIdTransactionRequest{
		Id: int32(id),
	}

	res, err := h.client.FindById(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to fetch transaction details", zap.Error(err))
		return h.handleGrpcError(err, "FindById")
	}

	so := h.mapping.ToApiResponseTransaction(res)

	h.cache.SetCachedTransactionCache(ctx, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active transactions
// @Tags Transaction
// @Description Retrieve a list of active transactions
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransactionDeleteAt "List of active transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/active [get]
func (h *transactionHandleApi) FindByActive(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllTransaction{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedTransactionActiveCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to fetch active transactions", zap.Error(err))
		return h.handleGrpcError(err, "FindByActive")
	}

	so := h.mapping.ToApiResponsePaginationTransactionDeleteAt(res)

	h.cache.SetCachedTransactionActiveCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// FindByTrashed retrieves a list of trashed transaction records.
// @Summary Retrieve trashed transactions
// @Tags Transaction
// @Description Retrieve a list of trashed transaction records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationTransactionDeleteAt "List of trashed transaction data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve transaction data"
// @Router /api/transaction/trashed [get]
func (h *transactionHandleApi) FindByTrashed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllTransaction{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedTransactionTrashedCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllTransactionRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to fetch archived transactions", zap.Error(err))
		return h.handleGrpcError(err, "FindByTrashed")
	}

	so := h.mapping.ToApiResponsePaginationTransactionDeleteAt(res)

	h.cache.SetCachedTransactionTrashedCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthStatusSuccess retrieves monthly successful transactions
// @Summary Get monthly successful transactions
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of successful transactions by month
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month in MM format (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year or month parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/monthly-success [get]
func (h *transactionHandleApi) FindMonthStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthAmountTransaction{
		Year:  year,
		Month: month,
	}

	if cached, found := h.cache.GetCachedMonthAmountSuccessCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthStatusSuccess(ctx, &pb.FindMonthlyTransactionStatus{
		Year:  int32(year),
		Month: int32(month),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction status success", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthStatusSuccess")
	}

	so := h.mapping.ToApiResponseTransactionMonthAmountSuccess(res)

	h.cache.SetCachedMonthAmountSuccessCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearStatusSuccess retrieves yearly successful transactions
// @Summary Get yearly successful transactions
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of successful transactions by year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/yearly-success [get]
func (h *transactionHandleApi) FindYearStatusSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedYearAmountSuccessCached(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearStatusSuccess(ctx, &pb.FindYearlyTransactionStatus{
		Year: int32(year),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction status success", zap.Error(err))
		return h.handleGrpcError(err, "FindYearStatusSuccess")
	}

	so := h.mapping.ToApiResponseTransactionYearAmountSuccess(res)

	h.cache.SetCachedYearAmountSuccessCached(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthStatusFailed retrieves monthly failed transactions
// @Summary Get monthly failed transactions
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of failed transactions by month
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month in MM format (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year or month parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/monthly-failed [get]
func (h *transactionHandleApi) FindMonthStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthAmountTransaction{
		Year:  year,
		Month: month,
	}

	if cached, found := h.cache.GetCachedMonthAmountFailedCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthStatusFailed(ctx, &pb.FindMonthlyTransactionStatus{
		Year:  int32(year),
		Month: int32(month),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction status failed", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthStatusFailed")
	}

	so := h.mapping.ToApiResponseTransactionMonthAmountFailed(res)

	h.cache.SetCachedMonthAmountFailedCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearStatusFailed retrieves yearly failed transactions
// @Summary Get yearly failed transactions
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of failed transactions by year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearFailed
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/yearly-failed [get]
func (h *transactionHandleApi) FindYearStatusFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedYearAmountFailedCached(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearStatusFailed(ctx, &pb.FindYearlyTransactionStatus{
		Year: int32(year),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction status failed", zap.Error(err))
		return h.handleGrpcError(err, "FindYearStatusFailed")
	}

	so := h.mapping.ToApiResponseTransactionYearAmountFailed(res)

	h.cache.SetCachedYearAmountFailedCached(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthStatusSuccessByMerchant retrieves monthly successful transactions by merchant
// @Summary Get monthly successful transactions by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of successful transactions by month for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month in MM format (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/monthly-success [get]
func (h *transactionHandleApi) FindMonthStatusSuccessByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	merchantIdStr := c.QueryParam("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthAmountTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedMonthAmountSuccessByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthStatusSuccessByMerchant(ctx, &pb.FindMonthlyTransactionStatusByMerchant{
		Year:       int32(year),
		Month:      int32(month),
		MerchantId: int32(merchant_id),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction status success", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthStatusSuccessByMerchant")
	}

	so := h.mapping.ToApiResponseTransactionMonthAmountSuccess(res)

	h.cache.SetCachedMonthAmountSuccessByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearStatusSuccessByMerchant retrieves yearly successful transactions by merchant
// @Summary Get yearly successful transactions by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of successful transactions by year for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearSuccess
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/yearly-success [get]
func (h *transactionHandleApi) FindYearStatusSuccessByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	merchantIdStr := c.QueryParam("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearAmountTransactionMerchant{
		Year:       year,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedYearAmountSuccessByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearStatusSuccessByMerchant(ctx, &pb.FindYearlyTransactionStatusByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant_id),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction status success", zap.Error(err))
		return h.handleGrpcError(err, "FindYearStatusSuccessByMerchant")
	}

	so := h.mapping.ToApiResponseTransactionYearAmountSuccess(res)

	h.cache.SetCachedYearAmountSuccessByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthStatusFailedByMerchant retrieves monthly failed transactions by merchant
// @Summary Get monthly failed transactions by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of failed transactions by month for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month in MM format (1-12)"
// @Success 200 {object} response.ApiResponsesTransactionMonthFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID, year or month parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/monthly-failed [get]
func (h *transactionHandleApi) FindMonthStatusFailedByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	monthStr := c.QueryParam("month")
	merchantIdStr := c.QueryParam("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthAmountTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedMonthAmountFailedByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthStatusFailedByMerchant(ctx, &pb.FindMonthlyTransactionStatusByMerchant{
		Year:       int32(year),
		Month:      int32(month),
		MerchantId: int32(merchant_id),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction status failed", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthStatusFailedByMerchant")
	}

	so := h.mapping.ToApiResponseTransactionMonthAmountFailed(res)

	h.cache.SetCachedMonthAmountFailedByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearStatusFailedByMerchant retrieves yearly failed transactions by merchant
// @Summary Get yearly failed transactions by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of failed transactions by year for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearFailed
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/yearly-failed [get]
func (h *transactionHandleApi) FindYearStatusFailedByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	merchantIdStr := c.QueryParam("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearAmountTransactionMerchant{
		Year:       year,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedYearAmountFailedByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearStatusFailedByMerchant(ctx, &pb.FindYearlyTransactionStatusByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant_id),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction status failed", zap.Error(err))
		return h.handleGrpcError(err, "FindYearStatusFailedByMerchant")
	}

	so := h.mapping.ToApiResponseTransactionYearAmountFailed(res)

	h.cache.SetCachedYearAmountFailedByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthMethod retrieves monthly payment method statistics
// @Summary Get monthly payment method distribution
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by month
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/monthly-method-success [get]
func (h *transactionHandleApi) FindMonthMethodSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthMethodTransaction{
		Year:  year,
		Month: month,
	}

	if cached, found := h.cache.GetCachedMonthMethodSuccessCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthMethodSuccess(ctx, &pb.MonthTransactionMethod{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthMethodSuccess")
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)

	h.cache.SetCachedMonthMethodSuccessCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearMethod retrieves yearly payment method statistics
// @Summary Get yearly payment method distribution
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/yearly-method-success [get]
func (h *transactionHandleApi) FindYearMethodSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedYearMethodSuccessCached(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearMethodSuccess(ctx, &pb.YearTransactionMethod{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindYearMethodSuccess")
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)

	h.cache.SetCachedYearMethodSuccessCached(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthMethodByMerchant retrieves monthly payment method statistics by merchant
// @Summary Get monthly payment method distribution by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by month for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/monthly-method-success/{merchant_id} [get]
func (h *transactionHandleApi) FindMonthMethodByMerchantSuccess(c echo.Context) error {
	merchantIdStr := c.Param("merchant_id")
	yearStr := c.QueryParam("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthMethodTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedMonthMethodSuccessByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthMethodByMerchantSuccess(ctx, &pb.MonthTransactionMethodByMerchant{
		Year:       int32(year),
		Month:      int32(month),
		MerchantId: int32(merchant_id),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthMethodByMerchantSuccess")
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)

	h.cache.SetCachedMonthMethodSuccessByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearMethodByMerchant retrieves yearly payment method statistics by merchant
// @Summary Get yearly payment method distribution by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by year for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/yearly-method-success/{merchant_id} [get]
func (h *transactionHandleApi) FindYearMethodByMerchantSuccess(c echo.Context) error {
	yearStr := c.QueryParam("year")
	merchantIdStr := c.Param("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearMethodTransactionMerchant{
		Year:       year,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedYearMethodSuccessByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearMethodByMerchantSuccess(ctx, &pb.YearTransactionMethodByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant_id),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindYearMethodByMerchantSuccess")
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)

	h.cache.SetCachedYearMethodSuccessByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthMethod retrieves monthly payment method statistics
// @Summary Get monthly payment method distribution
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by month
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/monthly-method-failed [get]
func (h *transactionHandleApi) FindMonthMethodFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthMethodTransaction{
		Year:  year,
		Month: month,
	}

	if cached, found := h.cache.GetCachedMonthMethodFailedCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthMethodFailed(ctx, &pb.MonthTransactionMethod{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthMethodFailed")
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)

	h.cache.SetCachedMonthMethodFailedCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearMethod retrieves yearly payment method statistics
// @Summary Get yearly payment method distribution
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/yearly-method-failed [get]
func (h *transactionHandleApi) FindYearMethodFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedYearMethodFailedCached(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearMethodFailed(ctx, &pb.YearTransactionMethod{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindYearMethodFailed")
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)

	h.cache.SetCachedYearMethodFailedCached(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthMethodByMerchant retrieves monthly payment method statistics by merchant
// @Summary Get monthly payment method distribution by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by month for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionMonthMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/monthly-method-failed/{merchant_id} [get]
func (h *transactionHandleApi) FindMonthMethodByMerchantFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	merchantIdStr := c.Param("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil {
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthMethodTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedMonthMethodFailedByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthMethodByMerchantFailed(ctx, &pb.MonthTransactionMethodByMerchant{
		Year:       int32(year),
		Month:      int32(month),
		MerchantId: int32(merchant_id),
	})

	if err != nil {
		h.logger.Debug("Failed to retrieve monthly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthMethodByMerchantFailed")
	}

	so := h.mapping.ToApiResponseTransactionMonthMethod(res)

	h.cache.SetCachedMonthMethodFailedByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearMethodByMerchant retrieves yearly payment method statistics by merchant
// @Summary Get yearly payment method distribution by merchant
// @Tags Transaction
// @Security Bearer
// @Description Retrieve statistics of payment methods used by year for specific merchant
// @Accept json
// @Produce json
// @Param merchant_id query int true "Merchant ID"
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponsesTransactionYearMethod
// @Failure 400 {object} response.ErrorResponse "Invalid merchant ID or year parameter"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 404 {object} response.ErrorResponse "Merchant not found"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/transaction/merchant/yearly-method-failed/{merchant_id} [get]
func (h *transactionHandleApi) FindYearMethodByMerchantFailed(c echo.Context) error {
	yearStr := c.QueryParam("year")
	merchantIdStr := c.Param("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)

	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearMethodTransactionMerchant{
		Year:       year,
		MerchantID: merchant_id,
	}

	if cached, found := h.cache.GetCachedYearMethodFailedByMerchantCached(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearMethodByMerchantFailed(ctx, &pb.YearTransactionMethodByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant_id),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly transaction methods", zap.Error(err))
		return h.handleGrpcError(err, "FindYearMethodByMerchantFailed")
	}

	so := h.mapping.ToApiResponseTransactionYearMethod(res)

	h.cache.SetCachedYearMethodFailedByMerchantCached(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create a new transaction
// @Tags Transaction
// @Description Create a new transaction record
// @Accept json
// @Produce json
// @Param request body requests.CreateTransactionRequest true "Transaction details"
// @Success 200 {object} response.ApiResponseTransaction "Successfully created transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create transaction"
// @Router /api/transaction/create [post]
func (h *transactionHandleApi) Create(c echo.Context) error {
	var body requests.CreateTransactionRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request format", zap.Error(err))
		return errors.NewBadRequestError("Invalid request format")
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return errors.NewBadRequestError("Validation failed: " + err.Error())
	}

	ctx := c.Request().Context()

	grpcReq := &pb.CreateTransactionRequest{
		OrderId:       int32(body.OrderID),
		CashierId:     int32(body.CashierID),
		PaymentMethod: body.PaymentMethod,
		Amount:        int32(body.Amount),
	}

	res, err := h.client.Create(ctx, grpcReq)

	if err != nil {
		h.logger.Error("transaction creation failed", zap.Error(err))
		return h.handleGrpcError(err, "Create")
	}

	so := h.mapping.ToApiResponseTransaction(res)

	h.cache.DeleteTransactionCache(ctx, int(res.Data.Id))

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Update an existing transaction
// @Tags Transaction
// @Description Update an existing transaction record
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Param request body requests.UpdateTransactionRequest true "Updated transaction details"
// @Success 200 {object} response.ApiResponseTransaction "Successfully updated transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update transaction"
// @Router /api/transaction/update/{id} [post]
func (h *transactionHandleApi) Update(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		h.logger.Debug("Invalid id parameter", zap.Error(err))
		return errors.NewBadRequestError("Invalid transaction ID")
	}

	var body requests.UpdateTransactionRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request format", zap.Error(err))
		return errors.NewBadRequestError("Invalid request format")
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return errors.NewBadRequestError("Validation failed: " + err.Error())
	}

	ctx := c.Request().Context()

	grpcReq := &pb.UpdateTransactionRequest{
		TransactionId: int32(idInt),
		OrderId:       int32(body.OrderID),
		CashierId:     int32(body.CashierID),
		PaymentMethod: body.PaymentMethod,
		Amount:        int32(body.Amount),
	}

	res, err := h.client.Update(ctx, grpcReq)

	if err != nil {
		h.logger.Error("transaction update failed", zap.Error(err))
		return h.handleGrpcError(err, "Update")
	}

	so := h.mapping.ToApiResponseTransaction(res)

	h.cache.DeleteTransactionCache(ctx, idInt)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// TrashedTransaction retrieves a trashed transaction record by its ID.
// @Summary Retrieve a trashed transaction
// @Tags Transaction
// @Description Retrieve a trashed transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransactionDeleteAt "Successfully retrieved trashed transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed transaction"
// @Router /api/transaction/trashed/{id} [get]
func (h *transactionHandleApi) TrashedTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid transaction ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid transaction ID")
	}

	ctx := c.Request().Context()

	grpcReq := &pb.FindByIdTransactionRequest{
		Id: int32(id),
	}

	res, err := h.client.TrashedTransaction(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to archive transaction", zap.Error(err))
		return h.handleGrpcError(err, "TrashedTransaction")
	}

	so := h.mapping.ToApiResponseTransactionDeleteAt(res)

	h.cache.DeleteTransactionCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreTransaction restores a transaction record from the trash by its ID.
// @Summary Restore a trashed transaction
// @Tags Transaction
// @Description Restore a trashed transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransactionDeleteAt "Successfully restored transaction"
// @Failure 400 {object} response.ErrorResponse "Invalid transaction ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore transaction"
// @Router /api/transaction/restore/{id} [post]
func (h *transactionHandleApi) RestoreTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid transaction ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid transaction ID")
	}

	ctx := c.Request().Context()

	grpcReq := &pb.FindByIdTransactionRequest{
		Id: int32(id),
	}

	res, err := h.client.RestoreTransaction(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to restore transaction", zap.Error(err))
		return h.handleGrpcError(err, "RestoreTransaction")
	}

	so := h.mapping.ToApiResponseTransactionDeleteAt(res)

	h.cache.DeleteTransactionCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteTransactionPermanent permanently deletes a transaction record by its ID.
// @Summary Permanently delete a transaction
// @Tags Transaction
// @Description Permanently delete a transaction record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} response.ApiResponseTransactionDelete "Successfully deleted transaction record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete transaction"
// @Router /api/transaction/delete/{id} [delete]
func (h *transactionHandleApi) DeleteTransactionPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid transaction ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid transaction ID")
	}

	ctx := c.Request().Context()

	grpcReq := &pb.FindByIdTransactionRequest{
		Id: int32(id),
	}

	res, err := h.client.DeleteTransactionPermanent(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to delete transaction", zap.Error(err))
		return h.handleGrpcError(err, "DeleteTransactionPermanent")
	}

	so := h.mapping.ToApiResponseTransactionDelete(res)

	h.cache.DeleteTransactionCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreAllTransaction restores all trashed transactions.
// @Summary Restore all trashed transactions
// @Tags Transaction
// @Description Restore all trashed transactions.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseTransactionAll "Successfully restored all transactions"
// @Failure 500 {object} response.ErrorResponse "Failed to restore transactions"
// @Router /api/transaction/restore/all [post]
func (h *transactionHandleApi) RestoreAllTransaction(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllTransaction(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Bulk transactions restoration failed", zap.Error(err))
		return h.handleGrpcError(err, "RestoreAllTransaction")
	}

	so := h.mapping.ToApiResponseTransactionAll(res)

	h.logger.Debug("Successfully restored all transactions")

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteAllTransactionPermanent permanently deletes all transactions.
// @Summary Permanently delete all transactions
// @Tags Transaction
// @Description Permanently delete all transactions.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseTransactionAll "Successfully deleted all transactions permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete transactions"
// @Router /api/transaction/delete/all [post]
func (h *transactionHandleApi) DeleteAllTransactionPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllTransactionPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		return h.handleGrpcError(err, "DeleteAllTransactionPermanent")
	}

	so := h.mapping.ToApiResponseTransactionAll(res)

	h.logger.Debug("Successfully deleted all transactions permanently")

	return c.JSON(http.StatusOK, so)
}

func (h *transactionHandleApi) handleGrpcError(err error, operation string) *errors.AppError {
	st, ok := status.FromError(err)
	if !ok {
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}

	switch st.Code() {
	case codes.NotFound:
		return errors.NewNotFoundError("Transaction").WithInternal(err)

	case codes.AlreadyExists:
		return errors.NewConflictError("Transaction already exists").WithInternal(err)

	case codes.InvalidArgument:
		return errors.NewBadRequestError(st.Message()).WithInternal(err)

	case codes.PermissionDenied:
		return errors.ErrForbidden.WithInternal(err)

	case codes.Unauthenticated:
		return errors.ErrUnauthorized.WithInternal(err)

	case codes.ResourceExhausted:
		return errors.ErrTooManyRequests.WithInternal(err)

	case codes.Unavailable:
		return errors.NewServiceUnavailableError("Transaction service").WithInternal(err)

	case codes.DeadlineExceeded:
		return errors.ErrTimeout.WithInternal(err)

	default:
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}
}

func (h *transactionHandleApi) parseValidationErrors(err error) []errors.ValidationError {
	var validationErrs []errors.ValidationError

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrs = append(validationErrs, errors.ValidationError{
				Field:   fe.Field(),
				Message: h.getValidationMessage(fe),
			})
		}
		return validationErrs
	}

	return []errors.ValidationError{
		{
			Field:   "general",
			Message: err.Error(),
		},
	}
}

func (h *transactionHandleApi) getValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("Must be at most %s", fe.Param())
	case "gte":
		return fmt.Sprintf("Must be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Must be less than or equal to %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("Must be one of: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on '%s' tag", fe.Tag())
	}
}
