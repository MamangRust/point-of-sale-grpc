package role_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type RoleQueryCache interface {
	SetCachedRoles(ctx context.Context, req *requests.FindAllRoles, res *response.ApiResponsePaginationRole)
	SetCachedRoleById(ctx context.Context, res *response.ApiResponseRole)
	SetCachedRoleByUserId(ctx context.Context, userId int, res *response.ApiResponsesRole)
	SetCachedRoleActive(ctx context.Context, req *requests.FindAllRoles, res *response.ApiResponsePaginationRoleDeleteAt)
	SetCachedRoleTrashed(ctx context.Context, req *requests.FindAllRoles, res *response.ApiResponsePaginationRoleDeleteAt)

	GetCachedRoles(ctx context.Context, req *requests.FindAllRoles) (*response.ApiResponsePaginationRole, bool)
	GetCachedRoleByUserId(ctx context.Context, userId int) (*response.ApiResponsesRole, bool)
	GetCachedRoleById(ctx context.Context, id int) (*response.ApiResponseRole, bool)
	GetCachedRoleActive(ctx context.Context, req *requests.FindAllRoles) (*response.ApiResponsePaginationRoleDeleteAt, bool)
	GetCachedRoleTrashed(ctx context.Context, req *requests.FindAllRoles) (*response.ApiResponsePaginationRoleDeleteAt, bool)
}

type RoleCommandCache interface {
	DeleteCachedRole(ctx context.Context, id int)
}
