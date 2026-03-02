package service

import (
	"context"
	cashier_cache "pointofsale/internal/cache/cashier"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/errorhandler"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/cashier_errors"
	"pointofsale/pkg/errors/merchant_errors"
	"pointofsale/pkg/errors/user_errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

type cashierService struct {
	merchantRepository repository.MerchantRepository
	userRepository     repository.UserRepository
	cashierRepository  repository.CashierRepository
	logger             logger.LoggerInterface
	observability      observability.TraceLoggerObservability
	cache              cashier_cache.CashierMencache
}

type CashierServiceDeps struct {
	MerchantRepo  repository.MerchantRepository
	UserRepo      repository.UserRepository
	CashierRepo   repository.CashierRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
	Cache         cashier_cache.CashierMencache
}

func NewCashierService(deps CashierServiceDeps) *cashierService {
	return &cashierService{
		merchantRepository: deps.MerchantRepo,
		userRepository:     deps.UserRepo,
		cashierRepository:  deps.CashierRepo,
		logger:             deps.Logger,
		observability:      deps.Observability,
		cache:              deps.Cache,
	}
}

func (s *cashierService) FindAllCashiers(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersRow, *int, error) {
	const method = "FindAllCashiers"

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

	if data, total, found := s.cache.GetCachedCashiersCache(ctx, req); found {
		logSuccess("Successfully retrieved all cashier records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	cashiers, err := s.cashierRepository.FindAllCashiers(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCashiersRow](
			s.logger,
			cashier_errors.ErrFailedFindAllCashiers,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(cashiers) > 0 {
		totalCount = int(cashiers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCashiersCache(ctx, req, cashiers, &totalCount)

	logSuccess("Successfully fetched cashiers",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return cashiers, &totalCount, nil
}

func (s *cashierService) FindById(ctx context.Context, cashierID int) (*db.GetCashierByIdRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("cashier_id", cashierID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedCashier(ctx, cashierID); found {
		logSuccess("Successfully retrieved cashier record from cache",
			zap.Int("cashier_id", cashierID))
		return data, nil
	}

	cashier, err := s.cashierRepository.FindById(ctx, cashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetCashierByIdRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierById,
			method,
			span,
			zap.Int("cashier_id", cashierID))
	}

	s.cache.SetCachedCashier(ctx, cashier)

	logSuccess("Successfully fetched cashier",
		zap.Int("cashier_id", cashierID))

	return cashier, nil
}

func (s *cashierService) FindByActive(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersActiveRow, *int, error) {
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

	if data, total, found := s.cache.GetCachedCashiersActive(ctx, req); found {
		logSuccess("Successfully retrieved active cashier records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	cashiers, err := s.cashierRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCashiersActiveRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierByActive,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(cashiers) > 0 {
		totalCount = int(cashiers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCashiersActive(ctx, req, cashiers, &totalCount)

	logSuccess("Successfully fetched active cashiers",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return cashiers, &totalCount, nil
}

func (s *cashierService) FindByTrashed(ctx context.Context, req *requests.FindAllCashiers) ([]*db.GetCashiersTrashedRow, *int, error) {
	const method = "FindByTrashed"

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", req.Page),
		attribute.Int("pageSize", req.PageSize),
		attribute.String("search", req.Search))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedCashiersTrashed(ctx, req); found {
		logSuccess("Successfully retrieved trashed cashier records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize))
		return data, total, nil
	}

	cashiers, err := s.cashierRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCashiersTrashedRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierByTrashed,
			method,
			span,
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize))
	}

	var totalCount int

	if len(cashiers) > 0 {
		totalCount = int(cashiers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCashiersTrashed(ctx, req, cashiers, &totalCount)

	logSuccess("Successfully fetched trashed cashiers",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return cashiers, &totalCount, nil
}

func (s *cashierService) FindByMerchant(ctx context.Context, req *requests.FindAllCashierMerchant) ([]*db.GetCashiersByMerchantRow, *int, error) {
	const method = "FindByMerchant"

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", req.Page),
		attribute.Int("pageSize", req.PageSize),
		attribute.String("search", req.Search),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedCashiersByMerchant(ctx, req); found {
		logSuccess("Successfully retrieved merchant cashier records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.Int("merchant_id", req.MerchantID))
		return data, total, nil
	}

	cashiers, err := s.cashierRepository.FindByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetCashiersByMerchantRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierByMerchant,
			method,
			span,
			zap.Int("merchant_id", req.MerchantID),
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize))
	}

	var totalCount int

	if len(cashiers) > 0 {
		totalCount = int(cashiers[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedCashiersByMerchant(ctx, req, cashiers, &totalCount)

	logSuccess("Successfully fetched merchant cashiers",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize),
		zap.Int("merchant_id", req.MerchantID))

	return cashiers, &totalCount, nil
}

func (s *cashierService) FindMonthlyTotalSales(ctx context.Context, req *requests.MonthTotalSales) ([]*db.GetMonthlyTotalSalesCashierRow, error) {
	const method = "FindMonthlyTotalSales"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTotalSalesCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total sales from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.cashierRepository.GetMonthlyTotalSales(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalSalesCashierRow](
			s.logger,
			cashier_errors.ErrFailedFindMonthlyTotalSales,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
	}

	s.cache.SetMonthlyTotalSalesCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total sales",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindYearlyTotalSales(ctx context.Context, year int) ([]*db.GetYearlyTotalSalesCashierRow, error) {
	const method = "FindYearlyTotalSales"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTotalSalesCache(ctx, year); found {
		logSuccess("Successfully retrieved yearly total sales from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.cashierRepository.GetYearlyTotalSales(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalSalesCashierRow](
			s.logger,
			cashier_errors.ErrFailedFindYearlyTotalSales,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetYearlyTotalSalesCache(ctx, year, res)

	logSuccess("Successfully fetched yearly total sales",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindMonthlyTotalSalesById(ctx context.Context, req *requests.MonthTotalSalesCashier) ([]*db.GetMonthlyTotalSalesByIdRow, error) {
	const method = "FindMonthlyTotalSalesById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("cashier_id", req.CashierID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTotalSalesByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total sales by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("cashier_id", req.CashierID))
		return data, nil
	}

	res, err := s.cashierRepository.GetMonthlyTotalSalesById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalSalesByIdRow](
			s.logger,
			cashier_errors.ErrFailedFindMonthlyTotalSalesById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("cashier_id", req.CashierID))
	}

	s.cache.SetMonthlyTotalSalesByIdCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total sales by ID",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("cashier_id", req.CashierID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindYearlyTotalSalesById(ctx context.Context, req *requests.YearTotalSalesCashier) ([]*db.GetYearlyTotalSalesByIdRow, error) {
	const method = "FindYearlyTotalSalesById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("cashier_id", req.CashierID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTotalSalesByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly total sales by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("cashier_id", req.CashierID))
		return data, nil
	}

	res, err := s.cashierRepository.GetYearlyTotalSalesById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalSalesByIdRow](
			s.logger,
			cashier_errors.ErrFailedFindYearlyTotalSalesById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("cashier_id", req.CashierID))
	}

	s.cache.SetYearlyTotalSalesByIdCache(ctx, req, res)

	logSuccess("Successfully fetched yearly total sales by ID",
		zap.Int("year", req.Year),
		zap.Int("cashier_id", req.CashierID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindMonthlyTotalSalesByMerchant(ctx context.Context, req *requests.MonthTotalSalesMerchant) ([]*db.GetMonthlyTotalSalesByMerchantRow, error) {
	const method = "FindMonthlyTotalSalesByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyTotalSalesByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly total sales by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.cashierRepository.GetMonthlyTotalSalesByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTotalSalesByMerchantRow](
			s.logger,
			cashier_errors.ErrFailedFindMonthlyTotalSalesByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetMonthlyTotalSalesByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched monthly total sales by merchant",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindYearlyTotalSalesByMerchant(ctx context.Context, req *requests.YearTotalSalesMerchant) ([]*db.GetYearlyTotalSalesByMerchantRow, error) {
	const method = "FindYearlyTotalSalesByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyTotalSalesByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly total sales by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.cashierRepository.GetYearlyTotalSalesByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTotalSalesByMerchantRow](
			s.logger,
			cashier_errors.ErrFailedFindYearlyTotalSalesByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetYearlyTotalSalesByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched yearly total sales by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindMonthyCashier(ctx context.Context, year int) ([]*db.GetMonthlyCashierRow, error) {
	const method = "FindMonthyCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlySalesCache(ctx, year); found {
		logSuccess("Successfully retrieved monthly cashier from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.cashierRepository.GetMonthyCashier(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyCashierRow](
			s.logger,
			cashier_errors.ErrFailedFindMonthlySales,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetMonthlySalesCache(ctx, year, res)

	logSuccess("Successfully fetched monthly cashier",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindYearlyCashier(ctx context.Context, year int) ([]*db.GetYearlyCashierRow, error) {
	const method = "FindYearlyCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlySalesCache(ctx, year); found {
		logSuccess("Successfully retrieved yearly cashier from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.cashierRepository.GetYearlyCashier(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyCashierRow](
			s.logger,
			cashier_errors.ErrFailedFindYearlySales,
			method,
			span,
			zap.Int("year", year))
	}

	s.cache.SetYearlySalesCache(ctx, year, res)

	logSuccess("Successfully fetched yearly cashier",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindMonthlyCashierByMerchant(ctx context.Context, req *requests.MonthCashierMerchant) ([]*db.GetMonthlyCashierByMerchantRow, error) {
	const method = "FindMonthlyCashierByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyCashierByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly cashier by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.cashierRepository.GetMonthlyCashierByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyCashierByMerchantRow](
			s.logger,
			cashier_errors.ErrFailedFindMonthlyCashierByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetMonthlyCashierByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched monthly cashier by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindYearlyCashierByMerchant(ctx context.Context, req *requests.YearCashierMerchant) ([]*db.GetYearlyCashierByMerchantRow, error) {
	const method = "FindYearlyCashierByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchant_id", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyCashierByMerchantCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly cashier by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
		return data, nil
	}

	res, err := s.cashierRepository.GetYearlyCashierByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyCashierByMerchantRow](
			s.logger,
			cashier_errors.ErrFailedFindYearlyCashierByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchant_id", req.MerchantID))
	}

	s.cache.SetYearlyCashierByMerchantCache(ctx, req, res)

	logSuccess("Successfully fetched yearly cashier by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchant_id", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindMonthlyCashierById(ctx context.Context, req *requests.MonthCashierId) ([]*db.GetMonthlyCashierByCashierIdRow, error) {
	const method = "FindMonthlyCashierById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("cashier_id", req.CashierID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetMonthlyCashierByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved monthly cashier by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("cashier_id", req.CashierID))
		return data, nil
	}

	res, err := s.cashierRepository.GetMonthlyCashierById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyCashierByCashierIdRow](
			s.logger,
			cashier_errors.ErrFailedFindMonthlyCashierById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("cashier_id", req.CashierID))
	}

	s.cache.SetMonthlyCashierByIdCache(ctx, req, res)

	logSuccess("Successfully fetched monthly cashier by ID",
		zap.Int("year", req.Year),
		zap.Int("cashier_id", req.CashierID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) FindYearlyCashierById(ctx context.Context, req *requests.YearCashierId) ([]*db.GetYearlyCashierByCashierIdRow, error) {
	const method = "FindYearlyCashierById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("cashier_id", req.CashierID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetYearlyCashierByIdCache(ctx, req); found {
		logSuccess("Successfully retrieved yearly cashier by ID from cache",
			zap.Int("year", req.Year),
			zap.Int("cashier_id", req.CashierID))
		return data, nil
	}

	res, err := s.cashierRepository.GetYearlyCashierById(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyCashierByCashierIdRow](
			s.logger,
			cashier_errors.ErrFailedFindYearlyCashierById,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("cashier_id", req.CashierID))
	}

	s.cache.SetYearlyCashierByIdCache(ctx, req, res)

	logSuccess("Successfully fetched yearly cashier by ID",
		zap.Int("year", req.Year),
		zap.Int("cashier_id", req.CashierID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *cashierService) CreateCashier(ctx context.Context, req *requests.CreateCashierRequest) (*db.CreateCashierRow, error) {
	const method = "CreateCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchantID", req.MerchantID),
		attribute.Int("userID", req.UserID))

	defer func() {
		end(status)
	}()

	_, err := s.merchantRepository.FindById(ctx, req.MerchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateCashierRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchant_id", req.MerchantID))
	}

	_, err = s.userRepository.FindById(ctx, req.UserID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateCashierRow](
			s.logger,
			user_errors.ErrFailedFindUserByID,
			method,
			span,
			zap.Int("user_id", req.UserID))
	}

	cashier, err := s.cashierRepository.CreateCashier(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateCashierRow](
			s.logger,
			cashier_errors.ErrFailedCreateCashier,
			method,
			span,
			zap.Any("request", req))
	}

	s.cache.DeleteCashierCache(ctx, int(cashier.CashierID))

	logSuccess("Successfully created cashier",
		zap.Int("cashier_id", int(cashier.CashierID)),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("userID", req.UserID))

	return cashier, nil
}

func (s *cashierService) UpdateCashier(ctx context.Context, req *requests.UpdateCashierRequest) (*db.UpdateCashierRow, error) {
	const method = "UpdateCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("cashier_id", *req.CashierID))

	defer func() {
		end(status)
	}()

	cashier, err := s.cashierRepository.UpdateCashier(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateCashierRow](
			s.logger,
			cashier_errors.ErrFailedUpdateCashier,
			method,
			span,
			zap.Int("cashier_id", *req.CashierID))
	}

	s.cache.DeleteCashierCache(ctx, int(cashier.CashierID))

	logSuccess("Successfully updated cashier",
		zap.Int("cashier_id", int(cashier.CashierID)))

	return cashier, nil
}

func (s *cashierService) TrashedCashier(ctx context.Context, cashierID int) (*db.Cashier, error) {
	const method = "TrashedCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("cashier_id", cashierID))

	defer func() {
		end(status)
	}()

	cashier, err := s.cashierRepository.TrashedCashier(ctx, cashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Cashier](
			s.logger,
			cashier_errors.ErrFailedTrashedCashier,
			method,
			span,
			zap.Int("cashier_id", cashierID))
	}

	s.cache.DeleteCashierCache(ctx, cashierID)

	logSuccess("Successfully trashed cashier",
		zap.Int("cashier_id", cashierID))

	return cashier, nil
}

func (s *cashierService) RestoreCashier(ctx context.Context, cashierID int) (*db.Cashier, error) {
	const method = "RestoreCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("cashier_id", cashierID))

	defer func() {
		end(status)
	}()

	cashier, err := s.cashierRepository.RestoreCashier(ctx, cashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Cashier](
			s.logger,
			cashier_errors.ErrFailedRestoreCashier,
			method,
			span,
			zap.Int("cashier_id", cashierID))
	}

	s.cache.DeleteCashierCache(ctx, cashierID)

	logSuccess("Successfully restored cashier",
		zap.Int("cashier_id", cashierID))

	return cashier, nil
}

func (s *cashierService) DeleteCashierPermanent(ctx context.Context, cashierID int) (bool, error) {
	const method = "DeleteCashierPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("cashier_id", cashierID))

	defer func() {
		end(status)
	}()

	success, err := s.cashierRepository.DeleteCashierPermanent(ctx, cashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			cashier_errors.ErrFailedDeleteCashierPermanent,
			method,
			span,
			zap.Int("cashier_id", cashierID))
	}

	s.cache.DeleteCashierCache(ctx, cashierID)

	logSuccess("Successfully permanently deleted cashier",
		zap.Int("cashier_id", cashierID))

	return success, nil
}

func (s *cashierService) RestoreAllCashier(ctx context.Context) (bool, error) {
	const method = "RestoreAllCashier"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.cashierRepository.RestoreAllCashier(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			cashier_errors.ErrFailedRestoreAllCashiers,
			method,
			span)
	}

	s.logger.Debug("All cashier caches should be invalidated after restore all operation")

	logSuccess("Successfully restored all trashed cashiers")

	return success, nil
}

func (s *cashierService) DeleteAllCashierPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllCashierPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.cashierRepository.DeleteAllCashierPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			cashier_errors.ErrFailedDeleteAllCashierPermanent,
			method,
			span)
	}

	s.logger.Debug("All cashier caches should be invalidated after delete all operation")

	logSuccess("Successfully permanently deleted all trashed cashiers")

	return success, nil
}
