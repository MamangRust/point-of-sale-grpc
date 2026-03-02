package repository

import (
	"context"
	"pointofsale/internal/domain/requests"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/category_errors"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type categoryRepository struct {
	db *db.Queries
}

func NewCategoryRepository(db *db.Queries) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) FindAllCategory(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCategoriesParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategories(ctx, reqDb)

	if err != nil {
		return nil, category_errors.ErrFindAllCategory
	}

	return res, nil
}

func (r *categoryRepository) FindById(ctx context.Context, category_id int) (*db.GetCategoryByIDRow, error) {
	res, err := r.db.GetCategoryByID(ctx, int32(category_id))

	if err != nil {
		return nil, category_errors.ErrFindById
	}

	return res, nil
}

func (r *categoryRepository) FindByName(ctx context.Context, name string) (*db.GetCategoryByNameRow, error) {
	res, err := r.db.GetCategoryByName(ctx, name)

	if err != nil {
		return nil, category_errors.ErrFindByName
	}

	return res, nil
}

func (r *categoryRepository) FindByNameAndId(ctx context.Context, req *requests.CategoryNameAndId) (*db.GetCategoryByNameAndIdRow, error) {
	res, err := r.db.GetCategoryByNameAndId(ctx, db.GetCategoryByNameAndIdParams{
		Name:       req.Name,
		CategoryID: int32(req.CategoryID),
	})

	if err != nil {
		return nil, category_errors.ErrFindByNameAndId
	}

	return res, nil
}

func (r *categoryRepository) FindByActive(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesActiveRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCategoriesActiveParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesActive(ctx, reqDb)

	if err != nil {
		return nil, category_errors.ErrFindByActive
	}

	return res, nil
}

