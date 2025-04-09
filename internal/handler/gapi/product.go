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

type productHandleGrpc struct {
	pb.UnimplementedProductServiceServer
	productService service.ProductService
	mapping        protomapper.ProductProtoMapper
}

func NewProductHandleGrpc(
	productService service.ProductService,
	mapping protomapper.ProductProtoMapper,
) *productHandleGrpc {
	return &productHandleGrpc{
		productService: productService,
		mapping:        mapping,
	}
}

func (s *productHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllProductRequest) (*pb.ApiResponsePaginationProduct, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	product, totalRecords, err := s.productService.FindAll(page, pageSize, search)

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

	so := s.mapping.ToProtoResponsePaginationProduct(paginationMeta, "success", "Successfully fetched product", product)
	return so, nil
}

func (s *productHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindAllProductMerchantRequest) (*pb.ApiResponsePaginationProduct, error) {
	req := requests.ProductByMerchantRequest{
		MerchantID: int(request.GetMerchantId()),
		Search:     request.GetSearch(),
		Page:       int(request.GetPage()),
		PageSize:   int(request.GetPageSize()),
	}

	if request.CategoryId != nil {
		catID := int(request.GetCategoryId().Value)
		req.CategoryID = &catID
	}
	if request.MinPrice != nil {
		minPrice := int(request.GetMinPrice().Value)
		req.MinPrice = &minPrice
	}
	if request.MaxPrice != nil {
		maxPrice := int(request.GetMaxPrice().Value)
		req.MaxPrice = &maxPrice
	}

	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	products, totalRecords, err := s.productService.FindByMerchant(&req)

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

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(req.PageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(req.Page),
		PageSize:     int32(req.PageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	return s.mapping.ToProtoResponsePaginationProduct(
		paginationMeta,
		"success",
		"Successfully fetched products",
		products,
	), nil
}

func (s *productHandleGrpc) FindByCategory(ctx context.Context, request *pb.FindAllProductCategoryRequest) (*pb.ApiResponsePaginationProduct, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	category_name := request.GetCategoryName()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	product, totalRecords, err := s.productService.FindByCategory(category_name, page, pageSize, search)

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

	so := s.mapping.ToProtoResponsePaginationProduct(paginationMeta, "success", "Successfully fetched product", product)
	return so, nil
}

func (s *productHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProduct, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_request",
				Message: "Valid Category ID is required",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	product, err := s.productService.FindById(id)

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

	so := s.mapping.ToProtoResponseProduct("success", "Successfully fetched product", product)

	return so, nil

}

func (s *productHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllProductRequest) (*pb.ApiResponsePaginationProductDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	product, totalRecords, err := s.productService.FindByActive(search, page, pageSize)

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
	so := s.mapping.ToProtoResponsePaginationProductDeleteAt(paginationMeta, "success", "Successfully fetched active product", product)

	return so, nil
}

func (s *productHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllProductRequest) (*pb.ApiResponsePaginationProductDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	users, totalRecords, err := s.productService.FindByTrashed(search, page, pageSize)

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

	so := s.mapping.ToProtoResponsePaginationProductDeleteAt(paginationMeta, "success", "Successfully fetched trashed product", users)

	return so, nil
}

func (s *productHandleGrpc) Create(ctx context.Context, request *pb.CreateProductRequest) (*pb.ApiResponseProduct, error) {
	req := &requests.CreateProductRequest{
		MerchantID:   int(request.GetMerchantId()),
		CategoryID:   int(request.GetCategoryId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Price:        int(request.GetPrice()),
		CountInStock: int(request.GetCountInStock()),
		Brand:        request.GetBrand(),
		Weight:       int(request.GetWeight()),
		ImageProduct: request.GetImageProduct(),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Unable to create new product. Please check your input.",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	product, err := s.productService.CreateProduct(req)

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

	so := s.mapping.ToProtoResponseProduct("success", "Successfully created product", product)
	return so, nil
}

func (s *productHandleGrpc) Update(ctx context.Context, request *pb.UpdateProductRequest) (*pb.ApiResponseProduct, error) {
	id := int(request.GetProductId())

	if request.GetProductId() == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Product ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	req := &requests.UpdateProductRequest{
		ProductID:    &id,
		MerchantID:   int(request.GetMerchantId()),
		CategoryID:   int(request.GetCategoryId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Price:        int(request.GetPrice()),
		CountInStock: int(request.GetCountInStock()),
		Brand:        request.GetBrand(),
		Weight:       int(request.GetWeight()),
		ImageProduct: request.GetImageProduct(),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Unable to process product update. Please review your data.",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	product, err := s.productService.UpdateProduct(req)

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

	so := s.mapping.ToProtoResponseProduct("success", "Successfully updated product", product)
	return so, nil
}

func (s *productHandleGrpc) TrashedProduct(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Product ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	product, err := s.productService.TrashProduct(id)

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

	so := s.mapping.ToProtoResponseProductDeleteAt("success", "Successfully trashed product", product)

	return so, nil
}

func (s *productHandleGrpc) RestoreProduct(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Product ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	product, err := s.productService.RestoreProduct(id)

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

	so := s.mapping.ToProtoResponseProductDeleteAt("success", "Successfully restored product", product)

	return so, nil
}

func (s *productHandleGrpc) DeleteProductPermanent(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Product ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	_, err := s.productService.DeleteProductPermanent(int(request.GetId()))

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

	so := s.mapping.ToProtoResponseProductDelete("success", "Successfully deleted Product permanently")

	return so, nil
}

func (s *productHandleGrpc) RestoreAllProduct(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseProductAll, error) {
	_, err := s.productService.RestoreAllProducts()

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

	so := s.mapping.ToProtoResponseProductAll("success", "Successfully restore all Product")

	return so, nil
}

func (s *productHandleGrpc) DeleteAllProductPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseProductAll, error) {
	_, err := s.productService.DeleteAllProductsPermanent()

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

	so := s.mapping.ToProtoResponseProductAll("success", "Successfully delete Product permanen")

	return so, nil
}
