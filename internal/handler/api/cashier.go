package api

import (
	"net/http"
	cashier_cache "pointofsale/internal/cache/api/cashier"
	"pointofsale/internal/domain/requests"
	response_api "pointofsale/internal/mapper"
	"pointofsale/internal/pb"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/logger"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cashierHandleApi struct {
	client     pb.CashierServiceClient
	logger     logger.LoggerInterface
	mapping    response_api.CashierResponseMapper
	apiHandler errors.ApiHandler
	cache      cashier_cache.CashierMencache
}

func NewHandlerCashier(
	router *echo.Echo,
	client pb.CashierServiceClient,
	logger logger.LoggerInterface,
	mapping response_api.CashierResponseMapper,
	apiHandler errors.ApiHandler,
	cache cashier_cache.CashierMencache,
) *cashierHandleApi {
	cashierHandler := &cashierHandleApi{
		client:     client,
		logger:     logger,
		mapping:    mapping,
		cache:      cache,
		apiHandler: apiHandler,
	}

	routerCashier := router.Group("/api/cashier")

	routerCashier.GET("", cashierHandler.FindAllCashier)
	routerCashier.GET("/:id", cashierHandler.FindById)
	routerCashier.GET("/active", cashierHandler.FindByActive)
	routerCashier.GET("/trashed", cashierHandler.FindByTrashed)

	routerCashier.GET("/monthly-total-sales", cashierHandler.FindMonthlyTotalSales)
	routerCashier.GET("/yearly-total-sales", cashierHandler.FindYearTotalSales)

	routerCashier.GET("/merchant/monthly-total-sales", cashierHandler.FindMonthlyTotalSalesByMerchant)
	routerCashier.GET("/merchant/yearly-total-sales", cashierHandler.FindYearTotalSalesByMerchant)

	routerCashier.GET("/mycashier/monthly-total-sales", cashierHandler.FindMonthlyTotalSalesById)
	routerCashier.GET("/mycashier/yearly-total-sales", cashierHandler.FindYearTotalSalesById)

	routerCashier.GET("/monthly-sales", cashierHandler.FindMonthSales)
	routerCashier.GET("/yearly-sales", cashierHandler.FindYearSales)
	routerCashier.GET("/merchant/monthly-sales", cashierHandler.FindMonthSalesByMerchant)
	routerCashier.GET("/merchant/yearly-sales", cashierHandler.FindYearSalesByMerchant)
	routerCashier.GET("/mycashier/monthly-sales", cashierHandler.FindMonthSalesById)
	routerCashier.GET("/mycashier/yearly-sales", cashierHandler.FindYearSalesById)

	routerCashier.POST("/create", apiHandler.Handle("create", cashierHandler.CreateCashier))
	routerCashier.POST("/update/:id", apiHandler.Handle("update", cashierHandler.UpdateCashier))

	routerCashier.POST("/trashed/:id", apiHandler.Handle("trashed", cashierHandler.TrashedCashier))
	routerCashier.POST("/restore/:id", apiHandler.Handle("restore", cashierHandler.RestoreCashier))
	routerCashier.DELETE("/permanent/:id", apiHandler.Handle("delete", cashierHandler.DeleteCashierPermanent))

	routerCashier.POST("/restore/all", apiHandler.Handle("restore-all", cashierHandler.RestoreAllCashier))
	routerCashier.POST("/permanent/all", apiHandler.Handle("delete-all", cashierHandler.DeleteAllCashierPermanent))

	return cashierHandler
}

// @Security Bearer
// @Summary Find all cashiers
// @Tags Cashier
// @Description Retrieve a list of all cashiers
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationCashier "List of cashiers"
// @Failure 500 {object} errors.ApiError "Failed to retrieve cashier data"
// @Router /api/cashier [get]
func (h *cashierHandleApi) FindAllCashier(c echo.Context) error {
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

	req := &requests.FindAllCashiers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	// Check cache first
	if cached, found := h.cache.GetCachedCashiersCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllCashierRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch cashiers", zap.Error(err))
		return h.handleGrpcError(err, "FindAllCashier")
	}

	so := h.mapping.ToApiResponsePaginationCashier(res)

	// Set cache
	h.cache.SetCachedCashiersCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find cashier by ID
