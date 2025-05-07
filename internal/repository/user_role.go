package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
)

type userRoleRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.UserRoleRecordMapping
}

func NewUserRoleRepository(db *db.Queries, ctx context.Context, mapping recordmapper.UserRoleRecordMapping) *userRoleRepository {
	return &userRoleRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *userRoleRepository) AssignRoleToUser(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error) {
	res, err := r.db.AssignRoleToUser(r.ctx, db.AssignRoleToUserParams{
		UserID: int32(req.UserId),
		RoleID: int32(req.RoleId),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user ID %d or role ID %d not found", req.UserId, req.RoleId)
		}
		return nil, fmt.Errorf("failed to assign role to user: %w", err)
	}

	return r.mapping.ToUserRoleRecord(res), nil
}

func (r *userRoleRepository) RemoveRoleFromUser(req *requests.RemoveUserRoleRequest) error {
	err := r.db.RemoveRoleFromUser(r.ctx, db.RemoveRoleFromUserParams{
		UserID: int32(req.UserId),
		RoleID: int32(req.RoleId),
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no role assignment found for user ID %d and role ID %d", req.UserId, req.RoleId)
		}
		return fmt.Errorf("failed to remove role from user: %w", err)
	}

	return nil
}
