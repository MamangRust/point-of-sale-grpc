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

type roleRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.RoleRecordMapping
}

func NewRoleRepository(db *db.Queries, ctx context.Context, mapping recordmapper.RoleRecordMapping) *roleRepository {
	return &roleRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *roleRepository) FindAllRoles(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetRolesParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetRoles(r.ctx, reqDb)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch roles: invalid pagination (page %d, size %d) or search query '%s'", req.Page, req.PageSize, req.Search)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToRolesRecordAll(res), &totalCount, nil
}

func (r *roleRepository) FindById(id int) (*record.RoleRecord, error) {
	res, err := r.db.GetRole(r.ctx, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role not found with ID: %d", id)
		}
		return nil, fmt.Errorf("failed to retrieve role with ID %d: %w", id, err)
	}

	return r.mapping.ToRoleRecord(res), nil
}

func (r *roleRepository) FindByName(name string) (*record.RoleRecord, error) {
	res, err := r.db.GetRoleByName(r.ctx, name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role not found with name: '%s'", name)
		}
		return nil, fmt.Errorf("failed to retrieve role with name '%s': %w", name, err)
	}

	return r.mapping.ToRoleRecord(res), nil
}

func (r *roleRepository) FindByUserId(user_id int) ([]*record.RoleRecord, error) {
	res, err := r.db.GetUserRoles(r.ctx, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no roles found for user ID: %d", user_id)
		}
		return nil, fmt.Errorf("failed to retrieve roles for user ID %d: %w", user_id, err)
	}

	return r.mapping.ToRolesRecord(res), nil
}

func (r *roleRepository) FindByActiveRole(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetActiveRolesParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetActiveRoles(r.ctx, reqDb)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch active roles: invalid parameters (page %d, size %d, search '%s')", req.Page, req.PageSize, req.Search)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToRolesRecordActive(res), &totalCount, nil
}

func (r *roleRepository) FindByTrashedRole(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetTrashedRolesParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTrashedRoles(r.ctx, reqDb)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch trashed roles: invalid parameters (page %d, size %d, search '%s')", req.Page, req.PageSize, req.Search)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToRolesRecordTrashed(res), &totalCount, nil
}

func (r *roleRepository) CreateRole(req *requests.CreateRoleRequest) (*record.RoleRecord, error) {
	res, err := r.db.CreateRole(r.ctx, req.Name)

	if err != nil {
		return nil, fmt.Errorf("failed to create role: invalid name '%s' or duplicate role", req.Name)
	}

	return r.mapping.ToRoleRecord(res), nil
}

func (r *roleRepository) UpdateRole(req *requests.UpdateRoleRequest) (*record.RoleRecord, error) {
	res, err := r.db.UpdateRole(r.ctx, db.UpdateRoleParams{
		RoleID:   int32(*req.ID),
		RoleName: req.Name,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update role ID %d: role not found or invalid data", req.ID)
	}

	return r.mapping.ToRoleRecord(res), nil
}

func (r *roleRepository) TrashedRole(id int) (*record.RoleRecord, error) {
	res, err := r.db.TrashRole(r.ctx, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role ID %d not found or already trashed", id)
		}
		return nil, fmt.Errorf("failed to trash role ID %d: %w", id, err)
	}

	return r.mapping.ToRoleRecord(res), nil
}

func (r *roleRepository) RestoreRole(id int) (*record.RoleRecord, error) {
	res, err := r.db.RestoreRole(r.ctx, int32(id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("role ID %d not found in trash", id)
		}
		return nil, fmt.Errorf("failed to restore role ID %d: %w", id, err)
	}

	return r.mapping.ToRoleRecord(res), nil
}

func (r *roleRepository) DeleteRolePermanent(role_id int) (bool, error) {
	err := r.db.DeletePermanentRole(r.ctx, int32(role_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("role ID %d not found or already deleted", role_id)
		}
		return false, fmt.Errorf("failed to permanently delete role ID %d: %w", role_id, err)
	}

	return true, nil
}

func (r *roleRepository) RestoreAllRole() (bool, error) {
	err := r.db.RestoreAllRoles(r.ctx)

	if err != nil {
		return false, fmt.Errorf("no trashed roles available to restore")
	}

	return true, nil
}

func (r *roleRepository) DeleteAllRolePermanent() (bool, error) {
	err := r.db.DeleteAllPermanentRoles(r.ctx)

	if err != nil {
		return false, fmt.Errorf("cannot permanently delete all roles: operation disabled for system protection")
	}

	return true, nil
}
