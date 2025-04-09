package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors_custom"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cashierHandleGrpc struct {
	pb.UnimplementedCashierServiceServer
	cashierService service.CashierService
	mapping        protomapper.CashierProtoMapper
}

func NewCashierHandleGrpc(
	cashierService service.CashierService,
	mapping protomapper.CashierProtoMapper,
) *cashierHandleGrpc {
	return &cashierHandleGrpc{
		cashierService: cashierService,
		mapping:        mapping,
	}
}

func (s *cashierHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllCashierRequest) (*pb.ApiResponsePaginationCashier, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	cashier, totalRecords, err := s.cashierService.FindAll(page, pageSize, search)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCashier(paginationMeta, "success", "Successfully fetched cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashier, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_request",
				Message: "Valid cashier ID is required",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	cashier, err := s.cashierService.FindById(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashier("success", "Successfully fetched categories", cashier)

	return so, nil
}

func (s *cashierHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllCashierRequest) (*pb.ApiResponsePaginationCashierDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	cashier, totalRecords, err := s.cashierService.FindByActive(search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(0),
	}

	so := s.mapping.ToProtoResponsePaginationCashierDeleteAt(paginationMeta, "success", "Successfully fetched active cashier", cashier)

	return so, nil
}

func (s *cashierHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllCashierRequest) (*pb.ApiResponsePaginationCashierDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.cashierService.FindByTrashed(search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(0),
	}

	so := s.mapping.ToProtoResponsePaginationCashierDeleteAt(paginationMeta, "success", "Successfully fetched trashed cashier", users)

	return so, nil
}

func (s *cashierHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindByMerchantCashierRequest) (*pb.ApiResponsePaginationCashier, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	merchant_id := int(request.GetMerchantId())

	if merchant_id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid merchant id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	cashier, totalRecords, err := s.cashierService.FindByMerchant(merchant_id, search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCashier(paginationMeta, "success", "Successfully fetched cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) FindMonthlyTotalSales(ctx context.Context, req *pb.FindYearMonthTotalSales) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindMonthlyTotalSales(year, month)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoMonthlyTotalSales("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindYearlyTotalSales(ctx context.Context, req *pb.FindYearTotalSales) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindYearlyTotalSales(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoYearlyTotalSales("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindMonthlyTotalSalesById(ctx context.Context, req *pb.FindYearMonthTotalSalesById) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetCashierId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindMonthlyTotalSalesById(year, month, id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoMonthlyTotalSales("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindYearlyTotalSalesById(ctx context.Context, req *pb.FindYearTotalSalesById) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	year := int(req.GetYear())
	id := int(req.GetCashierId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindYearlyTotalSalesById(year, id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoYearlyTotalSales("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindMonthlyTotalSalesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalSalesByMerchant) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if merchantId <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindMonthlyTotalSalesById(year, month, merchantId)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoMonthlyTotalSales("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindYearlyTotalSalesByMerchant(ctx context.Context, req *pb.FindYearTotalSalesByMerchant) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if merchantId <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindYearlyTotalSalesByMerchant(year, merchantId)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoYearlyTotalSales("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindMonthSales(ctx context.Context, req *pb.FindYearCashier) (*pb.ApiResponseCashierMonthSales, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindMonthlySales(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthlyTotalSales("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindYearSales(ctx context.Context, req *pb.FindYearCashier) (*pb.ApiResponseCashierYearSales, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindYearlySales(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearlyTotalSales("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindMonthSalesByMerchant(ctx context.Context, req *pb.FindYearCashierByMerchant) (*pb.ApiResponseCashierMonthSales, error) {
	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if merchantId <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindMonthlyCashierByMerchant(
		year, merchantId,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthlyTotalSales("success", "Merchant monthly revenue retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindYearSalesByMerchant(ctx context.Context, req *pb.FindYearCashierByMerchant) (*pb.ApiResponseCashierYearSales, error) {
	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if merchantId <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindYearlyCashierByMerchant(
		year, merchantId,
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearlyTotalSales("success", "Merchant yearly payment methods retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindMonthSalesById(ctx context.Context, req *pb.FindYearCashierById) (*pb.ApiResponseCashierMonthSales, error) {
	year := int(req.GetYear())
	cashierId := int(req.GetCashierId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if cashierId <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindMonthlyCashierById(
		year, cashierId,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthlyTotalSales("success", "Cashier monthly sales retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) FindYearSalesById(ctx context.Context, req *pb.FindYearCashierById) (*pb.ApiResponseCashierYearSales, error) {
	year := int(req.GetYear())
	cashierId := int(req.GetCashierId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if cashierId <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.cashierService.FindYearlyCashierById(
		year,
		cashierId,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearlyTotalSales("success", "Cashier yearly sales retrieved successfully", methods), nil
}

func (s *cashierHandleGrpc) CreateCashier(ctx context.Context, request *pb.CreateCashierRequest) (*pb.ApiResponseCashier, error) {
	req := &requests.CreateCashierRequest{
		Name:       request.GetName(),
		MerchantID: int(request.GetMerchantId()),
		UserID:     int(request.GetUserId()),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Unable to create new cashier. Please check your input.",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	cashier, err := s.cashierService.CreateCashier(req)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashier("success", "Successfully created cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) UpdateCashier(ctx context.Context, request *pb.UpdateCashierRequest) (*pb.ApiResponseCashier, error) {
	id := int(request.GetCashierId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Cashier ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	req := &requests.UpdateCashierRequest{
		CashierID: &id,
		Name:      request.GetName(),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Unable to process cashier update. Please review your data.",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	cashier, err := s.cashierService.UpdateCashier(req)
	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashier("success", "Successfully updated cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) TrashedCashier(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Cashier ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	cashier, err := s.cashierService.TrashedCashier(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashierDeleteAt("success", "Successfully trashed cashier", cashier)

	return so, nil
}

func (s *cashierHandleGrpc) RestoreCashier(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Cashier ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	cashier, err := s.cashierService.RestoreCashier(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashierDeleteAt("success", "Successfully restored cashier", cashier)

	return so, nil
}

func (s *cashierHandleGrpc) DeleteCashierPermanent(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Cashier ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	_, err := s.cashierService.DeleteCashierPermanent(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashierDelete("success", "Successfully deleted cashier permanently")

	return so, nil
}

func (s *cashierHandleGrpc) RestoreAllCashier(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCashierAll, error) {
	_, err := s.cashierService.RestoreAllCashier()

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashierAll("success", "Successfully restore all cashier")

	return so, nil
}

func (s *cashierHandleGrpc) DeleteAllCashierPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCashierAll, error) {
	_, err := s.cashierService.DeleteAllCashierPermanent()

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseCashierAll("success", "Successfully delete cashier permanen")

	return so, nil
}
