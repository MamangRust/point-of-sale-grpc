package cashier_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

const (
	cashierStatsMonthTotalSalesCacheKey = "cashier:stats:month:%d:year:%d"
	cashierStatsYearTotalSalesCacheKey  = "cashier:stats:year:%d"

	cashierStatsMonthSalesCacheKey = "cashier:stats:month:%d"
	cashierStatsYearSalesCacheKey  = "cashier:stats:year:%d"
)

type cashierStatsCache struct {
	store *cache.CacheStore
}

func NewCashierStatsCache(store *cache.CacheStore) *cashierStatsCache {
	return &cashierStatsCache{store: store}
}

func (s *cashierStatsCache) GetMonthlyTotalSalesCache(ctx context.Context, req *requests.MonthTotalSales) (*response.ApiResponseCashierMonthlyTotalSales, bool) {
	key := fmt.Sprintf(cashierStatsMonthTotalSalesCacheKey, req.Month, req.Year)
	result, found := cache.GetFromCache[*response.ApiResponseCashierMonthlyTotalSales](ctx, s.store, key)
	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierStatsCache) SetMonthlyTotalSalesCache(ctx context.Context, req *requests.MonthTotalSales, res *response.ApiResponseCashierMonthlyTotalSales) {
	if res == nil {
		return
	}
	key := fmt.Sprintf(cashierStatsMonthTotalSalesCacheKey, req.Month, req.Year)
	// Langsung simpan objek ApiResponse
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierStatsCache) GetYearlyTotalSalesCache(ctx context.Context, year int) (*response.ApiResponseCashierYearlyTotalSales, bool) {
	key := fmt.Sprintf(cashierStatsYearTotalSalesCacheKey, year)
	result, found := cache.GetFromCache[*response.ApiResponseCashierYearlyTotalSales](ctx, s.store, key)
	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierStatsCache) SetYearlyTotalSalesCache(ctx context.Context, year int, res *response.ApiResponseCashierYearlyTotalSales) {
	if res == nil {
		return
	}
	key := fmt.Sprintf(cashierStatsYearTotalSalesCacheKey, year)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierStatsCache) GetMonthlySalesCache(ctx context.Context, year int) (*response.ApiResponseCashierMonthSales, bool) {
	key := fmt.Sprintf(cashierStatsMonthSalesCacheKey, year)
	result, found := cache.GetFromCache[*response.ApiResponseCashierMonthSales](ctx, s.store, key)
	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *cashierStatsCache) SetMonthlySalesCache(ctx context.Context, year int, res *response.ApiResponseCashierMonthSales) {
	if res == nil {
		return
	}
	key := fmt.Sprintf(cashierStatsMonthSalesCacheKey, year)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *cashierStatsCache) GetYearlySalesCache(ctx context.Context, year int) (*response.ApiResponseCashierYearSales, bool) {
	key := fmt.Sprintf(cashierStatsYearSalesCacheKey, year)
	result, found := cache.GetFromCache[*response.ApiResponseCashierYearSales](ctx, s.store, key)
	if !found || result == nil {
		return nil, false
	}
	return result, true
}

func (s *cashierStatsCache) SetYearlySalesCache(ctx context.Context, year int, res *response.ApiResponseCashierYearSales) {
	if res == nil {
		return
	}
	key := fmt.Sprintf(cashierStatsYearSalesCacheKey, year)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}
