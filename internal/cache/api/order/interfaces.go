package order_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type OrderStatsCache interface {
	GetMonthlyTotalRevenueCache(ctx context.Context, req *requests.MonthTotalRevenue) (*response.ApiResponseOrderMonthlyTotalRevenue, bool)
	SetMonthlyTotalRevenueCache(ctx context.Context, req *requests.MonthTotalRevenue, res *response.ApiResponseOrderMonthlyTotalRevenue)

	GetYearlyTotalRevenueCache(ctx context.Context, year int) (*response.ApiResponseOrderYearlyTotalRevenue, bool)
	SetYearlyTotalRevenueCache(ctx context.Context, year int, res *response.ApiResponseOrderYearlyTotalRevenue)

	GetMonthlyOrderCache(ctx context.Context, year int) (*response.ApiResponseOrderMonthly, bool)
	SetMonthlyOrderCache(ctx context.Context, year int, res *response.ApiResponseOrderMonthly)

	GetYearlyOrderCache(ctx context.Context, year int) (*response.ApiResponseOrderYearly, bool)
	SetYearlyOrderCache(ctx context.Context, year int, res *response.ApiResponseOrderYearly)
}

type OrderStatsByMerchantCache interface {
	GetMonthlyTotalRevenueByMerchantCache(ctx context.Context, req *requests.MonthTotalRevenueMerchant) (*response.ApiResponseOrderMonthlyTotalRevenue, bool)
	SetMonthlyTotalRevenueByMerchantCache(ctx context.Context, req *requests.MonthTotalRevenueMerchant, res *response.ApiResponseOrderMonthlyTotalRevenue)

	GetYearlyTotalRevenueByMerchantCache(ctx context.Context, req *requests.YearTotalRevenueMerchant) (*response.ApiResponseOrderYearlyTotalRevenue, bool)
	SetYearlyTotalRevenueByMerchantCache(ctx context.Context, req *requests.YearTotalRevenueMerchant, res *response.ApiResponseOrderYearlyTotalRevenue)

	GetMonthlyOrderByMerchantCache(ctx context.Context, req *requests.MonthOrderMerchant) (*response.ApiResponseOrderMonthly, bool)
	SetMonthlyOrderByMerchantCache(ctx context.Context, req *requests.MonthOrderMerchant, res *response.ApiResponseOrderMonthly)

	GetYearlyOrderByMerchantCache(ctx context.Context, req *requests.YearOrderMerchant) (*response.ApiResponseOrderYearly, bool)
	SetYearlyOrderByMerchantCache(ctx context.Context, req *requests.YearOrderMerchant, res *response.ApiResponseOrderYearly)
}

type OrderQueryCache interface {
	GetOrderAllCache(ctx context.Context, req *requests.FindAllOrders) (*response.ApiResponsePaginationOrder, bool)
	SetOrderAllCache(ctx context.Context, req *requests.FindAllOrders, res *response.ApiResponsePaginationOrder)

	GetCachedOrderCache(ctx context.Context, orderID int) (*response.ApiResponseOrder, bool)
	SetCachedOrderCache(ctx context.Context, res *response.ApiResponseOrder)

	GetCachedOrderMerchant(ctx context.Context, req *requests.FindAllOrderMerchant) (*response.ApiResponsePaginationOrder, bool)
	SetCachedOrderMerchant(ctx context.Context, req *requests.FindAllOrderMerchant, res *response.ApiResponsePaginationOrder)

	GetOrderActiveCache(ctx context.Context, req *requests.FindAllOrders) (*response.ApiResponsePaginationOrderDeleteAt, bool)
	SetOrderActiveCache(ctx context.Context, req *requests.FindAllOrders, res *response.ApiResponsePaginationOrderDeleteAt)

	GetOrderTrashedCache(ctx context.Context, req *requests.FindAllOrders) (*response.ApiResponsePaginationOrderDeleteAt, bool)
	SetOrderTrashedCache(ctx context.Context, req *requests.FindAllOrders, res *response.ApiResponsePaginationOrderDeleteAt)
}

type OrderCommandCache interface {
	DeleteOrderCache(ctx context.Context, id int)
}
