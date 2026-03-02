package cashier_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type cashierQueryCache struct {
	store *cache.CacheStore
}

func NewCashierQueryCache(store *cache.CacheStore) *cashierQueryCache {
	return &cashierQueryCache{store: store}
}

func (s *cashierQueryCache) GetCachedCashiersCache(ctx context.Context, req *requests.FindAllCashiers) (*response.ApiResponsePaginationCashier, bool) {
	key := fmt.Sprintf(cashierAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationCashier](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierQueryCache) SetCachedCashiersCache(ctx context.Context, req *requests.FindAllCashiers, res *response.ApiResponsePaginationCashier) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(cashierAllCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashier(ctx context.Context, cashierID int) (*response.ApiResponseCashier, bool) {
	key := fmt.Sprintf(cashierByIdCacheKey, cashierID)

	result, found := cache.GetFromCache[*response.ApiResponseCashier](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierQueryCache) SetCachedCashier(ctx context.Context, res *response.ApiResponseCashier) {
	if res == nil || res.Data == nil {
		return
	}

	key := fmt.Sprintf(cashierByIdCacheKey, res.Data.ID)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashiersActive(ctx context.Context, req *requests.FindAllCashiers) (*response.ApiResponsePaginationCashierDeleteAt, bool) {
	key := fmt.Sprintf(cashierActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationCashierDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierQueryCache) SetCachedCashiersActive(ctx context.Context, req *requests.FindAllCashiers, res *response.ApiResponsePaginationCashierDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(cashierActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashiersTrashed(ctx context.Context, req *requests.FindAllCashiers) (*response.ApiResponsePaginationCashierDeleteAt, bool) {
	key := fmt.Sprintf(cashierTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationCashierDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierQueryCache) SetCachedCashiersTrashed(ctx context.Context, req *requests.FindAllCashiers, res *response.ApiResponsePaginationCashierDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(cashierTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashiersByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant) (*response.ApiResponsePaginationCashier, bool) {
	key := fmt.Sprintf(cashierByMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationCashier](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierQueryCache) SetCachedCashiersByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant, res *response.ApiResponsePaginationCashier) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(cashierByMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}
