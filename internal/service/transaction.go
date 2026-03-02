package service

import (
	"context"
	transaction_cache "pointofsale/internal/cache/transaction"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/errorhandler"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/cashier_errors"
	"pointofsale/pkg/errors/merchant_errors"
	"pointofsale/pkg/errors/order_errors"
	orderitem_errors "pointofsale/pkg/errors/order_item_errors"
	"pointofsale/pkg/errors/transaction_errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

type transactionService struct {
	cashierRepository     repository.CashierRepository
	merchantRepository    repository.MerchantRepository
	transactionRepository repository.TransactionRepository
	orderRepository       repository.OrderRepository
	orderItemRepository   repository.OrderItemRepository
	logger                logger.LoggerInterface
	observability         observability.TraceLoggerObservability
	cache                 transaction_cache.TransactionMencache
}

type TransactionServiceDeps struct {
	CashierRepo     repository.CashierRepository
	MerchantRepo    repository.MerchantRepository
	TransactionRepo repository.TransactionRepository
	OrderRepo       repository.OrderRepository
	OrderItemRepo   repository.OrderItemRepository
	Logger          logger.LoggerInterface
	Observability   observability.TraceLoggerObservability
	Cache           transaction_cache.TransactionMencache
}

func NewTransactionService(deps TransactionServiceDeps) *transactionService {
	return &transactionService{
		cashierRepository:     deps.CashierRepo,
		merchantRepository:    deps.MerchantRepo,
		transactionRepository: deps.TransactionRepo,
		orderRepository:       deps.OrderRepo,
		orderItemRepository:   deps.OrderItemRepo,
		logger:                deps.Logger,
		cache:                 deps.Cache,
		observability:         deps.Observability,
	}
}

func (s *transactionService) FindAllTransactions(ctx context.Context, req *requests.FindAllTransaction) ([]*db.GetTransactionsRow, *int, error) {
	const method = "FindAllTransactions"

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

	if data, total, found := s.cache.GetCachedTransactionsCache(ctx, req); found {
		logSuccess("Successfully retrieved all transaction records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionRepository.FindAllTransactions(ctx, req)

	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransactionsRow](
			s.logger,
			transaction_errors.ErrFailedFindAllTransactions,
			method,
			span,
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("page_size", req.PageSize))
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionsCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched transactions",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return transactions, &totalCount, nil
}

func (s *transactionService) FindByMerchant(ctx context.Context, req *requests.FindAllTransactionByMerchant) ([]*db.GetTransactionByMerchantRow, *int, error) {
	const method = "FindByMerchant"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search
	merchantId := req.MerchantID

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search),
		attribute.Int("merchant_id", merchantId))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedTransactionByMerchant(ctx, req); found {
		logSuccess("Successfully retrieved merchant transaction records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.Int("merchant_id", merchantId))
		return data, total, nil
	}

	transactions, err := s.transactionRepository.FindByMerchant(ctx, req)

	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransactionByMerchantRow](
			s.logger,
			transaction_errors.ErrFailedFindTransactionsByMerchant,
			method,
			span,
			zap.Int("merchant_id", merchantId),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("page_size", pageSize))
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionByMerchant(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched merchant transactions",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.Int("merchant_id", merchantId))

	return transactions, &totalCount, nil
}

func (s *transactionService) FindByActive(ctx context.Context, req *requests.FindAllTransaction) ([]*db.GetTransactionsActiveRow, *int, error) {
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

	// Check cache first
	if data, total, found := s.cache.GetCachedTransactionActiveCache(ctx, req); found {
		logSuccess("Successfully retrieved active transaction records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionRepository.FindByActive(ctx, req)

	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransactionsActiveRow](
			s.logger,
			transaction_errors.ErrFailedFindTransactionsByActive,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("page_size", pageSize))
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionActiveCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched active transactions",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return transactions, &totalCount, nil
}

func (s *transactionService) FindByTrashed(ctx context.Context, req *requests.FindAllTransaction) ([]*db.GetTransactionsTrashedRow, *int, error) {
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

	if data, total, found := s.cache.GetCachedTransactionTrashedCache(ctx, req); found {
		logSuccess("Successfully retrieved trashed transaction records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	transactions, err := s.transactionRepository.FindByTrashed(ctx, req)

	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetTransactionsTrashedRow](
			s.logger,
			transaction_errors.ErrFailedFindTransactionsByTrashed,
			method,
			span,
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("page_size", req.PageSize))
	}

	var totalCount int

	if len(transactions) > 0 {
		totalCount = int(transactions[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedTransactionTrashedCache(ctx, req, transactions, &totalCount)

	logSuccess("Successfully fetched trashed transactions",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return transactions, &totalCount, nil
}

func (s *transactionService) FindById(ctx context.Context, transactionID int) (*db.GetTransactionByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transactionID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransactionCache(ctx, transactionID); found {
		logSuccess("Successfully retrieved transaction record from cache",
			zap.Int("transaction_id", transactionID))
		return data, nil
	}

	transaction, err := s.transactionRepository.FindById(ctx, transactionID)

	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetTransactionByIDRow](
			s.logger,
			transaction_errors.ErrFailedFindTransactionById,
			method,
			span,
			zap.Int("transaction_id", transactionID))
	}

	s.cache.SetCachedTransactionCache(ctx, transaction)

	logSuccess("Successfully fetched transaction",
		zap.Int("transaction_id", transactionID))

	return transaction, nil
}

func (s *transactionService) FindByOrderId(ctx context.Context, orderID int) (*db.GetTransactionByOrderIDRow, error) {
	const method = "FindByOrderId"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("order_id", orderID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedTransactionByOrderId(ctx, orderID); found {
		logSuccess("Successfully retrieved transaction by order ID from cache",
			zap.Int("order_id", orderID))
		return data, nil
	}

	transaction, err := s.transactionRepository.FindByOrderId(ctx, orderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetTransactionByOrderIDRow](
			s.logger,
			transaction_errors.ErrFailedFindTransactionByOrderId,
			method,
			span,
			zap.Int("order_id", orderID))
	}

	s.cache.SetCachedTransactionByOrderId(ctx, orderID, transaction)

	logSuccess("Successfully fetched transaction by order ID",
		zap.Int("order_id", orderID))

	return transaction, nil
}

func (s *transactionService) FindMonthlyAmountSuccess(ctx context.Context, req *requests.MonthAmountTransaction) ([]*db.GetMonthlyAmountTransactionSuccessRow, error) {
	const method = "FindMonthlyAmountSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthAmountSuccessCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly successful transaction amounts from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyAmountSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountTransactionSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyAmountSuccess,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Error(err))
	}

	s.cache.SetCachedMonthAmountSuccessCached(ctx, req, res)

	logSuccess("Successfully fetched monthly successful transaction amounts",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyAmountSuccess(ctx context.Context, year int) ([]*db.GetYearlyAmountTransactionSuccessRow, error) {
	const method = "FindYearlyAmountSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearAmountSuccessCached(ctx, year); found {
		logSuccess("Successfully retrieved yearly successful transaction amounts from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyAmountSuccess(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountTransactionSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyAmountSuccess,
			method,
			span,
			zap.Int("year", year),
			zap.Error(err))
	}

	s.cache.SetCachedYearAmountSuccessCached(ctx, year, res)

	logSuccess("Successfully fetched yearly successful transaction amounts",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyAmountFailed(ctx context.Context, req *requests.MonthAmountTransaction) ([]*db.GetMonthlyAmountTransactionFailedRow, error) {
	const method = "FindMonthlyAmountFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthAmountFailedCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly failed transaction amounts from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyAmountFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountTransactionFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyAmountFailed,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Error(err))
	}

	s.cache.SetCachedMonthAmountFailedCached(ctx, req, res)

	logSuccess("Successfully fetched monthly failed transaction amounts",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyAmountFailed(ctx context.Context, year int) ([]*db.GetYearlyAmountTransactionFailedRow, error) {
	const method = "FindYearlyAmountFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearAmountFailedCached(ctx, year); found {
		logSuccess("Successfully retrieved yearly failed transaction amounts from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyAmountFailed(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountTransactionFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyAmountFailed,
			method,
			span,
			zap.Int("year", year),
			zap.Error(err))
	}

	s.cache.SetCachedYearAmountFailedCached(ctx, year, res)

	logSuccess("Successfully fetched yearly failed transaction amounts",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyMethodSuccess(ctx context.Context, req *requests.MonthMethodTransaction) ([]*db.GetMonthlyTransactionMethodsSuccessRow, error) {
	const method = "FindMonthlyMethodSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthMethodSuccessCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly successful transaction methods from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyTransactionMethodSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransactionMethodsSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyMethod,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Error(err))
	}

	s.cache.SetCachedMonthMethodSuccessCached(ctx, req, res)

	logSuccess("Successfully fetched monthly successful transaction methods",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyMethodSuccess(ctx context.Context, year int) ([]*db.GetYearlyTransactionMethodsSuccessRow, error) {
	const method = "FindYearlyMethodSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearMethodSuccessCached(ctx, year); found {
		logSuccess("Successfully retrieved yearly successful transaction methods from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyTransactionMethodSuccess(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionMethodsSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyMethod,
			method,
			span,
			zap.Int("year", year),
			zap.Error(err))
	}

	s.cache.SetCachedYearMethodSuccessCached(ctx, year, res)

	logSuccess("Successfully fetched yearly successful transaction methods",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyMethodFailed(ctx context.Context, req *requests.MonthMethodTransaction) ([]*db.GetMonthlyTransactionMethodsFailedRow, error) {
	const method = "FindMonthlyMethodFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthMethodFailedCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly failed transaction methods from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyTransactionMethodFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransactionMethodsFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyMethod,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Error(err))
	}

	s.cache.SetCachedMonthMethodFailedCached(ctx, req, res)

	logSuccess("Successfully fetched monthly failed transaction methods",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyMethodFailed(ctx context.Context, year int) ([]*db.GetYearlyTransactionMethodsFailedRow, error) {
	const method = "FindYearlyMethodFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", year))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearMethodFailedCached(ctx, year); found {
		logSuccess("Successfully retrieved yearly failed transaction methods from cache",
			zap.Int("year", year))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyTransactionMethodFailed(ctx, year)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionMethodsFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyMethod,
			method,
			span,
			zap.Int("year", year),
			zap.Error(err))
	}

	s.cache.SetCachedYearMethodFailedCached(ctx, year, res)

	logSuccess("Successfully fetched yearly failed transaction methods",
		zap.Int("year", year),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyAmountSuccessByMerchant(ctx context.Context, req *requests.MonthAmountTransactionMerchant) ([]*db.GetMonthlyAmountTransactionSuccessByMerchantRow, error) {
	const method = "FindMonthlyAmountSuccessByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthAmountSuccessByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly successful transaction amounts by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyAmountSuccessByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountTransactionSuccessByMerchantRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyAmountSuccessByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedMonthAmountSuccessByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched monthly successful transaction amounts by merchant",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyAmountSuccessByMerchant(ctx context.Context, req *requests.YearAmountTransactionMerchant) ([]*db.GetYearlyAmountTransactionSuccessByMerchantRow, error) {
	const method = "FindYearlyAmountSuccessByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearAmountSuccessByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved yearly successful transaction amounts by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyAmountSuccessByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountTransactionSuccessByMerchantRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyAmountSuccessByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedYearAmountSuccessByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched yearly successful transaction amounts by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyAmountFailedByMerchant(ctx context.Context, req *requests.MonthAmountTransactionMerchant) ([]*db.GetMonthlyAmountTransactionFailedByMerchantRow, error) {
	const method = "FindMonthlyAmountFailedByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("month", req.Month),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthAmountFailedByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly failed transaction amounts by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyAmountFailedByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyAmountTransactionFailedByMerchantRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyAmountFailedByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("month", req.Month),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedMonthAmountFailedByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched monthly failed transaction amounts by merchant",
		zap.Int("year", req.Year),
		zap.Int("month", req.Month),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyAmountFailedByMerchant(ctx context.Context, req *requests.YearAmountTransactionMerchant) ([]*db.GetYearlyAmountTransactionFailedByMerchantRow, error) {
	const method = "FindYearlyAmountFailedByMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearAmountFailedByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved yearly failed transaction amounts by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyAmountFailedByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyAmountTransactionFailedByMerchantRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyAmountFailedByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedYearAmountFailedByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched yearly failed transaction amounts by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyMethodByMerchantSuccess(ctx context.Context, req *requests.MonthMethodTransactionMerchant) ([]*db.GetMonthlyTransactionMethodsByMerchantSuccessRow, error) {
	const method = "FindMonthlyMethodByMerchantSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthMethodSuccessByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly successful transaction methods by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyTransactionMethodByMerchantSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransactionMethodsByMerchantSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyMethodByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedMonthMethodSuccessByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched monthly successful transaction methods by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyMethodByMerchantSuccess(ctx context.Context, req *requests.YearMethodTransactionMerchant) ([]*db.GetYearlyTransactionMethodsByMerchantSuccessRow, error) {
	const method = "FindYearlyMethodByMerchantSuccess"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearMethodSuccessByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved yearly successful transaction methods by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyTransactionMethodByMerchantSuccess(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionMethodsByMerchantSuccessRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyMethodByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedYearMethodSuccessByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched yearly successful transaction methods by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindMonthlyMethodByMerchantFailed(ctx context.Context, req *requests.MonthMethodTransactionMerchant) ([]*db.GetMonthlyTransactionMethodsByMerchantFailedRow, error) {
	const method = "FindMonthlyMethodByMerchantFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedMonthMethodFailedByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved monthly failed transaction methods by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetMonthlyTransactionMethodByMerchantFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetMonthlyTransactionMethodsByMerchantFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindMonthlyMethodByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedMonthMethodFailedByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched monthly failed transaction methods by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) FindYearlyMethodByMerchantFailed(ctx context.Context, req *requests.YearMethodTransactionMerchant) ([]*db.GetYearlyTransactionMethodsByMerchantFailedRow, error) {
	const method = "FindYearlyMethodByMerchantFailed"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("year", req.Year),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedYearMethodFailedByMerchantCached(ctx, req); found {
		logSuccess("Successfully retrieved yearly failed transaction methods by merchant from cache",
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID))
		return data, nil
	}

	res, err := s.transactionRepository.GetYearlyTransactionMethodByMerchantFailed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[[]*db.GetYearlyTransactionMethodsByMerchantFailedRow](
			s.logger,
			transaction_errors.ErrFailedFindYearlyMethodByMerchant,
			method,
			span,
			zap.Int("year", req.Year),
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))
	}

	s.cache.SetCachedYearMethodFailedByMerchantCached(ctx, req, res)

	logSuccess("Successfully fetched yearly failed transaction methods by merchant",
		zap.Int("year", req.Year),
		zap.Int("merchantID", req.MerchantID),
		zap.Int("count", len(res)))

	return res, nil
}

func (s *transactionService) CreateTransaction(ctx context.Context, req *requests.CreateTransactionRequest) (*db.CreateTransactionRow, error) {
	const method = "CreateTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("orderID", req.OrderID),
		attribute.Int("cashierID", req.CashierID))

	defer func() {
		end(status)
	}()

	cashier, err := s.cashierRepository.FindById(ctx, req.CashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierById,
			method,
			span,
			zap.Int("cashierId", req.CashierID),
			zap.Error(err))
	}

	_, err = s.merchantRepository.FindById(ctx, int(cashier.MerchantID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchantId", int(cashier.MerchantID)),
			zap.Error(err))
	}

	req.MerchantID = int(cashier.MerchantID)

	_, err = s.orderRepository.FindById(ctx, req.OrderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			order_errors.ErrFailedFindOrderById,
			method,
			span,
			zap.Int("orderID", req.OrderID),
			zap.Error(err))
	}

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(ctx, req.OrderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			orderitem_errors.ErrFailedFindOrderItemByOrder,
			method,
			span,
			zap.Int("orderID", req.OrderID),
			zap.Error(err))
	}

	if len(orderItems) == 0 {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			orderitem_errors.ErrFailedOrderItemEmpty,
			method,
			span,
			zap.Int("orderID", req.OrderID))
	}

	var totalAmount int32
	for _, item := range orderItems {
		if item.Quantity <= 0 {
			status = "error"
			return errorhandler.HandleError[*db.CreateTransactionRow](
				s.logger,
				orderitem_errors.ErrFailedFindOrderItemByOrder,
				method,
				span,
				zap.Int("orderID", req.OrderID),
				zap.Int("itemID", int(item.OrderItemID)),
				zap.Int("quantity", int(item.Quantity)))
		}
		totalAmount += item.Price * item.Quantity
	}

	ppn := totalAmount * 11 / 100
	totalAmountWithTax := totalAmount + ppn

	if req.Amount < int(totalAmountWithTax) {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			transaction_errors.ErrFailedPaymentInsufficientBalance,
			method,
			span,
			zap.Int("paid", req.Amount),
			zap.Int("required", int(totalAmountWithTax)))
	}

	changeAmount := req.Amount - int(totalAmountWithTax)
	paymentStatus := "success"

	req.PaymentStatus = &paymentStatus
	req.ChangeAmount = &changeAmount
	req.Amount = int(totalAmountWithTax)

	transaction, err := s.transactionRepository.CreateTransaction(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateTransactionRow](
			s.logger,
			transaction_errors.ErrFailedCreateTransaction,
			method,
			span,
			zap.Error(err))
	}

	s.cache.DeleteTransactionCache(ctx, int(transaction.TransactionID))

	logSuccess("Successfully created transaction",
		zap.Int("transactionID", int(transaction.TransactionID)),
		zap.Int("orderID", req.OrderID),
		zap.Int("amount", req.Amount),
		zap.Int("changeAmount", changeAmount))

	return transaction, nil
}

func (s *transactionService) UpdateTransaction(ctx context.Context, req *requests.UpdateTransactionRequest) (*db.UpdateTransactionRow, error) {
	const method = "UpdateTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transactionID", *req.TransactionID))

	defer func() {
		end(status)
	}()

	cashier, err := s.cashierRepository.FindById(ctx, req.CashierID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			cashier_errors.ErrFailedFindCashierById,
			method,
			span,
			zap.Int("cashierId", req.CashierID),
			zap.Error(err))
	}

	existingTx, err := s.transactionRepository.FindById(ctx, *req.TransactionID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			transaction_errors.ErrFailedFindTransactionById,
			method,
			span,
			zap.Int("transactionID", *req.TransactionID),
			zap.Error(err))
	}

	if existingTx.PaymentStatus == "paid" || existingTx.PaymentStatus == "refunded" {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			transaction_errors.ErrFailedPaymentStatusCannotBeModified,
			method,
			span,
			zap.Int("transactionID", *req.TransactionID),
			zap.String("paymentStatus", existingTx.PaymentStatus))
	}

	_, err = s.merchantRepository.FindById(ctx, int(cashier.MerchantID))
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchantId", int(cashier.MerchantID)),
			zap.Error(err))
	}

	req.MerchantID = int(cashier.MerchantID)

	_, err = s.orderRepository.FindById(ctx, req.OrderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			order_errors.ErrFailedFindOrderById,
			method,
			span,
			zap.Int("orderID", req.OrderID),
			zap.Error(err))
	}

	orderItems, err := s.orderItemRepository.FindOrderItemByOrder(ctx, req.OrderID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			orderitem_errors.ErrFailedFindOrderItemByOrder,
			method,
			span,
			zap.Int("orderID", req.OrderID),
			zap.Error(err))
	}

	var totalAmount int32
	for _, item := range orderItems {
		totalAmount += item.Price * item.Quantity
	}

	ppn := totalAmount * 11 / 100
	totalAmountWithTax := totalAmount + ppn

	var paymentStatus string
	if req.Amount >= int(totalAmountWithTax) {
		paymentStatus = "success"
	} else {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			transaction_errors.ErrFailedPaymentInsufficientBalance,
			method,
			span,
			zap.Int("paid", req.Amount),
			zap.Int("required", int(totalAmountWithTax)))
	}

	changeAmount := req.Amount - int(totalAmountWithTax)
	req.Amount = int(totalAmountWithTax)
	req.PaymentStatus = &paymentStatus
	req.ChangeAmount = &changeAmount

	transaction, err := s.transactionRepository.UpdateTransaction(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateTransactionRow](
			s.logger,
			transaction_errors.ErrFailedUpdateTransaction,
			method,
			span,
			zap.Error(err))
	}

	s.cache.DeleteTransactionCache(ctx, int(transaction.TransactionID))

	logSuccess("Successfully updated transaction",
		zap.Int("transactionID", int(transaction.TransactionID)),
		zap.String("paymentStatus", paymentStatus),
		zap.Int("amount", req.Amount),
		zap.Int("changeAmount", changeAmount))

	return transaction, nil
}

