package merchant_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type MerchantQueryCache interface {
	GetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchant, bool)
	SetCachedMerchants(ctx context.Context, req *requests.FindAllMerchants, res *response.ApiResponsePaginationMerchant)

	GetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchantDeleteAt, bool)
	SetCachedMerchantActive(ctx context.Context, req *requests.FindAllMerchants, res *response.ApiResponsePaginationMerchantDeleteAt)

	GetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants) (*response.ApiResponsePaginationMerchantDeleteAt, bool)
	SetCachedMerchantTrashed(ctx context.Context, req *requests.FindAllMerchants, res *response.ApiResponsePaginationMerchantDeleteAt)

	GetCachedMerchant(ctx context.Context, id int) (*response.ApiResponseMerchant, bool)
	SetCachedMerchant(ctx context.Context, res *response.ApiResponseMerchant)

	GetCachedMerchantsByUserId(ctx context.Context, id int) (*response.ApiResponsesMerchant, bool)
	SetCachedMerchantsByUserId(ctx context.Context, userId int, res *response.ApiResponsesMerchant)
}

type MerchantCommandCache interface {
	DeleteCachedMerchant(ctx context.Context, id int)
}
