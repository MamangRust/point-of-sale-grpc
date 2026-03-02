package user_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type UserQueryCache interface {
	GetCachedUsersCache(ctx context.Context, req *requests.FindAllUsers) (*response.ApiResponsePaginationUser, bool)
	SetCachedUsersCache(ctx context.Context, req *requests.FindAllUsers, res *response.ApiResponsePaginationUser)

	GetCachedUserActiveCache(ctx context.Context, req *requests.FindAllUsers) (*response.ApiResponsePaginationUserDeleteAt, bool)
	SetCachedUserActiveCache(ctx context.Context, req *requests.FindAllUsers, res *response.ApiResponsePaginationUserDeleteAt)

	GetCachedUserTrashedCache(ctx context.Context, req *requests.FindAllUsers) (*response.ApiResponsePaginationUserDeleteAt, bool)
	SetCachedUserTrashedCache(ctx context.Context, req *requests.FindAllUsers, res *response.ApiResponsePaginationUserDeleteAt)

	GetCachedUserCache(ctx context.Context, id int) (*response.ApiResponseUser, bool)
	SetCachedUserCache(ctx context.Context, res *response.ApiResponseUser)
}

type UserCommandCache interface {
	DeleteUserCache(ctx context.Context, id int)
}