func (s *transactionService) TrashedTransaction(ctx context.Context, transaction_id int) (*db.Transaction, error) {
	const method = "TrashedTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transaction_id))

	defer func() {
		end(status)
	}()

	transaction, err := s.transactionRepository.TrashTransaction(ctx, transaction_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Transaction](
			s.logger,
			transaction_errors.ErrFailedTrashedTransaction,
			method,
			span,
			zap.Int("transaction_id", transaction_id),
			zap.Error(err))
	}

	s.cache.DeleteTransactionCache(ctx, transaction_id)

	logSuccess("Successfully trashed transaction", zap.Int("transaction_id", transaction_id))

	return transaction, nil
}

func (s *transactionService) RestoreTransaction(ctx context.Context, transaction_id int) (*db.Transaction, error) {
	const method = "RestoreTransaction"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transaction_id))

	defer func() {
		end(status)
	}()

	transaction, err := s.transactionRepository.RestoreTransaction(ctx, transaction_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Transaction](
			s.logger,
			transaction_errors.ErrFailedRestoreTransaction,
			method,
			span,
			zap.Int("transaction_id", transaction_id),
			zap.Error(err))
	}

	s.cache.DeleteTransactionCache(ctx, transaction_id)

	logSuccess("Successfully restored transaction", zap.Int("transaction_id", transaction_id))

	return transaction, nil
}

