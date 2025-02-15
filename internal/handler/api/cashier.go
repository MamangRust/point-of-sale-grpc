package api

import (
	"net/http"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_api "pointofsale/internal/mapper/response/api"
	"pointofsale/internal/pb"
	"pointofsale/pkg/logger"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cashierHandleApi struct {
	client  pb.CashierServiceClient
	logger  logger.LoggerInterface
	mapping response_api.CashierResponseMapper
}

func NewHandlerCashier(
	router *echo.Echo,
	client pb.CashierServiceClient,
	logger logger.LoggerInterface,
	mapping response_api.CashierResponseMapper,
) *cashierHandleApi {
	cashierHandler := &cashierHandleApi{
		client:  client,
		logger:  logger,
		mapping: mapping,
	}

	routerCashier := router.Group("/api/cashier")

	routerCashier.GET("", cashierHandler.FindAllCashier)
	routerCashier.GET("/:id", cashierHandler.FindById)
	routerCashier.GET("/active", cashierHandler.FindByActive)
	routerCashier.GET("/trashed", cashierHandler.FindByTrashed)

	routerCashier.POST("/create", cashierHandler.Create)
	routerCashier.POST("/update/:id", cashierHandler.Update)

	routerCashier.POST("/trashed/:id", cashierHandler.TrashedCashier)
	routerCashier.POST("/restore/:id", cashierHandler.RestoreCashier)
	routerCashier.DELETE("/permanent/:id", cashierHandler.DeleteCashierPermanent)

	routerCashier.POST("/restore/all", cashierHandler.RestoreAllCashier)
	routerCashier.POST("/permanent/all", cashierHandler.DeleteAllCashierPermanent)

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
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve cashier data"
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

	req := &pb.FindAllCashierRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve cashier data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashier data: ",
		})
	}

	so := h.mapping.ToApiResponsePaginationCashier(res)

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
// @Failure 400 {object} response.ErrorResponse "Invalid cashier ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve cashier data"
// @Router /api/cashier/{id} [get]
func (h *cashierHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid cashier ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCashierRequest{
		Id: int32(id),
	}

	cashier, err := h.client.FindById(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve cashier data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashier data: ",
		})
	}

	so := h.mapping.ToApiResponseCashier(cashier)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active cashier
// @Tags Cashier
// @Description Retrieve a list of active cashier
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponsePaginationCashierDeleteAt "List of active cashier"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve cashier data"
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

	req := &pb.FindAllCashierRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve cashier data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashier data: ",
		})
	}

	so := h.mapping.ToApiResponsePaginationCashierDeleteAt(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// FindByTrashed retrieves a list of trashed cashier records.
// @Summary Retrieve trashed cashier
// @Tags Cashier
// @Description Retrieve a list of trashed cashier records
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponsePaginationCashierDeleteAt "List of trashed cashier data"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve cashier data"
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

	req := &pb.FindAllCashierRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to retrieve cashier data", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashier data: ",
		})
	}

	so := h.mapping.ToApiResponsePaginationCashierDeleteAt(res)

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
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create cashier"
// @Router /api/cashier/create [post]
func (h *cashierHandleApi) Create(c echo.Context) error {
	var body requests.CreateCashierRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	req := &pb.CreateCashierRequest{
		MerchantId: int32(body.MerchantID),
		UserId:     int32(body.UserID),
		Name:       body.Name,
	}

	res, err := h.client.CreateCashier(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to create cashier", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create cashier: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseCashier(res)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// Update handles the update of an existing cashier record.
// @Summary Update an existing cashier
// @Tags Cashier
// @Description Update an existing cashier record with the provided details
// @Accept json
// @Produce json
// @Param UpdateCashierRequest body requests.UpdateCashierRequest true "Update cashier request"
// @Success 200 {object} pb.ApiResponseCashier "Successfully updated cashier"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to update cashier"
// @Router /api/cashier/update/{id} [post]
func (h *cashierHandleApi) Update(c echo.Context) error {
	var body requests.UpdateCashierRequest

	if err := c.Bind(&body); err != nil {
		h.logger.Debug("Invalid request body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	if err := body.Validate(); err != nil {
		h.logger.Debug("Validation Error", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Validation Error: " + err.Error(),
		})
	}

	ctx := c.Request().Context()

	req := &pb.UpdateCashierRequest{
		CashierId: int32(body.CashierID),
		Name:      body.Name,
	}

	res, err := h.client.UpdateCashier(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to update cashier", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update cashier: " + err.Error(),
		})
	}

	so := h.mapping.ToApiResponseCashier(res)

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
// @Success 200 {object} pb.ApiResponseCashierDeleteAt "Successfully retrieved trashed cashier"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed cashier"
// @Router /api/cashier/trashed/{id} [get]
func (h *cashierHandleApi) TrashedCashier(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid cashier ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCashierRequest{
		Id: int32(id),
	}

	cashier, err := h.client.TrashedCashier(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to trashed cashier", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed cashier: ",
		})
	}

	so := h.mapping.ToApiResponseCashierDeleteAt(cashier)

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
// @Success 200 {object} pb.ApiResponseCashierDeleteAt "Successfully restored cashier"
// @Failure 400 {object} response.ErrorResponse "Invalid cashier ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore cashier"
// @Router /api/cashier/restore/{id} [post]
func (h *cashierHandleApi) RestoreCashier(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid cashier ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCashierRequest{
		Id: int32(id),
	}

	cashier, err := h.client.RestoreCashier(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to restore cashier", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore cashier: ",
		})
	}

	so := h.mapping.ToApiResponseCashierDeleteAt(cashier)

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
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete cashier:"
// @Router /api/cashier/delete/{id} [delete]
func (h *cashierHandleApi) DeleteCashierPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid cashier ID", zap.Error(err))
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier ID",
		})
	}

	ctx := c.Request().Context()

	req := &pb.FindByIdCashierRequest{
		Id: int32(id),
	}

	cashier, err := h.client.DeleteCashierPermanent(ctx, req)

	if err != nil {
		h.logger.Debug("Failed to delete cashier", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete cashier: ",
		})
	}

	so := h.mapping.ToApiResponseCashierDelete(cashier)

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
// @Success 200 {object} pb.ApiResponseCashierAll "Successfully restored cashier all"
// @Failure 400 {object} response.ErrorResponse "Invalid cashier ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore cashier"
// @Router /api/cashier/restore/all [post]
func (h *cashierHandleApi) RestoreAllCashier(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllCashier(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to restore all cashier", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently restore all cashier",
		})
	}

	h.logger.Debug("Successfully restored all cashier")

	return c.JSON(http.StatusOK, res)
}

// @Security Bearer
// DeleteCashierPermanent permanently deletes a cashier record by its ID.
// @Summary Permanently delete a cashier
// @Tags Cashier
// @Description Permanently delete a cashier record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "cashier ID"
// @Success 200 {object} pb.ApiResponseCashierAll "Successfully deleted cashier record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete cashier:"
// @Router /api/cashier/delete/all [post]
func (h *cashierHandleApi) DeleteAllCashierPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllCashierPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Failed to permanently delete all cashier", zap.Error(err))

		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all cashier",
		})
	}

	h.logger.Debug("Successfully deleted all cashier permanently")

	return c.JSON(http.StatusOK, res)
}
