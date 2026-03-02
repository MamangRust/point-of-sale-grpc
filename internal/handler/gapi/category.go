package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/category_errors"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type categoryHandleGrpc struct {
	pb.UnimplementedCategoryServiceServer
	categoryService service.CategoryService
}

func NewCategoryHandleGrpc(
	categoryService service.CategoryService,
) *categoryHandleGrpc {
	return &categoryHandleGrpc{
		categoryService: categoryService,
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

	categories, totalRecords, err := s.categoryService.FindAllCategory(ctx, &reqService)
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

	var categoryResponses []*pb.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, &pb.CategoryResponse{
			Id:           int32(category.CategoryID),
			Name:         category.Name,
			Description:  *category.Description,
			SlugCategory: *category.SlugCategory,
			CreatedAt:    category.CreatedAt.Time.String(),
			UpdatedAt:    category.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationCategory{
		Status:     "success",
		Message:    "Successfully fetched categories",
		Data:       categoryResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *categoryHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategory, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryService.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCategory{
		Status:  "success",
		Message: "Successfully fetched category",
		Data: &pb.CategoryResponse{
			Id:           int32(category.CategoryID),
			Name:         category.Name,
			Description:  *category.Description,
			SlugCategory: *category.SlugCategory,
			CreatedAt:    category.CreatedAt.Time.String(),
			UpdatedAt:    category.UpdatedAt.Time.String(),
		},
	}, nil
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

	categories, totalRecords, err := s.categoryService.FindByActive(ctx, &reqService)
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

	var categoryResponses []*pb.CategoryResponseDeleteAt
	for _, category := range categories {
		var deletedAt string
		if category.DeletedAt.Valid {
			deletedAt = category.DeletedAt.Time.String()
		}

		categoryResponses = append(categoryResponses, &pb.CategoryResponseDeleteAt{
			Id:           int32(category.CategoryID),
			Name:         category.Name,
			Description:  *category.Description,
			SlugCategory: *category.SlugCategory,
			CreatedAt:    category.CreatedAt.Time.String(),
			UpdatedAt:    category.UpdatedAt.Time.String(),
			DeletedAt:    &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationCategoryDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active categories",
		Data:       categoryResponses,
		Pagination: paginationMeta,
	}, nil
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

	categories, totalRecords, err := s.categoryService.FindByTrashed(ctx, &reqService)
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

	var categoryResponses []*pb.CategoryResponseDeleteAt
	for _, category := range categories {
		var deletedAt string
		if category.DeletedAt.Valid {
			deletedAt = category.DeletedAt.Time.String()
		}

		categoryResponses = append(categoryResponses, &pb.CategoryResponseDeleteAt{
			Id:           int32(category.CategoryID),
			Name:         category.Name,
			Description:  *category.Description,
			SlugCategory: *category.SlugCategory,
			CreatedAt:    category.CreatedAt.Time.String(),
			UpdatedAt:    category.UpdatedAt.Time.String(),
			DeletedAt:    &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationCategoryDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed categories",
		Data:       categoryResponses,
		Pagination: paginationMeta,
	}, nil
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

	prices, err := s.categoryService.FindMonthlyTotalPrice(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoriesMonthlyTotalPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoriesMonthlyTotalPriceResponse{
			Year:         price.Year,
			Month:        price.Month,
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly total prices retrieved successfully",
		Data:    priceResponses,
	}, nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPrices(ctx context.Context, req *pb.FindYearTotalPrices) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	prices, err := s.categoryService.FindYearlyTotalPrice(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoriesYearlyTotalPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoriesYearlyTotalPriceResponse{
			Year:         price.Year,
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly total prices retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindMonthlyTotalPriceById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoriesMonthlyTotalPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoriesMonthlyTotalPriceResponse{
			Year:         price.Year,
			Month:        price.Month,
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly total prices by ID retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindYearlyTotalPriceById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoriesYearlyTotalPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoriesYearlyTotalPriceResponse{
			Year:         price.Year,
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly total prices by ID retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindMonthlyTotalPriceByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoriesMonthlyTotalPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoriesMonthlyTotalPriceResponse{
			Year:         price.Year,
			Month:        price.Month,
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryMonthlyTotalPrice{
		Status:  "success",
		Message: "Monthly total prices by merchant retrieved successfully",
		Data:    priceResponses,
	}, nil
}

func (s *categoryHandleGrpc) FindYearlyTotalPricesByMerchant(ctx context.Context, req *pb.FindYearTotalPriceByMerchant) (*pb.ApiResponseCategoryYearlyTotalPrice, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	reqService := requests.YearTotalPriceMerchant{
		Year:       year,
		MerchantID: id,
	}

	prices, err := s.categoryService.FindYearlyTotalPriceByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoriesYearlyTotalPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoriesYearlyTotalPriceResponse{
			Year:         price.Year,
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryYearlyTotalPrice{
		Status:  "success",
		Message: "Yearly total prices by merchant retrieved successfully",
		Data:    priceResponses,
	}, nil
}

func (s *categoryHandleGrpc) FindMonthPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryMonthPrice, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	prices, err := s.categoryService.FindMonthPrice(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoryMonthPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoryMonthPriceResponse{
			Month:        price.Month,
			CategoryId:   int32(price.CategoryID),
			CategoryName: price.CategoryName,
			OrderCount:   int32(price.OrderCount),
			ItemsSold:    int32(price.ItemsSold),
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly category prices retrieved successfully",
		Data:    priceResponses,
	}, nil
}

func (s *categoryHandleGrpc) FindYearPrice(ctx context.Context, req *pb.FindYearCategory) (*pb.ApiResponseCategoryYearPrice, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, category_errors.ErrGrpcFailedInvalidYear
	}

	prices, err := s.categoryService.FindYearPrice(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoryYearPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoryYearPriceResponse{
			Year:               price.Year,
			CategoryId:         int32(price.CategoryID),
			CategoryName:       price.CategoryName,
			OrderCount:         int32(price.OrderCount),
			ItemsSold:          int32(price.ItemsSold),
			TotalRevenue:       int32(price.TotalRevenue),
			UniqueProductsSold: int32(price.UniqueProductsSold),
		})
	}

	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly category prices retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindMonthPriceByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoryMonthPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoryMonthPriceResponse{
			Month:        price.Month,
			CategoryId:   int32(price.CategoryID),
			CategoryName: price.CategoryName,
			OrderCount:   int32(price.OrderCount),
			ItemsSold:    int32(price.ItemsSold),
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly category prices by merchant retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindYearPriceByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoryYearPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoryYearPriceResponse{
			Year:               price.Year,
			CategoryId:         int32(price.CategoryID),
			CategoryName:       price.CategoryName,
			OrderCount:         int32(price.OrderCount),
			ItemsSold:          int32(price.ItemsSold),
			TotalRevenue:       int32(price.TotalRevenue),
			UniqueProductsSold: int32(price.UniqueProductsSold),
		})
	}

	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly category prices by merchant retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindMonthPriceById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoryMonthPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoryMonthPriceResponse{
			Month:        price.Month,
			CategoryId:   int32(price.CategoryID),
			CategoryName: price.CategoryName,
			OrderCount:   int32(price.OrderCount),
			ItemsSold:    int32(price.ItemsSold),
			TotalRevenue: int32(price.TotalRevenue),
		})
	}

	return &pb.ApiResponseCategoryMonthPrice{
		Status:  "success",
		Message: "Monthly category prices by ID retrieved successfully",
		Data:    priceResponses,
	}, nil
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

	prices, err := s.categoryService.FindYearPriceById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var priceResponses []*pb.CategoryYearPriceResponse
	for _, price := range prices {
		priceResponses = append(priceResponses, &pb.CategoryYearPriceResponse{
			Year:               price.Year,
			CategoryId:         int32(price.CategoryID),
			CategoryName:       price.CategoryName,
			OrderCount:         int32(price.OrderCount),
			ItemsSold:          int32(price.ItemsSold),
			TotalRevenue:       int32(price.TotalRevenue),
			UniqueProductsSold: int32(price.UniqueProductsSold),
		})
	}

	return &pb.ApiResponseCategoryYearPrice{
		Status:  "success",
		Message: "Yearly category prices by ID retrieved successfully",
		Data:    priceResponses,
	}, nil
}

func (s *categoryHandleGrpc) Create(ctx context.Context, request *pb.CreateCategoryRequest) (*pb.ApiResponseCategory, error) {
	req := &requests.CreateCategoryRequest{
		Name:        request.GetName(),
		Description: request.GetDescription(),
	}

	if err := req.Validate(); err != nil {
		return nil, category_errors.ErrGrpcValidateCreateCategory
	}

	category, err := s.categoryService.CreateCategory(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	pbData := &pb.CategoryResponse{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  *category.Description,
		SlugCategory: *category.SlugCategory,
		CreatedAt:    category.CreatedAt.Time.String(),
		UpdatedAt:    category.UpdatedAt.Time.String(),
	}

	return &pb.ApiResponseCategory{
		Status:  "success",
		Message: "Successfully updated category",
		Data:    pbData,
	}, nil
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

	category, err := s.categoryService.UpdateCategory(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	pbData := &pb.CategoryResponse{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  *category.Description,
		SlugCategory: *category.SlugCategory,
		CreatedAt:    category.CreatedAt.Time.String(),
		UpdatedAt:    category.UpdatedAt.Time.String(),
	}

	return &pb.ApiResponseCategory{
		Status:  "success",
		Message: "Successfully updated category",
		Data:    pbData,
	}, nil
}

func (s *categoryHandleGrpc) TrashedCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryService.TrashedCategory(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if category.DeletedAt.Valid {
		deletedAt = category.DeletedAt.Time.String()
	}

	pbData := &pb.CategoryResponseDeleteAt{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  *category.Description,
		SlugCategory: *category.SlugCategory,
		CreatedAt:    category.CreatedAt.Time.String(),
		UpdatedAt:    category.UpdatedAt.Time.String(),
		DeletedAt:    &wrapperspb.StringValue{Value: deletedAt},
	}

	return &pb.ApiResponseCategoryDeleteAt{
		Status:  "success",
		Message: "Successfully trashed category",
		Data:    pbData,
	}, nil
}

func (s *categoryHandleGrpc) RestoreCategory(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	category, err := s.categoryService.RestoreCategory(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if category.DeletedAt.Valid {
		deletedAt = category.DeletedAt.Time.String()
	}

	pbData := &pb.CategoryResponseDeleteAt{
		Id:           int32(category.CategoryID),
		Name:         category.Name,
		Description:  *category.Description,
		SlugCategory: *category.SlugCategory,
		CreatedAt:    category.CreatedAt.Time.String(),
		UpdatedAt:    category.UpdatedAt.Time.String(),
		DeletedAt:    &wrapperspb.StringValue{Value: deletedAt},
	}

	return &pb.ApiResponseCategoryDeleteAt{
		Status:  "success",
		Message: "Successfully restored category",
		Data:    pbData,
	}, nil
}

func (s *categoryHandleGrpc) DeleteCategoryPermanent(ctx context.Context, request *pb.FindByIdCategoryRequest) (*pb.ApiResponseCategoryDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, category_errors.ErrGrpcFailedInvalidId
	}

	_, err := s.categoryService.DeleteCategoryPermanently(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCategoryDelete{
		Status:  "success",
		Message: "Successfully deleted category permanently",
	}, nil
}

func (s *categoryHandleGrpc) RestoreAllCategory(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.RestoreAllCategories(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCategoryAll{
		Status:  "success",
		Message: "Successfully restore all categories",
	}, nil
}

func (s *categoryHandleGrpc) DeleteAllCategoryPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCategoryAll, error) {
	_, err := s.categoryService.DeleteAllPermanentCategories(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCategoryAll{
		Status:  "success",
		Message: "Successfully delete category permanent",
	}, nil
}
