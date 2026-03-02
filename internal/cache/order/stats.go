package order_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
)

const (
	monthlyTotalRevenueCacheKey = "order:monthly:totalRevenue:month:%d:year:%d"
	yearlyTotalRevenueCacheKey  = "order:yearly:totalRevenue:year:%d"

	monthlyOrderCacheKey = "order:monthly:order:month:%d"
	yearlyOrderCacheKey  = "order:yearly:order:year:%d"
)

type orderStatsCache struct {
	store *cache.CacheStore
}

func NewOrderStatsCache(store *cache.CacheStore) *orderStatsCache {
	return &orderStatsCache{store: store}
}

func (s *orderStatsCache) GetMonthlyTotalRevenueCache(ctx context.Context, req *requests.MonthTotalRevenue) ([]*db.GetMonthlyTotalRevenueRow, bool) {
	key := fmt.Sprintf(monthlyTotalRevenueCacheKey, req.Month, req.Year)

	result, found := cache.GetFromCache[[]*db.GetMonthlyTotalRevenueRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	// Perbaikan: Mengembalikan 'result' langsung, bukan '*result'
	return result, true
}

func (s *orderStatsCache) SetMonthlyTotalRevenueCache(ctx context.Context, req *requests.MonthTotalRevenue, res []*db.GetMonthlyTotalRevenueRow) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(monthlyTotalRevenueCacheKey, req.Month, req.Year)

	// Perbaikan: Menggunakan 'cache.SetToCache' untuk konsistensi
	cache.SetToCache(ctx, s.store, key, &res, ttlDefault)
}

func (s *orderStatsCache) GetYearlyTotalRevenueCache(ctx context.Context, year int) ([]*db.GetYearlyTotalRevenueRow, bool) {
	key := fmt.Sprintf(yearlyTotalRevenueCacheKey, year)

	result, found := cache.GetFromCache[[]*db.GetYearlyTotalRevenueRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetYearlyTotalRevenueCache(ctx context.Context, year int, res []*db.GetYearlyTotalRevenueRow) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(yearlyTotalRevenueCacheKey, year)
	cache.SetToCache(ctx, s.store, key, &res, ttlDefault)
}

func (s *orderStatsCache) GetMonthlyOrderCache(ctx context.Context, year int) ([]*db.GetMonthlyOrderRow, bool) {
	key := fmt.Sprintf(monthlyOrderCacheKey, year)

	result, found := cache.GetFromCache[[]*db.GetMonthlyOrderRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetMonthlyOrderCache(ctx context.Context, year int, res []*db.GetMonthlyOrderRow) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(monthlyOrderCacheKey, year)
	cache.SetToCache(ctx, s.store, key, &res, ttlDefault)
}

func (s *orderStatsCache) GetYearlyOrderCache(ctx context.Context, year int) ([]*db.GetYearlyOrderRow, bool) {
	key := fmt.Sprintf(yearlyOrderCacheKey, year)

	result, found := cache.GetFromCache[[]*db.GetYearlyOrderRow](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *orderStatsCache) SetYearlyOrderCache(ctx context.Context, year int, res []*db.GetYearlyOrderRow) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(yearlyOrderCacheKey, year)
	cache.SetToCache(ctx, s.store, key, &res, ttlDefault)
}
