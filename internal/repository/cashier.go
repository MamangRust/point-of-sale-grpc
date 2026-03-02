package repository

import (
	"context"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/cashier_errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type cashierRepository struct {
	db *db.Queries
}

func NewCashierRepository(db *db.Queries) *cashierRepository {
	return &cashierRepository{
		db: db,
	}
}

func (r *cashierRepository) FindAllCashiers(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiers(ctx, reqDb)

	if err != nil {
		return nil, cashier_errors.ErrFindAllCashiers
	}

	return res, nil
}

func (r *cashierRepository) FindByActive(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersActiveParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiersActive(ctx, reqDb)

	if err != nil {
		return nil, cashier_errors.ErrFindActiveCashiers
	}

	return res, nil
}

func (r *cashierRepository) FindByTrashed(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersTrashedParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCashiersTrashed(ctx, reqDb)

	if err != nil {
		return nil, cashier_errors.ErrFindTrashedCashiers
	}

	return res, nil
}

func (r *cashierRepository) FindByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant) ([]*db.GetCashiersByMerchantRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCashiersByMerchantParams{
		MerchantID: int32(req.MerchantID),
		Column2:    req.Search,
		Limit:      int32(req.PageSize),
		Offset:     int32(offset),
	}

	res, err := r.db.GetCashiersByMerchant(ctx, reqDb)

	if err != nil {
		return nil, cashier_errors.ErrFindCashiersByMerchant
	}

	return res, nil
}

func (r *cashierRepository) FindById(ctx context.Context, cashier_id int) (*db.GetCashierByIdRow, error) {
	res, err := r.db.GetCashierById(ctx, int32(cashier_id))

	if err != nil {
		return nil, cashier_errors.ErrFindCashierById
	}

	return res, nil
}

func (r *cashierRepository) GetMonthlyTotalSales(ctx context.Context, req *requests.MonthTotalSales) ([]*db.GetMonthlyTotalSalesCashierRow, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)

	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	extractDate := pgtype.Date{
		Time:  currentMonthStart,
		Valid: true,
	}

	currentEnd := pgtype.Timestamp{
		Time:  currentMonthEnd,
		Valid: true,
	}

	prevStart := pgtype.Timestamp{
		Time:  prevMonthStart,
		Valid: true,
	}

	prevEnd := pgtype.Timestamp{
		Time:  prevMonthEnd,
		Valid: true,
	}

	params := db.GetMonthlyTotalSalesCashierParams{
		Extract:     extractDate,
		CreatedAt:   currentEnd,
		CreatedAt_2: prevStart,
		CreatedAt_3: prevEnd,
	}

	res, err := r.db.GetMonthlyTotalSalesCashier(ctx, params)

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyTotalSales
	}

	return res, nil
}

func (r *cashierRepository) GetYearlyTotalSales(ctx context.Context, year int) ([]*db.GetYearlyTotalSalesCashierRow, error) {
	res, err := r.db.GetYearlyTotalSalesCashier(ctx, int32(year))

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyTotalSales
	}

	return res, nil
}

