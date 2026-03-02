package order_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	"time"
)

const (
	orderAllCacheKey      = "order:all:page:%d:pageSize:%d:search:%s"
	orderByIdCacheKey     = "order:id:%d"
	orderActiveCacheKey   = "order:active:page:%d:pageSize:%d:search:%s"
	orderTrashedCacheKey  = "order:trashed:page:%d:pageSize:%d:search:%s"
	orderMerchantCacheKey = "order:merchant:%d:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type orderQueryCache struct {
	store *cache.CacheStore
}

func NewOrderQueryCache(store *cache.CacheStore) *orderQueryCache {
	return &orderQueryCache{store: store}
}

func (s *orderQueryCache) GetOrderAllCache(ctx context.Context, req *requests.FindAllOrders) (*response.ApiResponsePaginationOrder, bool) {
	key := fmt.Sprintf(orderAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationOrder](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetOrderAllCache(ctx context.Context, req *requests.FindAllOrders, res *response.ApiResponsePaginationOrder) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(orderAllCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *orderQueryCache) GetCachedOrderCache(ctx context.Context, orderID int) (*response.ApiResponseOrder, bool) {
	key := fmt.Sprintf(orderByIdCacheKey, orderID)

	result, found := cache.GetFromCache[*response.ApiResponseOrder](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetCachedOrderCache(ctx context.Context, res *response.ApiResponseOrder) {
	if res == nil || res.Data == nil {
		return
	}

	key := fmt.Sprintf(orderByIdCacheKey, res.Data.ID)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *orderQueryCache) GetCachedOrderMerchant(ctx context.Context, req *requests.FindAllOrderMerchant) (*response.ApiResponsePaginationOrder, bool) {
	key := fmt.Sprintf(orderMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationOrder](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetCachedOrderMerchant(ctx context.Context, req *requests.FindAllOrderMerchant, res *response.ApiResponsePaginationOrder) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(orderMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *orderQueryCache) GetOrderActiveCache(ctx context.Context, req *requests.FindAllOrders) (*response.ApiResponsePaginationOrderDeleteAt, bool) {
	key := fmt.Sprintf(orderActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationOrderDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetOrderActiveCache(ctx context.Context, req *requests.FindAllOrders, res *response.ApiResponsePaginationOrderDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(orderActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *orderQueryCache) GetOrderTrashedCache(ctx context.Context, req *requests.FindAllOrders) (*response.ApiResponsePaginationOrderDeleteAt, bool) {
	key := fmt.Sprintf(orderTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationOrderDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetOrderTrashedCache(ctx context.Context, req *requests.FindAllOrders, res *response.ApiResponsePaginationOrderDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(orderTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}