func (s *transactionService) DeleteTransactionPermanently(ctx context.Context, transactionID int) (bool, error) {
	const method = "DeleteTransactionPermanently"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("transaction_id", transactionID))

	defer func() {
		end(status)
	}()

	success, err := s.transactionRepository.DeleteTransactionPermanently(ctx, transactionID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			transaction_errors.ErrFailedDeleteTransactionPermanently,
			method,
			span,
			zap.Int("transaction_id", transactionID),
			zap.Error(err))
	}

	s.cache.DeleteTransactionCache(ctx, transactionID)

	logSuccess("Successfully permanently deleted transaction", zap.Int("transaction_id", transactionID))

	return success, nil
}

func (s *transactionService) RestoreAllTransactions(ctx context.Context) (bool, error) {
	const method = "RestoreAllTransactions"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.transactionRepository.RestoreAllTransactions(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			transaction_errors.ErrFailedRestoreAllTransactions,
			method,
			span,
			zap.Error(err))
	}

	logSuccess("Successfully restored all trashed transactions")

	return success, nil
}

func (s *transactionService) DeleteAllTransactionPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllTransactionPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.transactionRepository.DeleteAllTransactionPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			transaction_errors.ErrFailedDeleteAllTransactionPermanent,
			method,
			span,
			zap.Error(err))
	}

	logSuccess("Successfully permanently deleted all trashed transactions")

	return success, nil
}
