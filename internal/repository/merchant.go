package repository

import (
	"context"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/merchant_errors"
)

type merchantRepository struct {
	db *db.Queries
}

func NewMerchantRepository(db *db.Queries) *merchantRepository {
	return &merchantRepository{
		db: db,
	}
}

func (r *merchantRepository) FindAllMerchants(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetMerchantsParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetMerchants(ctx, reqDb)

	if err != nil {
		return nil, merchant_errors.ErrFindAllMerchants
	}

	return res, nil
}

func (r *merchantRepository) FindByActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetMerchantsActiveParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetMerchantsActive(ctx, reqDb)

	if err != nil {
		return nil, merchant_errors.ErrFindByActive
	}

	return res, nil
}

func (r *merchantRepository) FindByTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetMerchantsTrashedParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetMerchantsTrashed(ctx, reqDb)

	if err != nil {
		return nil, merchant_errors.ErrFindByTrashed
	}

	return res, nil
}

func (r *merchantRepository) FindById(ctx context.Context, user_id int) (*db.GetMerchantByIDRow, error) {
	res, err := r.db.GetMerchantByID(ctx, int32(user_id))

	if err != nil {
		return nil, merchant_errors.ErrFindById
	}

	return res, nil
}

func (r *merchantRepository) CreateMerchant(ctx context.Context, request *requests.CreateMerchantRequest) (*db.CreateMerchantRow, error) {
	req := db.CreateMerchantParams{
		UserID:       int32(request.UserID),
		Name:         request.Name,
		Description:  &request.Description,
		Address:      &request.Address,
		ContactEmail: &request.ContactEmail,
		ContactPhone: &request.ContactPhone,
		Status:       request.Status,
	}

	merchant, err := r.db.CreateMerchant(ctx, req)

	if err != nil {
		return nil, merchant_errors.ErrCreateMerchant
	}

	return merchant, nil
}

func (r *merchantRepository) UpdateMerchant(ctx context.Context, request *requests.UpdateMerchantRequest) (*db.UpdateMerchantRow, error) {
	req := db.UpdateMerchantParams{
		MerchantID:   int32(*request.MerchantID),
		Name:         request.Name,
		Description:  &request.Description,
		Address:      &request.Address,
		ContactEmail: &request.ContactEmail,
		ContactPhone: &request.ContactPhone,
		Status:       request.Status,
	}

	res, err := r.db.UpdateMerchant(ctx, req)

	if err != nil {
		return nil, merchant_errors.ErrUpdateMerchant
	}

	return res, nil
}

func (r *merchantRepository) TrashedMerchant(ctx context.Context, merchant_id int) (*db.Merchant, error) {
	res, err := r.db.TrashMerchant(ctx, int32(merchant_id))

	if err != nil {
		return nil, merchant_errors.ErrTrashedMerchant
	}

	return res, nil
}

func (r *merchantRepository) RestoreMerchant(ctx context.Context, merchant_id int) (*db.Merchant, error) {
	res, err := r.db.RestoreMerchant(ctx, int32(merchant_id))

	if err != nil {
		return nil, merchant_errors.ErrRestoreMerchant
	}

	return res, nil
}

func (r *merchantRepository) DeleteMerchantPermanent(ctx context.Context, merchant_id int) (bool, error) {
	err := r.db.DeleteMerchantPermanently(ctx, int32(merchant_id))

	if err != nil {
		return false, merchant_errors.ErrDeleteMerchantPermanent
	}

	return true, nil
}

func (r *merchantRepository) RestoreAllMerchant(ctx context.Context) (bool, error) {
	err := r.db.RestoreAllMerchants(ctx)

	if err != nil {
		return false, merchant_errors.ErrRestoreAllMerchant
	}
	return true, nil
}

func (r *merchantRepository) DeleteAllMerchantPermanent(ctx context.Context) (bool, error) {
	err := r.db.DeleteAllPermanentMerchants(ctx)

	if err != nil {
		return false, merchant_errors.ErrDeleteAllMerchantPermanent
	}
	return true, nil
}
