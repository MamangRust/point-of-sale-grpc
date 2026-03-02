package cashier_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type CashierQueryCache interface {
	GetCachedCashiersCache(ctx context.Context, req *requests.FindAllCashiers) (*response.ApiResponsePaginationCashier, bool)
	SetCachedCashiersCache(ctx context.Context, req *requests.FindAllCashiers, res *response.ApiResponsePaginationCashier)

	GetCachedCashier(ctx context.Context, cashierID int) (*response.ApiResponseCashier, bool)
	SetCachedCashier(ctx context.Context, res *response.ApiResponseCashier)

	GetCachedCashiersActive(ctx context.Context, req *requests.FindAllCashiers) (*response.ApiResponsePaginationCashierDeleteAt, bool)
	SetCachedCashiersActive(ctx context.Context, req *requests.FindAllCashiers, res *response.ApiResponsePaginationCashierDeleteAt)

	GetCachedCashiersTrashed(ctx context.Context, req *requests.FindAllCashiers) (*response.ApiResponsePaginationCashierDeleteAt, bool)
	SetCachedCashiersTrashed(ctx context.Context, req *requests.FindAllCashiers, res *response.ApiResponsePaginationCashierDeleteAt)

	GetCachedCashiersByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant) (*response.ApiResponsePaginationCashier, bool)
	SetCachedCashiersByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant, res *response.ApiResponsePaginationCashier)
}

type CashierCommandCache interface {
	DeleteCashierCache(ctx context.Context, id int)
}

type CashierStatsCache interface {
	GetMonthlyTotalSalesCache(ctx context.Context, req *requests.MonthTotalSales) (*response.ApiResponseCashierMonthlyTotalSales, bool)
	SetMonthlyTotalSalesCache(ctx context.Context, req *requests.MonthTotalSales, res *response.ApiResponseCashierMonthlyTotalSales)

	GetYearlyTotalSalesCache(ctx context.Context, year int) (*response.ApiResponseCashierYearlyTotalSales, bool)
	SetYearlyTotalSalesCache(ctx context.Context, year int, res *response.ApiResponseCashierYearlyTotalSales)

	GetMonthlySalesCache(ctx context.Context, year int) (*response.ApiResponseCashierMonthSales, bool)
	SetMonthlySalesCache(ctx context.Context, year int, res *response.ApiResponseCashierMonthSales)

	GetYearlySalesCache(ctx context.Context, year int) (*response.ApiResponseCashierYearSales, bool)
	SetYearlySalesCache(ctx context.Context, year int, res *response.ApiResponseCashierYearSales)
}

type CashierStatsByIdCache interface {
	GetMonthlyTotalSalesByIdCache(ctx context.Context, req *requests.MonthTotalSalesCashier) (*response.ApiResponseCashierMonthlyTotalSales, bool)
	SetMonthlyTotalSalesByIdCache(ctx context.Context, req *requests.MonthTotalSalesCashier, res *response.ApiResponseCashierMonthlyTotalSales)

	GetYearlyTotalSalesByIdCache(ctx context.Context, req *requests.YearTotalSalesCashier) (*response.ApiResponseCashierYearlyTotalSales, bool)
	SetYearlyTotalSalesByIdCache(ctx context.Context, req *requests.YearTotalSalesCashier, res *response.ApiResponseCashierYearlyTotalSales)

	GetMonthlyCashierByIdCache(ctx context.Context, req *requests.MonthCashierId) (*response.ApiResponseCashierMonthSales, bool)
	SetMonthlyCashierByIdCache(ctx context.Context, req *requests.MonthCashierId, res *response.ApiResponseCashierMonthSales)

	GetYearlyCashierByIdCache(ctx context.Context, req *requests.YearCashierId) (*response.ApiResponseCashierYearSales, bool)
	SetYearlyCashierByIdCache(ctx context.Context, req *requests.YearCashierId, res *response.ApiResponseCashierYearSales)
}

type CashierStatsByMerchantCache interface {
	GetMonthlyTotalSalesByMerchantCache(ctx context.Context, req *requests.MonthTotalSalesMerchant) (*response.ApiResponseCashierMonthlyTotalSales, bool)
	SetMonthlyTotalSalesByMerchantCache(ctx context.Context, req *requests.MonthTotalSalesMerchant, res *response.ApiResponseCashierMonthlyTotalSales)

	GetYearlyTotalSalesByMerchantCache(ctx context.Context, req *requests.YearTotalSalesMerchant) (*response.ApiResponseCashierYearlyTotalSales, bool)
	SetYearlyTotalSalesByMerchantCache(ctx context.Context, req *requests.YearTotalSalesMerchant, res *response.ApiResponseCashierYearlyTotalSales)

	GetMonthlyCashierByMerchantCache(ctx context.Context, req *requests.MonthCashierMerchant) (*response.ApiResponseCashierMonthSales, bool)
	SetMonthlyCashierByMerchantCache(ctx context.Context, req *requests.MonthCashierMerchant, res *response.ApiResponseCashierMonthSales)

	GetYearlyCashierByMerchantCache(ctx context.Context, req *requests.YearCashierMerchant) (*response.ApiResponseCashierYearSales, bool)
	SetYearlyCashierByMerchantCache(ctx context.Context, req *requests.YearCashierMerchant, res *response.ApiResponseCashierYearSales)
}
