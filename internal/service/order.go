package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type orderService struct {
	orderRepository     repository.OrderRepository
	orderItemRepository repository.OrderItemRepository
	productRepository   repository.ProductRepository
	cashierRepository   repository.CashierRepository
	merchantRepository  repository.MerchantRepository
	logger              logger.LoggerInterface
	mapping             response_service.OrderResponseMapper
}

func NewOrderServiceMapper(
	orderRepository repository.OrderRepository,
	orderItemRepository repository.OrderItemRepository,
	cashierRepository repository.CashierRepository,
	merchantRepository repository.MerchantRepository,
	productRepository repository.ProductRepository,
	logger logger.LoggerInterface,
	mapping response_service.OrderResponseMapper) *orderService {
	return &orderService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
		productRepository:   productRepository,
		cashierRepository:   cashierRepository,
		merchantRepository:  merchantRepository,
		logger:              logger,
		mapping:             mapping,
	}
}

func (s *orderService) FindAll(page int, pageSize int, search string) ([]*response.OrderResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching orders",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orders, totalRecords, err := s.orderRepository.FindAllOrders(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve order list",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve order list",
			Code:    http.StatusInternalServerError,
		}
	}

	orderResponse := s.mapping.ToOrdersResponse(orders)

	s.logger.Debug("Successfully fetched order",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return orderResponse, &totalRecords, nil
}

func (s *orderService) FindById(order_id int) (*response.OrderResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching order by ID", zap.Int("order_id", order_id))

	order, err := s.orderRepository.FindById(order_id)

	if err != nil {
		s.logger.Error("Failed to retrieve order details",
			zap.Error(err),
			zap.Int("order_id", order_id))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Order with ID %d not found", order_id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve order details",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderResponse(order), nil
}

func (s *orderService) FindByActive(page int, pageSize int, search string) ([]*response.OrderResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching orders",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orders, totalRecords, err := s.orderRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to retrieve active orders from database",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve active orders",
			Code:    http.StatusInternalServerError,
		}
	}

	orderResponse := s.mapping.ToOrdersResponseDeleteAt(orders)

	s.logger.Debug("Successfully fetched order",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return orderResponse, &totalRecords, nil
}

func (s *orderService) FindByTrashed(page int, pageSize int, search string) ([]*response.OrderResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching orders",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orders, totalRecords, err := s.orderRepository.FindByTrashed(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to retrieve trashed orders from database",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve trashed orders",
			Code:    http.StatusInternalServerError,
		}
	}

	orderResponse := s.mapping.ToOrdersResponseDeleteAt(orders)

	s.logger.Debug("Successfully fetched order",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return orderResponse, &totalRecords, nil
}

func (s *orderService) FindMonthlyTotalRevenue(year int, month int) ([]*response.OrderMonthlyTotalRevenueResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if month <= 0 || month >= 12 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Month must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetMonthlyTotalRevenue(year, month)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderMonthlyTotalRevenues(res), nil
}

func (s *orderService) FindYearlyTotalRevenue(year int) ([]*response.OrderYearlyTotalRevenueResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetYearlyTotalRevenue(year)
	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderYearlyTotalRevenues(res), nil
}

func (s *orderService) FindMonthlyTotalRevenueById(year int, month int, order_id int) ([]*response.OrderMonthlyTotalRevenueResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if month <= 0 || month >= 12 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Month must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetMonthlyTotalRevenueById(year, month, order_id)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderMonthlyTotalRevenues(res), nil
}

func (s *orderService) FindYearlyTotalRevenueById(year int, order_id int) ([]*response.OrderYearlyTotalRevenueResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetYearlyTotalRevenueById(year, order_id)
	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderYearlyTotalRevenues(res), nil
}

func (s *orderService) FindMonthlyTotalRevenueByMerchant(year int, month int, merchant_id int) ([]*response.OrderMonthlyTotalRevenueResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if month <= 0 || month >= 12 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Month must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetMonthlyTotalRevenueByMerchant(year, month, merchant_id)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderMonthlyTotalRevenues(res), nil
}

func (s *orderService) FindYearlyTotalRevenueByMerchant(year int, merchant_id int) ([]*response.OrderYearlyTotalRevenueResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetYearlyTotalRevenueByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderYearlyTotalRevenues(res), nil
}

