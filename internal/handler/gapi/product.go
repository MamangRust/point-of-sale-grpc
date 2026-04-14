package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/product_errors"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type productHandleGrpc struct {
	pb.UnimplementedProductServiceServer
	productService service.ProductService
}

func NewProductHandleGrpc(
	productService service.ProductService,
) *productHandleGrpc {
	return &productHandleGrpc{
		productService: productService,
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

	reqService := requests.FindAllProducts{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	products, totalRecords, err := s.productService.FindAllProducts(ctx, &reqService)
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

	var productResponses []*pb.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &pb.ProductResponse{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationProduct{
		Status:     "success",
		Message:    "Successfully fetched products",
		Data:       productResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *productHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindAllProductMerchantRequest) (*pb.ApiResponsePaginationProduct, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	merchant_id := int(request.GetMerchantId())
	min_price := int(request.GetMinPrice())
	max_price := int(request.GetMaxPrice())

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	if min_price <= 0 {
		min_price = 0
	}

	if max_price <= 0 {
		max_price = 0
	}

	reqService := requests.ProductByMerchantRequest{
		MerchantID: merchant_id,
		Page:       page,
		PageSize:   pageSize,
		Search:     search,
		MinPrice:   min_price,
		MaxPrice:   max_price,
	}

	products, totalRecords, err := s.productService.FindByMerchant(ctx, &reqService)
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

	var productResponses []*pb.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &pb.ProductResponse{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationProduct{
		Status:     "success",
		Message:    "Successfully fetched products by merchant",
		Data:       productResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *productHandleGrpc) FindByCategory(ctx context.Context, request *pb.FindAllProductCategoryRequest) (*pb.ApiResponsePaginationProduct, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	category_name := request.GetCategoryName()
	max_price := int(request.GetMaxprice())
	min_price := int(request.GetMinprice())

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.ProductByCategoryRequest{
		Page:         page,
		PageSize:     pageSize,
		Search:       search,
		CategoryName: category_name,
		MaxPrice:     max_price,
		MinPrice:     min_price,
	}

	products, totalRecords, err := s.productService.FindByCategory(ctx, &reqService)
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

	var productResponses []*pb.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, &pb.ProductResponse{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationProduct{
		Status:     "success",
		Message:    "Successfully fetched products by category",
		Data:       productResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *productHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProduct, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	product, err := s.productService.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseProduct{
		Status:  "success",
		Message: "Successfully fetched product",
		Data: &pb.ProductResponse{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
		},
	}, nil
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

	reqService := requests.FindAllProducts{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	products, totalRecords, err := s.productService.FindByActive(ctx, &reqService)
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

	var productResponses []*pb.ProductResponseDeleteAt
	for _, product := range products {
		var deletedAt *wrapperspb.StringValue
		if product.DeletedAt.Valid {
			deletedAt = &wrapperspb.StringValue{Value: product.DeletedAt.Time.String()}
		}

		productResponses = append(productResponses, &pb.ProductResponseDeleteAt{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		})
	}

	return &pb.ApiResponsePaginationProductDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active products",
		Data:       productResponses,
		Pagination: paginationMeta,
	}, nil
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

	reqService := requests.FindAllProducts{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	products, totalRecords, err := s.productService.FindByTrashed(ctx, &reqService)
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

	var productResponses []*pb.ProductResponseDeleteAt
	for _, product := range products {
		var deletedAt *wrapperspb.StringValue
		if product.DeletedAt.Valid {
			deletedAt = &wrapperspb.StringValue{Value: product.DeletedAt.Time.String()}
		}

		productResponses = append(productResponses, &pb.ProductResponseDeleteAt{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		})
	}

	return &pb.ApiResponsePaginationProductDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed products",
		Data:       productResponses,
		Pagination: paginationMeta,
	}, nil
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
		return nil, product_errors.ErrGrpcValidateCreateProduct
	}

	product, err := s.productService.CreateProduct(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseProduct{
		Status:  "success",
		Message: "Successfully created product",
		Data: &pb.ProductResponse{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *productHandleGrpc) Update(ctx context.Context, request *pb.UpdateProductRequest) (*pb.ApiResponseProduct, error) {
	id := int(request.GetProductId())

	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
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
		return nil, product_errors.ErrGrpcValidateUpdateProduct
	}

	product, err := s.productService.UpdateProduct(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseProduct{
		Status:  "success",
		Message: "Successfully updated product",
		Data: &pb.ProductResponse{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *productHandleGrpc) TrashedProduct(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	product, err := s.productService.TrashedProduct(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt *wrapperspb.StringValue
	if product.DeletedAt.Valid {
		deletedAt = &wrapperspb.StringValue{Value: product.DeletedAt.Time.String()}
	}

	return &pb.ApiResponseProductDeleteAt{
		Status:  "success",
		Message: "Successfully trashed product",
		Data: &pb.ProductResponseDeleteAt{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		},
	}, nil
}

func (s *productHandleGrpc) RestoreProduct(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	product, err := s.productService.RestoreProduct(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt *wrapperspb.StringValue
	if product.DeletedAt.Valid {
		deletedAt = &wrapperspb.StringValue{Value: product.DeletedAt.Time.String()}
	}

	return &pb.ApiResponseProductDeleteAt{
		Status:  "success",
		Message: "Successfully restored product",
		Data: &pb.ProductResponseDeleteAt{
			Id:           int32(product.ProductID),
			MerchantId:   int32(product.MerchantID),
			CategoryId:   int32(product.CategoryID),
			Name:         product.Name,
			Description:  *product.Description,
			Price:        int32(product.Price),
			CountInStock: int32(product.CountInStock),
			Brand:        *product.Brand,
			Weight:       int32(*product.Weight),
			SlugProduct:  *product.SlugProduct,
			ImageProduct: *product.ImageProduct,
			Barcode:      *product.Barcode,
			CreatedAt:    product.CreatedAt.Time.String(),
			UpdatedAt:    product.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		},
	}, nil
}

func (s *productHandleGrpc) DeleteProductPermanent(ctx context.Context, request *pb.FindByIdProductRequest) (*pb.ApiResponseProductDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, product_errors.ErrGrpcInvalidID
	}

	_, err := s.productService.DeleteProductPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseProductDelete{
		Status:  "success",
		Message: "Successfully deleted product permanently",
	}, nil
}

func (s *productHandleGrpc) RestoreAllProduct(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseProductAll, error) {
	_, err := s.productService.RestoreAllProducts(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseProductAll{
		Status:  "success",
		Message: "Successfully restore all products",
	}, nil
}

func (s *productHandleGrpc) DeleteAllProductPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseProductAll, error) {
	_, err := s.productService.DeleteAllProductPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseProductAll{
		Status:  "success",
		Message: "Successfully delete all products permanently",
	}, nil
}
