package orderitem_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
	"time"
)

const (
	orderItemAllCacheKey     = "order_item:all:page:%d:pageSize:%d:search:%s"
	orderItemActiveCacheKey  = "order_item:active:page:%d:pageSize:%d:search:%s"
	orderItemTrashedCacheKey = "order_item:trashed:page:%d:pageSize:%d:search:%s"

	orderItemByIdCacheKey = "order_item:id:%d"

	ttlDefault = 5 * time.Minute
)

type orderItemListCacheResponse[T any] struct {
	Data         []T  `json:"data"`
	TotalRecords *int `json:"total_records"`
}

type orderItemQueryCache struct {
	store *cache.CacheStore
}

func NewOrderItemQueryCache(store *cache.CacheStore) *orderItemQueryCache {
	return &orderItemQueryCache{store: store}
}

func (o *orderItemQueryCache) GetCachedOrderItemsAll(ctx context.Context, req *requests.FindAllOrderItems) ([]*db.GetOrderItemsRow, *int, bool) {
	key := fmt.Sprintf(orderItemAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[orderItemListCacheResponse[*db.GetOrderItemsRow]](ctx, o.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (o *orderItemQueryCache) SetCachedOrderItemsAll(ctx context.Context, req *requests.FindAllOrderItems, data []*db.GetOrderItemsRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetOrderItemsRow{}
	}

	key := fmt.Sprintf(orderItemAllCacheKey, req.Page, req.PageSize, req.Search)
	payload := &orderItemListCacheResponse[*db.GetOrderItemsRow]{Data: data, TotalRecords: total}
	cache.SetToCache(ctx, o.store, key, payload, ttlDefault)
}

func (o *orderItemQueryCache) GetCachedOrderItemActive(ctx context.Context, req *requests.FindAllOrderItems) ([]*db.GetOrderItemsActiveRow, *int, bool) {
	key := fmt.Sprintf(orderItemActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[orderItemListCacheResponse[*db.GetOrderItemsActiveRow]](ctx, o.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (o *orderItemQueryCache) SetCachedOrderItemActive(ctx context.Context, req *requests.FindAllOrderItems, data []*db.GetOrderItemsActiveRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetOrderItemsActiveRow{}
	}

	key := fmt.Sprintf(orderItemActiveCacheKey, req.Page, req.PageSize, req.Search)
	payload := &orderItemListCacheResponse[*db.GetOrderItemsActiveRow]{Data: data, TotalRecords: total}
	cache.SetToCache(ctx, o.store, key, payload, ttlDefault)
}

func (o *orderItemQueryCache) GetCachedOrderItemTrashed(ctx context.Context, req *requests.FindAllOrderItems) ([]*db.GetOrderItemsTrashedRow, *int, bool) {
	key := fmt.Sprintf(orderItemTrashedCacheKey, req.Page, req.PageSize, req.Search)
	result, found := cache.GetFromCache[orderItemListCacheResponse[*db.GetOrderItemsTrashedRow]](ctx, o.store, key)

	if !found || result.Data == nil {
		return nil, nil, false
	}

	return result.Data, result.TotalRecords, true
}

func (o *orderItemQueryCache) SetCachedOrderItemTrashed(ctx context.Context, req *requests.FindAllOrderItems, data []*db.GetOrderItemsTrashedRow, total *int) {
	if total == nil {
		zero := 0
		total = &zero
	}

	if data == nil {
		data = []*db.GetOrderItemsTrashedRow{}
	}

	key := fmt.Sprintf(orderItemTrashedCacheKey, req.Page, req.PageSize, req.Search)
	payload := &orderItemListCacheResponse[*db.GetOrderItemsTrashedRow]{Data: data, TotalRecords: total}
	cache.SetToCache(ctx, o.store, key, payload, ttlDefault)
}

func (o *orderItemQueryCache) GetCachedOrderItems(ctx context.Context, orderID int) ([]*db.GetOrderItemsByOrderRow, bool) {
	key := fmt.Sprintf(orderItemByIdCacheKey, orderID)
	result, found := cache.GetFromCache[[]*db.GetOrderItemsByOrderRow](ctx, o.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (o *orderItemQueryCache) SetCachedOrderItems(ctx context.Context, data []*db.GetOrderItemsByOrderRow) {
	if len(data) == 0 {
		return
	}

	key := fmt.Sprintf(orderItemByIdCacheKey, data[0].OrderID)
	cache.SetToCache(ctx, o.store, key, &data, ttlDefault)
}