func (r *cashierRepository) GetMonthlyTotalSalesById(ctx context.Context, req *requests.MonthTotalSalesCashier) ([]*db.GetMonthlyTotalSalesByIdRow, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)

	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	extractDate := pgtype.Date{
		Time:  currentMonthStart,
		Valid: true,
	}

	currentEnd := pgtype.Timestamp{
		Time:  currentMonthEnd,
		Valid: true,
	}

	prevStart := pgtype.Timestamp{
		Time:  prevMonthStart,
		Valid: true,
	}

	prevEnd := pgtype.Timestamp{
		Time:  prevMonthEnd,
		Valid: true,
	}

	res, err := r.db.GetMonthlyTotalSalesById(ctx, db.GetMonthlyTotalSalesByIdParams{
		Extract:     extractDate,
		CreatedAt:   currentEnd,
		CreatedAt_2: prevStart,
		CreatedAt_3: prevEnd,
		CashierID:   int32(req.CashierID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyTotalSalesById
	}

	return res, nil
}

func (r *cashierRepository) GetYearlyTotalSalesById(ctx context.Context, req *requests.YearTotalSalesCashier) ([]*db.GetYearlyTotalSalesByIdRow, error) {
	res, err := r.db.GetYearlyTotalSalesById(ctx, db.GetYearlyTotalSalesByIdParams{
		Column1:   int32(req.Year),
		CashierID: int32(req.CashierID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyTotalSalesById
	}

	return res, nil
}

func (r *cashierRepository) GetMonthlyTotalSalesByMerchant(ctx context.Context, req *requests.MonthTotalSalesMerchant) ([]*db.GetMonthlyTotalSalesByMerchantRow, error) {
	currentMonthStart := time.Date(req.Year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	extractDate := pgtype.Date{
		Time:  currentMonthStart,
		Valid: true,
	}

	currentEnd := pgtype.Timestamp{
		Time:  currentMonthEnd,
		Valid: true,
	}

	prevStart := pgtype.Timestamp{
		Time:  prevMonthStart,
		Valid: true,
	}

	prevEnd := pgtype.Timestamp{
		Time:  prevMonthEnd,
		Valid: true,
	}

	res, err := r.db.GetMonthlyTotalSalesByMerchant(ctx, db.GetMonthlyTotalSalesByMerchantParams{
		Extract:     extractDate,
		CreatedAt:   currentEnd,
		CreatedAt_2: prevStart,
		CreatedAt_3: prevEnd,
		MerchantID:  int32(req.MerchantID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyTotalSalesByMerchant
	}

	return res, nil
}

func (r *cashierRepository) GetYearlyTotalSalesByMerchant(ctx context.Context, req *requests.YearTotalSalesMerchant) ([]*db.GetYearlyTotalSalesByMerchantRow, error) {
	res, err := r.db.GetYearlyTotalSalesByMerchant(ctx, db.GetYearlyTotalSalesByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyTotalSalesByMerchant
	}

	return res, nil
}

func (r *cashierRepository) GetMonthyCashier(ctx context.Context, year int) ([]*db.GetMonthlyCashierRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashier(ctx, yearStart)

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyCashier
	}

	return res, nil

}

func (r *cashierRepository) GetYearlyCashier(ctx context.Context, year int) ([]*db.GetYearlyCashierRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashier(ctx, yearStart)

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyCashier
	}

	return res, nil
}

func (r *cashierRepository) GetMonthlyCashierByMerchant(ctx context.Context, req *requests.MonthCashierMerchant) ([]*db.GetMonthlyCashierByMerchantRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashierByMerchant(ctx, db.GetMonthlyCashierByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyCashierByMerchant
	}

	return res, nil

}

func (r *cashierRepository) GetYearlyCashierByMerchant(ctx context.Context, req *requests.YearCashierMerchant) ([]*db.GetYearlyCashierByMerchantRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashierByMerchant(ctx, db.GetYearlyCashierByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(req.Year),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyCashierByMerchant
	}

	return res, nil
}

func (r *cashierRepository) GetMonthlyCashierById(ctx context.Context, req *requests.MonthCashierId) ([]*db.GetMonthlyCashierByCashierIdRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCashierByCashierId(ctx, db.GetMonthlyCashierByCashierIdParams{
		Column1:   yearStart,
		CashierID: int32(req.Year),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetMonthlyCashierById
	}

	return res, nil
}

func (r *cashierRepository) GetYearlyCashierById(ctx context.Context, req *requests.YearCashierId) ([]*db.GetYearlyCashierByCashierIdRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCashierByCashierId(ctx, db.GetYearlyCashierByCashierIdParams{
		Column1:   yearStart,
		CashierID: int32(req.CashierID),
	})

	if err != nil {
		return nil, cashier_errors.ErrGetYearlyCashierById
	}

	return res, nil
}

func (r *cashierRepository) CreateCashier(ctx context.Context, request *requests.CreateCashierRequest) (*db.CreateCashierRow, error) {
	req := db.CreateCashierParams{
		MerchantID: int32(request.MerchantID),
		UserID:     int32(request.UserID),
		Name:       request.Name,
	}

	cashier, err := r.db.CreateCashier(ctx, req)

	if err != nil {
		return nil, cashier_errors.ErrCreateCashier
	}

	return cashier, nil
}

func (r *cashierRepository) UpdateCashier(ctx context.Context, request *requests.UpdateCashierRequest) (*db.UpdateCashierRow, error) {
	req := db.UpdateCashierParams{
		CashierID: int32(*request.CashierID),
		Name:      request.Name,
	}

	res, err := r.db.UpdateCashier(ctx, req)

	if err != nil {
		return nil, cashier_errors.ErrUpdateCashier
	}

	return res, nil
}

func (r *cashierRepository) TrashedCashier(ctx context.Context, cashier_id int) (*db.Cashier, error) {
	res, err := r.db.TrashCashier(ctx, int32(cashier_id))

	if err != nil {
		return nil, cashier_errors.ErrTrashedCashier
	}

	return res, nil
}

func (r *cashierRepository) RestoreCashier(ctx context.Context, cashier_id int) (*db.Cashier, error) {
	res, err := r.db.RestoreCashier(ctx, int32(cashier_id))

	if err != nil {
		return nil, cashier_errors.ErrRestoreCashier
	}

	return res, nil
}

func (r *cashierRepository) DeleteCashierPermanent(ctx context.Context, cashier_id int) (bool, error) {
	err := r.db.DeleteCashierPermanently(ctx, int32(cashier_id))

	if err != nil {
		return false, cashier_errors.ErrDeleteCashierPermanent
	}

	return true, nil
}

func (r *cashierRepository) RestoreAllCashier(ctx context.Context) (bool, error) {
	err := r.db.RestoreAllCashiers(ctx)

	if err != nil {
		return false, cashier_errors.ErrRestoreAllCashiers
	}

	return true, nil
}

func (r *cashierRepository) DeleteAllCashierPermanent(ctx context.Context) (bool, error) {
	err := r.db.DeleteAllPermanentCashiers(ctx)

	if err != nil {
		return false, cashier_errors.ErrDeleteAllCashiersPermanent
	}

	return true, nil
}
