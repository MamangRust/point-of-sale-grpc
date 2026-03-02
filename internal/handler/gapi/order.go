package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/order_errors"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type orderHandleGrpc struct {
	pb.UnimplementedOrderServiceServer
	orderService service.OrderService
}

func NewOrderHandleGrpc(
	orderService service.OrderService,
) *orderHandleGrpc {
	return &orderHandleGrpc{
		orderService: orderService,
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

	reqService := requests.FindAllOrders{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	orders, totalRecords, err := s.orderService.FindAllOrders(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var orderResponses []*pb.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, &pb.OrderResponse{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationOrder{
		Status:     "success",
		Message:    "Successfully fetched order",
		Data:       orderResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *orderHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindAllOrderMerchantRequest) (*pb.ApiResponsePaginationOrder, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	merchant_id := request.GetMerchantId()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllOrderMerchant{
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
		MerchantID: int(merchant_id),
	}

	orders, totalRecords, err := s.orderService.FindByMerchant(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var orderResponses []*pb.OrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, &pb.OrderResponse{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationOrder{
		Status:     "success",
		Message:    "Successfully fetched order",
		Data:       orderResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *orderHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrder, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	order, err := s.orderService.FindById(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrder{
		Status:  "success",
		Message: "Successfully fetched order",
		Data: &pb.OrderResponse{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *orderHandleGrpc) FindMonthlyTotalRevenue(ctx context.Context, req *pb.FindYearMonthTotalRevenue) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, order_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthTotalRevenue{
		Year:  year,
		Month: month,
	}

	methods, err := s.orderService.FindMonthlyTotalRevenue(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyRevenueResponses []*pb.OrderMonthlyTotalRevenueResponse
	for _, method := range methods {
		monthlyRevenueResponses = append(monthlyRevenueResponses, &pb.OrderMonthlyTotalRevenueResponse{
			Year:         method.Year,
			Month:        method.Month,
			TotalRevenue: int32(method.TotalRevenue),
		})
	}

	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    monthlyRevenueResponses,
	}, nil
}

func (s *orderHandleGrpc) FindYearlyTotalRevenue(ctx context.Context, req *pb.FindYearTotalRevenue) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	methods, err := s.orderService.FindYearlyTotalRevenue(ctx, year)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyRevenueResponses []*pb.OrderYearlyTotalRevenueResponse
	for _, method := range methods {
		yearlyRevenueResponses = append(yearlyRevenueResponses, &pb.OrderYearlyTotalRevenueResponse{
			Year:         method.Year,
			TotalRevenue: int32(method.TotalRevenue),
		})
	}

	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    yearlyRevenueResponses,
	}, nil
}

func (s *orderHandleGrpc) FindMonthlyTotalRevenueById(ctx context.Context, req *pb.FindYearMonthTotalRevenueById) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetOrderId())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, order_errors.ErrGrpcInvalidMonth
	}

	if id <= 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthTotalRevenueOrder{
		OrderID: id,
		Month:   month,
		Year:    year,
	}

	methods, err := s.orderService.FindMonthlyTotalRevenueById(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyRevenueResponses []*pb.OrderMonthlyTotalRevenueResponse
	for _, method := range methods {
		monthlyRevenueResponses = append(monthlyRevenueResponses, &pb.OrderMonthlyTotalRevenueResponse{
			Year:         method.Year,
			Month:        method.Month,
			TotalRevenue: int32(method.TotalRevenue),
		})
	}

	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    monthlyRevenueResponses,
	}, nil
}

func (s *orderHandleGrpc) FindYearlyTotalRevenueById(ctx context.Context, req *pb.FindYearTotalRevenueById) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	year := int(req.GetYear())
	id := int(req.GetOrderId())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalRevenueOrder{
		OrderID: id,
		Year:    year,
	}

	methods, err := s.orderService.FindYearlyTotalRevenueById(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyRevenueResponses []*pb.OrderYearlyTotalRevenueResponse
	for _, method := range methods {
		yearlyRevenueResponses = append(yearlyRevenueResponses, &pb.OrderYearlyTotalRevenueResponse{
			Year:         method.Year,
			TotalRevenue: method.TotalRevenue,
		})
	}

	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    yearlyRevenueResponses,
	}, nil
}

func (s *orderHandleGrpc) FindMonthlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearMonthTotalRevenueByMerchant) (*pb.ApiResponseOrderMonthlyTotalRevenue, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, order_errors.ErrGrpcInvalidMonth
	}

	if id <= 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthTotalRevenueMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	methods, err := s.orderService.FindMonthlyTotalRevenueByMerchant(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyRevenueResponses []*pb.OrderMonthlyTotalRevenueResponse
	for _, method := range methods {
		monthlyRevenueResponses = append(monthlyRevenueResponses, &pb.OrderMonthlyTotalRevenueResponse{
			Year:         method.Year,
			Month:        method.Month,
			TotalRevenue: int32(method.TotalRevenue),
		})
	}

	return &pb.ApiResponseOrderMonthlyTotalRevenue{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    monthlyRevenueResponses,
	}, nil
}

func (s *orderHandleGrpc) FindYearlyTotalRevenueByMerchant(ctx context.Context, req *pb.FindYearTotalRevenueByMerchant) (*pb.ApiResponseOrderYearlyTotalRevenue, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalRevenueMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.orderService.FindYearlyTotalRevenueByMerchant(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyRevenueResponses []*pb.OrderYearlyTotalRevenueResponse
	for _, method := range methods {
		yearlyRevenueResponses = append(yearlyRevenueResponses, &pb.OrderYearlyTotalRevenueResponse{
			Year:         method.Year,
			TotalRevenue: method.TotalRevenue,
		})
	}

	return &pb.ApiResponseOrderYearlyTotalRevenue{
		Status:  "success",
		Message: "Yearly payment methods retrieved successfully",
		Data:    yearlyRevenueResponses,
	}, nil
}

func (s *orderHandleGrpc) FindMonthlyRevenue(ctx context.Context, request *pb.FindYearOrder) (*pb.ApiResponseOrderMonthly, error) {
	year := int(request.GetYear())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	res, err := s.orderService.FindMonthlyOrder(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyResponses []*pb.OrderMonthlyResponse
	for _, month := range res {
		monthlyResponses = append(monthlyResponses, &pb.OrderMonthlyResponse{
			Month:          month.Month,
			OrderCount:     int32(month.OrderCount),
			TotalRevenue:   int32(month.TotalRevenue),
			TotalItemsSold: int32(month.TotalItemsSold),
		})
	}

	return &pb.ApiResponseOrderMonthly{
		Status:  "success",
		Message: "Monthly revenue data retrieved",
		Data:    monthlyResponses,
	}, nil
}

func (s *orderHandleGrpc) FindYearlyRevenue(ctx context.Context, request *pb.FindYearOrder) (*pb.ApiResponseOrderYearly, error) {
	year := int(request.GetYear())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	res, err := s.orderService.FindYearlyOrder(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyResponses []*pb.OrderYearlyResponse
	for _, yearData := range res {
		yearlyResponses = append(yearlyResponses, &pb.OrderYearlyResponse{
			Year:               yearData.Year,
			OrderCount:         int32(yearData.OrderCount),
			TotalRevenue:       int32(yearData.TotalRevenue),
			TotalItemsSold:     int32(yearData.TotalItemsSold),
			ActiveCashiers:     int32(yearData.ActiveCashiers),
			UniqueProductsSold: int32(yearData.UniqueProductsSold),
		})
	}

	return &pb.ApiResponseOrderYearly{
		Status:  "success",
		Message: "Yearly revenue data retrieved",
		Data:    yearlyResponses,
	}, nil
}

func (s *orderHandleGrpc) FindMonthlyRevenueByMerchant(ctx context.Context, request *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderMonthly, error) {
	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, order_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthOrderMerchant{
		Year:       year,
		MerchantID: id,
	}

	res, err := s.orderService.FindMonthlyOrderByMerchant(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyResponses []*pb.OrderMonthlyResponse
	for _, month := range res {
		monthlyResponses = append(monthlyResponses, &pb.OrderMonthlyResponse{
			Month:          month.Month,
			OrderCount:     int32(month.OrderCount),
			TotalRevenue:   int32(month.TotalRevenue),
			TotalItemsSold: int32(month.TotalItemsSold),
		})
	}

	return &pb.ApiResponseOrderMonthly{
		Status:  "success",
		Message: "Monthly revenue by merchant data retrieved",
		Data:    monthlyResponses,
	}, nil
}

func (s *orderHandleGrpc) FindYearlyRevenueByMerchant(ctx context.Context, request *pb.FindYearOrderByMerchant) (*pb.ApiResponseOrderYearly, error) {
	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, order_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearOrderMerchant{
		Year:       year,
		MerchantID: id,
	}

	res, err := s.orderService.FindYearlyOrderByMerchant(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyResponses []*pb.OrderYearlyResponse
	for _, yearData := range res {
		yearlyResponses = append(yearlyResponses, &pb.OrderYearlyResponse{
			Year:               yearData.Year,
			OrderCount:         int32(yearData.OrderCount),
			TotalRevenue:       int32(yearData.TotalRevenue),
			TotalItemsSold:     int32(yearData.TotalItemsSold),
			ActiveCashiers:     int32(yearData.ActiveCashiers),
			UniqueProductsSold: int32(yearData.UniqueProductsSold),
		})
	}

	return &pb.ApiResponseOrderYearly{
		Status:  "success",
		Message: "Yearly revenue by merchant data retrieved",
		Data:    yearlyResponses,
	}, nil
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

	reqService := requests.FindAllOrders{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	orders, totalRecords, err := s.orderService.FindByActive(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var orderResponses []*pb.OrderResponseDeleteAt
	for _, order := range orders {
		var deletedAt string
		if order.DeletedAt.Valid {
			deletedAt = order.DeletedAt.Time.Format("2006-01-02")
		}

		orderResponses = append(orderResponses, &pb.OrderResponseDeleteAt{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationOrderDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active order",
		Data:       orderResponses,
		Pagination: paginationMeta,
	}, nil
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

	reqService := requests.FindAllOrders{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	orders, totalRecords, err := s.orderService.FindByTrashed(ctx, &reqService)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var orderResponses []*pb.OrderResponseDeleteAt
	for _, order := range orders {
		var deletedAt string
		if order.DeletedAt.Valid {
			deletedAt = order.DeletedAt.Time.Format("2006-01-02")
		}
		orderResponses = append(orderResponses, &pb.OrderResponseDeleteAt{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationOrderDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed order",
		Data:       orderResponses,
		Pagination: paginationMeta,
	}, nil
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
		return nil, order_errors.ErrGrpcValidateCreateOrder
	}

	order, err := s.orderService.CreateOrder(ctx, req)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrder{
		Status:  "success",
		Message: "Successfully created order",
		Data: &pb.OrderResponse{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *orderHandleGrpc) Update(ctx context.Context, request *pb.UpdateOrderRequest) (*pb.ApiResponseOrder, error) {
	id := int(request.GetOrderId())

	if id == 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
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
		return nil, order_errors.ErrGrpcValidateUpdateOrder
	}

	order, err := s.orderService.UpdateOrder(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrder{
		Status:  "success",
		Message: "Successfully updated order",
		Data: &pb.OrderResponse{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *orderHandleGrpc) TrashedOrder(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrderDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	order, err := s.orderService.TrashedOrder(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrderDeleteAt{
		Status:  "success",
		Message: "Successfully trashed order",
		Data: &pb.OrderResponseDeleteAt{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: order.DeletedAt.Time.String()},
		},
	}, nil
}

func (s *orderHandleGrpc) RestoreOrder(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrderDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	order, err := s.orderService.RestoreOrder(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrderDeleteAt{
		Status:  "success",
		Message: "Successfully restored order",
		Data: &pb.OrderResponseDeleteAt{
			Id:         int32(order.OrderID),
			MerchantId: int32(order.MerchantID),
			CashierId:  int32(order.CashierID),
			TotalPrice: int32(order.TotalPrice),
			CreatedAt:  order.CreatedAt.Time.String(),
			UpdatedAt:  order.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: order.DeletedAt.Time.String()},
		},
	}, nil
}

func (s *orderHandleGrpc) DeleteOrderPermanent(ctx context.Context, request *pb.FindByIdOrderRequest) (*pb.ApiResponseOrderDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, order_errors.ErrGrpcFailedInvalidId
	}

	_, err := s.orderService.DeleteOrderPermanent(ctx, id)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrderDelete{
		Status:  "success",
		Message: "Successfully deleted order permanently",
	}, nil
}

func (s *orderHandleGrpc) RestoreAllOrder(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseOrderAll, error) {
	_, err := s.orderService.RestoreAllOrder(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrderAll{
		Status:  "success",
		Message: "Successfully restore all order",
	}, nil
}

func (s *orderHandleGrpc) DeleteAllOrderPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseOrderAll, error) {
	_, err := s.orderService.DeleteAllOrderPermanent(ctx)

	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseOrderAll{
		Status:  "success",
		Message: "Successfully delete order permanen",
	}, nil
}
