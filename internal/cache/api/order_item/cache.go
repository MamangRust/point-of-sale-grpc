package orderitem_cache

import "pointofsale/internal/cache"

type OrderItemCache interface {
	OrderItemQueryCache
}

type orderItemCache struct {
	OrderItemQueryCache
}

func NewOrderItemCache(store *cache.CacheStore) OrderItemCache {
	return &orderItemCache{
		OrderItemQueryCache: NewOrderItemQueryCache(store),
	}
}
