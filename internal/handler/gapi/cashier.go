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
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch cashier: ",
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCashier(paginationMeta, "success", "Successfully fetched cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashier, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier id",
		})
	}

	cashier, err := s.cashierService.FindById(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch cashier: " + err.Message,
		})
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
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch active categories: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

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
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch trashed cashier: " + err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

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
	merchant_id := request.GetMerchantId()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	cashier, totalRecords, err := s.cashierService.FindByMerchant(int(merchant_id), search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch cashier: ",
		})
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationCashier(paginationMeta, "success", "Successfully fetched cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) Create(ctx context.Context, request *pb.CreateCashierRequest) (*pb.ApiResponseCashier, error) {
	req := &requests.CreateCashierRequest{
		Name:       request.GetName(),
		MerchantID: int(request.GetMerchantId()),
		UserID:     int(request.GetUserId()),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create cashier: " + err.Error(),
		})
	}

	cashier, err := s.cashierService.CreateCashier(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create cashier: ",
		})
	}

	so := s.mapping.ToProtoResponseCashier("success", "Successfully created cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) Update(ctx context.Context, request *pb.UpdateCashierRequest) (*pb.ApiResponseCashier, error) {
	if request.GetCashierId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier id",
		})
	}

	req := &requests.UpdateCashierRequest{
		CashierID: int(request.GetCashierId()),
		Name:      request.GetName(),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update cashier: " + err.Error(),
		})
	}

	cashier, err := s.cashierService.UpdateCashier(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update cashier: ",
		})
	}

	so := s.mapping.ToProtoResponseCashier("success", "Successfully updated cashier", cashier)
	return so, nil
}

func (s *cashierHandleGrpc) TrashedCashier(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier id",
		})
	}

	cashier, err := s.cashierService.TrashedCashier(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed cashier: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCashierDeleteAt("success", "Successfully trashed cashier", cashier)

	return so, nil
}

func (s *cashierHandleGrpc) RestoreCashier(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier id",
		})
	}

	cashier, err := s.cashierService.RestoreCashier(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore cashier: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCashierDeleteAt("success", "Successfully restored cashier", cashier)

	return so, nil
}

func (s *cashierHandleGrpc) DeleteCashierPermanent(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDelete, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid cashier id",
		})
	}

	_, err := s.cashierService.DeleteCashierPermanent(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete cashier permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCashierDelete("success", "Successfully deleted cashier permanently")

	return so, nil
}

func (s *cashierHandleGrpc) RestoreAllCashier(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCashierAll, error) {
	_, err := s.cashierService.RestoreAllCashier()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all cashier: ",
		})
	}

	so := s.mapping.ToProtoResponseCashierAll("success", "Successfully restore all cashier")

	return so, nil
}

func (s *cashierHandleGrpc) DeleteAllCashierPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCashierAll, error) {
	_, err := s.cashierService.DeleteAllCashierPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete cashier permanent: ",
		})
	}

	so := s.mapping.ToProtoResponseCashierAll("success", "Successfully delete cashier permanen")

	return so, nil
}
