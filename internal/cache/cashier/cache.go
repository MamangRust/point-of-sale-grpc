package cashier_cache

import "pointofsale/internal/cache"

type cashierMencache struct {
	CashierQueryCache
	CashierCommandCache
	CashierStatsCache
	CashierStatsByIdCache
	CashierStatsByMerchantCache
}

type CashierMencache interface {
	CashierQueryCache
	CashierCommandCache
	CashierStatsCache
	CashierStatsByIdCache
	CashierStatsByMerchantCache
}

func NewCashierMencache(store *cache.CacheStore) CashierMencache {
	return &cashierMencache{
		CashierQueryCache:           NewCashierQueryCache(store),
		CashierCommandCache:         NewCashierCommandCache(store),
		CashierStatsCache:           NewCashierStatsCache(store),
		CashierStatsByIdCache:       NewCashierStatsByIdCache(store),
		CashierStatsByMerchantCache: NewCashierStatsByMerchantCache(store),
	}
}
