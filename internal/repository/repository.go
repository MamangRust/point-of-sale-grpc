package repository

import (
	"context"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
)

type Repositories struct {
	User         UserRepository
	Role         RoleRepository
	UserRole     UserRoleRepository
	RefreshToken RefreshTokenRepository
}

type Deps struct {
	DB           *db.Queries
	Ctx          context.Context
	MapperRecord *recordmapper.RecordMapper
}

func NewRepositories(deps Deps) *Repositories {
	return &Repositories{
		User:         NewUserRepository(deps.DB, deps.Ctx, deps.MapperRecord.UserRecordMapper),
		Role:         NewRoleRepository(deps.DB, deps.Ctx, deps.MapperRecord.RoleRecordMapper),
		UserRole:     NewUserRoleRepository(deps.DB, deps.Ctx, deps.MapperRecord.UserRoleRecordMapper),
		RefreshToken: NewRefreshTokenRepository(deps.DB, deps.Ctx, deps.MapperRecord.RefreshTokenRecordMapper),
	}
}