func (r *categoryRepository) FindByTrashed(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesTrashedRow, error) {
	offset := (req.Page - 1) * req.PageSize

	reqDb := db.GetCategoriesTrashedParams{
		Column1: req.Search,
		Limit:   int32(req.PageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetCategoriesTrashed(ctx, reqDb)

	if err != nil {
		return nil, category_errors.ErrFindByTrashed
	}

	return res, nil
}

func (r *categoryRepository) GetMonthlyTotalPrice(ctx context.Context, req *requests.MonthTotalPrice) ([]*db.GetMonthlyTotalPriceRow, error) {
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

	res, err := r.db.GetMonthlyTotalPrice(ctx, db.GetMonthlyTotalPriceParams{
		Extract:     extractDate,
		CreatedAt:   currentEnd,
		CreatedAt_2: prevStart,
		CreatedAt_3: prevEnd,
	})

	if err != nil {
		return nil, category_errors.ErrGetMonthlyTotalPrice
	}

	return res, nil
}

func (r *categoryRepository) GetYearlyTotalPrices(ctx context.Context, year int) ([]*db.GetYearlyTotalPriceRow, error) {
	res, err := r.db.GetYearlyTotalPrice(ctx, int32(year))

	if err != nil {
		return nil, category_errors.ErrGetYearlyTotalPrices
	}

	return res, nil
}

func (r *categoryRepository) GetMonthlyTotalPriceById(ctx context.Context, req *requests.MonthTotalPriceCategory) ([]*db.GetMonthlyTotalPriceByIdRow, error) {
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

	res, err := r.db.GetMonthlyTotalPriceById(ctx, db.GetMonthlyTotalPriceByIdParams{
		Extract:     extractDate,
		CreatedAt:   currentEnd,
		CreatedAt_2: prevStart,
		CreatedAt_3: prevEnd,
		CategoryID:  int32(req.CategoryID),
	})

	if err != nil {
		return nil, category_errors.ErrGetMonthlyTotalPriceById
	}

	return res, nil
}

func (r *categoryRepository) GetYearlyTotalPricesById(ctx context.Context, req *requests.YearTotalPriceCategory) ([]*db.GetYearlyTotalPriceByIdRow, error) {
	res, err := r.db.GetYearlyTotalPriceById(ctx, db.GetYearlyTotalPriceByIdParams{
		Column1:    int32(req.Year),
		CategoryID: int32(req.CategoryID),
	})

	if err != nil {
		return nil, category_errors.ErrGetYearlyTotalPricesById
	}

	return res, nil
}

func (r *categoryRepository) GetMonthlyTotalPriceByMerchant(ctx context.Context, req *requests.MonthTotalPriceMerchant) ([]*db.GetMonthlyTotalPriceByMerchantRow, error) {
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

	res, err := r.db.GetMonthlyTotalPriceByMerchant(ctx, db.GetMonthlyTotalPriceByMerchantParams{
		Extract:     extractDate,
		CreatedAt:   currentEnd,
		CreatedAt_2: prevStart,
		CreatedAt_3: prevEnd,
		MerchantID:  int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrGetMonthlyTotalPriceByMerchant
	}

	return res, nil
}

func (r *categoryRepository) GetYearlyTotalPricesByMerchant(ctx context.Context, req *requests.YearTotalPriceMerchant) ([]*db.GetYearlyTotalPriceByMerchantRow, error) {
	res, err := r.db.GetYearlyTotalPriceByMerchant(ctx, db.GetYearlyTotalPriceByMerchantParams{
		Column1:    int32(req.Year),
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrGetYearlyTotalPricesByMerchant
	}

	return res, nil
}

func (r *categoryRepository) GetMonthPrice(ctx context.Context, year int) ([]*db.GetMonthlyCategoryRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCategory(ctx, yearStart)

	if err != nil {
		return nil, category_errors.ErrGetMonthPrice
	}

	return res, nil
}

func (r *categoryRepository) GetYearPrice(ctx context.Context, year int) ([]*db.GetYearlyCategoryRow, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCategory(ctx, yearStart)

	if err != nil {
		return nil, category_errors.ErrGetYearPrice
	}

	return res, nil
}

func (r *categoryRepository) GetMonthPriceByMerchant(ctx context.Context, req *requests.MonthPriceMerchant) ([]*db.GetMonthlyCategoryByMerchantRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCategoryByMerchant(ctx, db.GetMonthlyCategoryByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(req.MerchantID),
	})
	if err != nil {
		return nil, category_errors.ErrGetMonthPriceByMerchant
	}

	return res, nil
}

func (r *categoryRepository) GetYearPriceByMerchant(ctx context.Context, req *requests.YearPriceMerchant) ([]*db.GetYearlyCategoryByMerchantRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCategoryByMerchant(ctx, db.GetYearlyCategoryByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(req.MerchantID),
	})

	if err != nil {
		return nil, category_errors.ErrGetYearPriceByMerchant
	}

	return res, nil
}

func (r *categoryRepository) GetMonthPriceById(ctx context.Context, req *requests.MonthPriceId) ([]*db.GetMonthlyCategoryByIdRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyCategoryById(ctx, db.GetMonthlyCategoryByIdParams{
		Column1:    yearStart,
		CategoryID: int32(req.CategoryID),
	})
	if err != nil {
		return nil, category_errors.ErrGetMonthPriceById
	}

	return res, nil
}

func (r *categoryRepository) GetYearPriceById(ctx context.Context, req *requests.YearPriceId) ([]*db.GetYearlyCategoryByIdRow, error) {
	yearStart := time.Date(req.Year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyCategoryById(ctx, db.GetYearlyCategoryByIdParams{
		Column1:    yearStart,
		CategoryID: int32(req.CategoryID),
	})

	if err != nil {
		return nil, category_errors.ErrGetYearPriceById
	}

	return res, nil
}

func (r *categoryRepository) CreateCategory(ctx context.Context, request *requests.CreateCategoryRequest) (*db.CreateCategoryRow, error) {
	req := db.CreateCategoryParams{
		Name:         request.Name,
		Description:  &request.Description,
		SlugCategory: request.SlugCategory,
	}

	category, err := r.db.CreateCategory(ctx, req)

	if err != nil {
		return nil, category_errors.ErrCreateCategory
	}

	return category, nil
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, request *requests.UpdateCategoryRequest) (*db.UpdateCategoryRow, error) {
	req := db.UpdateCategoryParams{
		CategoryID:   int32(*request.CategoryID),
		Name:         request.Name,
		Description:  &request.Description,
		SlugCategory: request.SlugCategory,
	}

	res, err := r.db.UpdateCategory(ctx, req)

	if err != nil {
		return nil, category_errors.ErrUpdateCategory
	}

	return res, nil
}

func (r *categoryRepository) TrashedCategory(ctx context.Context, category_id int) (*db.Category, error) {
	res, err := r.db.TrashCategory(ctx, int32(category_id))
	if err != nil {
		return nil, category_errors.ErrTrashedCategory
	}
	return res, nil
}

func (r *categoryRepository) RestoreCategory(ctx context.Context, category_id int) (*db.Category, error) {
	res, err := r.db.RestoreCategory(ctx, int32(category_id))
	if err != nil {
		return nil, category_errors.ErrRestoreCategory
	}
	return res, nil
}

func (r *categoryRepository) DeleteCategoryPermanently(ctx context.Context, category_id int) (bool, error) {
	err := r.db.DeleteCategoryPermanently(ctx, int32(category_id))
	if err != nil {
		return false, category_errors.ErrDeleteCategoryPermanently
	}
	return true, nil
}

func (r *categoryRepository) RestoreAllCategories(ctx context.Context) (bool, error) {
	err := r.db.RestoreAllCategories(ctx)

	if err != nil {
		return false, category_errors.ErrRestoreAllCategories
	}
	return true, nil
}

func (r *categoryRepository) DeleteAllPermanentCategories(ctx context.Context) (bool, error) {
	err := r.db.DeleteAllPermanentCategories(ctx)

	if err != nil {
		return false, category_errors.ErrDeleteAllPermanentCategories
	}
	return true, nil
}
