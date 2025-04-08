package service

import (
	"database/sql"
	"errors"
	"net/http"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type orderItemService struct {
	orderItemRepository repository.OrderItemRepository
	logger              logger.LoggerInterface
	mapping             response_service.OrderItemResponseMapper
}

func NewOrderItemService(
	orderItemRepository repository.OrderItemRepository,
	logger logger.LoggerInterface,
	mapping response_service.OrderItemResponseMapper,
) *orderItemService {
	return &orderItemService{
		orderItemRepository: orderItemRepository,
		logger:              logger,
		mapping:             mapping,
	}
}

func (s *orderItemService) FindAllOrderItems(search string, page, pageSize int) ([]*response.OrderItemResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching all order items",
		zap.String("search", search),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orderItems, total, err := s.orderItemRepository.FindAllOrderItems(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve order items list",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Unable to retrieve order items at this time",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderItemsResponse(orderItems), &total, nil
}

func (s *orderItemService) FindByActive(search string, page, pageSize int) ([]*response.OrderItemResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active order items",
		zap.String("search", search),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orderItems, total, err := s.orderItemRepository.FindByActive(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve active order items",
			zap.Error(err),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Unable to retrieve active order items",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderItemsResponseDeleteAt(orderItems), &total, nil
}

func (s *orderItemService) FindByTrashed(search string, page, pageSize int) ([]*response.OrderItemResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching archived order items",
		zap.String("search", search),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orderItems, total, err := s.orderItemRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve archived order items",
			zap.Error(err),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Unable to retrieve archived order items",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderItemsResponseDeleteAt(orderItems), &total, nil
}

func (s *orderItemService) FindOrderItemByOrder(orderID int) ([]*response.OrderItemResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching order items for order",
		zap.Int("order_id", orderID))

	if orderID <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Invalid order ID",
			Code:    http.StatusBadRequest,
		}
	}

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(orderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "error",
				Message: "No items found for the specified order",
				Code:    http.StatusNotFound,
			}
		}
		s.logger.Error("Failed to retrieve order items",
			zap.Error(err),
			zap.Int("order_id", orderID))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Unable to retrieve order items",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToOrderItemsResponse(orderItems), nil
}
