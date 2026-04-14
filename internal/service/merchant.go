package service

import (
	"context"
	merchant_cache "pointofsale/internal/cache/merchant"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/errorhandler"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/merchant_errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

type merchantService struct {
	merchantRepository repository.MerchantRepository
	logger             logger.LoggerInterface
	observability      observability.TraceLoggerObservability
	Cache              merchant_cache.MerchantMenCache
}

type MerchantServiceDeps struct {
	MerchantRepo  repository.MerchantRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
	Cache         merchant_cache.MerchantMenCache
}

func NewMerchantService(deps MerchantServiceDeps) *merchantService {
	return &merchantService{
		merchantRepository: deps.MerchantRepo,
		logger:             deps.Logger,
		observability:      deps.Observability,
		Cache:              deps.Cache,
	}
}

func (s *merchantService) FindAllMerchants(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsRow, *int, error) {
	const method = "FindAllMerchants"

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

	if data, total, found := s.Cache.GetCachedMerchants(ctx, req); found {
		logSuccess("Successfully retrieved all merchant records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	merchants, err := s.merchantRepository.FindAllMerchants(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetMerchantsRow](
			s.logger,
			merchant_errors.ErrFailedFindAllMerchants,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(merchants) > 0 {
		totalCount = int(merchants[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.Cache.SetCachedMerchants(ctx, req, merchants, &totalCount)

	logSuccess("Successfully fetched merchants",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return merchants, &totalCount, nil
}

func (s *merchantService) FindByActive(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsActiveRow, *int, error) {
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

	if data, total, found := s.Cache.GetCachedMerchantActive(ctx, req); found {
		logSuccess("Successfully retrieved active merchant records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	merchants, err := s.merchantRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetMerchantsActiveRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantsByActive,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(merchants) > 0 {
		totalCount = int(merchants[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.Cache.SetCachedMerchantActive(ctx, req, merchants, &totalCount)

	logSuccess("Successfully fetched active merchants",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return merchants, &totalCount, nil
}

func (s *merchantService) FindByTrashed(ctx context.Context, req *requests.FindAllMerchants) ([]*db.GetMerchantsTrashedRow, *int, error) {
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

	if data, total, found := s.Cache.GetCachedMerchantTrashed(ctx, req); found {
		logSuccess("Successfully retrieved trashed merchant records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	merchants, err := s.merchantRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetMerchantsTrashedRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantsByTrashed,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(merchants) > 0 {
		totalCount = int(merchants[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.Cache.SetCachedMerchantTrashed(ctx, req, merchants, &totalCount)

	logSuccess("Successfully fetched trashed merchants",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return merchants, &totalCount, nil
}

func (s *merchantService) FindById(ctx context.Context, merchantID int) (*db.GetMerchantByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", merchantID))

	defer func() {
		end(status)
	}()

	if data, found := s.Cache.GetCachedMerchant(ctx, merchantID); found {
		logSuccess("Successfully retrieved merchant record from cache",
			zap.Int("merchant_id", merchantID))
		return data, nil
	}

	merchant, err := s.merchantRepository.FindById(ctx, merchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetMerchantByIDRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchant_id", merchantID))
	}

	s.Cache.SetCachedMerchant(ctx, merchant)

	logSuccess("Successfully fetched merchant",
		zap.Int("merchant_id", merchantID))

	return merchant, nil
}

func (s *merchantService) CreateMerchant(ctx context.Context, req *requests.CreateMerchantRequest) (*db.CreateMerchantRow, error) {
	const method = "CreateMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("name", req.Name),
		attribute.String("email", req.ContactEmail))

	defer func() {
		end(status)
	}()

	merchant, err := s.merchantRepository.CreateMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateMerchantRow](
			s.logger,
			merchant_errors.ErrFailedCreateMerchant,
			method,
			span,
			zap.Any("request", req))
	}

	s.Cache.DeleteCachedMerchant(ctx, int(merchant.MerchantID))

	logSuccess("Successfully created merchant",
		zap.Int("merchant_id", int(merchant.MerchantID)),
		zap.String("name", req.Name))

	return merchant, nil
}

func (s *merchantService) UpdateMerchant(ctx context.Context, req *requests.UpdateMerchantRequest) (*db.UpdateMerchantRow, error) {
	const method = "UpdateMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", *req.MerchantID),
		attribute.String("name", req.Name))

	defer func() {
		end(status)
	}()

	_, err := s.merchantRepository.FindById(ctx, *req.MerchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateMerchantRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchant_id", *req.MerchantID))
	}

	merchant, err := s.merchantRepository.UpdateMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateMerchantRow](
			s.logger,
			merchant_errors.ErrFailedUpdateMerchant,
			method,
			span,
			zap.Any("request", req))
	}

	s.Cache.DeleteCachedMerchant(ctx, int(merchant.MerchantID))

	logSuccess("Successfully updated merchant",
		zap.Int("merchant_id", int(merchant.MerchantID)),
		zap.String("name", req.Name))

	return merchant, nil
}

func (s *merchantService) TrashedMerchant(ctx context.Context, merchantID int) (*db.Merchant, error) {
	const method = "TrashedMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", merchantID))

	defer func() {
		end(status)
	}()

	merchant, err := s.merchantRepository.TrashedMerchant(ctx, merchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Merchant](
			s.logger,
			merchant_errors.ErrFailedTrashMerchant,
			method,
			span,
			zap.Int("merchant_id", merchantID))
	}

	s.Cache.DeleteCachedMerchant(ctx, merchantID)

	logSuccess("Successfully trashed merchant",
		zap.Int("merchant_id", merchantID))

	return merchant, nil
}

func (s *merchantService) RestoreMerchant(ctx context.Context, merchantID int) (*db.Merchant, error) {
	const method = "RestoreMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", merchantID))

	defer func() {
		end(status)
	}()

	merchant, err := s.merchantRepository.RestoreMerchant(ctx, merchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Merchant](
			s.logger,
			merchant_errors.ErrFailedRestoreMerchant,
			method,
			span,
			zap.Int("merchant_id", merchantID))
	}

	s.Cache.DeleteCachedMerchant(ctx, merchantID)

	logSuccess("Successfully restored merchant",
		zap.Int("merchant_id", merchantID))

	return merchant, nil
}

func (s *merchantService) DeleteMerchantPermanent(ctx context.Context, merchantID int) (bool, error) {
	const method = "DeleteMerchantPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("merchant_id", merchantID))

	defer func() {
		end(status)
	}()

	success, err := s.merchantRepository.DeleteMerchantPermanent(ctx, merchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			merchant_errors.ErrFailedDeleteMerchantPermanent,
			method,
			span,
			zap.Int("merchant_id", merchantID))
	}

	s.Cache.DeleteCachedMerchant(ctx, merchantID)

	logSuccess("Successfully permanently deleted merchant",
		zap.Int("merchant_id", merchantID))

	return success, nil
}

func (s *merchantService) RestoreAllMerchant(ctx context.Context) (bool, error) {
	const method = "RestoreAllMerchant"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.merchantRepository.RestoreAllMerchant(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			merchant_errors.ErrFailedRestoreAllMerchants,
			method,
			span)
	}

	s.logger.Debug("All merchant caches should be invalidated after restore all operation")

	logSuccess("Successfully restored all trashed merchants")

	return success, nil
}

func (s *merchantService) DeleteAllMerchantPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllMerchantPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.merchantRepository.DeleteAllMerchantPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			merchant_errors.ErrFailedDeleteAllMerchantsPermanent,
			method,
			span)
	}

	s.logger.Debug("All merchant caches should be invalidated after delete all operation")

	logSuccess("Successfully permanently deleted all merchants")

	return success, nil
}
