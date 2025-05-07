package service

import (
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/errors/cashier_errors"
	"pointofsale/pkg/errors/merchant_errors"
	"pointofsale/pkg/errors/user_errors"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type cashierService struct {
	merchantRepository repository.MerchantRepository
	userRepository     repository.UserRepository
	cashierRepository  repository.CashierRepository
	logger             logger.LoggerInterface
	mapping            response_service.CashierResponseMapper
}

func NewCashierService(
	merchantRepository repository.MerchantRepository,
	userRepository repository.UserRepository,
	cashierRepository repository.CashierRepository,
	logger logger.LoggerInterface,
	mapping response_service.CashierResponseMapper,
) *cashierService {
	return &cashierService{
		merchantRepository: merchantRepository,
		userRepository:     userRepository,
		cashierRepository:  cashierRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *cashierService) FindAll(req *requests.FindAllCashiers) ([]*response.CashierResponse, *int, *response.ErrorResponse) {
	page := req.Page
	pageSize := req.PageSize
	search := req.Search

	s.logger.Debug("Fetching all merchant",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	cashier, totalRecords, err := s.cashierRepository.FindAllCashiers(req)

	if err != nil {
		s.logger.Error("Failed to fetch cashier",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, cashier_errors.ErrFailedFindAllCashiers
	}

	cashierResponse := s.mapping.ToCashiersResponse(cashier)

	s.logger.Debug("Successfully fetched merchants",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return cashierResponse, totalRecords, nil
}

func (s *cashierService) FindById(cashierID int) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashier by ID", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.FindById(cashierID)

	if err != nil {
		s.logger.Error("Failed to retrieve cashier details",
			zap.Error(err),
			zap.Int("cashier_id", cashierID))

		return nil, cashier_errors.ErrFailedFindCashierById
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) FindByActive(req *requests.FindAllCashiers) ([]*response.CashierResponseDeleteAt, *int, *response.ErrorResponse) {
	page := req.Page
	pageSize := req.PageSize
	search := req.Search

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

	cashiers, totalRecords, err := s.cashierRepository.FindByActive(req)

	if err != nil {
		s.logger.Error("Failed to retrieve active cashiers",
			zap.Error(err),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.String("search", req.Search))

		return nil, nil, cashier_errors.ErrFailedFindCashierByActive
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return s.mapping.ToCashiersResponseDeleteAt(cashiers), totalRecords, nil
}

func (s *cashierService) FindByTrashed(req *requests.FindAllCashiers) ([]*response.CashierResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashiers",
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize),
		zap.String("search", req.Search))

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	cashiers, totalRecords, err := s.cashierRepository.FindByTrashed(req)

	if err != nil {
		s.logger.Error("Failed to retrieve trashed cashiers",
			zap.Error(err),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.String("search", req.Search))

		return nil, nil, cashier_errors.ErrFailedFindCashierByTrashed
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return s.mapping.ToCashiersResponseDeleteAt(cashiers), totalRecords, nil
}

func (s *cashierService) FindByMerchant(req *requests.FindAllCashierMerchant) ([]*response.CashierResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashiers",
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize),
		zap.String("search", req.Search))

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	cashiers, totalRecords, err := s.cashierRepository.FindByMerchant(req)

	if err != nil {
		s.logger.Error("Failed to retrieve merchant's cashiers",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize),
			zap.String("search", req.Search))
		return nil, nil, cashier_errors.ErrFailedFindCashierByMerchant
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", *totalRecords),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return s.mapping.ToCashiersResponse(cashiers), totalRecords, nil
}

func (s *cashierService) FindMonthlyTotalSales(req *requests.MonthTotalSales) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse) {
	year := req.Year
	month := req.Month

	res, err := s.cashierRepository.GetMonthlyTotalSales(req)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Int("month", month),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindMonthlyTotalSales
	}

	return s.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (s *cashierService) FindYearlyTotalSales(year int) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse) {
	res, err := s.cashierRepository.GetYearlyTotalSales(year)

	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))

		return nil, cashier_errors.ErrFailedFindYearlyTotalSales
	}

	return s.mapping.ToCashierYearlyTotalSales(res), nil
}

