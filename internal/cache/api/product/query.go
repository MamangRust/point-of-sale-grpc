package product_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	"time"
)

const (
	productAllCacheKey      = "product:all:page:%d:pageSize:%d:search:%s"
	productCategoryCacheKey = "product:category:%s:page:%d:pageSize:%d:search:%s"
	productMerchantCacheKey = "product:merchant:%d:page:%d:pageSize:%d:search:%s"

	productActiveCacheKey  = "product:active:page:%d:pageSize:%d:search:%s"
	productTrashedCacheKey = "product:trashed:page:%d:pageSize:%d:search:%s"
	productByIdCacheKey    = "product:id:%d"

	ttlDefault = 5 * time.Minute
)

type productQueryCache struct {
	store *cache.CacheStore
}

func NewProductQueryCache(store *cache.CacheStore) *productQueryCache {
	return &productQueryCache{store: store}
}

func (p *productQueryCache) GetCachedProducts(ctx context.Context, req *requests.FindAllProducts) (*response.ApiResponsePaginationProduct, bool) {
	key := fmt.Sprintf(productAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (p *productQueryCache) SetCachedProducts(ctx context.Context, req *requests.FindAllProducts, res *response.ApiResponsePaginationProduct) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(productAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, p.store, key, res, ttlDefault)
}

func (p *productQueryCache) GetCachedProductsByMerchant(ctx context.Context, req *requests.ProductByMerchantRequest) (*response.ApiResponsePaginationProduct, bool) {
	key := fmt.Sprintf(productMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductsByMerchant(ctx context.Context, req *requests.ProductByMerchantRequest, res *response.ApiResponsePaginationProduct) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(productMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, p.store, key, res, ttlDefault)
}

func (p *productQueryCache) GetCachedProductsByCategory(ctx context.Context, req *requests.ProductByCategoryRequest) (*response.ApiResponsePaginationProduct, bool) {
	key := fmt.Sprintf(productCategoryCacheKey, req.CategoryName, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductsByCategory(ctx context.Context, req *requests.ProductByCategoryRequest, res *response.ApiResponsePaginationProduct) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(productCategoryCacheKey, req.CategoryName, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, p.store, key, res, ttlDefault)
}

func (p *productQueryCache) GetCachedProductActive(ctx context.Context, req *requests.FindAllProducts) (*response.ApiResponsePaginationProductDeleteAt, bool) {
	key := fmt.Sprintf(productActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationProductDeleteAt](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductActive(ctx context.Context, req *requests.FindAllProducts, res *response.ApiResponsePaginationProductDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(productActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, p.store, key, res, ttlDefault)
}

func (p *productQueryCache) GetCachedProductTrashed(ctx context.Context, req *requests.FindAllProducts) (*response.ApiResponsePaginationProductDeleteAt, bool) {
	key := fmt.Sprintf(productTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationProductDeleteAt](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProductTrashed(ctx context.Context, req *requests.FindAllProducts, res *response.ApiResponsePaginationProductDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(productTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, p.store, key, res, ttlDefault)
}

func (p *productQueryCache) GetCachedProduct(ctx context.Context, productID int) (*response.ApiResponseProduct, bool) {
	key := fmt.Sprintf(productByIdCacheKey, productID)

	result, found := cache.GetFromCache[*response.ApiResponseProduct](ctx, p.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (p *productQueryCache) SetCachedProduct(ctx context.Context, res *response.ApiResponseProduct) {
	if res == nil || res.Data == nil {
		return
	}

	key := fmt.Sprintf(productByIdCacheKey, res.Data.ID)
	cache.SetToCache(ctx, p.store, key, res, ttlDefault)
}
