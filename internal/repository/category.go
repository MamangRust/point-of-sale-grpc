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

type categoryRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.CategoryRecordMapper
}

func NewCategoryRepository(db *db.Queries, ctx context.Context, mapping recordmapper.CategoryRecordMapper) *categoryRepository {
	return &categoryRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *categoryRepository) FindAllCategory(search string, page, pageSize int) ([]*record.CategoriesRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCategoriesParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategories(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find categories: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategoriesRecordPagination(res), totalCount, nil
}

func (r *categoryRepository) FindById(category_id int) (*record.CategoriesRecord, error) {
	res, err := r.db.GetCategoryByID(r.ctx, int32(category_id))

	if err != nil {
		fmt.Printf("Error fetching cashier: %v\n", err)

		return nil, fmt.Errorf("failed to find cashier: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) FindByActive(search string, page, pageSize int) ([]*record.CategoriesRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCategoriesActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cashier: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategoriesRecordActivePagination(res), totalCount, nil
}

func (r *categoryRepository) FindByTrashed(search string, page, pageSize int) ([]*record.CategoriesRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetCategoriesTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find cashiers: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToCategoriesRecordTrashedPagination(res), totalCount, nil
}

func (r *categoryRepository) GetMonthlyTotalPrice(year int, month int) ([]*record.CategoriesMonthlyTotalPriceRecord, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalPrice(r.ctx, db.GetMonthlyTotalPriceParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total price category: %w", err)
	}

	so := r.mapping.ToCategoryMonthlyTotalPrices(res)

	return so, nil
}

func (r *categoryRepository) GetYearlyTotalPrices(year int) ([]*record.CategoriesYearlyTotalPriceRecord, error) {
	res, err := r.db.GetYearlyTotalPrice(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total price category: %w", err)
	}

	so := r.mapping.ToCategoryYearlyTotalPrices(res)

	return so, nil
}

func (r *categoryRepository) GetMonthlyTotalPriceById(year int, month int, order_id int) ([]*record.CategoriesMonthlyTotalPriceRecord, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalPriceById(r.ctx, db.GetMonthlyTotalPriceByIdParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		OrderID:     int32(order_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total price category: %w", err)
	}

	so := r.mapping.ToCategoryMonthlyTotalPricesById(res)

	return so, nil
}

func (r *categoryRepository) GetYearlyTotalPricesById(year int, order_id int) ([]*record.CategoriesYearlyTotalPriceRecord, error) {
	res, err := r.db.GetYearlyTotalPriceById(r.ctx, db.GetYearlyTotalPriceByIdParams{
		Column1: int32(year),
		OrderID: int32(order_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total price category: %w", err)
	}

	so := r.mapping.ToCategoryYearlyTotalPricesById(res)

	return so, nil
}

func (r *categoryRepository) GetMonthlyTotalPriceByMerchant(year int, month int, merchant_id int) ([]*record.CategoriesMonthlyTotalPriceRecord, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalPriceByMerchant(r.ctx, db.GetMonthlyTotalPriceByMerchantParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		MerchantID:  int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly total price category: %w", err)
	}

	so := r.mapping.ToCategoryMonthlyTotalPricesByMerchant(res)

	return so, nil
}

func (r *categoryRepository) GetYearlyTotalPricesByMerchant(year int, merchant_id int) ([]*record.CategoriesYearlyTotalPriceRecord, error) {
	res, err := r.db.GetYearlyTotalPriceByMerchant(r.ctx, db.GetYearlyTotalPriceByMerchantParams{
		Column1:    int32(year),
		MerchantID: int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total price category: %w", err)
	}

	so := r.mapping.ToCategoryYearlyTotalPricesByMerchant(res)

	return so, nil
}

func (r *categoryRepository) GetMonthPrice(year int) ([]*record.CategoriesMonthPriceRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCategory(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly category prices: %w", err)
	}

	return r.mapping.ToCategoryMonthlyPrices(res), nil
}

func (r *categoryRepository) GetYearPrice(year int) ([]*record.CategoriesYearPriceRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCategory(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly category prices: %w", err)
	}

	return r.mapping.ToCategoryYearlyPrices(res), nil
}

func (r *categoryRepository) GetMonthPriceByMerchant(year int, merchant_id int) ([]*record.CategoriesMonthPriceRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCategoryByMerchant(r.ctx, db.GetMonthlyCategoryByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly category prices for merchant %d: %w", merchant_id, err)
	}

	return r.mapping.ToCategoryMonthlyPricesByMerchant(res), nil
}

func (r *categoryRepository) GetYearPriceByMerchant(year int, merchant_id int) ([]*record.CategoriesYearPriceRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCategoryByMerchant(r.ctx, db.GetYearlyCategoryByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly category prices for merchant %d: %w", merchant_id, err)
	}

	return r.mapping.ToCategoryYearlyPricesByMerchant(res), nil
}

func (r *categoryRepository) GetMonthPriceById(year int, category_id int) ([]*record.CategoriesMonthPriceRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCategoryById(r.ctx, db.GetMonthlyCategoryByIdParams{
		Column1:    yearStart,
		CategoryID: int32(category_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly prices for category %d: %w", category_id, err)
	}

	return r.mapping.ToCategoryMonthlyPricesById(res), nil
}

func (r *categoryRepository) GetYearPriceById(year int, category_id int) ([]*record.CategoriesYearPriceRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCategoryById(r.ctx, db.GetYearlyCategoryByIdParams{
		Column1:    yearStart,
		CategoryID: int32(category_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly prices for category %d: %w", category_id, err)
	}

	return r.mapping.ToCategoryYearlyPricesById(res), nil
}

func (r *categoryRepository) CreateCategory(request *requests.CreateCategoryRequest) (*record.CategoriesRecord, error) {
	fmt.Println("hello des", request.Description)

	req := db.CreateCategoryParams{
		Name: request.Name,
		Description: sql.NullString{
			String: request.Description,
			Valid:  true,
		},
		SlugCategory: sql.NullString{
			String: *request.SlugCategory,
			Valid:  true,
		},
	}

	category, err := r.db.CreateCategory(r.ctx, req)

	if err != nil {
		return nil, errors.New("failed to create category")
	}

	return r.mapping.ToCategoryRecord(category), nil
}

func (r *categoryRepository) UpdateCategory(request *requests.UpdateCategoryRequest) (*record.CategoriesRecord, error) {
	req := db.UpdateCategoryParams{
		CategoryID: int32(*request.CategoryID),
		Name:       request.Name,
		Description: sql.NullString{
			String: request.Description,
			Valid:  true,
		},
		SlugCategory: sql.NullString{
			String: *request.SlugCategory,
			Valid:  true,
		},
	}

	res, err := r.db.UpdateCategory(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) TrashedCategory(category_id int) (*record.CategoriesRecord, error) {
	res, err := r.db.TrashCategory(r.ctx, int32(category_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash category: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) RestoreCategory(category_id int) (*record.CategoriesRecord, error) {
	res, err := r.db.RestoreCategory(r.ctx, int32(category_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore category: %w", err)
	}

	return r.mapping.ToCategoryRecord(res), nil
}

func (r *categoryRepository) DeleteCategoryPermanently(category_id int) (bool, error) {
	err := r.db.DeleteCategoryPermanently(r.ctx, int32(category_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete category: %w", err)
	}

	return true, nil
}

func (r *categoryRepository) RestoreAllCategories() (bool, error) {
	err := r.db.RestoreAllCategories(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all category: %w", err)
	}
	return true, nil
}

func (r *categoryRepository) DeleteAllPermanentCategories() (bool, error) {
	err := r.db.DeleteAllPermanentCategories(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all category permanently: %w", err)
	}
	return true, nil
}
