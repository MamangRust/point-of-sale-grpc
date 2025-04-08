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

func (s *cashierService) FindAll(page int, pageSize int, search string) ([]*response.CashierResponse, *int, *response.ErrorResponse) {
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

		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashier list",
			Code:    http.StatusInternalServerError,
		}
	}

	cashierResponse := s.mapping.ToCashiersResponse(cashier)

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return cashierResponse, &totalRecords, nil
}

func (s *cashierService) FindById(cashierID int) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching cashier by ID", zap.Int("cashierID", cashierID))

	cashier, err := s.cashierRepository.FindById(cashierID)

	if err != nil {
		s.logger.Error("Failed to retrieve cashier details",
			zap.Error(err),
			zap.Int("cashier_id", cashierID))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Cashier with ID %d not found", cashierID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashier details",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) FindByActive(search string, page, pageSize int) ([]*response.CashierResponseDeleteAt, *int, *response.ErrorResponse) {
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
		s.logger.Error("Failed to retrieve active cashiers",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve active cashiers",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCashiersResponseDeleteAt(cashiers), &totalRecords, nil
}

func (s *cashierService) FindByTrashed(search string, page, pageSize int) ([]*response.CashierResponseDeleteAt, *int, *response.ErrorResponse) {
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
		s.logger.Error("Failed to retrieve trashed cashiers",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))

		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve trashed cashiers",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCashiersResponseDeleteAt(cashiers), &totalRecords, nil
}

func (s *cashierService) FindByMerchant(merchantID int, search string, page, pageSize int) ([]*response.CashierResponse, *int, *response.ErrorResponse) {
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
		s.logger.Error("Failed to retrieve merchant's cashiers",
			zap.Error(err),
			zap.Int("merchant_id", merchantID),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve cashiers for merchant",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched cashier",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCashiersResponse(cashiers), &totalRecords, nil
}

func (s *cashierService) FindMonthlyTotalSales(year int, month int) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if month <= 0 || month >= 12 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Month must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetMonthlyTotalSales(year, month)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (s *cashierService) FindYearlyTotalSales(year int) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetYearlyTotalSales(year)
	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierYearlyTotalSales(res), nil
}

func (s *cashierService) FindMonthlyTotalSalesById(year int, month int, cashier_id int) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if month <= 0 || month >= 12 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Month must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetMonthlyTotalSalesById(year, month, cashier_id)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (s *cashierService) FindYearlyTotalSalesById(year int, cashier_id int) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if cashier_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Cashier ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetYearlyTotalSalesById(year, cashier_id)
	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierYearlyTotalSales(res), nil
}

func (s *cashierService) FindMonthlyTotalSalesByMerchant(year int, month int, merchant_id int) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if month <= 0 || month >= 12 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Month must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetMonthlyTotalSalesById(year, month, merchant_id)

	if err != nil {
		s.logger.Error("failed to get monthly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierMonthlyTotalSales(res), nil
}

func (s *cashierService) FindYearlyTotalSalesByMerchant(year int, merchant_id int) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	if merchant_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Merchant ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetYearlyTotalSalesByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get yearly total sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly total sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierYearlyTotalSales(res), nil
}

func (s *cashierService) FindMonthlySales(year int) ([]*response.CashierResponseMonthSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetMonthyCashier(year)

	if err != nil {
		s.logger.Error("failed to get monthly cashier sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly cashier sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierMonthlySales(res), nil
}

func (s *cashierService) FindYearlySales(year int) ([]*response.CashierResponseYearSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetYearlyCashier(year)
	if err != nil {
		s.logger.Error("failed to get yearly cashier sales",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly cashier sales data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierYearlySales(res), nil
}

func (s *cashierService) FindMonthlyCashierByMerchant(year int, merchant_id int) ([]*response.CashierResponseMonthSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}
	if merchant_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Merchant ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetMonthlyCashierByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get monthly cashier sales by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve monthly cashier sales for merchant %d", merchant_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierMonthlySales(res), nil
}

func (s *cashierService) FindYearlyCashierByMerchant(year int, merchant_id int) ([]*response.CashierResponseYearSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}
	if merchant_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Merchant ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetYearlyCashierByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get yearly cashier sales by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve yearly cashier sales for merchant %d", merchant_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierYearlySales(res), nil
}

func (s *cashierService) FindMonthlyCashierById(year int, cashier_id int) ([]*response.CashierResponseMonthSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}
	if cashier_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Cashier ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetMonthlyCashierById(year, cashier_id)
	if err != nil {
		s.logger.Error("failed to get monthly cashier sales by ID",
			zap.Int("year", year),
			zap.Int("cashier_id", cashier_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve monthly cashier sales for cashier %d", cashier_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierMonthlySales(res), nil
}

func (s *cashierService) FindYearlyCashierById(year int, cashier_id int) ([]*response.CashierResponseYearSales, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}
	if cashier_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Cashier ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.cashierRepository.GetYearlyCashierById(year, cashier_id)
	if err != nil {
		s.logger.Error("failed to get yearly cashier sales by ID",
			zap.Int("year", year),
			zap.Int("cashier_id", cashier_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve yearly cashier sales for cashier %d", cashier_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierYearlySales(res), nil
}

func (s *cashierService) CreateCashier(req *requests.CreateCashierRequest) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new cashier")

	cashier, err := s.cashierRepository.CreateCashier(req)

	if err != nil {
		s.logger.Error("Failed to create new cashier",
			zap.Error(err),
			zap.Any("request", req))
		return nil, &response.ErrorResponse{
			Status:  "database_error",
			Message: "Failed to create new cashier record",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCashierResponse(cashier), nil
}

func (s *cashierService) UpdateCashier(req *requests.UpdateCashierRequest) (*response.CashierResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating cashier", zap.Int("cashierID", *req.CashierID))

	cashier, err := s.cashierRepository.UpdateCashier(req)

	if err != nil {
		s.logger.Error("Failed to update cashier",
			zap.Error(err),
			zap.Any("request", req))

		if errors.Is(err, sql.ErrNoRows) {

			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: "Cashier not found for update",
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update cashier record",
			Code:    http.StatusInternalServerError,
		}
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Cashier with ID %d not found", cashierID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move cashier to trash",
			Code:    http.StatusInternalServerError,
		}
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Cashier with ID %d not found in trash", cashierID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore cashier from trash",
			Code:    http.StatusInternalServerError,
		}
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
		if errors.Is(err, sql.ErrNoRows) {
			return false, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Cashier with ID %d not found", cashierID),
				Code:    http.StatusNotFound,
			}
		}
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete cashier",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *cashierService) RestoreAllCashier() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed cashiers")

	success, err := s.cashierRepository.RestoreAllCashier()

	if err != nil {
		s.logger.Error("Failed to restore all trashed cashiers",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all trashed cashiers",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *cashierService) DeleteAllCashierPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all cashiers")

	success, err := s.cashierRepository.DeleteAllCashierPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all trashed cashiers",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all trashed cashiers",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}