func (s *orderService) FindMonthlyOrder(year int) ([]*response.OrderMonthlyResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid year parameter",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetMonthlyOrder(year)
	if err != nil {
		s.logger.Error("failed to get monthly orders",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly orders data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderMonthlyPrices(res), nil
}

func (s *orderService) FindYearlyOrder(year int) ([]*response.OrderYearlyResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid year parameter",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetYearlyOrder(year)
	if err != nil {
		s.logger.Error("failed to get yearly orders",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly orders data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderYearlyPrices(res), nil
}

func (s *orderService) FindMonthlyOrderByMerchant(year int, merchant_id int) ([]*response.OrderMonthlyResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid year parameter",
			Code:    http.StatusBadRequest,
		}
	}
	if merchant_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid merchant ID",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetMonthlyOrderByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get monthly orders by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve monthly orders for merchant %d", merchant_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderMonthlyPrices(res), nil
}

func (s *orderService) FindYearlyOrderByMerchant(year int, merchant_id int) ([]*response.OrderYearlyResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid year parameter",
			Code:    http.StatusBadRequest,
		}
	}
	if merchant_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Invalid merchant ID",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.orderRepository.GetYearlyOrderByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get yearly orders by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve yearly orders for merchant %d", merchant_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderYearlyPrices(res), nil
}

