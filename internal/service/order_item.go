package service

import (
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	orderitem_errors "pointofsale/pkg/errors/order_item_errors"
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

func (s *orderItemService) FindAllOrderItems(req *requests.FindAllOrderItems) ([]*response.OrderItemResponse, *int, *response.ErrorResponse) {
	page := req.Page
	pageSize := req.PageSize
	search := req.Search

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

	orderItems, totalRecords, err := s.orderItemRepository.FindAllOrderItems(req)
	if err != nil {
		s.logger.Error("Failed to fetch all order items",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, orderitem_errors.ErrFailedFindAllOrderItems
	}
	s.logger.Debug("Successfully fetched order-item",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToOrderItemsResponse(orderItems), totalRecords, nil
}

func (s *orderItemService) FindByActive(req *requests.FindAllOrderItems) ([]*response.OrderItemResponseDeleteAt, *int, *response.ErrorResponse) {
	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	s.logger.Debug("Fetching all order items active",
		zap.String("search", search),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orderItems, totalRecords, err := s.orderItemRepository.FindByActive(req)
	if err != nil {
		s.logger.Error("Failed to retrieve order-item active",
			zap.Error(err),
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize))

		return nil, nil, orderitem_errors.ErrFailedFindOrderItemsByActive
	}

	s.logger.Debug("Successfully fetched order-items",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToOrderItemsResponseDeleteAt(orderItems), totalRecords, nil
}

func (s *orderItemService) FindByTrashed(req *requests.FindAllOrderItems) ([]*response.OrderItemResponseDeleteAt, *int, *response.ErrorResponse) {
	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	s.logger.Debug("Fetching all order items trashed",
		zap.String("search", search),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	orderItems, totalRecords, err := s.orderItemRepository.FindByTrashed(req)

	if err != nil {
		s.logger.Error("Failed to retrieve trashed order-items",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, orderitem_errors.ErrFailedFindOrderItemsByTrashed
	}

	s.logger.Debug("Successfully fetched order-items trashed",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToOrderItemsResponseDeleteAt(orderItems), totalRecords, nil
}

func (s *orderItemService) FindOrderItemByOrder(orderID int) ([]*response.OrderItemResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching order items for order",
		zap.Int("order_id", orderID))

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(orderID)
	if err != nil {
		s.logger.Error("Failed to retrieve order items",
			zap.Error(err),
			zap.Int("order_id", orderID))
		return nil, orderitem_errors.ErrFailedFindOrderItemByOrder
	}

	return s.mapping.ToOrderItemsResponse(orderItems), nil
}
