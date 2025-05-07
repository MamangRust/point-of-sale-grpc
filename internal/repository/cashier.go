package repository

import (
	"context"
	"database/sql"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/cashier_errors"
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

func (r *cashierRepository) FindAllCashiers(req *requests.FindAllCashiers) ([]*record.CashierRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiers(r.ctx, reqDb)

	if err != nil {
		return nil, nil, cashier_errors.ErrFindAllCashiers
	}

	var totalCount int

	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersRecordPagination(res), &totalCount, nil
}

func (r *cashierRepository) FindByActive(req *requests.FindAllCashiers) ([]*record.CashierRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersActiveParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiersActive(r.ctx, reqDb)

	if err != nil {
		return nil, nil, cashier_errors.ErrFindActiveCashiers
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersRecordActivePagination(res), &totalCount, nil
}

func (r *cashierRepository) FindByTrashed(req *requests.FindAllCashiers) ([]*record.CashierRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersTrashedParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiersTrashed(r.ctx, reqDb)

	if err != nil {
		return nil, nil, cashier_errors.ErrFindTrashedCashiers
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersRecordTrashedPagination(res), &totalCount, nil
}

func (r *cashierRepository) FindByMerchant(req *requests.FindAllCashierMerchant) ([]*record.CashierRecord, *int, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersByMerchantParams{
		MerchantID: int32(req.MerchantID),
		Column2:    req.Search,
		Limit:      int32(req.PageSize),
		Offset:     int32(offset),
	}

	res, err := r.db.GetCashiersByMerchant(r.ctx, reqDb)

	if err != nil {
		return nil, nil, cashier_errors.ErrFindCashiersByMerchant
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCashiersMerchantRecordPagination(res), &totalCount, nil
}

func (r *cashierRepository) FindById(cashier_id int) (*record.CashierRecord, error) {
	res, err := r.db.GetCashierById(r.ctx, int32(cashier_id))

	if err != nil {
		return nil, cashier_errors.ErrFindCashierById
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) GetMonthlyTotalSales(req *requests.MonthTotalSales) ([]*record.CashierRecordMonthTotalSales, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
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
		return nil, cashier_errors.ErrGetMonthlyTotalSales
	}

	return r.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (r *cashierRepository) GetYearlyTotalSales(year int) ([]*record.CashierRecordYearTotalSales, error) {
	res, err := r.db.GetYearlyTotalSalesCashier(r.ctx, int32(year))

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyTotalSales
	}

	so := r.mapping.ToCashierYearlyTotalSales(res)

	return so, nil
}

func (r *cashierRepository) GetMonthlyTotalSalesById(req *requests.MonthTotalSalesCashier) ([]*record.CashierRecordMonthTotalSales, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)

	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalSalesById(r.ctx, db.GetMonthlyTotalSalesByIdParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		CashierID:   int32(req.CashierID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyTotalSalesById
	}

	so := r.mapping.ToCashierMonthlyTotalSalesById(res)

	return so, nil
}

func (r *cashierRepository) GetYearlyTotalSalesById(req *requests.YearTotalSalesCashier) ([]*record.CashierRecordYearTotalSales, error) {
	res, err := r.db.GetYearlyTotalSalesById(r.ctx, db.GetYearlyTotalSalesByIdParams{
		Column1:   int32(req.Year),
		CashierID: int32(req.CashierID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyTotalSalesById
	}

	so := r.mapping.ToCashierYearlyTotalSalesById(res)

	return so, nil
}

func (r *cashierRepository) GetMonthlyTotalSalesByMerchant(req *requests.MonthTotalSalesMerchant) ([]*record.CashierRecordMonthTotalSales, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalSalesByMerchant(r.ctx, db.GetMonthlyTotalSalesByMerchantParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		MerchantID:  int32(req.MerchantID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyTotalSalesByMerchant
	}

	so := r.mapping.ToCashierMonthlyTotalSalesByMerchant(res)

	return so, nil
}

func (r *cashierRepository) GetYearlyTotalSalesByMerchant(req *requests.YearTotalSalesMerchant) ([]*record.CashierRecordYearTotalSales, error) {
	res, err := r.db.GetYearlyTotalSalesByMerchant(r.ctx, db.GetYearlyTotalSalesByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyTotalSalesByMerchant
	}

	so := r.mapping.ToCashierYearlyTotalSalesByMerchant(res)

	return so, nil
}

func (r *cashierRepository) GetMonthyCashier(year int) ([]*record.CashierRecordMonthSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashier(r.ctx, yearStart)

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyCashier
	}

	return r.mapping.ToCashierMonthlySales(res), nil

}

func (r *cashierRepository) GetYearlyCashier(year int) ([]*record.CashierRecordYearSales, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashier(r.ctx, yearStart)

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyCashier
	}

	return r.mapping.ToCashierYearlySales(res), nil
}

func (r *cashierRepository) GetMonthlyCashierByMerchant(req *requests.MonthCashierMerchant) ([]*record.CashierRecordMonthSales, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashierByMerchant(r.ctx, db.GetMonthlyCashierByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyCashierByMerchant
	}

	return r.mapping.ToCashierMonthlySalesByMerchant(res), nil

}

func (r *cashierRepository) GetYearlyCashierByMerchant(req *requests.YearCashierMerchant) ([]*record.CashierRecordYearSales, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashierByMerchant(r.ctx, db.GetYearlyCashierByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(req.Year),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyCashierByMerchant
	}

	return r.mapping.ToCashierYearlySalesByMerchant(res), nil
}

func (r *cashierRepository) GetMonthlyCashierById(req *requests.MonthCashierId) ([]*record.CashierRecordMonthSales, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashierByCashierId(r.ctx, db.GetMonthlyCashierByCashierIdParams{
		Column1:   yearStart,
		CashierID: int32(req.Year),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyCashierById
	}

	return r.mapping.ToCashierMonthlySalesById(res), nil
}

func (r *cashierRepository) GetYearlyCashierById(req *requests.YearCashierId) ([]*record.CashierRecordYearSales, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashierByCashierId(r.ctx, db.GetYearlyCashierByCashierIdParams{
		Column1:   yearStart,
		CashierID: int32(req.CashierID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyCashierById
	}

	return r.mapping.ToCashierYearlySalesById(res), nil
}

func (r *cashierRepository) CreateCashier(request *requests.CreateCashierRequest) (*record.CashierRecord, error) {
	req := db.CreateCashierParams{
		MerchantID: int32(request.MerchantID),
		UserID:     int32(request.UserID),
		Name:       request.Name,
	}

	cashier, err := r.db.CreateCashier(r.ctx, req)

	if err != nil {
		return nil, cashier_errors.ErrCreateCashier
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
		return nil, cashier_errors.ErrUpdateCashier
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) TrashedCashier(cashier_id int) (*record.CashierRecord, error) {
	res, err := r.db.TrashCashier(r.ctx, int32(cashier_id))

	if err != nil {
		return nil, cashier_errors.ErrTrashedCashier
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) RestoreCashier(cashier_id int) (*record.CashierRecord, error) {
	res, err := r.db.RestoreCashier(r.ctx, int32(cashier_id))

	if err != nil {
		return nil, cashier_errors.ErrRestoreCashier
	}

	return r.mapping.ToCashierRecord(res), nil
}

func (r *cashierRepository) DeleteCashierPermanent(cashier_id int) (bool, error) {
	err := r.db.DeleteCashierPermanently(r.ctx, int32(cashier_id))

	if err != nil {
		return false, cashier_errors.ErrDeleteCashierPermanent
	}

	return true, nil
}

func (r *cashierRepository) RestoreAllCashier() (bool, error) {
	err := r.db.RestoreAllCashiers(r.ctx)

	if err != nil {
		return false, cashier_errors.ErrRestoreAllCashiers
	}

	return true, nil
}

func (r *cashierRepository) DeleteAllCashierPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentCashiers(r.ctx)

	if err != nil {
		return false, cashier_errors.ErrDeleteAllCashiersPermanent
	}

	return true, nil
}
