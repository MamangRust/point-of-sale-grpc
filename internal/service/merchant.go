package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type merchantService struct {
	merchantRepository repository.MerchantRepository
	logger             logger.LoggerInterface
	mapping            response_service.MerchantResponseMapper
}

func NewMerchantService(
	merchantRepository repository.MerchantRepository,
	logger logger.LoggerInterface,
	mapping response_service.MerchantResponseMapper,
) *merchantService {
	return &merchantService{
		merchantRepository: merchantRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *merchantService) FindAll(page, pageSize int, search string) ([]*response.MerchantResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching all merchants", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindAllMerchants(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve merchant list",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant list",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantsResponse(merchants), &totalRecords, nil
}

func (s *merchantService) FindByActive(search string, page, pageSize int) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active merchants", zap.String("search", search), zap.Int("page", page), zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindByActive(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve active merchants",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve active merchants",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantsResponseDeleteAt(merchants), &totalRecords, nil
}

func (s *merchantService) FindByTrashed(search string, page, pageSize int) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed merchants", zap.String("search", search), zap.Int("page", page), zap.Int("pageSize", pageSize))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	merchants, totalRecords, err := s.merchantRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve trashed merchants",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve trashed merchants",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantsResponseDeleteAt(merchants), &totalRecords, nil
}

func (s *merchantService) FindById(merchantID int) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching merchant by ID", zap.Int("merchantID", merchantID))

	merchant, err := s.merchantRepository.FindById(merchantID)
	if err != nil {
		s.logger.Error("Failed to retrieve merchant details",
			zap.Error(err),
			zap.Int("merchant_id", merchantID))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Merchant with ID %d not found", merchantID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant details",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantResponse(merchant), nil
}

func (s *merchantService) CreateMerchant(req *requests.CreateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new merchant")

	merchant, err := s.merchantRepository.CreateMerchant(req)
	if err != nil {
		s.logger.Error("Failed to create new merchant record",
			zap.Error(err),
			zap.Any("request", req))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create new merchant",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantResponse(merchant), nil
}

func (s *merchantService) UpdateMerchant(req *requests.UpdateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating merchant", zap.Int("merchantID", *req.MerchantID))

	merchant, err := s.merchantRepository.UpdateMerchant(req)

	if err != nil {
		s.logger.Error("Failed to update merchant record",
			zap.Error(err),
			zap.Any("request", req))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: "Merchant not found for update",
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update merchant",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantResponse(merchant), nil
}

func (s *merchantService) TrashedMerchant(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Trashing merchant", zap.Int("merchantID", merchantID))

	merchant, err := s.merchantRepository.TrashedMerchant(merchantID)
	if err != nil {
		s.logger.Error("Failed to move merchant to trash",
			zap.Error(err),
			zap.Int("merchant_id", merchantID))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Merchant with ID %d not found", merchantID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move merchant to trash",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantResponseDeleteAt(merchant), nil
}

func (s *merchantService) RestoreMerchant(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring merchant", zap.Int("merchantID", merchantID))

	merchant, err := s.merchantRepository.RestoreMerchant(merchantID)
	if err != nil {
		s.logger.Error("Failed to restore merchant from trash",
			zap.Error(err),
			zap.Int("merchant_id", merchantID))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Merchant with ID %d not found in trash", merchantID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore merchant from trash",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToMerchantResponseDeleteAt(merchant), nil
}

func (s *merchantService) DeleteMerchantPermanent(merchantID int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Deleting merchant permanently", zap.Int("merchantID", merchantID))

	success, err := s.merchantRepository.DeleteMerchantPermanent(merchantID)
	if err != nil {
		s.logger.Error("Failed to permanently delete merchant",
			zap.Error(err),
			zap.Int("merchant_id", merchantID))
		if errors.Is(err, sql.ErrNoRows) {
			return false, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Merchant with ID %d not found", merchantID),
				Code:    http.StatusNotFound,
			}
		}
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete merchant",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *merchantService) RestoreAllMerchant() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed merchants")

	success, err := s.merchantRepository.RestoreAllMerchant()
	if err != nil {
		s.logger.Error("Failed to restore all trashed merchants",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all merchants",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *merchantService) DeleteAllMerchantPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all merchants")

	success, err := s.merchantRepository.DeleteAllMerchantPermanent()
	if err != nil {
		s.logger.Error("Failed to permanently delete all merchants",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all merchants",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}
