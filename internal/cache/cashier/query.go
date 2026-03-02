package cashier_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
)

type cashierListCacheResponse[T any] struct {
	Data         []T  `json:"data"`
	TotalRecords *int `json:"totalRecords"`
}

type cashierQueryCache struct {
	store *cache.CacheStore
}

func NewCashierQueryCache(store *cache.CacheStore) *cashierQueryCache {
	return &cashierQueryCache{store: store}
}

func (s *cashierQueryCache) GetCachedCashiersCache(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersRow, *int, bool) {
	key := fmt.Sprintf(cashierAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[cashierListCacheResponse[*db.GetCashiersRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *cashierQueryCache) SetCachedCashiersCache(ctx context.Context, req *requests.FindAllCashiers, res []*db.GetCashiersRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if res == nil {
		res = []*db.GetCashiersRow{}
	}

	key := fmt.Sprintf(cashierAllCacheKey, req.Page, req.PageSize, req.Search)
	payload := &cashierListCacheResponse[*db.GetCashiersRow]{Data: res, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashiersByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant) ([]*db.GetCashiersByMerchantRow, *int, bool) {
	key := fmt.Sprintf(cashierByMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[cashierListCacheResponse[*db.GetCashiersByMerchantRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *cashierQueryCache) SetCachedCashiersByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant, res []*db.GetCashiersByMerchantRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if res == nil {
		res = []*db.GetCashiersByMerchantRow{}
	}

	key := fmt.Sprintf(cashierByMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)
	payload := &cashierListCacheResponse[*db.GetCashiersByMerchantRow]{Data: res, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashiersActive(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersActiveRow, *int, bool) {
	key := fmt.Sprintf(cashierActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[cashierListCacheResponse[*db.GetCashiersActiveRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *cashierQueryCache) SetCachedCashiersActive(ctx context.Context, req *requests.FindAllCashiers, res []*db.GetCashiersActiveRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if res == nil {
		res = []*db.GetCashiersActiveRow{}
	}

	key := fmt.Sprintf(cashierActiveCacheKey, req.Page, req.PageSize, req.Search)
	payload := &cashierListCacheResponse[*db.GetCashiersActiveRow]{Data: res, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashiersTrashed(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersTrashedRow, *int, bool) {
	key := fmt.Sprintf(cashierTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[cashierListCacheResponse[*db.GetCashiersTrashedRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *cashierQueryCache) SetCachedCashiersTrashed(ctx context.Context, req *requests.FindAllCashiers, res []*db.GetCashiersTrashedRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if res == nil {
		res = []*db.GetCashiersTrashedRow{}
	}

	key := fmt.Sprintf(cashierTrashedCacheKey, req.Page, req.PageSize, req.Search)
	payload := &cashierListCacheResponse[*db.GetCashiersTrashedRow]{Data: res, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *cashierQueryCache) GetCachedCashier(ctx context.Context, cashierID int) (*db.GetCashierByIdRow, bool) {
	key := fmt.Sprintf(cashierByIdCacheKey, cashierID)

	result, found := cache.GetFromCache[*db.GetCashierByIdRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierQueryCache) SetCachedCashier(ctx context.Context, res *db.GetCashierByIdRow) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(cashierByIdCacheKey, res.CashierID)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}
