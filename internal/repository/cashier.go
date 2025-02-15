package repository

import (
	"context"
	"errors"
	"fmt"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
)

type cashierRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.CashierRecordMapping
}

func NewCashierRepository(db *db.Queries, ctx context.Context, mapping recordmapper.CashierRecordMapping) *cashierRepository {
	return &cashierRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *cashierRepository) FindAllCashiers(search string, page, pageSize int) ([]*record.CashierRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCashiersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiers(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cashiers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersRecordPagination(res), totalCount, nil
}

func (r *cashierRepository) FindById(user_id int) (*record.CashierRecord, error) {
	fmt.Printf("Searching for cashier with ID: %d\n", user_id)
	res, err := r.db.GetCashierByID(r.ctx, int32(user_id))

	if err != nil {
		fmt.Printf("Error fetching cashier: %v\n", err)

		return nil, fmt.Errorf("failed to find cashier: %w", err)
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) FindByActive(search string, page, pageSize int) ([]*record.CashierRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCashiersActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiersActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cashier: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersRecordActivePagination(res), totalCount, nil
}

func (r *cashierRepository) FindByTrashed(search string, page, pageSize int) ([]*record.CashierRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCashiersTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiersTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cashiers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersRecordTrashedPagination(res), totalCount, nil
}

func (r *cashierRepository) FindByMerchant(merchant_id int, search string, page, pageSize int) ([]*record.CashierRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCashiersByMerchantParams{
		MerchantID: int32(merchant_id),
		Column2:    search,
		Limit:      int32(pageSize),
		Offset:     int32(offset),
	}

	res, err := r.db.GetCashiersByMerchant(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cashiers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersMerchantRecordPagination(res), totalCount, nil
}

func (r *cashierRepository) CreateCashier(request *requests.CreateCashierRequest) (*record.CashierRecord, error) {
	req := db.CreateCashierParams{
		MerchantID: int32(request.MerchantID),
		UserID:     int32(request.UserID),
		Name:       request.Name,
	}

	cashier, err := r.db.CreateCashier(r.ctx, req)

	if err != nil {
		return nil, errors.New("failed create user")
	}

	return r.mapping.ToCashierRecord(cashier), nil
}

func (r *cashierRepository) UpdateCashier(request *requests.UpdateCashierRequest) (*record.CashierRecord, error) {
	req := db.UpdateCashierParams{
		CashierID: int32(request.CashierID),
		Name:      request.Name,
	}

	res, err := r.db.UpdateCashier(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) TrashedCashier(cashier_id int) (*record.CashierRecord, error) {
	res, err := r.db.TrashCashier(r.ctx, int32(cashier_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash user: %w", err)
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) RestoreCashier(cashier_id int) (*record.CashierRecord, error) {
	res, err := r.db.RestoreCashier(r.ctx, int32(cashier_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore cashier: %w", err)
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) DeleteCashierPermanent(cashier_id int) (bool, error) {
	err := r.db.DeleteCashierPermanently(r.ctx, int32(cashier_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete cashier: %w", err)
	}

	return true, nil
}

func (r *cashierRepository) RestoreAllCashier() (bool, error) {
	err := r.db.RestoreAllCashiers(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all cashier: %w", err)
	}
	return true, nil
}

func (r *cashierRepository) DeleteAllCashierPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentCashiers(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all cashier permanently: %w", err)
	}
	return true, nil
}