func (s *cashierService) FindMonthlyTotalSalesById(req *requests.MonthTotalSalesCashier) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse) {
	month := req.Month
	year := req.Year

	res, err := s.cashierRepository.GetMonthlyTotalSalesById(req)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Int("month", month),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindMonthlyTotalSalesById
	}

	return s.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (s *cashierService) FindYearlyTotalSalesById(req *requests.YearTotalSalesCashier) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse) {
	year := req.Year
	cashier_id := req.CashierID

	res, err := s.cashierRepository.GetYearlyTotalSalesById(req)

	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("cashier_id", cashier_id),
			zap.Int("year", year),
			zap.Error(err))

		return nil, cashier_errors.ErrFailedFindYearlyTotalSalesById
	}

	return s.mapping.ToCashierYearlyTotalSales(res), nil
}

func (s *cashierService) FindMonthlyTotalSalesByMerchant(req *requests.MonthTotalSalesMerchant) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse) {
	year := req.Year
	month := req.Month
	merchant_id := req.MerchantID

	res, err := s.cashierRepository.GetMonthlyTotalSalesByMerchant(req)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Int("month", month),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindMonthlyTotalSalesByMerchant
	}

	return s.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (s *cashierService) FindYearlyTotalSalesByMerchant(req *requests.YearTotalSalesMerchant) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse) {
	year := req.Year
	merchant_id := req.MerchantID

	res, err := s.cashierRepository.GetYearlyTotalSalesByMerchant(req)

	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))

		return nil, cashier_errors.ErrFailedFindYearlyTotalSalesByMerchant
	}

	return s.mapping.ToCashierYearlyTotalSales(res), nil
}

