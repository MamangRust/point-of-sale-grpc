package order_cache

import "pointofsale/internal/cache"

type OrderMencache interface {
	OrderQueryCache
	OrderCommandCache
	OrderStatsCache
	OrderStatsByMerchantCache
}

type orderMencache struct {
	OrderQueryCache
	OrderCommandCache
	OrderStatsCache
	OrderStatsByMerchantCache
}

func NewOrderMencache(store *cache.CacheStore) OrderMencache {
	return &orderMencache{
		OrderQueryCache:           NewOrderQueryCache(store),
		OrderCommandCache:         NewOrderCommandCache(store),
		OrderStatsCache:           NewOrderStatsCache(store),
		OrderStatsByMerchantCache: NewOrderStatsByMerchantCache(store),
	}
}
