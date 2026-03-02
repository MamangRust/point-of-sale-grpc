package order_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
	"time"
)

const (
	orderAllCacheKey     = "order:all:page:%d:pageSize:%d:search:%s"
	orderByIdCacheKey    = "order:id:%d"
	orderActiveCacheKey  = "order:active:page:%d:pageSize:%d:search:%s"
	orderTrashedCacheKey = "order:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

// ... (definisi cache key dan ttlDefault tetap sama) ...

// Asumsi ada cache key khusus untuk query berdasarkan merchant
const orderMerchantCacheKey = "order:merchant:%d:page:%d:pageSize:%d:search:%s"

// Menggunakan struct generik untuk membungkus data cache yang berupa list.
// Ini mengurangi pengulangan kode untuk setiap tipe db yang berbeda.
type orderListCacheResponse[T any] struct {
	Data         []T  `json:"data"`
	TotalRecords *int `json:"total_records"`
}

type orderQueryCache struct {
	store *cache.CacheStore
}

func NewOrderQueryCache(store *cache.CacheStore) *orderQueryCache {
	return &orderQueryCache{store: store}
}

func (s *orderQueryCache) GetOrderAllCache(ctx context.Context, req *requests.FindAllOrders) ([]*db.GetOrdersRow, *int, bool) {
	key := fmt.Sprintf(orderAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[orderListCacheResponse[*db.GetOrdersRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *orderQueryCache) SetOrderAllCache(ctx context.Context, req *requests.FindAllOrders, data []*db.GetOrdersRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetOrdersRow{}
	}

	key := fmt.Sprintf(orderAllCacheKey, req.Page, req.PageSize, req.Search)
	payload := &orderListCacheResponse[*db.GetOrdersRow]{Data: data, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *orderQueryCache) GetCachedOrderCache(ctx context.Context, orderID int) (*db.GetOrderByIDRow, bool) {
	key := fmt.Sprintf(orderByIdCacheKey, orderID)

	result, found := cache.GetFromCache[*db.GetOrderByIDRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderQueryCache) SetCachedOrderCache(ctx context.Context, data *db.GetOrderByIDRow) {
	if data == nil {
		return
	}

	key := fmt.Sprintf(orderByIdCacheKey, data.OrderID)
	cache.SetToCache(ctx, s.store, key, data, ttlDefault)
}

func (s *orderQueryCache) GetCachedOrderMerchant(ctx context.Context, req *requests.FindAllOrderMerchant) ([]*db.GetOrdersByMerchantRow, *int, bool) {
	key := fmt.Sprintf(orderMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[orderListCacheResponse[*db.GetOrdersByMerchantRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *orderQueryCache) SetCachedOrderMerchant(ctx context.Context, req *requests.FindAllOrderMerchant, res []*db.GetOrdersByMerchantRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if res == nil {
		res = []*db.GetOrdersByMerchantRow{}
	}

	key := fmt.Sprintf(orderMerchantCacheKey, req.MerchantID, req.Page, req.PageSize, req.Search)
	payload := &orderListCacheResponse[*db.GetOrdersByMerchantRow]{Data: res, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *orderQueryCache) GetOrderActiveCache(ctx context.Context, req *requests.FindAllOrders) ([]*db.GetOrdersActiveRow, *int, bool) {
	key := fmt.Sprintf(orderActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[orderListCacheResponse[*db.GetOrdersActiveRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *orderQueryCache) SetOrderActiveCache(ctx context.Context, req *requests.FindAllOrders, data []*db.GetOrdersActiveRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetOrdersActiveRow{}
	}

	key := fmt.Sprintf(orderActiveCacheKey, req.Page, req.PageSize, req.Search)
	payload := &orderListCacheResponse[*db.GetOrdersActiveRow]{Data: data, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}

func (s *orderQueryCache) GetOrderTrashedCache(ctx context.Context, req *requests.FindAllOrders) ([]*db.GetOrdersTrashedRow, *int, bool) {
	key := fmt.Sprintf(orderTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[orderListCacheResponse[*db.GetOrdersTrashedRow]](ctx, s.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (s *orderQueryCache) SetOrderTrashedCache(ctx context.Context, req *requests.FindAllOrders, data []*db.GetOrdersTrashedRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetOrdersTrashedRow{}
	}

	key := fmt.Sprintf(orderTrashedCacheKey, req.Page, req.PageSize, req.Search)
	payload := &orderListCacheResponse[*db.GetOrdersTrashedRow]{Data: data, TotalRecords: total}
	cache.SetToCache(ctx, s.store, key, payload, ttlDefault)
}
