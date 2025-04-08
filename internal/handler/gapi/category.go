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

type categoryHandleGrpc struct {
	pb.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
	mapping         protomapper.CategoryProtoMapper
}

func NewCategoryHandleGrpc(
	categoryService service.CategoryService,
	mapping protomapper.CategoryProtoMapper,
) *categoryHandleGrpc {
	return &categoryHandleGrpc{
		categoryService: categoryService,
		mapping:         mapping,
	}
}

func (s *categoryHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategory, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	category, totalRecords, err := s.categoryService.FindAll(page, pageSize, search)

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

	so := s.mapping.ToProtoResponsePaginationCategory(paginationMeta, "success", "Successfully fetched categories", category)
	return so, nil
}

func (s *categoryHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategory, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Category ID parameter cannot be empty and must be a positive number",
		})
	}

	category, err := s.categoryService.FindById(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully fetched categories", category)

	return so, nil

}

func (s *categoryHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategoryDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.categoryService.FindByActive(search, page, pageSize)

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
	so := s.mapping.ToProtoResponsePaginationCategoryDeleteAt(paginationMeta, "success", "Successfully fetched active categories", users)

	return so, nil
}

func (s *categoryHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllCategoryRequest) (*pb.ApiResponsePaginationCategoryDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.categoryService.FindByTrashed(search, page, pageSize)

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

	so := s.mapping.ToProtoResponsePaginationCategoryDeleteAt(paginationMeta, "success", "Successfully fetched trashed categories", users)

	return so, nil
}

func (s *categoryHandleGrpc) FindMonthlyTotalPrices(ctx context.Context, req *pb.FindYearMonthTotalPrices) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
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

	methods, err := s.categoryService.FindMonthlyTotalPrice(int(req.GetYear()), int(req.GetMonth()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseMonthlyTotalPrice("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.categoryService.FindYearlyTotalPrice(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseYearlyTotalPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
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

	methods, err := s.categoryService.FindMonthlyTotalPriceById(int(req.GetYear()), int(req.GetMonth()), int(req.GetCategoryId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseMonthlyTotalPrice("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.categoryService.FindYearlyTotalPriceById(int(req.GetYear()), int(req.GetCategoryId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseYearlyTotalPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
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

	methods, err := s.categoryService.FindMonthlyTotalPriceByMerchant(int(req.GetYear()), int(req.GetMonth()), int(req.GetMerchantId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseMonthlyTotalPrice("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.categoryService.FindYearlyTotalPriceByMerchant(int(req.GetYear()), int(req.GetMerchantId()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	return s.mapping.ToProtoResponseYearlyTotalPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryMonthPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.categoryService.FindMonthPrice(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve monthly revenue",
		})
	}

	return s.mapping.ToProtoResponseCategoryMonthlyPrice("success", "Monthly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryYearPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}

	methods, err := s.categoryService.FindYearPrice(int(req.GetYear()))
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve yearly revenue",
		})
	}

	return s.mapping.ToProtoResponseCategoryYearlyPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryMonthPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}
	if req.GetMerchantId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid merchant ID",
		})
	}

	methods, err := s.categoryService.FindMonthPriceByMerchant(
		int(req.GetYear()),
		int(req.GetMerchantId()),
	)
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve category monthly revenue",
		})
	}

	return s.mapping.ToProtoResponseCategoryMonthlyPrice("success", "Merchant monthly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryYearPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}
	if req.GetMerchantId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid merchant ID",
		})
	}

	methods, err := s.categoryService.FindYearPriceByMerchant(
		int(req.GetYear()),
		int(req.GetMerchantId()),
	)
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve category yearly revenue",
		})
	}

	return s.mapping.ToProtoResponseCategoryYearlyPrice("success", "Merchant yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryMonthPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}
	if req.GetCategoryId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid category ID",
		})
	}

	methods, err := s.categoryService.FindMonthPriceByMerchant(
		int(req.GetYear()),
		int(req.GetCategoryId()),
	)
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve category monthly revenue",
		})
	}

	return s.mapping.ToProtoResponseCategoryMonthlyPrice("success", "Merchant monthly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryYearPrice, error) {
	if req.GetYear() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid year parameter",
		})
	}
	if req.GetCategoryId() <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid category ID",
		})
	}

	methods, err := s.categoryService.FindYearPriceByMerchant(
		int(req.GetYear()),
		int(req.GetCategoryId()),
	)
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve category yearly revenue",
		})
	}

	return s.mapping.ToProtoResponseCategoryYearlyPrice("success", "Merchant yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) Create(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.ApiResponseCategory, error) {
	req := &requests.CreateCategoryRequest{
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Category data is invalid. Please check your input.",
		})
	}

	category, err := s.categoryService.CreateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to create category: ",
		})
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully created category", category)
	return so, nil
}

func (s *categoryHandleGrpc) Update(ctx context.Context, request *pb.UpdateCategoryRequest) (*pb.ApiResponseCategory, error) {
	id := int(request.GetCategoryId())

	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Category ID parameter cannot be empty and must be a positive number",
		})
	}

	req := &requests.UpdateCategoryRequest{
		CategoryID:  &id,
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "invalid_input",
			Message: "Invalid category update data. Please verify the fields.",
		})
	}

	category, err := s.categoryService.UpdateCategory(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to update category: ",
		})
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully updated category", category)
	return so, nil
}

func (s *categoryHandleGrpc) TrashedCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Invalid category id",
		})
	}

	category, err := s.categoryService.TrashedCategory(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to trashed category: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully trashed category", category)

	return so, nil
}

func (s *categoryHandleGrpc) RestoreCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Category ID parameter cannot be empty and must be a positive number",
		})
	}

	category, err := s.categoryService.RestoreCategory(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore category: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully restored category", category)

	return so, nil
}

func (s *categoryHandleGrpc) DeleteCategoryPermanent(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDelete, error) {
	if request.GetId() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "%v", &pb.ErrorResponse{
			Status:  "validation_error",
			Message: "Category ID parameter cannot be empty and must be a positive number",
		})
	}

	_, err := s.categoryService.DeleteCategoryPermanent(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete category permanently: " + err.Message,
		})
	}

	so := s.mapping.ToProtoResponseCategoryDelete("success", "Successfully deleted category permanently")

	return so, nil
}

func (s *categoryHandleGrpc) RestoreAllCategory(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.RestoreAllCategories()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all category: ",
		})
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully restore all category")

	return so, nil
}

func (s *categoryHandleGrpc) DeleteAllCategoryPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.DeleteAllCategoriesPermanent()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", &pb.ErrorResponse{
			Status:  "error",
			Message: "Failed to delete category permanent: ",
		})
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully delete category permanen")

	return so, nil
}
