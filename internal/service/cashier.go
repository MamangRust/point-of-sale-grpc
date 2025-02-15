package service

import (
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type cashierService struct {
	cashierRepository repository.CashierRepository
	logger            logger.LoggerInterface
	mapping           response_service.CashierResponseMapper
}

func NewCashierService(
	cashierRepository repository.CashierRepository,
	logger logger.LoggerInterface,
	mapping response_service.CashierResponseMapper,
) *cashierService {
	return &cashierService{
		cashierRepository: cashierRepository,
		logger:            logger,
		mapping:           mapping,
	}
}

func (s *cashierService) FindAll(page int, pageSize int, search string) ([]*response.CashierResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashier",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	cashier, totalRecords, err := s.cashierRepository.FindAllCashiers(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch cashier",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, 0, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to fetch cashier",
		}
	}

	cashierResponse := s.mapping.ToCashiersResponse(cashier)

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return cashierResponse, int(totalRecords), nil
}

func (s *cashierService) FindById(cashierID int) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashier by ID", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.FindById(cashierID)
	if err != nil {
		s.logger.Error("Failed to fetch cashier", zap.Error(err))
		return nil, &response.ErrorResponse{Status: "error", Message: "Cashier not found"}
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) FindByActive(search string, page, pageSize int) ([]*response.CashierResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashiers",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	cashiers, totalRecords, err := s.cashierRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to fetch cashier",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, 0, &response.ErrorResponse{Status: "error", Message: "Failed to fetch active cashiers"}
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCashiersResponseDeleteAt(cashiers), totalRecords, nil
}

func (s *cashierService) FindByTrashed(search string, page, pageSize int) ([]*response.CashierResponseDeleteAt, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashiers",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	cashiers, totalRecords, err := s.cashierRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch cashier",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, 0, &response.ErrorResponse{Status: "error", Message: "Failed to fetch trashed cashiers"}
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCashiersResponseDeleteAt(cashiers), totalRecords, nil
}

func (s *cashierService) FindByMerchant(merchantID int, search string, page, pageSize int) ([]*response.CashierResponse, int, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashiers",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	cashiers, totalRecords, err := s.cashierRepository.FindByMerchant(merchantID, search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to fetch cashier",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, 0, &response.ErrorResponse{Status: "error", Message: "Failed to fetch cashiers by merchant"}
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCashiersResponse(cashiers), totalRecords, nil
}

func (s *cashierService) CreateCashier(req *requests.CreateCashierRequest) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new cashier")

	cashier, err := s.cashierRepository.CreateCashier(req)
	if err != nil {
		s.logger.Error("Failed to create cashier", zap.Error(err))
		return nil, &response.ErrorResponse{Status: "error", Message: "Failed to create cashier"}
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) UpdateCashier(req *requests.UpdateCashierRequest) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating cashier", zap.Int("cashierID", req.CashierID))

	cashier, err := s.cashierRepository.UpdateCashier(req)
	if err != nil {
		s.logger.Error("Failed to update cashier", zap.Error(err))
		return nil, &response.ErrorResponse{Status: "error", Message: "Failed to update cashier"}
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) TrashedCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Trashing cashier", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.TrashedCashier(cashierID)
	if err != nil {
		s.logger.Error("Failed to trash cashier", zap.Error(err))
		return nil, &response.ErrorResponse{Status: "error", Message: "Failed to trash cashier"}
	}

	return s.mapping.ToCashierResponseDeleteAt(cashier), nil
}

func (s *cashierService) RestoreCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring cashier", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.RestoreCashier(cashierID)
	if err != nil {
		s.logger.Error("Failed to restore cashier", zap.Error(err))
		return nil, &response.ErrorResponse{Status: "error", Message: "Failed to restore cashier"}
	}

	return s.mapping.ToCashierResponseDeleteAt(cashier), nil
}

func (s *cashierService) DeleteCashierPermanent(cashierID int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting cashier", zap.Int("cashierID", cashierID))

	success, err := s.cashierRepository.DeleteCashierPermanent(cashierID)
	if err != nil {
		s.logger.Error("Failed to permanently delete cashier", zap.Error(err))
		return false, &response.ErrorResponse{Status: "error", Message: "Failed to permanently delete cashier"}
	}

	return success, nil
}

func (s *cashierService) RestoreAllCashier() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed cashiers")

	success, err := s.cashierRepository.RestoreAllCashier()
	if err != nil {
		s.logger.Error("Failed to restore all cashiers", zap.Error(err))
		return false, &response.ErrorResponse{Status: "error", Message: "Failed to restore all cashiers"}
	}

	return success, nil
}

func (s *cashierService) DeleteAllCashierPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all cashiers")

	success, err := s.cashierRepository.DeleteAllCashierPermanent()
	if err != nil {
		s.logger.Error("Failed to permanently delete all cashiers", zap.Error(err))
		return false, &response.ErrorResponse{Status: "error", Message: "Failed to permanently delete all cashiers"}
	}

	return success, nil
}
