package role_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
)

type RoleQueryCache interface {
	SetCachedRoles(ctx context.Context, req *requests.FindAllRoles, data []*db.GetRolesRow, total *int)
	SetCachedRoleById(ctx context.Context, data *db.GetRoleRow)
	SetCachedRoleByUserId(ctx context.Context, userId int, data []*db.GetUserRolesRow)
	SetCachedRoleActive(ctx context.Context, req *requests.FindAllRoles, data []*db.GetActiveRolesRow, total *int)
	SetCachedRoleTrashed(ctx context.Context, req *requests.FindAllRoles, data []*db.GetTrashedRolesRow, total *int)

	GetCachedRoles(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetRolesRow, *int, bool)
	GetCachedRoleByUserId(ctx context.Context, userId int) ([]*db.GetUserRolesRow, bool)
	GetCachedRoleById(ctx context.Context, id int) (*db.GetRoleRow, bool)
	GetCachedRoleActive(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetActiveRolesRow, *int, bool)
	GetCachedRoleTrashed(ctx context.Context, req *requests.FindAllRoles) ([]*db.GetTrashedRolesRow, *int, bool)
}

type RoleCommandCache interface {
	DeleteCachedRole(ctx context.Context, id int)
}
