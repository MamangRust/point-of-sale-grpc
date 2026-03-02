package service

import (
	"context"
	category_cache "pointofsale/internal/cache/category"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/errorhandler"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/category_errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/pkg/utils"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
	logger             logger.LoggerInterface
	observability      observability.TraceLoggerObservability
	cache              category_cache.CategoryMencache
}

type CategoryServiceDeps struct {
	CategoryRepo  repository.CategoryRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
	cache         category_cache.CategoryMencache
}

func NewCategoryService(deps CategoryServiceDeps) *categoryService {
	return &categoryService{
		categoryRepository: deps.CategoryRepo,
		logger:             deps.Logger,
		observability:      deps.Observability,
		cache:              deps.cache,
	}
}

func (s *categoryService) FindAllCategory(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesRow, *int, error) {
	const method = "FindAllCategory"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedCategoriesCache(ctx, req); found {
		logSuccess("Successfully retrieved all category records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	categories, err := s.categoryRepository.FindAllCategory(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCategoriesRow](
			s.logger,
			category_errors.ErrFailedFindAllCategories,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(categories) > 0 {
		totalCount = int(categories[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCategoriesCache(ctx, req, categories, &totalCount)

	logSuccess("Successfully fetched categories",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return categories, &totalCount, nil
}

func (s *categoryService) FindByActive(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesActiveRow, *int, error) {
	const method = "FindByActive"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedCategoryActiveCache(ctx, req); found {
		logSuccess("Successfully retrieved active category records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	categories, err := s.categoryRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCategoriesActiveRow](
			s.logger,
			category_errors.ErrFailedFindCategoryByActive,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(categories) > 0 {
		totalCount = int(categories[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCategoryActiveCache(ctx, req, categories, &totalCount)

	logSuccess("Successfully fetched active categories",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return categories, &totalCount, nil
}

func (s *categoryService) FindByTrashed(ctx context.Context, req *requests.FindAllCategory) ([]*db.GetCategoriesTrashedRow, *int, error) {
	const method = "FindByTrashed"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedCategoryTrashedCache(ctx, req); found {
		logSuccess("Successfully retrieved trashed category records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	categories, err := s.categoryRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCategoriesTrashedRow](
			s.logger,
			category_errors.ErrFailedFindCategoryByTrashed,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(categories) > 0 {
		totalCount = int(categories[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCategoryTrashedCache(ctx, req, categories, &totalCount)

	logSuccess("Successfully fetched trashed categories",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return categories, &totalCount, nil
}

func (s *categoryService) FindById(ctx context.Context, category_id int) (*db.GetCategoryByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("category_id", category_id))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedCategoryCache(ctx, category_id); found {
		logSuccess("Successfully retrieved category record from cache",
			zap.Int("category_id", category_id))
		return data, nil
	}

	category, err := s.categoryRepository.FindById(ctx, category_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetCategoryByIDRow](
			s.logger,
			category_errors.ErrFailedFindCategoryById,
			method,
			span,
			zap.Int("category_id", category_id))
	}

	s.cache.SetCachedCategoryCache(ctx, category)

	logSuccess("Successfully fetched category",
		zap.Int("category_id", category_id))

	return category, nil
}

func (s *categoryService) FindMonthlyTotalPrice(ctx context.Context, req *requests.MonthTotalPrice) ([]*db.GetMonthlyTotalPriceRow, error) {
	const method = "FindMonthlyTotalPrice"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthTotalPriceCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total price from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.categoryRepository.GetMonthlyTotalPrice(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalPriceRow](
			s.logger,
			category_errors.ErrFailedFindMonthlyTotalPrice,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
	}

	s.cache.SetCachedMonthTotalPriceCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total price",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindYearlyTotalPrice(ctx context.Context, year int) ([]*db.GetYearlyTotalPriceRow, error) {
	const method = "FindYearlyTotalPrice"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearTotalPriceCache(ctx, year); found {
		logSuccess("Successfully retrieved yearly total price from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.categoryRepository.GetYearlyTotalPrices(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalPriceRow](
			s.logger,
			category_errors.ErrFailedFindYearlyTotalPrice,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetCachedYearTotalPriceCache(ctx, year, res)

	logSuccess("Successfully fetched yearly total price",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindMonthlyTotalPriceById(ctx context.Context, req *requests.MonthTotalPriceCategory) ([]*db.GetMonthlyTotalPriceByIdRow, error) {
	const method = "FindMonthlyTotalPriceById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("category_id", req.CategoryID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthTotalPriceByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total price by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("category_id", req.CategoryID))
		return data, nil
	}

	res, err := s.categoryRepository.GetMonthlyTotalPriceById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalPriceByIdRow](
			s.logger,
			category_errors.ErrFailedFindMonthlyTotalPriceById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("category_id", req.CategoryID))
	}

	s.cache.SetCachedMonthTotalPriceByIdCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total price by ID",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("category_id", req.CategoryID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindYearlyTotalPriceById(ctx context.Context, req *requests.YearTotalPriceCategory) ([]*db.GetYearlyTotalPriceByIdRow, error) {
	const method = "FindYearlyTotalPriceById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("category_id", req.CategoryID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearTotalPriceByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly total price by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("category_id", req.CategoryID))
		return data, nil
	}

	res, err := s.categoryRepository.GetYearlyTotalPricesById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalPriceByIdRow](
			s.logger,
			category_errors.ErrFailedFindYearlyTotalPriceById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("category_id", req.CategoryID))
	}

	s.cache.SetCachedYearTotalPriceByIdCache(ctx, req, res)

	logSuccess("Successfully fetched yearly total price by ID",
		zap.Int("year", req.Year),
		zap.Int("category_id", req.CategoryID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindMonthlyTotalPriceByMerchant(ctx context.Context, req *requests.MonthTotalPriceMerchant) ([]*db.GetMonthlyTotalPriceByMerchantRow, error) {
	const method = "FindMonthlyTotalPriceByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthTotalPriceByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total price by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.categoryRepository.GetMonthlyTotalPriceByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalPriceByMerchantRow](
			s.logger,
			category_errors.ErrFailedFindMonthlyTotalPriceByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetCachedMonthTotalPriceByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total price by merchant",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindYearlyTotalPriceByMerchant(ctx context.Context, req *requests.YearTotalPriceMerchant) ([]*db.GetYearlyTotalPriceByMerchantRow, error) {
	const method = "FindYearlyTotalPriceByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearTotalPriceByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly total price by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.categoryRepository.GetYearlyTotalPricesByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalPriceByMerchantRow](
			s.logger,
			category_errors.ErrFailedFindYearlyTotalPriceByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetCachedYearTotalPriceByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched yearly total price by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindMonthPrice(ctx context.Context, year int) ([]*db.GetMonthlyCategoryRow, error) {
	const method = "FindMonthPrice"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthPriceCache(ctx, year); found {
		logSuccess("Successfully retrieved month price from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.categoryRepository.GetMonthPrice(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyCategoryRow](
			s.logger,
			category_errors.ErrFailedFindMonthPrice,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetCachedMonthPriceCache(ctx, year, res)

	logSuccess("Successfully fetched month price",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindYearPrice(ctx context.Context, year int) ([]*db.GetYearlyCategoryRow, error) {
	const method = "FindYearPrice"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearPriceCache(ctx, year); found {
		logSuccess("Successfully retrieved year price from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.categoryRepository.GetYearPrice(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyCategoryRow](
			s.logger,
			category_errors.ErrFailedFindYearPrice,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetCachedYearPriceCache(ctx, year, res)

	logSuccess("Successfully fetched year price",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindMonthPriceByMerchant(ctx context.Context, req *requests.MonthPriceMerchant) ([]*db.GetMonthlyCategoryByMerchantRow, error) {
	const method = "FindMonthPriceByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthPriceByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved month price by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.categoryRepository.GetMonthPriceByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyCategoryByMerchantRow](
			s.logger,
			category_errors.ErrFailedFindMonthPriceByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetCachedMonthPriceByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched month price by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindYearPriceByMerchant(ctx context.Context, req *requests.YearPriceMerchant) ([]*db.GetYearlyCategoryByMerchantRow, error) {
	const method = "FindYearPriceByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearPriceByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved year price by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.categoryRepository.GetYearPriceByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyCategoryByMerchantRow](
			s.logger,
			category_errors.ErrFailedFindYearPriceByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetCachedYearPriceByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched year price by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindMonthPriceById(ctx context.Context, req *requests.MonthPriceId) ([]*db.GetMonthlyCategoryByIdRow, error) {
	const method = "FindMonthPriceById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("category_id", req.CategoryID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthPriceByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved month price by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("category_id", req.CategoryID))
		return data, nil
	}

	res, err := s.categoryRepository.GetMonthPriceById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyCategoryByIdRow](
			s.logger,
			category_errors.ErrFailedFindMonthPriceById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("category_id", req.CategoryID))
	}

	s.cache.SetCachedMonthPriceByIdCache(ctx, req, res)

	logSuccess("Successfully fetched month price by ID",
		zap.Int("year", req.Year),
		zap.Int("category_id", req.CategoryID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) FindYearPriceById(ctx context.Context, req *requests.YearPriceId) ([]*db.GetYearlyCategoryByIdRow, error) {
	const method = "FindYearPriceById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("category_id", req.CategoryID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearPriceByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved year price by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("category_id", req.CategoryID))
		return data, nil
	}

	res, err := s.categoryRepository.GetYearPriceById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyCategoryByIdRow](
			s.logger,
			category_errors.ErrFailedFindYearPriceById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("category_id", req.CategoryID))
	}

	s.cache.SetCachedYearPriceByIdCache(ctx, req, res)

	logSuccess("Successfully fetched year price by ID",
		zap.Int("year", req.Year),
		zap.Int("category_id", req.CategoryID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *categoryService) CreateCategory(ctx context.Context, req *requests.CreateCategoryRequest) (*db.CreateCategoryRow, error) {
	const method = "CreateCategory"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("name", req.Name))

	defer func() {
		end(status)
	}()

	slug := utils.GenerateSlug(req.Name)
	req.SlugCategory = &slug

	_, err := s.categoryRepository.FindByName(ctx, req.Name)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateCategoryRow](
			s.logger,
			category_errors.ErrFailedFindCategoryByName,
			method,
			span,
			zap.String("name", req.Name))
	}

	category, err := s.categoryRepository.CreateCategory(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateCategoryRow](
			s.logger,
			category_errors.ErrFailedCreateCategory,
			method,
			span,
			zap.Any("request", req))
	}

	s.cache.DeleteCachedCategoryCache(ctx, int(category.CategoryID))

	logSuccess("Successfully created category",
		zap.Int("category_id", int(category.CategoryID)),
		zap.String("name", req.Name))

	return category, nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, req *requests.UpdateCategoryRequest) (*db.UpdateCategoryRow, error) {
	const method = "UpdateCategory"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("category_id", *req.CategoryID),
		attribute.String("name", req.Name))

	defer func() {
		end(status)
	}()

	slug := utils.GenerateSlug(req.Name)
	req.SlugCategory = &slug

	_, err := s.categoryRepository.FindByNameAndId(ctx, &requests.CategoryNameAndId{
		Name:       req.Name,
		CategoryID: *req.CategoryID,
	})
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateCategoryRow](
			s.logger,
			category_errors.ErrFailedFindCategoryByName,
			method,
			span,
			zap.Int("category_id", *req.CategoryID),
			zap.String("name", req.Name))
	}

	category, err := s.categoryRepository.UpdateCategory(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateCategoryRow](
			s.logger,
			category_errors.ErrFailedUpdateCategory,
			method,
			span,
			zap.Any("request", req))
	}

	s.cache.DeleteCachedCategoryCache(ctx, int(category.CategoryID))

	logSuccess("Successfully updated category",
		zap.Int("category_id", int(category.CategoryID)),
		zap.String("name", req.Name))

	return category, nil
}

func (s *categoryService) TrashedCategory(ctx context.Context, category_id int) (*db.Category, error) {
	const method = "TrashedCategory"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("category_id", category_id))

	defer func() {
		end(status)
	}()

	category, err := s.categoryRepository.TrashedCategory(ctx, category_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Category](
			s.logger,
			category_errors.ErrFailedTrashedCategory,
			method,
			span,
			zap.Int("category_id", category_id))
	}

	s.cache.DeleteCachedCategoryCache(ctx, category_id)

	logSuccess("Successfully trashed category",
		zap.Int("category_id", category_id))

	return category, nil
}

func (s *categoryService) RestoreCategory(ctx context.Context, categoryID int) (*db.Category, error) {
	const method = "RestoreCategory"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("category_id", categoryID))

	defer func() {
		end(status)
	}()

	category, err := s.categoryRepository.RestoreCategory(ctx, categoryID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Category](
			s.logger,
			category_errors.ErrFailedRestoreCategory,
			method,
			span,
			zap.Int("category_id", categoryID))
	}

	s.cache.DeleteCachedCategoryCache(ctx, categoryID)

	logSuccess("Successfully restored category",
		zap.Int("category_id", categoryID))

	return category, nil
}

func (s *categoryService) DeleteCategoryPermanently(ctx context.Context, categoryID int) (bool, error) {
	const method = "DeleteCategoryPermanently"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("category_id", categoryID))

	defer func() {
		end(status)
	}()

	success, err := s.categoryRepository.DeleteCategoryPermanently(ctx, categoryID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			category_errors.ErrFailedDeleteCategoryPermanent,
			method,
			span,
			zap.Int("category_id", categoryID))
	}

	s.cache.DeleteCachedCategoryCache(ctx, categoryID)

	logSuccess("Successfully permanently deleted category",
		zap.Int("category_id", categoryID))

	return success, nil
}

func (s *categoryService) RestoreAllCategories(ctx context.Context) (bool, error) {
	const method = "RestoreAllCategories"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.categoryRepository.RestoreAllCategories(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			category_errors.ErrFailedRestoreAllCategories,
			method,
			span)
	}

	s.logger.Debug("All category caches should be invalidated after restore all operation")

	logSuccess("Successfully restored all trashed categories")

	return success, nil
}

func (s *categoryService) DeleteAllPermanentCategories(ctx context.Context) (bool, error) {
	const method = "DeleteAllPermanentCategories"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.categoryRepository.DeleteAllPermanentCategories(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			category_errors.ErrFailedDeleteAllCategoriesPermanent,
			method,
			span)
	}

	s.logger.Debug("All category caches should be invalidated after delete all operation")

	logSuccess("Successfully permanently deleted all trashed categories")

	return success, nil
}
