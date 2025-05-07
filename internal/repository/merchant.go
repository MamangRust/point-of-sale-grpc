package repository

import (
	"context"
	"database/sql"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/merchant_errors"
)

type merchantRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.MerchantRecordMapping
}

func NewMerchantRepository(db *db.Queries, ctx context.Context, mapping recordmapper.MerchantRecordMapping) *merchantRepository {
	return &merchantRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *merchantRepository) FindAllMerchants(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetMerchantsParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetMerchants(r.ctx, reqDb)

	if err != nil {
		return nil, nil, merchant_errors.ErrFindAllMerchants
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToMerchantsRecordPagination(res), &totalCount, nil
}

func (r *merchantRepository) FindByActive(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetMerchantsActiveParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetMerchantsActive(r.ctx, reqDb)

	if err != nil {
		return nil, nil, merchant_errors.ErrFindByActive
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToMerchantsRecordActivePagination(res), &totalCount, nil
}

func (r *merchantRepository) FindByTrashed(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetMerchantsTrashedParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetMerchantsTrashed(r.ctx, reqDb)

	if err != nil {
		return nil, nil, merchant_errors.ErrFindByTrashed
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToMerchantsRecordTrashedPagination(res), &totalCount, nil
}

func (r *merchantRepository) FindById(user_id int) (*record.MerchantRecord, error) {
	res, err := r.db.GetMerchantByID(r.ctx, int32(user_id))

	if err != nil {
		return nil, merchant_errors.ErrFindById
	}

	return r.mapping.ToMerchantRecord(res), nil
}

func (r *merchantRepository) CreateMerchant(request *requests.CreateMerchantRequest) (*record.MerchantRecord, error) {
	req := db.CreateMerchantParams{
		UserID:       int32(request.UserID),
		Name:         request.Name,
		Description:  sql.NullString{String: request.Description, Valid: request.Description != ""},
		Address:      sql.NullString{String: request.Address, Valid: request.Address != ""},
		ContactEmail: sql.NullString{String: request.ContactEmail, Valid: request.ContactEmail != ""},
		ContactPhone: sql.NullString{String: request.ContactPhone, Valid: request.ContactPhone != ""},
		Status:       request.Status,
	}

	merchant, err := r.db.CreateMerchant(r.ctx, req)

	if err != nil {
		return nil, merchant_errors.ErrCreateMerchant
	}

	return r.mapping.ToMerchantRecord(merchant), nil
}

func (r *merchantRepository) UpdateMerchant(request *requests.UpdateMerchantRequest) (*record.MerchantRecord, error) {
	req := db.UpdateMerchantParams{
		MerchantID:   int32(*request.MerchantID),
		Name:         request.Name,
		Description:  sql.NullString{String: request.Description, Valid: request.Description != ""},
		Address:      sql.NullString{String: request.Address, Valid: request.Address != ""},
		ContactEmail: sql.NullString{String: request.ContactEmail, Valid: request.ContactEmail != ""},
		ContactPhone: sql.NullString{String: request.ContactPhone, Valid: request.ContactPhone != ""},
		Status:       request.Status,
	}

	res, err := r.db.UpdateMerchant(r.ctx, req)

	if err != nil {
		return nil, merchant_errors.ErrUpdateMerchant
	}

	return r.mapping.ToMerchantRecord(res), nil
}

func (r *merchantRepository) TrashedMerchant(merchant_id int) (*record.MerchantRecord, error) {
	res, err := r.db.TrashMerchant(r.ctx, int32(merchant_id))

	if err != nil {
		return nil, merchant_errors.ErrTrashedMerchant
	}

	return r.mapping.ToMerchantRecord(res), nil
}

func (r *merchantRepository) RestoreMerchant(merchant_id int) (*record.MerchantRecord, error) {
	res, err := r.db.RestoreMerchant(r.ctx, int32(merchant_id))

	if err != nil {
		return nil, merchant_errors.ErrRestoreMerchant
	}

	return r.mapping.ToMerchantRecord(res), nil
}

func (r *merchantRepository) DeleteMerchantPermanent(merchant_id int) (bool, error) {
	err := r.db.DeleteMerchantPermanently(r.ctx, int32(merchant_id))

	if err != nil {
		return false, merchant_errors.ErrDeleteMerchantPermanent
	}

	return true, nil
}

func (r *merchantRepository) RestoreAllMerchant() (bool, error) {
	err := r.db.RestoreAllMerchants(r.ctx)

	if err != nil {
		return false, merchant_errors.ErrRestoreAllMerchant
	}
	return true, nil
}

func (r *merchantRepository) DeleteAllMerchantPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentMerchants(r.ctx)

	if err != nil {
		return false, merchant_errors.ErrDeleteAllMerchantPermanent
	}
	return true, nil
}
