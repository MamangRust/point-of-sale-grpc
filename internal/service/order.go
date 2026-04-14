package service

import (
	"context"
	order_cache "pointofsale/internal/cache/order"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/errorhandler"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/cashier_errors"
	"pointofsale/pkg/errors/merchant_errors"
	"pointofsale/pkg/errors/order_errors"
	orderitem_errors "pointofsale/pkg/errors/order_item_errors"
	"pointofsale/pkg/errors/product_errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

type orderService struct {
	orderRepository     repository.OrderRepository
	orderItemRepository repository.OrderItemRepository
	productRepository   repository.ProductRepository
	cashierRepository   repository.CashierRepository
	merchantRepository  repository.MerchantRepository
	logger              logger.LoggerInterface
	observability       observability.TraceLoggerObservability
	cache               order_cache.OrderMencache
}

type OrderServiceDeps struct {
	OrderRepo     repository.OrderRepository
	OrderItemRepo repository.OrderItemRepository
	ProductRepo   repository.ProductRepository
	CashierRepo   repository.CashierRepository
	MerchantRepo  repository.MerchantRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
	Cache         order_cache.OrderMencache
}

func NewOrderService(deps OrderServiceDeps) *orderService {
	return &orderService{
		orderRepository:     deps.OrderRepo,
		orderItemRepository: deps.OrderItemRepo,
		productRepository:   deps.ProductRepo,
		cashierRepository:   deps.CashierRepo,
		merchantRepository:  deps.MerchantRepo,
		logger:              deps.Logger,
		observability:       deps.Observability,
		cache:               deps.Cache,
	}
}

func (s *orderService) FindAllOrders(ctx context.Context, req *requests.FindAllOrders) ([]*db.GetOrdersRow, *int, error) {
	const method = "FindAllOrders"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetOrderAllCache(ctx, req); found {
		logSuccess("Successfully retrieved all order records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	orders, err := s.orderRepository.FindAllOrders(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetOrdersRow](
			s.logger,
			order_errors.ErrFailedFindAllOrders,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(orders) > 0 {
		totalCount = int(orders[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetOrderAllCache(ctx, req, orders, &totalCount)

	logSuccess("Successfully fetched orders",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return orders, &totalCount, nil
}

func (s *orderService) FindById(ctx context.Context, order_id int) (*db.GetOrderByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("order_id", order_id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedOrderCache(ctx, order_id); found {
		logSuccess("Successfully retrieved order record from cache",
			zap.Int("order_id", order_id))
		return data, nil
	}

	order, err := s.orderRepository.FindById(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetOrderByIDRow](
			s.logger,
			order_errors.ErrFailedFindOrderById,
			method,
			span,
			zap.Int("order_id", order_id))
	}

	s.cache.SetCachedOrderCache(ctx, order)

	logSuccess("Successfully fetched order",
		zap.Int("order_id", order_id))

	return order, nil
}

func (s *orderService) FindByActive(ctx context.Context, req *requests.FindAllOrders) ([]*db.GetOrdersActiveRow, *int, error) {
	const method = "FindByActive"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetOrderActiveCache(ctx, req); found {
		logSuccess("Successfully retrieved active order records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	orders, err := s.orderRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetOrdersActiveRow](
			s.logger,
			order_errors.ErrFailedFindOrdersByActive,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(orders) > 0 {
		totalCount = int(orders[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetOrderActiveCache(ctx, req, orders, &totalCount)

	logSuccess("Successfully fetched active orders",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return orders, &totalCount, nil
}

func (s *orderService) FindByTrashed(ctx context.Context, req *requests.FindAllOrders) ([]*db.GetOrdersTrashedRow, *int, error) {
	const method = "FindByTrashed"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetOrderTrashedCache(ctx, req); found {
		logSuccess("Successfully retrieved trashed order records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	orders, err := s.orderRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetOrdersTrashedRow](
			s.logger,
			order_errors.ErrFailedFindOrdersByTrashed,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(orders) > 0 {
		totalCount = int(orders[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetOrderTrashedCache(ctx, req, orders, &totalCount)

	logSuccess("Successfully fetched trashed orders",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return orders, &totalCount, nil
}

func (s *orderService) FindByMerchant(ctx context.Context, req *requests.FindAllOrderMerchant) ([]*db.GetOrdersByMerchantRow, *int, error) {
	const method = "FindByMerchant"

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", req.Page),
		attribute.Int("pageSize", req.PageSize),
		attribute.String("search", req.Search),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedOrderMerchant(ctx, req); found {
		logSuccess("Successfully retrieved merchant order records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.Int("merchant_id", req.MerchantID))
		return data, total, nil
	}

	orders, err := s.orderRepository.FindByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetOrdersByMerchantRow](
			s.logger,
			order_errors.ErrFailedFindOrdersByMerchant,
			method,
			span,
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.Int("merchant_id", req.MerchantID))
	}

	var totalCount int

	if len(orders) > 0 {
		totalCount = int(orders[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedOrderMerchant(ctx, req, orders, &totalCount)

	logSuccess("Successfully fetched merchant orders",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize),
		zap.Int("merchant_id", req.MerchantID))

	return orders, &totalCount, nil
}

func (s *orderService) FindMonthlyTotalRevenue(ctx context.Context, req *requests.MonthTotalRevenue) ([]*db.GetMonthlyTotalRevenueRow, error) {
	const method = "FindMonthlyTotalRevenue"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTotalRevenueCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total revenue from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.orderRepository.GetMonthlyTotalRevenue(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalRevenueRow](
			s.logger,
			order_errors.ErrFailedFindMonthlyTotalRevenue,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
	}

	s.cache.SetMonthlyTotalRevenueCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total revenue",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindYearlyTotalRevenue(ctx context.Context, year int) ([]*db.GetYearlyTotalRevenueRow, error) {
	const method = "FindYearlyTotalRevenue"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTotalRevenueCache(ctx, year); found {
		logSuccess("Successfully retrieved yearly total revenue from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.orderRepository.GetYearlyTotalRevenue(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalRevenueRow](
			s.logger,
			order_errors.ErrFailedFindYearlyTotalRevenue,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetYearlyTotalRevenueCache(ctx, year, res)

	logSuccess("Successfully fetched yearly total revenue",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindMonthlyTotalRevenueById(ctx context.Context, req *requests.MonthTotalRevenueOrder) ([]*db.GetMonthlyTotalRevenueByIdRow, error) {
	const method = "FindMonthlyTotalRevenueById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("order_id", req.OrderID))

	defer func() {
		end(status)
	}()

	res, err := s.orderRepository.GetMonthlyTotalRevenueById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalRevenueByIdRow](
			s.logger,
			order_errors.ErrFailedFindMonthlyTotalRevenueById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("order_id", req.OrderID))
	}

	logSuccess("Successfully fetched monthly total revenue by order",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("order_id", req.OrderID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindYearlyTotalRevenueById(ctx context.Context, req *requests.YearTotalRevenueOrder) ([]*db.GetYearlyTotalRevenueByIdRow, error) {
	const method = "FindYearlyTotalRevenueById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("order_id", req.OrderID))

	defer func() {
		end(status)
	}()

	res, err := s.orderRepository.GetYearlyTotalRevenueById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalRevenueByIdRow](
			s.logger,
			order_errors.ErrFailedFindYearlyTotalRevenueById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("order_id", req.OrderID))
	}

	logSuccess("Successfully fetched yearly total revenue by order",
		zap.Int("year", req.Year),
		zap.Int("order_id", req.OrderID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *requests.MonthTotalRevenueMerchant) ([]*db.GetMonthlyTotalRevenueByMerchantRow, error) {
	const method = "FindMonthlyTotalRevenueByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTotalRevenueByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total revenue by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.orderRepository.GetMonthlyTotalRevenueByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalRevenueByMerchantRow](
			s.logger,
			order_errors.ErrFailedFindMonthlyTotalRevenueByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetMonthlyTotalRevenueByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total revenue by merchant",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *requests.YearTotalRevenueMerchant) ([]*db.GetYearlyTotalRevenueByMerchantRow, error) {
	const method = "FindYearlyTotalRevenueByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTotalRevenueByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly total revenue by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.orderRepository.GetYearlyTotalRevenueByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalRevenueByMerchantRow](
			s.logger,
			order_errors.ErrFailedFindYearlyTotalRevenueByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetYearlyTotalRevenueByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched yearly total revenue by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindMonthlyOrder(ctx context.Context, year int) ([]*db.GetMonthlyOrderRow, error) {
	const method = "FindMonthlyOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyOrderCache(ctx, year); found {
		logSuccess("Successfully retrieved monthly order from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.orderRepository.GetMonthlyOrder(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyOrderRow](
			s.logger,
			order_errors.ErrFailedFindMonthlyOrder,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetMonthlyOrderCache(ctx, year, res)

	logSuccess("Successfully fetched monthly order",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindYearlyOrder(ctx context.Context, year int) ([]*db.GetYearlyOrderRow, error) {
	const method = "FindYearlyOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyOrderCache(ctx, year); found {
		logSuccess("Successfully retrieved yearly order from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.orderRepository.GetYearlyOrder(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyOrderRow](
			s.logger,
			order_errors.ErrFailedFindYearlyOrder,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetYearlyOrderCache(ctx, year, res)

	logSuccess("Successfully fetched yearly order",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindMonthlyOrderByMerchant(ctx context.Context, req *requests.MonthOrderMerchant) ([]*db.GetMonthlyOrderByMerchantRow, error) {
	const method = "FindMonthlyOrderByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyOrderByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly order by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.orderRepository.GetMonthlyOrderByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyOrderByMerchantRow](
			s.logger,
			order_errors.ErrFailedFindMonthlyOrderByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetMonthlyOrderByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched monthly order by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) FindYearlyOrderByMerchant(ctx context.Context, req *requests.YearOrderMerchant) ([]*db.GetYearlyOrderByMerchantRow, error) {
	const method = "FindYearlyOrderByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyOrderByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly order by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.orderRepository.GetYearlyOrderByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyOrderByMerchantRow](
			s.logger,
			order_errors.ErrFailedFindYearlyOrderByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetYearlyOrderByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched yearly order by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *orderService) CreateOrder(ctx context.Context, req *requests.CreateOrderRequest) (*db.UpdateOrderRow, error) {
	const method = "CreateOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", req.MerchantID),
		attribute.Int("cashier_id", req.CashierID))

	defer func() {
		end(status)
	}()

	_, err := s.merchantRepository.FindById(ctx, req.MerchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchant_id", req.MerchantID))
	}

	_, err = s.cashierRepository.FindById(ctx, req.CashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierById,
			method,
			span,
			zap.Int("cashier_id", req.CashierID))
	}

	order, err := s.orderRepository.CreateOrder(ctx, &requests.CreateOrderRecordRequest{
		MerchantID: req.MerchantID,
		CashierID:  req.CashierID,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			order_errors.ErrFailedCreateOrder,
			method,
			span,
			zap.Int("order_id", int(order.OrderID)))
	}

	for _, item := range req.Items {
		product, err := s.productRepository.FindById(ctx, item.ProductID)
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.UpdateOrderRow](
				s.logger,
				product_errors.ErrFailedFindProductById,
				method,
				span,
				zap.Int("product_id", item.ProductID))
		}

		if product.CountInStock < int32(item.Quantity) {
			status = "error"
			return errorhandler.HandleError[*db.UpdateOrderRow](
				s.logger,
				order_errors.ErrFailedInvalidCountInStock,
				method,
				span,
				zap.Int("product_id", item.ProductID),
				zap.Int("requested", item.Quantity),
				zap.Int("available", int(product.CountInStock)))
		}

		_, err = s.orderItemRepository.CreateOrderItem(ctx, &requests.CreateOrderItemRecordRequest{
			OrderID:   int(order.OrderID),
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     int(product.Price),
		})
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.UpdateOrderRow](
				s.logger,
				orderitem_errors.ErrFailedCreateOrderItem,
				method,
				span,
				zap.Int("order_id", int(order.OrderID)),
				zap.Int("product_id", item.ProductID))
		}

		product.CountInStock -= int32(item.Quantity)
		_, err = s.productRepository.UpdateProductCountStock(ctx, int(product.ProductID), int(product.CountInStock))
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.UpdateOrderRow](
				s.logger,
				product_errors.ErrFailedUpdateProduct,
				method,
				span,
				zap.Int("product_id", int(product.ProductID)))
		}
	}

	totalPrice, err := s.orderItemRepository.CalculateTotalPrice(ctx, int(order.OrderID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			orderitem_errors.ErrFailedCalculateTotal,
			method,
			span,
			zap.Int("order_id", int(order.OrderID)))
	}

	res, err := s.orderRepository.UpdateOrder(ctx, &requests.UpdateOrderRecordRequest{
		OrderID:    int(order.OrderID),
		TotalPrice: int(*totalPrice),
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			order_errors.ErrFailedUpdateOrder,
			method,
			span,
			zap.Int("order_id", int(order.OrderID)))
	}

	logSuccess("Successfully created order",
		zap.Int("order_id", int(order.OrderID)))

	return res, nil
}

func (s *orderService) UpdateOrder(ctx context.Context, req *requests.UpdateOrderRequest) (*db.UpdateOrderRow, error) {
	const method = "UpdateOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("order_id", *req.OrderID))

	defer func() {
		end(status)
	}()

	_, err := s.orderRepository.FindById(ctx, *req.OrderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			order_errors.ErrFailedFindOrderById,
			method,
			span,
			zap.Int("order_id", *req.OrderID))
	}

	for _, item := range req.Items {
		product, err := s.productRepository.FindById(ctx, item.ProductID)
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.UpdateOrderRow](
				s.logger,
				product_errors.ErrFailedFindProductById,
				method,
				span,
				zap.Int("product_id", item.ProductID))
		}

		if item.OrderItemID > 0 {
			_, err := s.orderItemRepository.UpdateOrderItem(ctx, &requests.UpdateOrderItemRecordRequest{
				OrderItemID: item.OrderItemID,
				ProductID:   item.ProductID,
				Quantity:    item.Quantity,
				Price:       int(product.Price),
			})
			if err != nil {
				status = "error"
				return errorhandler.HandleError[*db.UpdateOrderRow](
					s.logger,
					orderitem_errors.ErrFailedUpdateOrderItem,
					method,
					span,
					zap.Int("order_item_id", item.OrderItemID))
			}
		} else {
			if product.CountInStock < int32(item.Quantity) {
				status = "error"
				return errorhandler.HandleError[*db.UpdateOrderRow](
					s.logger,
					order_errors.ErrFailedInvalidCountInStock,
					method,
					span,
					zap.Int("product_id", item.ProductID),
					zap.Int("requested", item.Quantity),
					zap.Int("available", int(product.CountInStock)))
			}

			_, err := s.orderItemRepository.CreateOrderItem(ctx, &requests.CreateOrderItemRecordRequest{
				OrderID:   *req.OrderID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     int(product.Price),
			})
			if err != nil {
				status = "error"
				return errorhandler.HandleError[*db.UpdateOrderRow](
					s.logger,
					orderitem_errors.ErrFailedCreateOrderItem,
					method,
					span,
					zap.Int("order_id", *req.OrderID),
					zap.Int("product_id", item.ProductID))
			}

			product.CountInStock -= int32(item.Quantity)
			_, err = s.productRepository.UpdateProductCountStock(ctx, int(product.ProductID), int(product.CountInStock))
			if err != nil {
				status = "error"
				return errorhandler.HandleError[*db.UpdateOrderRow](
					s.logger,
					product_errors.ErrFailedUpdateProduct,
					method,
					span,
					zap.Int("product_id", int(product.ProductID)))
			}
		}
	}

	totalPrice, err := s.orderItemRepository.CalculateTotalPrice(ctx, *req.OrderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			orderitem_errors.ErrFailedCalculateTotal,
			method,
			span,
			zap.Int("order_id", *req.OrderID))
	}

	res, err := s.orderRepository.UpdateOrder(ctx, &requests.UpdateOrderRecordRequest{
		OrderID:    *req.OrderID,
		TotalPrice: int(*totalPrice),
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateOrderRow](
			s.logger,
			order_errors.ErrFailedUpdateOrder,
			method,
			span,
			zap.Int("order_id", *req.OrderID))
	}

	logSuccess("Successfully updated order",
		zap.Int("order_id", *req.OrderID))

	return res, nil
}

func (s *orderService) TrashedOrder(ctx context.Context, order_id int) (*db.Order, error) {
	const method = "TrashedOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("order_id", order_id))

	defer func() {
		end(status)
	}()

	_, err := s.orderRepository.FindById(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Order](
			s.logger,
			order_errors.ErrFailedFindOrderById,
			method,
			span,
			zap.Int("order_id", order_id),
		)
	}

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Order](
			s.logger,
			order_errors.ErrFailedTrashOrder,
			method,
			span,
			zap.Int("order_id", order_id),
		)
	}

	for _, item := range orderItems {
		_, err := s.orderItemRepository.TrashedOrderItem(ctx, int(item.OrderItemID))
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.Order](
				s.logger,
				orderitem_errors.ErrFailedTrashedOrderItem,
				method,
				span,
				zap.Int("order_item_id", int(item.OrderItemID)),
			)
		}
	}

	trashedOrder, err := s.orderRepository.TrashedOrder(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Order](
			s.logger,
			order_errors.ErrFailedTrashOrder,
			method,
			span,
			zap.Int("order_id", order_id),
		)
	}

	logSuccess("Order moved to trash successfully",
		zap.Int("order_id", order_id),
		zap.String("deleted_at", trashedOrder.DeletedAt.Time.String()))

	return trashedOrder, nil
}

func (s *orderService) RestoreOrder(ctx context.Context, order_id int) (*db.Order, error) {
	const method = "RestoreOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("order_id", order_id))

	defer func() {
		end(status)
	}()

	orderItems, err := s.orderItemRepository.FindOrderItemByOrderTrashed(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Order](
			s.logger,
			orderitem_errors.ErrFailedFindOrderItemByOrder,
			method,
			span,
			zap.Int("order_id", order_id))
	}

	for _, item := range orderItems {
		_, err := s.orderItemRepository.RestoreOrderItem(ctx, int(item.OrderItemID))
		if err != nil {
			status = "error"
			return errorhandler.HandleError[*db.Order](
				s.logger,
				orderitem_errors.ErrFailedRestoreOrderItem,
				method,
				span,
				zap.Int("order_item_id", int(item.OrderItemID)))
		}
	}

	order, err := s.orderRepository.RestoreOrder(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Order](
			s.logger,
			order_errors.ErrFailedRestoreOrder,
			method,
			span,
			zap.Int("order_id", order_id))
	}

	s.cache.DeleteOrderCache(ctx, order_id)

	logSuccess("Successfully restored order",
		zap.Int("order_id", order_id))

	return order, nil
}

func (s *orderService) DeleteOrderPermanent(ctx context.Context, order_id int) (bool, error) {
	const method = "DeleteOrderPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("order_id", order_id))

	defer func() {
		end(status)
	}()

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			orderitem_errors.ErrFailedFindOrderItemByOrder,
			method,
			span,
			zap.Int("order_id", order_id))
	}

	for _, item := range orderItems {
		_, err := s.orderItemRepository.DeleteOrderItemPermanent(ctx, int(item.OrderItemID))
		if err != nil {
			status = "error"
			return errorhandler.HandleError[bool](
				s.logger,
				orderitem_errors.ErrFailedDeleteOrderItem,
				method,
				span,
				zap.Int("order_item_id", int(item.OrderItemID)))
		}
	}

	success, err := s.orderRepository.DeleteOrderPermanent(ctx, order_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			order_errors.ErrFailedDeleteOrderPermanent,
			method,
			span,
			zap.Int("order_id", order_id))
	}

	s.cache.DeleteOrderCache(ctx, order_id)

	logSuccess("Successfully permanently deleted order",
		zap.Int("order_id", order_id))

	return success, nil
}

func (s *orderService) RestoreAllOrder(ctx context.Context) (bool, error) {
	const method = "RestoreAllOrder"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	successItems, err := s.orderItemRepository.RestoreAllOrderItem(ctx)
	if err != nil || !successItems {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			orderitem_errors.ErrFailedRestoreAllOrderItem,
			method,
			span)
	}

	success, err := s.orderRepository.RestoreAllOrder(ctx)
	if err != nil || !success {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			order_errors.ErrFailedRestoreAllOrder,
			method,
			span)
	}

	s.logger.Debug("All order caches should be invalidated after restore all operation")

	logSuccess("Successfully restored all trashed orders")

	return success, nil
}

func (s *orderService) DeleteAllOrderPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllOrderPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	successItems, err := s.orderItemRepository.DeleteAllOrderPermanent(ctx)
	if err != nil || !successItems {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			orderitem_errors.ErrFailedDeleteAllOrderItem,
			method,
			span)
	}

	success, err := s.orderRepository.DeleteAllOrderPermanent(ctx)
	if err != nil || !success {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			order_errors.ErrFailedDeleteAllOrderPermanent,
			method,
			span)
	}

	s.logger.Debug("All order caches should be invalidated after delete all operation")

	logSuccess("Successfully permanently deleted all trashed orders")

	return success, nil
}
