package category_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type CategoryQueryCache interface {
	GetCachedCategoriesCache(ctx context.Context, req *requests.FindAllCategory) (*response.ApiResponsePaginationCategory, bool)
	SetCachedCategoriesCache(ctx context.Context, req *requests.FindAllCategory, res *response.ApiResponsePaginationCategory)

	GetCachedCategoryActiveCache(ctx context.Context, req *requests.FindAllCategory) (*response.ApiResponsePaginationCategoryDeleteAt, bool)
	SetCachedCategoryActiveCache(ctx context.Context, req *requests.FindAllCategory, res *response.ApiResponsePaginationCategoryDeleteAt)

	GetCachedCategoryTrashedCache(ctx context.Context, req *requests.FindAllCategory) (*response.ApiResponsePaginationCategoryDeleteAt, bool)
	SetCachedCategoryTrashedCache(ctx context.Context, req *requests.FindAllCategory, res *response.ApiResponsePaginationCategoryDeleteAt)

	GetCachedCategoryCache(ctx context.Context, id int) (*response.ApiResponseCategory, bool)
	SetCachedCategoryCache(ctx context.Context, res *response.ApiResponseCategory)
}

type CategoryCommandCache interface {
	DeleteCachedCategoryCache(ctx context.Context, id int)
}

type CategoryStatsCache interface {
	GetCachedMonthTotalPriceCache(ctx context.Context, req *requests.MonthTotalPrice) (*response.ApiResponseCategoryMonthlyTotalPrice, bool)
	SetCachedMonthTotalPriceCache(ctx context.Context, req *requests.MonthTotalPrice, res *response.ApiResponseCategoryMonthlyTotalPrice)

	GetCachedYearTotalPriceCache(ctx context.Context, year int) (*response.ApiResponseCategoryYearlyTotalPrice, bool)
	SetCachedYearTotalPriceCache(ctx context.Context, year int, res *response.ApiResponseCategoryYearlyTotalPrice)

	GetCachedMonthPriceCache(ctx context.Context, year int) (*response.ApiResponseCategoryMonthPrice, bool)
	SetCachedMonthPriceCache(ctx context.Context, year int, res *response.ApiResponseCategoryMonthPrice)

	GetCachedYearPriceCache(ctx context.Context, year int) (*response.ApiResponseCategoryYearPrice, bool)
	SetCachedYearPriceCache(ctx context.Context, year int, res *response.ApiResponseCategoryYearPrice)
}

type CategoryStatsByIdCache interface {
	GetCachedMonthTotalPriceByIdCache(ctx context.Context, req *requests.MonthTotalPriceCategory) (*response.ApiResponseCategoryMonthlyTotalPrice, bool)
	SetCachedMonthTotalPriceByIdCache(ctx context.Context, req *requests.MonthTotalPriceCategory, res *response.ApiResponseCategoryMonthlyTotalPrice)

	GetCachedYearTotalPriceByIdCache(ctx context.Context, req *requests.YearTotalPriceCategory) (*response.ApiResponseCategoryYearlyTotalPrice, bool)
	SetCachedYearTotalPriceByIdCache(ctx context.Context, req *requests.YearTotalPriceCategory, res *response.ApiResponseCategoryYearlyTotalPrice)

	GetCachedMonthPriceByIdCache(ctx context.Context, req *requests.MonthPriceId) (*response.ApiResponseCategoryMonthPrice, bool)
	SetCachedMonthPriceByIdCache(ctx context.Context, req *requests.MonthPriceId, res *response.ApiResponseCategoryMonthPrice)

	GetCachedYearPriceByIdCache(ctx context.Context, req *requests.YearPriceId) (*response.ApiResponseCategoryYearPrice, bool)
	SetCachedYearPriceByIdCache(ctx context.Context, req *requests.YearPriceId, res *response.ApiResponseCategoryYearPrice)
}

type CategoryStatsByMerchantCache interface {
	GetCachedMonthTotalPriceByMerchantCache(ctx context.Context, req *requests.MonthTotalPriceMerchant) (*response.ApiResponseCategoryMonthlyTotalPrice, bool)
	SetCachedMonthTotalPriceByMerchantCache(ctx context.Context, req *requests.MonthTotalPriceMerchant, res *response.ApiResponseCategoryMonthlyTotalPrice)

	GetCachedYearTotalPriceByMerchantCache(ctx context.Context, req *requests.YearTotalPriceMerchant) (*response.ApiResponseCategoryYearlyTotalPrice, bool)
	SetCachedYearTotalPriceByMerchantCache(ctx context.Context, req *requests.YearTotalPriceMerchant, res *response.ApiResponseCategoryYearlyTotalPrice)

	GetCachedMonthPriceByMerchantCache(ctx context.Context, req *requests.MonthPriceMerchant) (*response.ApiResponseCategoryMonthPrice, bool)
	SetCachedMonthPriceByMerchantCache(ctx context.Context, req *requests.MonthPriceMerchant, res *response.ApiResponseCategoryMonthPrice)

	GetCachedYearPriceByMerchantCache(ctx context.Context, req *requests.YearPriceMerchant) (*response.ApiResponseCategoryYearPrice, bool)
	SetCachedYearPriceByMerchantCache(ctx context.Context, req *requests.YearPriceMerchant, res *response.ApiResponseCategoryYearPrice)
}
