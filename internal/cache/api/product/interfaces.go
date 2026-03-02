package product_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type ProductQueryCache interface {
	GetCachedProducts(ctx context.Context, req *requests.FindAllProducts) (*response.ApiResponsePaginationProduct, bool)
	SetCachedProducts(ctx context.Context, req *requests.FindAllProducts, res *response.ApiResponsePaginationProduct)

	GetCachedProductsByMerchant(ctx context.Context, req *requests.ProductByMerchantRequest) (*response.ApiResponsePaginationProduct, bool)
	SetCachedProductsByMerchant(ctx context.Context, req *requests.ProductByMerchantRequest, res *response.ApiResponsePaginationProduct)

	GetCachedProductsByCategory(ctx context.Context, req *requests.ProductByCategoryRequest) (*response.ApiResponsePaginationProduct, bool)
	SetCachedProductsByCategory(ctx context.Context, req *requests.ProductByCategoryRequest, res *response.ApiResponsePaginationProduct)

	GetCachedProductActive(ctx context.Context, req *requests.FindAllProducts) (*response.ApiResponsePaginationProductDeleteAt, bool)
	SetCachedProductActive(ctx context.Context, req *requests.FindAllProducts, res *response.ApiResponsePaginationProductDeleteAt)

	GetCachedProductTrashed(ctx context.Context, req *requests.FindAllProducts) (*response.ApiResponsePaginationProductDeleteAt, bool)
	SetCachedProductTrashed(ctx context.Context, req *requests.FindAllProducts, res *response.ApiResponsePaginationProductDeleteAt)

	GetCachedProduct(ctx context.Context, productID int) (*response.ApiResponseProduct, bool)
	SetCachedProduct(ctx context.Context, res *response.ApiResponseProduct)
}

type ProductCommandCache interface {
	DeleteCachedProduct(ctx context.Context, productID int)
}
