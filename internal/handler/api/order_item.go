package api

import (
	"net/http"
	orderitem_cache "pointofsale/internal/cache/api/order_item"
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
)

type orderItemHandleApi struct {
	client     pb.OrderItemServiceClient
	logger     logger.LoggerInterface
	mapping    response_api.OrderItemResponseMapper
	apiHandler errors.ApiHandler
	cache      orderitem_cache.OrderItemCache
}

func NewHandlerOrderItem(
	router *echo.Echo,
	client pb.OrderItemServiceClient,
	logger logger.LoggerInterface,
	mapping response_api.OrderItemResponseMapper,
	apiHandler errors.ApiHandler,
	cache orderitem_cache.OrderItemCache,
) *orderItemHandleApi {
	orderItemHandler := &orderItemHandleApi{
		client:     client,
		logger:     logger,
		mapping:    mapping,
		apiHandler: apiHandler,
		cache:      cache,
	}

	routerOrderItem := router.Group("/api/order-item")

	routerOrderItem.GET("", orderItemHandler.FindAllOrderItems)
	routerOrderItem.GET("/:order_id", orderItemHandler.FindOrderItemByOrder)
	routerOrderItem.GET("/active", orderItemHandler.FindByActive)
	routerOrderItem.GET("/trashed", orderItemHandler.FindByTrashed)

	return orderItemHandler
}

// @Security Bearer
// @Summary Find all order items
// @Tags Order-Item
// @Description Retrieve a list of all order items
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationOrderItem "List of order items"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve order item data"
// @Router /api/order-item [get]
func (h *orderItemHandleApi) FindAllOrderItems(c echo.Context) error {
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

	req := &requests.FindAllOrderItems{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedOrderItemsAll(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllOrderItemRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch order-items", zap.Error(err))
		return h.handleGrpcError(err, "FindAllOrderItems")
	}

	so := h.mapping.ToApiResponsePaginationOrderItem(res)

	h.cache.SetCachedOrderItemsAll(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active order items
// @Tags Order-Item
// @Description Retrieve a list of active order items
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationOrderItemDeleteAt "List of active order items"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve order item data"
// @Router /api/order-item/active [get]
func (h *orderItemHandleApi) FindByActive(c echo.Context) error {
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

	req := &requests.FindAllOrderItems{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedOrderItemActive(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllOrderItemRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch active order-items", zap.Error(err))
		return h.handleGrpcError(err, "FindByActive")
	}

	so := h.mapping.ToApiResponsePaginationOrderItemDeleteAt(res)

	h.cache.SetCachedOrderItemActive(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed order items
// @Tags Order-Item
// @Description Retrieve a list of trashed order items
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationOrderItemDeleteAt "List of trashed order items"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve order item data"
// @Router /api/order-item/trashed [get]
func (h *orderItemHandleApi) FindByTrashed(c echo.Context) error {
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

	req := &requests.FindAllOrderItems{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedOrderItemTrashed(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllOrderItemRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch archived order-items", zap.Error(err))
		return h.handleGrpcError(err, "FindByTrashed")
	}

	so := h.mapping.ToApiResponsePaginationOrderItemDeleteAt(res)

	h.cache.SetCachedOrderItemTrashed(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find order items by order ID
// @Tags Order-Item
// @Description Retrieve order items by order ID
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} response.ApiResponsesOrderItem "List of order items by order ID"
// @Failure 400 {object} response.ErrorResponse "Invalid order ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve order item data"
// @Router /api/order-item/order/{order_id} [get]
func (h *orderItemHandleApi) FindOrderItemByOrder(c echo.Context) error {
	orderID, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		h.logger.Debug("Invalid order ID format", zap.Error(err))
		return errors.NewBadRequestError("Please provide a valid order ID")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedOrderItems(ctx, orderID); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindByIdOrderItemRequest{
		Id: int32(orderID),
	}

	res, err := h.client.FindOrderItemByOrder(ctx, grpcReq)
	if err != nil {
		h.logger.Error("Failed to fetch order item details", zap.Error(err))
		return h.handleGrpcError(err, "FindOrderItemByOrder")
	}

	so := h.mapping.ToApiResponsesOrderItem(res)

	h.cache.SetCachedOrderItems(ctx, so)

	return c.JSON(http.StatusOK, so)
}

func (h *orderItemHandleApi) handleGrpcError(err error, operation string) *errors.AppError {
	st, ok := status.FromError(err)
	if !ok {
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}

	switch st.Code() {
	case codes.NotFound:
		return errors.NewNotFoundError("OrderItem").WithInternal(err)

	case codes.AlreadyExists:
		return errors.NewConflictError("OrderItem already exists").WithInternal(err)

	case codes.InvalidArgument:
		return errors.NewBadRequestError(st.Message()).WithInternal(err)

	case codes.PermissionDenied:
		return errors.ErrForbidden.WithInternal(err)

	case codes.Unauthenticated:
		return errors.ErrUnauthorized.WithInternal(err)

	case codes.ResourceExhausted:
		return errors.ErrTooManyRequests.WithInternal(err)

	case codes.Unavailable:
		return errors.NewServiceUnavailableError("OrderItem service").WithInternal(err)

	case codes.DeadlineExceeded:
		return errors.ErrTimeout.WithInternal(err)

	default:
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}
}
