package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderHandleGrpc struct {
	pb.UnimplementedOrderServiceServer
	orderService service.OrderService
	mapping      protomapper.OrderProtoMapper
}

func NewOrderHandleGrpc(
	orderService service.OrderService,
	mapping protomapper.OrderProtoMapper,
) *orderHandleGrpc {
	return &orderHandleGrpc{
		orderService: orderService,
		mapping:      mapping,
	}
}

func (s *orderHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllOrderRequest) (*pb.ApiResponsePaginationOrder, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchant, totalRecords, err := s.orderService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationOrder(paginationMeta, "success", "Successfully fetched order", merchant)
	return so, nil
}

func (s *orderHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrder, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order ID parameter cannot be empty and must be a positive number",
		})
	}

	merchant, err := s.orderService.FindById(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseOrder("success", "Successfully fetched order", merchant)

	return so, nil
}

func (s *orderHandleGrpc) FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	if req.GetMonth() <= 0 || req.GetMonth() >= 12 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_month",
			Message: "Month must be between 1 and 12",
		})
	}

	methods, err := s.orderService.FindMonthlyTotalRevenue(int(req.GetYear()), int(req.GetMonth()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseMonthlyTotalRevenue("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *orderHandleGrpc) FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.orderService.FindYearlyTotalRevenue(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseYearlyTotalRevenue("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *orderHandleGrpc) FindMonthlyTotalRevenueById(ctx context.Context, req *pb.FindYearMonthTotalRevenueById) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	if req.GetMonth() <= 0 || req.GetMonth() >= 12 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_month",
			Message: "Month must be between 1 and 12",
		})
	}

	methods, err := s.orderService.FindMonthlyTotalRevenueById(int(req.GetYear()), int(req.GetMonth()), int(req.GetOrderId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseMonthlyTotalRevenue("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *orderHandleGrpc) FindYearlyTotalRevenueById(ctx context.Context, req *pb.FindYearTotalRevenueById) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.orderService.FindYearlyTotalRevenueById(int(req.GetYear()), int(req.GetOrderId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseYearlyTotalRevenue("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *orderHandleGrpc) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	if req.GetMonth() <= 0 || req.GetMonth() >= 12 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_month",
			Message: "Month must be between 1 and 12",
		})
	}

	methods, err := s.orderService.FindMonthlyTotalRevenueByMerchant(int(req.GetYear()), int(req.GetMonth()), int(req.GetMerchantId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseMonthlyTotalRevenue("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *orderHandleGrpc) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.orderService.FindYearlyTotalRevenueByMerchant(int(req.GetYear()), int(req.GetMerchantId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseYearlyTotalRevenue("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *orderHandleGrpc) FindMonthlyRevenue(ctx context.Context, request *pb.FindYearOrder) (*pb.ApiResponseOrderMonthly, error) {
	res, err := s.orderService.FindMonthlyOrder(int(request.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseMonthlyRevenue("success", "Monthly revenue data retrieved", res)
	return so, nil
}

func (s *orderHandleGrpc) FindYearlyRevenue(ctx context.Context, request *pb.FindYearOrder) (*pb.ApiResponseOrderYearly, error) {
	res, err := s.orderService.FindYearlyOrder(int(request.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseYearlyRevenue("success", "Yearly revenue data retrieved", res)
	return so, nil
}

func (s *orderHandleGrpc) FindMonthlyRevenueByMerchant(ctx context.Context, request *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderMonthly, error) {
	res, err := s.orderService.FindMonthlyOrderByMerchant(int(request.GetMerchantId()), int(request.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseMonthlyRevenue("success", "Monthly revenue by merchant data retrieved", res)
	return so, nil
}

func (s *orderHandleGrpc) FindYearlyRevenueByMerchant(ctx context.Context, request *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderYearly, error) {
	res, err := s.orderService.FindYearlyOrderByMerchant(int(request.GetMerchantId()), int(request.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseYearlyRevenue("success", "Yearly revenue by merchant data retrieved", res)
	return so, nil
}

func (s *orderHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllOrderRequest) (*pb.ApiResponsePaginationOrderDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchant, totalRecords, err := s.orderService.FindByActive(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(0),
	}
	so := s.mapping.ToProtoResponsePaginationOrderDeleteAt(paginationMeta, "success", "Successfully fetched active order", merchant)

	return so, nil
}

func (s *orderHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllOrderRequest) (*pb.ApiResponsePaginationOrderDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.orderService.FindByTrashed(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(0),
	}

	so := s.mapping.ToProtoResponsePaginationOrderDeleteAt(paginationMeta, "success", "Successfully fetched trashed order", users)

	return so, nil
}

func (s *orderHandleGrpc) Create(ctx context.Context, request *pb.CreateOrderRequest) (*pb.ApiResponseOrder, error) {
	req := &requests.CreateOrderRequest{
		MerchantID: int(request.GetMerchantId()),
		CashierID:  int(request.GetCashierId()),
	}

	for _, item := range request.GetItems() {
		req.Items = append(req.Items, requests.CreateOrderItemRequest{
			ProductID: int(item.GetProductId()),
			Quantity:  int(item.GetQuantity()),
		})
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order creation failed. Please check your order details.",
		})
	}

	order, err := s.orderService.CreateOrder(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create order: ",
		})
	}

	so := s.mapping.ToProtoResponseOrder("success", "Successfully created order", order)
	return so, nil
}

func (s *orderHandleGrpc) Update(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.ApiResponseOrder, error) {
	id := int(request.GetOrderId())

	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order ID parameter cannot be empty and must be a positive number",
		})
	}

	req := &requests.UpdateOrderRequest{
		OrderID: &id,
	}

	for _, item := range request.GetItems() {
		req.Items = append(req.Items, requests.UpdateOrderItemRequest{
			OrderItemID: int(item.GetOrderItemId()),
			ProductID:   int(item.GetProductId()),
			Quantity:    int(item.GetQuantity()),
		})
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order update failed. Please verify your changes.",
		})
	}

	order, err := s.orderService.UpdateOrder(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update order: ",
		})
	}

	so := s.mapping.ToProtoResponseOrder("success", "Successfully updated order", order)
	return so, nil
}

func (s *orderHandleGrpc) TrashedOrder(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrderDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order ID parameter cannot be empty and must be a positive number",
		})
	}

	merchant, err := s.orderService.TrashedOrder(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseOrderDeleteAt("success", "Successfully trashed order", merchant)

	return so, nil
}

func (s *orderHandleGrpc) RestoreOrder(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrderDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order ID parameter cannot be empty and must be a positive number",
		})
	}

	merchant, err := s.orderService.RestoreOrder(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseOrderDeleteAt("success", "Successfully restored order", merchant)

	return so, nil
}

func (s *orderHandleGrpc) DeleteOrderPermanent(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrderDelete, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Order ID parameter cannot be empty and must be a positive number",
		})
	}

	_, err := s.orderService.DeleteOrderPermanent(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseOrderDelete("success", "Successfully deleted order permanently")

	return so, nil
}

func (s *orderHandleGrpc) RestoreAllOrder(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseOrderAll, error) {
	_, err := s.orderService.RestoreAllOrder()

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseOrderAll("success", "Successfully restore all order")

	return so, nil
}

func (s *orderHandleGrpc) DeleteAllOrderPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseOrderAll, error) {
	_, err := s.orderService.DeleteAllOrderPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseOrderAll("success", "Successfully delete order permanen")

	return so, nil
}