// @Tags Cashier
// @Description Retrieve a cashier by ID
// @Accept json
// @Produce json
// @Param id path int true "cashier ID"
// @Success 200 {object} response.ApiResponseCashier "cashier data"
// @Failure 400 {object} errors.ApiError "Invalid cashier ID"
// @Failure 500 {object} errors.ApiError "Failed to retrieve cashier data"
// @Router /api/cashier/{id} [get]
func (h *cashierHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid cashier ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid cashier ID")
	}

	ctx := c.Request().Context()

	// Check cache first
	if cached, found := h.cache.GetCachedCashier(ctx, id); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindByIdCashierRequest{Id: int32(id)}

	cashier, err := h.client.FindById(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch cashier details", zap.Error(err))
		return h.handleGrpcError(err, "FindById")
	}

	so := h.mapping.ToApiResponseCashier(cashier)

	// Set cache
	h.cache.SetCachedCashier(ctx, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active cashier
// @Tags Cashier
// @Description Retrieve a list of active cashier
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationCashierDeleteAt "List of active cashier"
// @Failure 500 {object} errors.ApiError "Failed to retrieve cashier data"
// @Router /api/cashier/active [get]
func (h *cashierHandleApi) FindByActive(c echo.Context) error {
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

	req := &requests.FindAllCashiers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	// Check cache first
	if cached, found := h.cache.GetCachedCashiersActive(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllCashierRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch active cashiers", zap.Error(err))
		return h.handleGrpcError(err, "FindByActive")
	}

	so := h.mapping.ToApiResponsePaginationCashierDeleteAt(res)

	// Set cache
	h.cache.SetCachedCashiersActive(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// FindByTrashed retrieves a list of trashed cashier records.
// @Summary Retrieve trashed cashier
// @Tags Cashier
// @Description Retrieve a list of trashed cashier records
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationCashierDeleteAt "List of trashed cashier data"
// @Failure 500 {object} errors.ApiError "Failed to retrieve cashier data"
// @Router /api/cashier/trashed [get]
func (h *cashierHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &requests.FindAllCashiers{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	// Check cache first
	if cached, found := h.cache.GetCachedCashiersTrashed(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllCashierRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch archived cashiers", zap.Error(err))
		return h.handleGrpcError(err, "FindByTrashed")
	}

	so := h.mapping.ToApiResponsePaginationCashierDeleteAt(res)

	// Set cache
	h.cache.SetCachedCashiersTrashed(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTotalSales retrieves the monthly cashiers for a specific year.
// @Summary Get monthly cashiers statistics
// @Tags Cashier
// @Security Bearer
// @Description Retrieve monthly cashiers statistics for a given year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month"
// @Success 200 {object} response.ApiResponseCashierMonthSales "Successfully retrieved monthly sales data"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/monthly-total-sales [get]
func (h *cashierHandleApi) FindMonthlyTotalSales(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	monthStr := c.QueryParam("month")
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		h.logger.Debug("Invalid month parameter", zap.Error(err))
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthTotalSales{
		Year:  year,
		Month: month,
	}

	// Check cache first
	if cached, found := h.cache.GetMonthlyTotalSalesCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthlyTotalSales(ctx, &pb.FindYearMonthTotalSales{
		Year:  int32(year),
		Month: int32(month),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthlyTotalSales")
	}

	so := h.mapping.ToApiResponseMonthlyTotalSales(res)

	// Set cache
	h.cache.SetMonthlyTotalSalesCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearTotalSales retrieves the yearly cashiers for a specific year.
// @Summary Get yearly cashiers
// @Tags Cashier
// @Security Bearer
// @Description Retrieve the yearly cashiers for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponseCashierYearSales "Yearly cashiers"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 500 {object} errors.ApiError "Failed to retrieve yearly cashiers"
// @Router /api/cashier/yearly-total-sales [get]
func (h *cashierHandleApi) FindYearTotalSales(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	// Check cache first
	if cached, found := h.cache.GetYearlyTotalSalesCache(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearlyTotalSales(ctx, &pb.FindYearTotalSales{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindYearTotalSales")
	}

	so := h.mapping.ToApiResponseYearlyTotalSales(res)

	// Set cache
	h.cache.SetYearlyTotalSalesCache(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTotalSalesById retrieves the monthly cashiers for a specific year.
// @Summary Get monthly cashiers statistics
// @Tags Cashier
// @Security Bearer
// @Description Retrieve monthly cashiers statistics for a given year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month"
// @Param cashier_id query int true "Cashier id"
// @Success 200 {object} response.ApiResponseCashierMonthSales "Successfully retrieved monthly sales data"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/mycashier/monthly-total-sales [get]
func (h *cashierHandleApi) FindMonthlyTotalSalesById(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	monthStr := c.QueryParam("month")
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		h.logger.Debug("Invalid month parameter", zap.Error(err))
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	cashierStr := c.QueryParam("cashier_id")
	cashier, err := strconv.Atoi(cashierStr)
	if err != nil {
		h.logger.Debug("Invalid cashier id parameter", zap.Error(err))
		return errors.NewBadRequestError("cashier_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthTotalSalesCashier{
		Year:      year,
		Month:     month,
		CashierID: cashier,
	}

	// Check cache first
	if cached, found := h.cache.GetMonthlyTotalSalesByIdCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthlyTotalSalesById(ctx, &pb.FindYearMonthTotalSalesById{
		Year:      int32(year),
		Month:     int32(month),
		CashierId: int32(cashier),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthlyTotalSalesById")
	}

	so := h.mapping.ToApiResponseMonthlyTotalSales(res)

	// Set cache
	h.cache.SetMonthlyTotalSalesByIdCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearTotalSalesById retrieves the yearly cashiers for a specific year.
// @Summary Get yearly cashiers
// @Tags Cashier
// @Security Bearer
// @Description Retrieve the yearly cashiers for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param cashier_id query int true "Cashier ID"
// @Success 200 {object} response.ApiResponseCashierYearSales "Yearly cashiers"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 500 {object} errors.ApiError "Failed to retrieve yearly cashiers"
// @Router /api/cashier/mycashier/yearly-total-sales [get]
func (h *cashierHandleApi) FindYearTotalSalesById(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	cashierStr := c.QueryParam("cashier_id")
	cashier, err := strconv.Atoi(cashierStr)
	if err != nil {
		h.logger.Debug("Invalid cashier id parameter", zap.Error(err))
		return errors.NewBadRequestError("cashier_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearTotalSalesCashier{
		Year:      year,
		CashierID: cashier,
	}

	// Check cache first
	if cached, found := h.cache.GetYearlyTotalSalesByIdCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearlyTotalSalesById(ctx, &pb.FindYearTotalSalesById{
		Year:      int32(year),
		CashierId: int32(cashier),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindYearTotalSalesById")
	}

	so := h.mapping.ToApiResponseYearlyTotalSales(res)

	// Set cache
	h.cache.SetYearlyTotalSalesByIdCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthlyTotalSalesByMerchant retrieves the monthly cashiers for a specific year.
// @Summary Get monthly cashiers statistics
// @Tags Cashier
// @Security Bearer
// @Description Retrieve monthly cashiers statistics for a given year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param month query int true "Month"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCashierMonthSales "Successfully retrieved monthly sales data"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/merchant/monthly-total-sales [get]
func (h *cashierHandleApi) FindMonthlyTotalSalesByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	monthStr := c.QueryParam("month")
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		h.logger.Debug("Invalid month parameter", zap.Error(err))
		return errors.NewBadRequestError("month is required and must be a valid number")
	}

	merchantStr := c.QueryParam("merchant_id")
	merchant, err := strconv.Atoi(merchantStr)
	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthTotalSalesMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchant,
	}

	// Check cache first
	if cached, found := h.cache.GetMonthlyTotalSalesByMerchantCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthlyTotalSalesByMerchant(ctx, &pb.FindYearMonthTotalSalesByMerchant{
		Year:       int32(year),
		Month:      int32(month),
		MerchantId: int32(merchant),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthlyTotalSalesByMerchant")
	}

	so := h.mapping.ToApiResponseMonthlyTotalSales(res)

	// Set cache
	h.cache.SetMonthlyTotalSalesByMerchantCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearTotalSalesByMerchant retrieves the yearly cashiers for a specific year.
// @Summary Get yearly cashiers
// @Tags Cashier
// @Security Bearer
// @Description Retrieve the yearly cashiers for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCashierYearSales "Yearly cashiers"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 500 {object} errors.ApiError "Failed to retrieve yearly cashiers"
// @Router /api/cashier/merchant/yearly-total-sales [get]
func (h *cashierHandleApi) FindYearTotalSalesByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchantStr := c.QueryParam("merchant_id")
	merchant, err := strconv.Atoi(merchantStr)
	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearTotalSalesMerchant{
		Year:       year,
		MerchantID: merchant,
	}

	// Check cache first
	if cached, found := h.cache.GetYearlyTotalSalesByMerchantCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearlyTotalSalesByMerchant(ctx, &pb.FindYearTotalSalesByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindYearTotalSalesByMerchant")
	}

	so := h.mapping.ToApiResponseYearlyTotalSales(res)

	// Set cache
	h.cache.SetYearlyTotalSalesByMerchantCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthSales retrieves the monthly cashiers for a specific year.
// @Summary Get monthly cashiers statistics
// @Tags Cashier
// @Security Bearer
// @Description Retrieve monthly cashiers statistics for a given year
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponseCashierMonthSales "Successfully retrieved monthly sales data"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/monthly-sales [get]
func (h *cashierHandleApi) FindMonthSales(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	// Check cache first
	if cached, found := h.cache.GetMonthlySalesCache(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthSales(ctx, &pb.FindYearCashier{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthSales")
	}

	so := h.mapping.ToApiResponseCashierMonthlySale(res)

	// Set cache
	h.cache.SetMonthlySalesCache(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearSales retrieves the yearly cashiers for a specific year.
// @Summary Get yearly cashiers
// @Tags Cashier
// @Security Bearer
// @Description Retrieve the yearly cashiers for a specific year.
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Success 200 {object} response.ApiResponseCashierYearSales "Yearly cashiers"
// @Failure 400 {object} errors.ApiError "Invalid year parameter"
// @Failure 500 {object} errors.ApiError "Failed to retrieve yearly cashiers"
// @Router /api/cashier/yearly-sales [get]
func (h *cashierHandleApi) FindYearSales(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	ctx := c.Request().Context()

	// Check cache first
	if cached, found := h.cache.GetYearlySalesCache(ctx, year); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearSales(ctx, &pb.FindYearCashier{
		Year: int32(year),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindYearSales")
	}

	so := h.mapping.ToApiResponseCashierYearlySale(res)

	// Set cache
	h.cache.SetYearlySalesCache(ctx, year, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthSalesByMerchant retrieves monthly cashiers for a specific merchant.
// @Summary Get monthly sales by merchant
// @Tags Cashier
// @Security Bearer
// @Description Retrieve monthly cashiers statistics for a specific merchant
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCashierMonthSales "Successfully retrieved monthly sales by merchant"
// @Failure 400 {object} errors.ApiError "Invalid merchant ID or year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 404 {object} errors.ApiError "Merchant not found"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/merchant/monthly-sales [get]
func (h *cashierHandleApi) FindMonthSalesByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	merchantIdStr := c.QueryParam("merchant_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchant_id, err := strconv.Atoi(merchantIdStr)
	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthCashierMerchant{
		Year:       year,
		MerchantID: merchant_id,
	}

	// Check cache first
	if cached, found := h.cache.GetMonthlyCashierByMerchantCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthSalesByMerchant(ctx, &pb.FindYearCashierByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant_id),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthSalesByMerchant")
	}

	so := h.mapping.ToApiResponseCashierMonthlySale(res)

	// Set cache
	h.cache.SetMonthlyCashierByMerchantCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearSalesByMerchant retrieves yearly cashier for a specific merchant.
// @Summary Get yearly sales by merchant
// @Tags Cashier
// @Security Bearer
// @Description Retrieve yearly cashier statistics for a specific merchant
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param merchant_id query int true "Merchant ID"
// @Success 200 {object} response.ApiResponseCashierYearSales "Successfully retrieved yearly sales by merchant"
// @Failure 400 {object} errors.ApiError "Invalid merchant ID or year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 404 {object} errors.ApiError "Merchant not found"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/merchant/yearly-sales [get]
func (h *cashierHandleApi) FindYearSalesByMerchant(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	merchantIdStr := c.QueryParam("merchant_id")
	merchant_id, err := strconv.Atoi(merchantIdStr)
	if err != nil {
		h.logger.Debug("Invalid merchant id parameter", zap.Error(err))
		return errors.NewBadRequestError("merchant_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearCashierMerchant{
		Year:       year,
		MerchantID: merchant_id,
	}

	// Check cache first
	if cached, found := h.cache.GetYearlyCashierByMerchantCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearSalesByMerchant(ctx, &pb.FindYearCashierByMerchant{
		Year:       int32(year),
		MerchantId: int32(merchant_id),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindYearSalesByMerchant")
	}

	so := h.mapping.ToApiResponseCashierYearlySale(res)

	// Set cache
	h.cache.SetYearlyCashierByMerchantCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindMonthSalesById retrieves monthly cashier for a specific cashier.
// @Summary Get monthly sales by cashier
// @Tags Cashier
// @Security Bearer
// @Description Retrieve monthly cashier statistics for a specific cashier
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param cashier_id query int true "Cashier ID"
// @Success 200 {object} response.ApiResponseCashierMonthSales "Successfully retrieved monthly sales by cashier"
// @Failure 400 {object} errors.ApiError "Invalid cashier ID or year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 404 {object} errors.ApiError "Cashier not found"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/mycashier/monthly-sales [get]
func (h *cashierHandleApi) FindMonthSalesById(c echo.Context) error {
	yearStr := c.QueryParam("year")
	cashierIdStr := c.QueryParam("cashier_id")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	cashier_id, err := strconv.Atoi(cashierIdStr)
	if err != nil {
		h.logger.Debug("Invalid cashier id parameter", zap.Error(err))
		return errors.NewBadRequestError("cashier_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.MonthCashierId{
		Year:      year,
		CashierID: cashier_id,
	}

	// Check cache first
	if cached, found := h.cache.GetMonthlyCashierByIdCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindMonthSalesById(ctx, &pb.FindYearCashierById{
		Year:      int32(year),
		CashierId: int32(cashier_id),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve monthly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindMonthSalesById")
	}

	so := h.mapping.ToApiResponseCashierMonthlySale(res)

	// Set cache
	h.cache.SetMonthlyCashierByIdCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// FindYearSalesById retrieves yearly cashier for a specific cashier.
// @Summary Get yearly sales by cashier
// @Tags Cashier
// @Security Bearer
// @Description Retrieve yearly cashier statistics for a specific cashier
// @Accept json
// @Produce json
// @Param year query int true "Year in YYYY format (e.g., 2023)"
// @Param cashier_id query int true "Cashier ID"
// @Success 200 {object} response.ApiResponseCashierYearSales "Successfully retrieved yearly sales by cashier"
// @Failure 400 {object} errors.ApiError "Invalid cashier ID or year parameter"
// @Failure 401 {object} errors.ApiError "Unauthorized"
// @Failure 404 {object} errors.ApiError "Cashier not found"
// @Failure 500 {object} errors.ApiError "Internal server error"
// @Router /api/cashier/mycashier/yearly-sales [get]
func (h *cashierHandleApi) FindYearSalesById(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		h.logger.Debug("Invalid year parameter", zap.Error(err))
		return errors.NewBadRequestError("year is required and must be a valid number")
	}

	cashierIdStr := c.QueryParam("cashier_id")
	cashier_id, err := strconv.Atoi(cashierIdStr)
	if err != nil {
		h.logger.Debug("Invalid cashier id parameter", zap.Error(err))
		return errors.NewBadRequestError("cashier_id is required and must be a valid number")
	}

	ctx := c.Request().Context()

	req := &requests.YearCashierId{
		Year:      year,
		CashierID: cashier_id,
	}

	// Check cache first
	if cached, found := h.cache.GetYearlyCashierByIdCache(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	res, err := h.client.FindYearSalesById(ctx, &pb.FindYearCashierById{
		Year:      int32(year),
		CashierId: int32(cashier_id),
	})
	if err != nil {
		h.logger.Debug("Failed to retrieve yearly cashier sales", zap.Error(err))
		return h.handleGrpcError(err, "FindYearSalesById")
	}

	so := h.mapping.ToApiResponseCashierYearlySale(res)

	// Set cache
	h.cache.SetYearlyCashierByIdCache(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// Create handles the creation of a new cashier.
// @Summary Create a new cashier
// @Tags Cashier
// @Description Create a new cashier with the provided details
// @Accept json
// @Produce json
// @Param request body requests.CreateCashierRequest true "Create cashier request"
// @Success 200 {object} response.ApiResponseCashier "Successfully created cashier"
// @Failure 400 {object} errors.ApiError "Invalid request body or validation error"
// @Failure 500 {object} errors.ApiError "Failed to create cashier"
// @Router /api/cashier/create [post]
func (h *cashierHandleApi) CreateCashier(c echo.Context) error {
	var body requests.CreateCashierRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request format", zap.Error(err))
		return errors.NewBadRequestError("Invalid request format")
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return errors.NewBadRequestError("Validation failed: " + err.Error())
	}

	ctx := c.Request().Context()
	grpcReq := &pb.CreateCashierRequest{
		MerchantId: int32(body.MerchantID),
		UserId:     int32(body.UserID),
		Name:       body.Name,
	}

	res, err := h.client.CreateCashier(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Cashier creation failed", zap.Error(err))
		return h.handleGrpcError(err, "CreateCashier")
	}

	so := h.mapping.ToApiResponseCashier(res)

	// Invalidate cache for related data
	h.cache.DeleteCashierCache(ctx, int(res.Data.Id))

	return c.JSON(http.StatusCreated, so)
}

// @Security Bearer
// Update handles the update of an existing cashier record.
// @Summary Update an existing cashier
// @Tags Cashier
// @Description Update an existing cashier record with the provided details
// @Accept json
// @Produce json
// @Param id path int true "Cashier ID"
// @Param UpdateCashierRequest body requests.UpdateCashierRequest true "Update cashier request"
// @Success 200 {object} response.ApiResponseCashier "Successfully updated cashier"
// @Failure 400 {object} errors.ApiError "Invalid request body or validation error"
// @Failure 500 {object} errors.ApiError "Failed to update cashier"
// @Router /api/cashier/update/{id} [post]
func (h *cashierHandleApi) UpdateCashier(c echo.Context) error {
	id := c.Param("id")
	idStr, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Debug("Invalid id parameter", zap.Error(err))
		return errors.NewBadRequestError("Invalid cashier ID")
	}

	var body requests.UpdateCashierRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request format", zap.Error(err))
		return errors.NewBadRequestError("Invalid request format")
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation failed", zap.Error(err))
		return errors.NewBadRequestError("Validation failed: " + err.Error())
	}

	ctx := c.Request().Context()
	grpcReq := &pb.UpdateCashierRequest{
		CashierId: int32(idStr),
		Name:      body.Name,
	}

	res, err := h.client.UpdateCashier(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Cashier update failed", zap.Error(err))
		return h.handleGrpcError(err, "UpdateCashier")
	}

	so := h.mapping.ToApiResponseCashier(res)

	// Invalidate cache for related data
	h.cache.DeleteCashierCache(ctx, idStr)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// TrashedCasher retrieves a trashed casher record by its ID.
// @Summary Retrieve a trashed casher
// @Tags Cashier
// @Description Retrieve a trashed casher record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Cashier ID"
// @Success 200 {object} response.ApiResponseCashierDeleteAt "Successfully retrieved trashed cashier"
// @Failure 400 {object} errors.ApiError "Invalid request body or validation error"
// @Failure 500 {object} errors.ApiError "Failed to retrieve trashed cashier"
// @Router /api/cashier/trashed/{id} [get]
func (h *cashierHandleApi) TrashedCashier(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid cashier ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid cashier ID")
	}

	ctx := c.Request().Context()
	grpcReq := &pb.FindByIdCashierRequest{Id: int32(id)}

	cashier, err := h.client.TrashedCashier(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to archive cashier", zap.Error(err))
		return h.handleGrpcError(err, "TrashedCashier")
	}

	so := h.mapping.ToApiResponseCashierDeleteAt(cashier)

	// Invalidate cache for related data
	h.cache.DeleteCashierCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreCashier restores a cashier record from the trash by its ID.
// @Summary Restore a trashed cashier
// @Tags Cashier
// @Description Restore a trashed cashier record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Cashier ID"
// @Success 200 {object} response.ApiResponseCashierDeleteAt "Successfully restored cashier"
// @Failure 400 {object} errors.ApiError "Invalid cashier ID"
// @Failure 500 {object} errors.ApiError "Failed to restore cashier"
// @Router /api/cashier/restore/{id} [post]
func (h *cashierHandleApi) RestoreCashier(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid cashier ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid cashier ID")
	}

	ctx := c.Request().Context()
	grpcReq := &pb.FindByIdCashierRequest{Id: int32(id)}

	cashier, err := h.client.RestoreCashier(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to restore cashier", zap.Error(err))
		return h.handleGrpcError(err, "RestoreCashier")
	}

	so := h.mapping.ToApiResponseCashierDeleteAt(cashier)

	// Invalidate cache for related data
	h.cache.DeleteCashierCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteCashierPermanent permanently deletes a cashier record by its ID.
// @Summary Permanently delete a cashier
// @Tags Cashier
// @Description Permanently delete a cashier record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "cashier ID"
// @Success 200 {object} response.ApiResponseCashierDelete "Successfully deleted cashier record permanently"
// @Failure 400 {object} errors.ApiError "Bad Request: Invalid ID"
// @Failure 500 {object} errors.ApiError "Failed to delete cashier:"
// @Router /api/cashier/delete/{id} [delete]
func (h *cashierHandleApi) DeleteCashierPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Debug("Invalid cashier ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid cashier ID")
	}

	ctx := c.Request().Context()
	grpcReq := &pb.FindByIdCashierRequest{Id: int32(id)}

	cashier, err := h.client.DeleteCashierPermanent(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to delete cashier", zap.Error(err))
		return h.handleGrpcError(err, "DeleteCashierPermanent")
	}

	so := h.mapping.ToApiResponseCashierDelete(cashier)

	// Invalidate cache for related data
	h.cache.DeleteCashierCache(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreCashier restores a cashier record from the trash by its ID.
// @Summary Restore a trashed cashier
// @Tags Cashier
// @Description Restore a trashed cashier record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Cashier ID"
// @Success 200 {object} response.ApiResponseCashierAll "Successfully restored cashier all"
// @Failure 400 {object} errors.ApiError "Invalid cashier ID"
// @Failure 500 {object} errors.ApiError "Failed to restore cashier"
// @Router /api/cashier/restore/all [post]
func (h *cashierHandleApi) RestoreAllCashier(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllCashier(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Error("Bulk cashier restoration failed", zap.Error(err))
		return h.handleGrpcError(err, "RestoreAllCashier")
	}

	h.logger.Info("All cashier accounts restored successfully")

	so := h.mapping.ToApiResponseCashierAll(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteCashierPermanent permanently deletes a cashier record by its ID.
// @Summary Permanently delete a cashier
// @Tags Cashier
// @Description Permanently delete a cashier record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "cashier ID"
// @Success 200 {object} response.ApiResponseCashierAll "Successfully deleted cashier record permanently"
// @Failure 400 {object} errors.ApiError "Bad Request: Invalid ID"
// @Failure 500 {object} errors.ApiError "Failed to delete cashier:"
// @Router /api/cashier/delete/all [post]
func (h *cashierHandleApi) DeleteAllCashierPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllCashierPermanent(ctx, &emptypb.Empty{})
	if err != nil {
		h.logger.Error("Bulk cashier deletion failed", zap.Error(err))
		return h.handleGrpcError(err, "DeleteAllCashierPermanent")
	}

	h.logger.Info("All cashier accounts permanently deleted")

	so := h.mapping.ToApiResponseCashierAll(res)

	return c.JSON(http.StatusOK, so)
}

func (h *cashierHandleApi) handleGrpcError(err error, operation string) *errors.AppError {
	st, ok := status.FromError(err)
	if !ok {
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}

	switch st.Code() {
	case codes.NotFound:
		return errors.NewNotFoundError("Cashier").WithInternal(err)

	case codes.AlreadyExists:
		return errors.NewConflictError("Cashier already exists").WithInternal(err)

	case codes.InvalidArgument:
		return errors.NewBadRequestError(st.Message()).WithInternal(err)

	case codes.PermissionDenied:
		return errors.ErrForbidden.WithInternal(err)

	case codes.Unauthenticated:
		return errors.ErrUnauthorized.WithInternal(err)

	case codes.ResourceExhausted:
		return errors.ErrTooManyRequests.WithInternal(err)

	case codes.Unavailable:
		return errors.NewServiceUnavailableError("Cashier service").WithInternal(err)

	case codes.DeadlineExceeded:
		return errors.ErrTimeout.WithInternal(err)

	default:
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}
}
