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
	"pointofsale/pkg/utils"

	"go.uber.org/zap"
)

type categoryService struct {
	categoryRepository repository.CategoryRepository
	logger             logger.LoggerInterface
	mapping            response_service.CategoryResponseMapper
}

func NewCategoryService(
	categoryRepository repository.CategoryRepository,
	logger logger.LoggerInterface,
	mapping response_service.CategoryResponseMapper,
) *categoryService {
	return &categoryService{
		categoryRepository: categoryRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *categoryService) FindAll(page int, pageSize int, search string) ([]*response.CategoryResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching category",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	category, totalRecords, err := s.categoryRepository.FindAllCategory(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to retrieve category list",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve category list",
			Code:    http.StatusInternalServerError,
		}
	}

	categoriesResponse := s.mapping.ToCategorysResponse(category)

	s.logger.Debug("Successfully fetched category",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return categoriesResponse, &totalRecords, nil
}

func (s *categoryService) FindById(category_id int) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching category by ID", zap.Int("category_id", category_id))

	category, err := s.categoryRepository.FindById(category_id)

	if err != nil {
		s.logger.Error("Failed to retrieve category details",
			zap.Error(err),
			zap.Int("category_id", category_id))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Category with ID %d not found", category_id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve category details",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryResponse(category), nil
}

func (s *categoryService) FindByActive(search string, page, pageSize int) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching categories",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	categories, totalRecords, err := s.categoryRepository.FindByActive(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to retrieve active categories",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve active categories",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched categories",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCategoryResponsesDeleteAt(categories), &totalRecords, nil
}

func (s *categoryService) FindByTrashed(search string, page, pageSize int) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching categories",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	categories, totalRecords, err := s.categoryRepository.FindByTrashed(search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to retrieve trashed categories",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve trashed categories",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched categories",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToCategoryResponsesDeleteAt(categories), &totalRecords, nil
}

func (s *categoryService) FindMonthlyTotalPrice(year int, month int) ([]*response.CategoriesMonthlyTotalPriceResponse, *response.ErrorResponse) {
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

	res, err := s.categoryRepository.GetMonthlyTotalPrice(year, month)

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

	return s.mapping.ToCategoryMonthlyTotalPrices(res), nil
}

func (s *categoryService) FindYearlyTotalPrice(year int) ([]*response.CategoriesYearlyTotalPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetYearlyTotalPrices(year)
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

	return s.mapping.ToCategoryYearlyTotalPrices(res), nil
}

func (s *categoryService) FindMonthlyTotalPriceById(year int, month int, category_id int) ([]*response.CategoriesMonthlyTotalPriceResponse, *response.ErrorResponse) {
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

	res, err := s.categoryRepository.GetMonthlyTotalPriceById(year, month, category_id)

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

	return s.mapping.ToCategoryMonthlyTotalPrices(res), nil
}

func (s *categoryService) FindYearlyTotalPriceById(year int, category_id int) ([]*response.CategoriesYearlyTotalPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetYearlyTotalPricesById(year, category_id)
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

	return s.mapping.ToCategoryYearlyTotalPrices(res), nil
}

func (s *categoryService) FindMonthlyTotalPriceByMerchant(year int, month int, merchant_id int) ([]*response.CategoriesMonthlyTotalPriceResponse, *response.ErrorResponse) {
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

	res, err := s.categoryRepository.GetMonthlyTotalPriceByMerchant(year, month, merchant_id)

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

	return s.mapping.ToCategoryMonthlyTotalPrices(res), nil
}

func (s *categoryService) FindYearlyTotalPriceByMerchant(year int, merchant_id int) ([]*response.CategoriesYearlyTotalPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetYearlyTotalPricesByMerchant(year, merchant_id)
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

	return s.mapping.ToCategoryYearlyTotalPrices(res), nil
}

