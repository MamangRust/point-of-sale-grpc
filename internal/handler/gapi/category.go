package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors/category_errors"

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

	reqService := requests.FindAllCategory{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	category, totalRecords, err := s.categoryService.FindAll(&reqService)

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

	so := s.mapping.ToProtoResponsePaginationCategory(paginationMeta, "success", "Successfully fetched categories", category)
	return so, nil
}

func (s *categoryHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategory, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryService.FindById(id)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
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

	reqService := requests.FindAllCategory{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	users, totalRecords, err := s.categoryService.FindByActive(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
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

	reqService := requests.FindAllCategory{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	users, totalRecords, err := s.categoryService.FindByTrashed(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
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
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, category_errors.ErrGrpcFailedInvalidMonth
	}

	reqService := requests.MonthTotalPrice{
		Year:  year,
		Month: month,
	}

	methods, err := s.categoryService.FindMonthlyTotalPrice(&reqService)
	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseMonthlyTotalPrice("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.categoryService.FindYearlyTotalPrice(year)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseYearlyTotalPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthlyTotalPricesById(ctx context.Context, req *pb.FindYearMonthTotalPriceById) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, category_errors.ErrGrpcFailedInvalidMonth
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthTotalPriceCategory{
		Year:       year,
		Month:      month,
		CategoryID: id,
	}

	methods, err := s.categoryService.FindMonthlyTotalPriceById(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseMonthlyTotalPrice("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPricesById(ctx context.Context, req *pb.FindYearTotalPriceById) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalPriceCategory{
		Year:       year,
		CategoryID: id,
	}

	methods, err := s.categoryService.FindYearlyTotalPriceById(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseYearlyTotalPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalPriceByMerchant) (*pb.ApiResponseCategoryMonthlyTotalPrice, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, category_errors.ErrGrpcFailedInvalidMonth
	}

	reqService := requests.MonthTotalPriceMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	methods, err := s.categoryService.FindMonthlyTotalPriceByMerchant(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseMonthlyTotalPrice("success", "Monthly sales retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.categoryService.FindYearlyTotalPriceByMerchant(&reqService)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseYearlyTotalPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryMonthPrice, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.categoryService.FindMonthPrice(year)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseCategoryMonthlyPrice("success", "Monthly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryYearPrice, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	methods, err := s.categoryService.FindYearPrice(year)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseCategoryYearlyPrice("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryMonthPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.categoryService.FindMonthPriceByMerchant(
		&reqService,
	)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseCategoryMonthlyPrice("success", "Merchant monthly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearPriceByMerchant(ctx context.Context, req *pb.FindYearCategoryByMerchant) (*pb.ApiResponseCategoryYearPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.categoryService.FindYearPriceByMerchant(
		&reqService,
	)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseCategoryYearlyPrice("success", "Merchant yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindMonthPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryMonthPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthPriceId{
		Year:       year,
		CategoryID: id,
	}

	methods, err := s.categoryService.FindMonthPriceById(
		&reqService,
	)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseCategoryMonthlyPrice("success", "Merchant monthly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) FindYearPriceById(ctx context.Context, req *pb.FindYearCategoryById) (*pb.ApiResponseCategoryYearPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetCategoryId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearPriceId{
		Year:       year,
		CategoryID: id,
	}

	methods, err := s.categoryService.FindYearPriceById(
		&reqService,
	)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	return s.mapping.ToProtoResponseCategoryYearlyPrice("success", "Merchant yearly payment methods retrieved successfully", methods), nil
}

func (s *categoryHandleGrpc) Create(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.ApiResponseCategory, error) {
	req := &requests.CreateCategoryRequest{
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateCreateCategory
	}

	category, err := s.categoryService.CreateCategory(req)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully created category", category)
	return so, nil
}

func (s *categoryHandleGrpc) Update(ctx context.Context, request *pb.UpdateCategoryRequest) (*pb.ApiResponseCategory, error) {
	id := int(request.GetCategoryId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	req := &requests.UpdateCategoryRequest{
		CategoryID:  &id,
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateUpdateCategory
	}

	category, err := s.categoryService.UpdateCategory(req)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategory("success", "Successfully updated category", category)
	return so, nil
}

func (s *categoryHandleGrpc) TrashedCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryService.TrashedCategory(id)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully trashed category", category)

	return so, nil
}

func (s *categoryHandleGrpc) RestoreCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryService.RestoreCategory(id)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategoryDeleteAt("success", "Successfully restored category", category)

	return so, nil
}

func (s *categoryHandleGrpc) DeleteCategoryPermanent(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	_, err := s.categoryService.DeleteCategoryPermanent(id)

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategoryDelete("success", "Successfully deleted category permanently")

	return so, nil
}

func (s *categoryHandleGrpc) RestoreAllCategory(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.RestoreAllCategories()

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully restore all category")

	return so, nil
}

func (s *categoryHandleGrpc) DeleteAllCategoryPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.DeleteAllCategoriesPermanent()

	if err != nil {
		return nil, response.ToGrpcErrorFromErrorResponse(err)
	}

	so := s.mapping.ToProtoResponseCategoryAll("success", "Successfully delete category permanen")

	return so, nil
}