func (s *cashierService) FindMonthlySales(year int) ([]*response.CashierResponseMonthSales, *response.ErrorResponse) {
	res, err := s.cashierRepository.GetMonthyCashier(year)

	if err != nil {
		s.logger.Error("failed to get monthly cashier sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindMonthlySales
	}

	return s.mapping.ToCashierMonthlySales(res), nil
}

func (s *cashierService) FindYearlySales(year int) ([]*response.CashierResponseYearSales, *response.ErrorResponse) {
	res, err := s.cashierRepository.GetYearlyCashier(year)

	if err != nil {
		s.logger.Error("failed to get yearly cashier sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindYearlySales
	}

	return s.mapping.ToCashierYearlySales(res), nil
}

func (s *cashierService) FindMonthlyCashierByMerchant(req *requests.MonthCashierMerchant) ([]*response.CashierResponseMonthSales, *response.ErrorResponse) {
	year := req.Year
	merchant_id := req.MerchantID

	res, err := s.cashierRepository.GetMonthlyCashierByMerchant(req)

	if err != nil {
		s.logger.Error("failed to get monthly cashier sales by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindMonthlyCashierByMerchant
	}

	return s.mapping.ToCashierMonthlySales(res), nil
}

func (s *cashierService) FindYearlyCashierByMerchant(req *requests.YearCashierMerchant) ([]*response.CashierResponseYearSales, *response.ErrorResponse) {
	year := req.Year
	merchant_id := req.MerchantID

	res, err := s.cashierRepository.GetYearlyCashierByMerchant(req)

	if err != nil {
		s.logger.Error("failed to get yearly cashier sales by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindYearlyCashierByMerchant
	}

	return s.mapping.ToCashierYearlySales(res), nil
}

func (s *cashierService) FindMonthlyCashierById(req *requests.MonthCashierId) ([]*response.CashierResponseMonthSales, *response.ErrorResponse) {
	year := req.Year
	cashier_id := req.CashierID

	res, err := s.cashierRepository.GetMonthlyCashierById(req)

	if err != nil {
		s.logger.Error("failed to get monthly cashier sales by ID",
			zap.Int("year", year),
			zap.Int("cashier_id", cashier_id),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindMonthlyCashierById
	}

	return s.mapping.ToCashierMonthlySales(res), nil
}

func (s *cashierService) FindYearlyCashierById(req *requests.YearCashierId) ([]*response.CashierResponseYearSales, *response.ErrorResponse) {
	year := req.Year
	cashier_id := req.CashierID

	res, err := s.cashierRepository.GetYearlyCashierById(req)

	if err != nil {
		s.logger.Error("failed to get yearly cashier sales by ID",
			zap.Int("year", year),
			zap.Int("cashier_id", cashier_id),
			zap.Error(err))
		return nil, cashier_errors.ErrFailedFindYearlyCashierById
	}

	return s.mapping.ToCashierYearlySales(res), nil
}

func (s *cashierService) CreateCashier(req *requests.CreateCashierRequest) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new cashier")

	_, err := s.merchantRepository.FindById(req.MerchantID)

	if err != nil {
		s.logger.Error("Failed to retrieve merchant details",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID))

		return nil, merchant_errors.ErrFailedFindMerchantById
	}

	_, err = s.userRepository.FindById(req.UserID)

	if err != nil {
		s.logger.Error("Failed to retrieve user details",
			zap.Error(err),
			zap.Int("user_id", req.UserID))

		return nil, user_errors.ErrFailedFindUserByID
	}

	cashier, err := s.cashierRepository.CreateCashier(req)

	if err != nil {
		s.logger.Error("Failed to create new cashier",
			zap.Error(err),
			zap.Any("request", req))
		return nil, cashier_errors.ErrFailedCreateCashier
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) UpdateCashier(req *requests.UpdateCashierRequest) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating cashier", zap.Int("cashier_id", *req.CashierID))

	cashier, err := s.cashierRepository.UpdateCashier(req)

	if err != nil {
		s.logger.Error("Failed to update cashier",
			zap.Error(err),
			zap.Int("cashier_id", *req.CashierID))

		return nil, cashier_errors.ErrFailedUpdateCashier
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) TrashedCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Trashing cashier", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.TrashedCashier(cashierID)

	if err != nil {
		s.logger.Error("Failed to move cashier to trash",
			zap.Error(err),
			zap.Int("cashier_id", cashierID))

		return nil, cashier_errors.ErrFailedTrashedCashier
	}

	return s.mapping.ToCashierResponseDeleteAt(cashier), nil
}

func (s *cashierService) RestoreCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring cashier", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.RestoreCashier(cashierID)

	if err != nil {
		s.logger.Error("Failed to restore cashier from trash",
			zap.Error(err),
			zap.Int("cashier_id", cashierID))

		return nil, cashier_errors.ErrFailedRestoreCashier
	}

	return s.mapping.ToCashierResponseDeleteAt(cashier), nil
}

func (s *cashierService) DeleteCashierPermanent(cashierID int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting cashier", zap.Int("cashierID", cashierID))

	success, err := s.cashierRepository.DeleteCashierPermanent(cashierID)

	if err != nil {
		s.logger.Error("Failed to permanently delete cashier",
			zap.Error(err),
			zap.Int("cashier_id", cashierID))

		return false, cashier_errors.ErrFailedDeleteCashierPermanent
	}

	return success, nil
}

func (s *cashierService) RestoreAllCashier() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed cashiers")

	success, err := s.cashierRepository.RestoreAllCashier()

	if err != nil {
		s.logger.Error("Failed to restore all trashed cashiers",
			zap.Error(err))
		return false, cashier_errors.ErrFailedRestoreAllCashiers
	}

	return success, nil
}

func (s *cashierService) DeleteAllCashierPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all cashiers")

	success, err := s.cashierRepository.DeleteAllCashierPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all trashed cashiers",
			zap.Error(err))
		return false, cashier_errors.ErrFailedDeleteAllCashierPermanent
	}

	return success, nil
}
