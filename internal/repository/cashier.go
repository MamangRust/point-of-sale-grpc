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
	"time"
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
	res, err := r.db.GetCashierByID(r.ctx, int32(user_id))

	if err != nil {
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

func (r *cashierRepository) GetMonthlyTotalSales(year int, month int) ([]*record.CashierRecordMonthTotalSales, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)

	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	params := db.GetMonthlyTotalSalesCashierParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
	}

	res, err := r.db.GetMonthlyTotalSalesCashier(r.ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total sales cashier: %w", err)
	}

	return r.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (r *cashierRepository) GetYearlyTotalSales(year int) ([]*record.CashierRecordYearTotalSales, error) {

	res, err := r.db.GetYearlyTotalSalesCashier(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total sales cashier: %w", err)
	}

	so := r.mapping.ToCashierYearlyTotalSales(res)

	return so, nil
}

func (r *cashierRepository) GetMonthlyTotalSalesById(year int, month int, cashier_id int) ([]*record.CashierRecordMonthTotalSales, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)

	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalSalesById(r.ctx, db.GetMonthlyTotalSalesByIdParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		CashierID:   int32(cashier_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total sales cashier: %w", err)
	}

	so := r.mapping.ToCashierMonthlyTotalSalesById(res)

	return so, nil
}

func (r *cashierRepository) GetYearlyTotalSalesById(year int, cashier_id int) ([]*record.CashierRecordYearTotalSales, error) {
	res, err := r.db.GetYearlyTotalSalesById(r.ctx, db.GetYearlyTotalSalesByIdParams{
		Column1:   int32(year),
		CashierID: int32(cashier_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total sales cashier: %w", err)
	}

	so := r.mapping.ToCashierYearlyTotalSalesById(res)

	return so, nil
}

func (r *cashierRepository) GetMonthlyTotalSalesByMerchant(year int, month int, merchant_id int) ([]*record.CashierRecordMonthTotalSales, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalSalesByMerchant(r.ctx, db.GetMonthlyTotalSalesByMerchantParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		MerchantID:  int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total sales cashier: %w", err)
	}

	so := r.mapping.ToCashierMonthlyTotalSalesByMerchant(res)

	return so, nil
}

func (r *cashierRepository) GetYearlyTotalSalesByMerchant(year int, merchant_id int) ([]*record.CashierRecordYearTotalSales, error) {
	res, err := r.db.GetYearlyTotalSalesByMerchant(r.ctx, db.GetYearlyTotalSalesByMerchantParams{
		Column1:    int32(year),
		MerchantID: int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total sales cashier: %w", err)
	}

	so := r.mapping.ToCashierYearlyTotalSalesByMerchant(res)

	return so, nil
}

func (r *cashierRepository) GetMonthyCashier(year int) ([]*record.CashierRecordMonthSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashier(r.ctx, yearStart)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve monthly sales data for year %d: %w", year, err)
	}

	return r.mapping.ToCashierMonthlySales(res), nil

}

func (r *cashierRepository) GetYearlyCashier(year int) ([]*record.CashierRecordYearSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashier(r.ctx, yearStart)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve yearly sales data for year %d: %w", year, err)
	}

	return r.mapping.ToCashierYearlySales(res), nil
}

func (r *cashierRepository) GetMonthlyCashierByMerchant(year int, merchant_id int) ([]*record.CashierRecordMonthSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashierByMerchant(r.ctx, db.GetMonthlyCashierByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve monthly sales data for merchant %d in year %d: %w", merchant_id, year, err)
	}

	return r.mapping.ToCashierMonthlySalesByMerchant(res), nil

}

func (r *cashierRepository) GetYearlyCashierByMerchant(year int, merchant_id int) ([]*record.CashierRecordYearSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashierByMerchant(r.ctx, db.GetYearlyCashierByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve monthly sales data for merchant %d in year %d: %w", merchant_id, year, err)
	}

	return r.mapping.ToCashierYearlySalesByMerchant(res), nil
}

func (r *cashierRepository) GetMonthlyCashierById(year int, cashier_id int) ([]*record.CashierRecordMonthSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashierByCashierId(r.ctx, db.GetMonthlyCashierByCashierIdParams{
		Column1:   yearStart,
		CashierID: int32(cashier_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve yearly sales data for cashier ID %d in year %d: %w", cashier_id, year, err)
	}

	return r.mapping.ToCashierMonthlySalesById(res), nil
}

func (r *cashierRepository) GetYearlyCashierById(year int, cashier_id int) ([]*record.CashierRecordYearSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashierByCashierId(r.ctx, db.GetYearlyCashierByCashierIdParams{
		Column1:   yearStart,
		CashierID: int32(cashier_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve yearly sales data for cashier ID %d in year %d: %w", cashier_id, year, err)
	}

	return r.mapping.ToCashierYearlySalesById(res), nil
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
		CashierID: int32(*request.CashierID),
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