func (s *categoryService) FindMonthPrice(year int) ([]*response.CategoryMonthPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetMonthPrice(year)
	if err != nil {
		s.logger.Error("failed to get monthly category prices",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve monthly category data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryMonthlyPrices(res), nil
}

func (s *categoryService) FindYearPrice(year int) ([]*response.CategoryYearPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetYearPrice(year)
	if err != nil {
		s.logger.Error("failed to get yearly category prices",
			zap.Int("year", year),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: "Failed to retrieve yearly category data",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryYearlyPrices(res), nil
}

func (s *categoryService) FindMonthPriceByMerchant(year int, merchant_id int) ([]*response.CategoryMonthPriceResponse, *response.ErrorResponse) {
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

	res, err := s.categoryRepository.GetMonthPriceByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get monthly category prices by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve monthly category data for merchant %d", merchant_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryMonthlyPrices(res), nil
}

func (s *categoryService) FindYearPriceByMerchant(year int, merchant_id int) ([]*response.CategoryYearPriceResponse, *response.ErrorResponse) {
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

	res, err := s.categoryRepository.GetYearPriceByMerchant(year, merchant_id)
	if err != nil {
		s.logger.Error("failed to get yearly category prices by merchant",
			zap.Int("year", year),
			zap.Int("merchant_id", merchant_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve yearly category data for merchant %d", merchant_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryYearlyPrices(res), nil
}

func (s *categoryService) FindMonthPriceById(year int, category_id int) ([]*response.CategoryMonthPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}
	if category_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Category ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetMonthPriceById(year, category_id)
	if err != nil {
		s.logger.Error("failed to get monthly category prices by ID",
			zap.Int("year", year),
			zap.Int("category_id", category_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve monthly category data for category %d", category_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryMonthlyPrices(res), nil
}

func (s *categoryService) FindYearPriceById(year int, category_id int) ([]*response.CategoryYearPriceResponse, *response.ErrorResponse) {
	if year <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Year must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}
	if category_id <= 0 {
		return nil, &response.ErrorResponse{
			Status:  "invalid_request",
			Message: "Category ID must be a positive number",
			Code:    http.StatusBadRequest,
		}
	}

	res, err := s.categoryRepository.GetYearPriceById(year, category_id)
	if err != nil {
		s.logger.Error("failed to get yearly category prices by ID",
			zap.Int("year", year),
			zap.Int("category_id", category_id),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "server_error",
			Message: fmt.Sprintf("Failed to retrieve yearly category data for category %d", category_id),
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryYearlyPrices(res), nil
}

func (s *categoryService) CreateCategory(req *requests.CreateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("hello", zap.Any("request", req))

	s.logger.Debug("Creating new category")

	slug := utils.GenerateSlug(req.Name)

	req.SlugCategory = &slug

	category, err := s.categoryRepository.CreateCategory(req)

	if err != nil {
		s.logger.Error("Failed to create new category",
			zap.Error(err),
			zap.Any("request", req))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create new category",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryResponse(category), nil
}

func (s *categoryService) UpdateCategory(req *requests.UpdateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating category", zap.Int("category_id", *req.CategoryID))

	slug := utils.GenerateSlug(req.Name)

	req.SlugCategory = &slug

	category, err := s.categoryRepository.UpdateCategory(req)

	if err != nil {
		s.logger.Error("Failed to update category",
			zap.Error(err),
			zap.Any("request", req))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: "Category not found for update",
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update category",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryResponse(category), nil
}

func (s *categoryService) TrashedCategory(category_id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Trashing category", zap.Int("category", category_id))

	category, err := s.categoryRepository.TrashedCategory(category_id)

	if err != nil {
		s.logger.Error("Failed to move category to trash",
			zap.Error(err),
			zap.Int("category_id", category_id))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Category with ID %d not found", category_id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move category to trash",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryResponseDeleteAt(category), nil
}

func (s *categoryService) RestoreCategory(categoryID int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring category", zap.Int("categoryID", categoryID))

	category, err := s.categoryRepository.RestoreCategory(categoryID)
	if err != nil {
		s.logger.Error("Failed to restore category from trash",
			zap.Error(err),
			zap.Int("category_id", categoryID))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Category with ID %d not found in trash", categoryID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore category from trash",
			Code:    http.StatusInternalServerError,
		}
	}

	return s.mapping.ToCategoryResponseDeleteAt(category), nil
}

func (s *categoryService) DeleteCategoryPermanent(categoryID int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting category", zap.Int("categoryID", categoryID))

	success, err := s.categoryRepository.DeleteCategoryPermanently(categoryID)

	if err != nil {
		s.logger.Error("Failed to permanently delete category",
			zap.Error(err),
			zap.Int("category_id", categoryID))
		if errors.Is(err, sql.ErrNoRows) {
			return false, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Category with ID %d not found", categoryID),
				Code:    http.StatusNotFound,
			}
		}
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete category",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *categoryService) RestoreAllCategories() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed categories")

	success, err := s.categoryRepository.RestoreAllCategories()
	if err != nil {
		s.logger.Error("Failed to restore all trashed categories",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all trashed categories",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}

func (s *categoryService) DeleteAllCategoriesPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all categories")

	success, err := s.categoryRepository.DeleteAllPermanentCategories()

	if err != nil {
		s.logger.Error("Failed to permanently delete all trashed categories",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all trashed categories",
			Code:    http.StatusInternalServerError,
		}
	}

	return success, nil
}
