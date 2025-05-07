package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	orderitem_errors "pointofsale/pkg/errors/order_item_errors"
)

type orderItemHandleGrpc struct {
	pb.UnimplementedOrderItemServiceServer
	orderItemService service.OrderItemService
	mapping          protomapper.OrderItemProtoMapper
}

func NewOrderItemHandleGrpc(
	orderItemService service.OrderItemService,
	mapping protomapper.OrderItemProtoMapper,
) *orderItemHandleGrpc {
	return &orderItemHandleGrpc{
		orderItemService: orderItemService,
		mapping:          mapping,
	}
}

func (s *orderItemHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllOrderItemRequest) (*pb.ApiResponsePaginationOrderItem, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllOrderItems{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	orderItems, totalRecords, err := s.orderItemService.FindAllOrderItems(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationOrderItem(paginationMeta, "success", "Successfully fetched order items", orderItems)
	return so, nil
}

func (s *orderItemHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllOrderItemRequest) (*pb.ApiResponsePaginationOrderItemDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllOrderItems{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	orderItems, totalRecords, err := s.orderItemService.FindByActive(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationOrderItemDeleteAt(paginationMeta, "success", "Successfully fetched active order items", orderItems)
	return so, nil
}

func (s *orderItemHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllOrderItemRequest) (*pb.ApiResponsePaginationOrderItemDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllOrderItems{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	orderItems, totalRecords, err := s.orderItemService.FindByTrashed(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationOrderItemDeleteAt(paginationMeta, "success", "Successfully fetched trashed order items", orderItems)
	return so, nil
}

func (s *orderItemHandleGrpc) FindOrderItemByOrder(ctx context.Context, request *pb.FindByIdOrderItemRequest) (*pb.ApiResponsesOrderItem, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, orderitem_errors.ErrGrpcInvalidID
	}

	orderItems, err := s.orderItemService.FindOrderItemByOrder(id)
	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponsesOrderItem("success", "Successfully fetched order items by order", orderItems)
	return so, nil
}