func (s *orderService) CreateOrder(req *requests.CreateOrderRequest) (*response.OrderResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new order",
		zap.Int("merchantID", req.MerchantID),
		zap.Int("cashierID", req.CashierID))

	_, err := s.merchantRepository.FindById(req.MerchantID)
	if err != nil {
		s.logger.Error("Merchant not found for order creation",
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "not_found",
			Message: fmt.Sprintf("Merchant with ID %d not found", req.MerchantID),
			Code:    http.StatusNotFound,
		}
	}

	_, err = s.cashierRepository.FindById(req.CashierID)

	if err != nil {
		s.logger.Error("Cashier not found for order creation",
			zap.Int("cashierID", req.CashierID),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "not_found",
			Message: fmt.Sprintf("Cashier with ID %d not found", req.CashierID),
			Code:    http.StatusNotFound,
		}
	}

	order, err := s.orderRepository.CreateOrder(&requests.CreateOrderRecordRequest{
		MerchantID: req.MerchantID,
		CashierID:  req.CashierID,
	})
	if err != nil {
		s.logger.Error("Failed to create order record",
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create order record",
			Code:    http.StatusInternalServerError,
		}
	}

	for _, item := range req.Items {
		product, err := s.productRepository.FindById(item.ProductID)
		if err != nil {
			s.logger.Error("Product not found for order item",
				zap.Int("productID", item.ProductID),
				zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found", item.ProductID),
				Code:    http.StatusNotFound,
			}
		}

		if product.CountInStock < item.Quantity {
			s.logger.Error("Insufficient product stock",
				zap.Int("productID", item.ProductID),
				zap.Int("requested", item.Quantity),
				zap.Int("available", product.CountInStock))
			return nil, &response.ErrorResponse{
				Status: "invalid_request",
				Message: fmt.Sprintf("Insufficient stock for product %d (requested %d, available %d)",
					item.ProductID, item.Quantity, product.CountInStock),
				Code: http.StatusBadRequest,
			}
		}

		_, err = s.orderItemRepository.CreateOrderItem(&requests.CreateOrderItemRecordRequest{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
		if err != nil {
			s.logger.Error("Failed to create order item",
				zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to create order item",
				Code:    http.StatusInternalServerError,
			}
		}

		product.CountInStock -= item.Quantity
		_, err = s.productRepository.UpdateProductCountStock(product.ID, product.CountInStock)
		if err != nil {
			s.logger.Error("Failed to update product stock",
				zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "Failed to update product stock",
				Code:    http.StatusInternalServerError,
			}
		}
	}

	totalPrice, err := s.orderItemRepository.CalculateTotalPrice(order.ID)
	if err != nil {
		s.logger.Error("Failed to calculate order total price",
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to calculate order total",
			Code:    http.StatusInternalServerError,
		}
	}

	res, err := s.orderRepository.UpdateOrder(&requests.UpdateOrderRecordRequest{
		OrderID:    order.ID,
		TotalPrice: int(*totalPrice),
	})
	if err != nil {
		s.logger.Error("Failed to update order total price",
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update order total price",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderResponse(res), nil
}

func (s *orderService) UpdateOrder(req *requests.UpdateOrderRequest) (*response.OrderResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating order",
		zap.Int("orderID", *req.OrderID))

	_, err := s.orderRepository.FindById(*req.OrderID)

	if err != nil {
		s.logger.Error("Order not found for update",
			zap.Int("orderID", *req.OrderID),
			zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "not_found",
			Message: fmt.Sprintf("Order with ID %d not found", req.OrderID),
			Code:    http.StatusNotFound,
		}
	}

	for _, item := range req.Items {
		product, err := s.productRepository.FindById(item.ProductID)

		if err != nil {
			s.logger.Error("Product not found for order update",
				zap.Int("productID", item.ProductID),
				zap.Error(err))

			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found", item.ProductID),
				Code:    http.StatusNotFound,
			}
		}

		if item.OrderItemID > 0 {
			_, err := s.orderItemRepository.UpdateOrderItem(&requests.UpdateOrderItemRecordRequest{
				OrderItemID: item.OrderItemID,
				ProductID:   item.ProductID,
				Quantity:    item.Quantity,
				Price:       product.Price,
			})
			if err != nil {
				s.logger.Error("Failed to update order item",
					zap.Error(err))
				return nil, &response.ErrorResponse{
					Status:  "error",
					Message: "Failed to update order item",
					Code:    http.StatusInternalServerError,
				}
			}
		} else {
			if product.CountInStock < item.Quantity {
				s.logger.Error("Insufficient product stock for new order item",
					zap.Int("productID", item.ProductID),
					zap.Int("requested", item.Quantity),
					zap.Int("available", product.CountInStock))
				return nil, &response.ErrorResponse{
					Status: "invalid_request",
					Message: fmt.Sprintf("Insufficient stock for product %d (requested %d, available %d)",
						item.ProductID, item.Quantity, product.CountInStock),
					Code: http.StatusBadRequest,
				}
			}

			_, err := s.orderItemRepository.CreateOrderItem(&requests.CreateOrderItemRecordRequest{
				OrderID:   *req.OrderID,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})
			if err != nil {
				s.logger.Error("Failed to add new order item",
					zap.Error(err))
				return nil, &response.ErrorResponse{
					Status:  "error",
					Message: "Failed to add new order item",
					Code:    http.StatusInternalServerError,
				}
			}

			product.CountInStock -= item.Quantity
			_, err = s.productRepository.UpdateProductCountStock(product.ID, product.CountInStock)
			if err != nil {
				s.logger.Error("Failed to update product stock",
					zap.Error(err))
				return nil, &response.ErrorResponse{
					Status:  "error",
					Message: "Failed to update product stock",
					Code:    http.StatusInternalServerError,
				}
			}
		}
	}

	totalPrice, err := s.orderItemRepository.CalculateTotalPrice(*req.OrderID)

	if err != nil {
		s.logger.Error("Failed to calculate updated order total",
			zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to calculate order total",
			Code:    http.StatusInternalServerError,
		}
	}

	res, err := s.orderRepository.UpdateOrder(&requests.UpdateOrderRecordRequest{
		OrderID:    *req.OrderID,
		TotalPrice: int(*totalPrice),
	})
	if err != nil {
		s.logger.Error("Failed to update order total price",
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update order total price",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderResponse(res), nil
}

func (s *orderService) TrashedOrder(orderID int) (*response.OrderResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Moving order to trash",
		zap.Int("order_id", orderID))

	order, err := s.orderRepository.FindById(orderID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Order with ID %d not found", orderID),
				Code:    http.StatusNotFound,
			}
		}
		s.logger.Error("Failed to fetch order",
			zap.Int("order_id", orderID),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to verify order existence",
			Code:    http.StatusInternalServerError,
		}
	}

	if order.DeletedAt != nil {
		return nil, &response.ErrorResponse{
			Status:  "already_trashed",
			Message: fmt.Sprintf("Order with ID %d is already trashed", orderID),
			Code:    http.StatusBadRequest,
		}
	}

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(orderID)
	if err != nil {
		s.logger.Error("Failed to retrieve order items for trashing",
			zap.Int("order_id", orderID),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve order items",
			Code:    http.StatusInternalServerError,
		}
	}

	for _, item := range orderItems {
		if item.DeletedAt != nil {
			s.logger.Debug("Order item already trashed, skipping",
				zap.Int("order_item_id", item.ID))
			continue
		}

		trashedItem, err := s.orderItemRepository.TrashedOrderItem(item.ID)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				s.logger.Debug("Order item not found - may have been deleted",
					zap.Int("order_item_id", item.ID))
				continue
			}

			s.logger.Error("Failed to move order item to trash",
				zap.Int("order_item_id", item.ID),
				zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: fmt.Sprintf("Failed to move order item %d to trash", item.ID),
				Code:    http.StatusInternalServerError,
			}
		}

		s.logger.Debug("Order item trashed successfully",
			zap.Int("order_item_id", trashedItem.ID),
			zap.String("deleted_at", *trashedItem.DeletedAt))
	}

	trashedOrder, err := s.orderRepository.TrashedOrder(orderID)

	if err != nil {
		s.logger.Error("Failed to move order to trash",
			zap.Int("order_id", orderID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Order with ID %d not found", orderID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move order to trash",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Order moved to trash successfully",
		zap.Int("order_id", orderID),
		zap.String("deleted_at", *trashedOrder.DeletedAt))

	return s.mapping.ToOrderResponseDeleteAt(trashedOrder), nil
}

func (s *orderService) RestoreOrder(order_id int) (*response.OrderResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring order from trash",
		zap.Int("order_id", order_id))

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(order_id)
	if err != nil {
		s.logger.Error("Failed to retrieve order items for restoration",
			zap.Int("order_id", order_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve order items",
			Code:    http.StatusInternalServerError,
		}
	}

	for _, item := range orderItems {
		_, err := s.orderItemRepository.RestoreOrderItem(item.ID)
		if err != nil {
			s.logger.Error("Failed to restore order item from trash",
				zap.Int("order_item_id", item.ID),
				zap.Error(err))
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: fmt.Sprintf("Failed to restore order item %d", item.ID),
				Code:    http.StatusInternalServerError,
			}
		}
	}

	order, err := s.orderRepository.RestoreOrder(order_id)
	if err != nil {
		s.logger.Error("Failed to restore order from trash",
			zap.Int("order_id", order_id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Order with ID %d not found in trash", order_id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore order from trash",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderResponseDeleteAt(order), nil
}

func (s *orderService) DeleteOrderPermanent(order_id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting order",
		zap.Int("order_id", order_id))

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(order_id)
	if err != nil {
		s.logger.Error("Failed to retrieve order items for permanent deletion",
			zap.Int("order_id", order_id),
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve order items",
			Code:    http.StatusInternalServerError,
		}
	}

	for _, item := range orderItems {
		_, err := s.orderItemRepository.DeleteOrderItemPermanent(item.ID)
		if err != nil {
			s.logger.Error("Failed to permanently delete order item",
				zap.Int("order_item_id", item.ID),
				zap.Error(err))
			return false, &response.ErrorResponse{
				Status:  "error",
				Message: fmt.Sprintf("Failed to permanently delete order item %d", item.ID),
				Code:    http.StatusInternalServerError,
			}
		}
	}

	success, err := s.orderRepository.DeleteOrderPermanent(order_id)
	if err != nil {
		s.logger.Error("Failed to permanently delete order",
			zap.Int("order_id", order_id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return false, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Order with ID %d not found", order_id),
				Code:    http.StatusNotFound,
			}
		}
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete order",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *orderService) RestoreAllOrder() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed orders")

	successItems, err := s.orderItemRepository.RestoreAllOrderItem()
	if err != nil || !successItems {
		s.logger.Error("Failed to restore all order items",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all order items",
			Code:    http.StatusInternalServerError,
		}
	}

	success, err := s.orderRepository.RestoreAllOrder()
	if err != nil || !success {
		s.logger.Error("Failed to restore all orders",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all orders",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *orderService) DeleteAllOrderPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all trashed orders")

	successItems, err := s.orderItemRepository.DeleteAllOrderPermanent()
	if err != nil || !successItems {
		s.logger.Error("Failed to permanently delete all order items",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all order items",
			Code:    http.StatusInternalServerError,
		}
	}

	success, err := s.orderRepository.DeleteAllOrderPermanent()
	if err != nil || !success {
		s.logger.Error("Failed to permanently delete all orders",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all orders",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}
