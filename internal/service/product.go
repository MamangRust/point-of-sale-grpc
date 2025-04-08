package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/utils"

	"go.uber.org/zap"
)

type productService struct {
	categoryRepository repository.CategoryRepository
	merchantRepository repository.MerchantRepository
	productRepository  repository.ProductRepository
	logger             logger.LoggerInterface
	mapping            response_service.ProductResponseMapper
}

func NewProductService(
	categoryRepository repository.CategoryRepository,
	merchantRepository repository.MerchantRepository,
	productRepository repository.ProductRepository,
	logger logger.LoggerInterface,
	mapping response_service.ProductResponseMapper,
) *productService {
	return &productService{
		categoryRepository: categoryRepository,
		merchantRepository: merchantRepository,
		productRepository:  productRepository,
		logger:             logger,
		mapping:            mapping,
	}
}

func (s *productService) FindAll(page, pageSize int, search string) ([]*response.ProductResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching products",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	products, totalRecords, err := s.productRepository.FindAllProducts(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve product list from database",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve product list",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched products",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToProductsResponse(products), &totalRecords, nil
}

func (s *productService) FindByMerchant(req *requests.ProductByMerchantRequest) ([]*response.ProductResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching products",
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize),
		zap.String("search", req.Search), zap.Any("req", req))

	if req.Page <= 0 {
		req.Page = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	products, totalRecords, err := s.productRepository.FindByMerchant(req)

	s.logger.Debug("Hello Products", zap.Any("products", products))

	if err != nil {
		s.logger.Error("Failed to retrieve merchant's products from database",
			zap.Error(err),
			zap.Int("merchant_id", req.MerchantID),
			zap.String("search", req.Search),
			zap.Int("page", req.Page),
			zap.Int("pageSize", req.PageSize))

		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve merchant's products",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched products",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", req.Page),
		zap.Int("pageSize", req.PageSize))

	return s.mapping.ToProductsResponse(products), &totalRecords, nil
}

func (s *productService) FindByCategory(category_name string, page, pageSize int, search string) ([]*response.ProductResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching products",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	products, totalRecords, err := s.productRepository.FindByCategory(category_name, search, page, pageSize)

	if err != nil {
		s.logger.Error("Failed to retrieve products by category from database",
			zap.Error(err),
			zap.String("category", category_name),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))

		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: fmt.Sprintf("Failed to retrieve products in category '%s'", category_name),
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched products",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToProductsResponse(products), &totalRecords, nil
}

func (s *productService) FindById(productID int) (*response.ProductResponse, *response.ErrorResponse) {
	s.logger.Debug("Retrieving product details",
		zap.Int("product_id", productID))

	product, err := s.productRepository.FindById(productID)
	if err != nil {
		s.logger.Error("Failed to retrieve product details",
			zap.Int("product_id", productID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found", productID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve product details",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully retrieved product details",
		zap.Int("product_id", productID))
	return s.mapping.ToProductResponse(product), nil
}

func (s *productService) FindByActive(search string, page, pageSize int) ([]*response.ProductResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching products",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	products, totalRecords, err := s.productRepository.FindByActive(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve active products from database",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve active products",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched products",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToProductsResponseDeleteAt(products), &totalRecords, nil
}

func (s *productService) FindByTrashed(search string, page, pageSize int) ([]*response.ProductResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching products",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	products, totalRecords, err := s.productRepository.FindByTrashed(search, page, pageSize)
	if err != nil {
		s.logger.Error("Failed to retrieve trashed products from database",
			zap.Error(err),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve trashed products",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched products",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return s.mapping.ToProductsResponseDeleteAt(products), &totalRecords, nil
}

func (s *productService) CreateProduct(req *requests.CreateProductRequest) (*response.ProductResponse, *response.ErrorResponse) {
	s.logger.Debug("Creating new product",
		zap.String("name", req.Name),
		zap.Int("categoryID", req.CategoryID),
		zap.Int("merchantID", req.MerchantID))

	_, err := s.categoryRepository.FindById(req.CategoryID)

	if err != nil {
		s.logger.Error("Category not found for product creation",
			zap.Int("categoryID", req.CategoryID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Category with ID %d not found", req.CategoryID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to verify category",
			Code:    http.StatusInternalServerError,
		}
	}

	_, err = s.merchantRepository.FindById(req.MerchantID)

	if err != nil {
		s.logger.Error("Merchant not found for product creation",
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Merchant with ID %d not found", req.MerchantID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to verify merchant",
			Code:    http.StatusInternalServerError,
		}
	}

	barcode := utils.GenerateBarcode(req.Name)
	slug := utils.GenerateSlug(req.Name)
	req.Barcode = &barcode
	req.SlugProduct = &slug

	product, err := s.productRepository.CreateProduct(req)

	if err != nil {
		s.logger.Error("Failed to create product record",
			zap.Error(err))

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create product",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Product created successfully",
		zap.Int("productID", product.ID))

	return s.mapping.ToProductResponse(product), nil
}

func (s *productService) UpdateProduct(req *requests.UpdateProductRequest) (*response.ProductResponse, *response.ErrorResponse) {
	s.logger.Debug("Updating product",
		zap.Int("productID", *req.ProductID),
		zap.Int("categoryID", req.CategoryID),
		zap.Int("merchantID", req.MerchantID))

	_, err := s.categoryRepository.FindById(req.CategoryID)

	if err != nil {
		s.logger.Error("Category not found for product update",
			zap.Int("categoryID", req.CategoryID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Category with ID %d not found", req.CategoryID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to verify category",
			Code:    http.StatusInternalServerError,
		}
	}

	_, err = s.merchantRepository.FindById(req.MerchantID)

	if err != nil {
		s.logger.Error("Merchant not found for product update",
			zap.Int("merchantID", req.MerchantID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Merchant with ID %d not found", req.MerchantID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to verify merchant",
			Code:    http.StatusInternalServerError,
		}
	}

	barcode := utils.GenerateBarcode(req.Name)
	slug := utils.GenerateSlug(req.Name)
	req.Barcode = &barcode
	req.SlugProduct = &slug

	product, err := s.productRepository.UpdateProduct(req)

	if err != nil {
		s.logger.Error("Failed to update product record",
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found", req.ProductID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update product",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Product updated successfully",
		zap.Int("productID", product.ID))

	return s.mapping.ToProductResponse(product), nil
}

func (s *productService) TrashProduct(productID int) (*response.ProductResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Moving product to trash",
		zap.Int("product_id", productID))

	product, err := s.productRepository.TrashedProduct(productID)

	if err != nil {
		s.logger.Error("Failed to move product to trash",
			zap.Int("product_id", productID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found", productID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move product to trash",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Product moved to trash successfully",
		zap.Int("product_id", productID))

	return s.mapping.ToProductResponseDeleteAt(product), nil
}

func (s *productService) RestoreProduct(productID int) (*response.ProductResponseDeleteAt, *response.ErrorResponse) {
	s.logger.Debug("Restoring product from trash",
		zap.Int("product_id", productID))

	product, err := s.productRepository.RestoreProduct(productID)

	if err != nil {
		s.logger.Error("Failed to restore product from trash",
			zap.Int("product_id", productID),
			zap.Error(err))

		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found in trash", productID),
				Code:    http.StatusNotFound,
			}
		}

		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore product from trash",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Product restored successfully",
		zap.Int("product_id", productID))
	return s.mapping.ToProductResponseDeleteAt(product), nil
}

func (s *productService) DeleteProductPermanent(productID int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting product",
		zap.Int("product_id", productID))

	res, err := s.productRepository.FindByIdTrashed(productID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Product with ID %d not found", productID),
				Code:    http.StatusNotFound,
			}
		}

		s.logger.Error("Failed to find product",
			zap.Int("product_id", productID),
			zap.Error(err))

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to verify product existence",
			Code:    http.StatusInternalServerError,
		}
	}

	if res.ImageProduct != "" {
		err := os.Remove(res.ImageProduct)
		if err != nil {
			if os.IsNotExist(err) {
				s.logger.Debug("Product image file not found, continuing with product deletion",
					zap.String("image_path", res.ImageProduct))
			} else {
				s.logger.Debug("Failed to delete product image",
					zap.String("image_path", res.ImageProduct),
					zap.Error(err))

				return false, &response.ErrorResponse{
					Status:  "error",
					Message: "Failed to delete product image",
					Code:    http.StatusInternalServerError,
				}
			}
		} else {
			s.logger.Debug("Successfully deleted product image",
				zap.String("image_path", res.ImageProduct))
		}
	}

	success, err := s.productRepository.DeleteProductPermanent(productID)

	if err != nil {
		s.logger.Error("Failed to permanently delete product",
			zap.Int("product_id", productID),
			zap.Error(err))

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete product",
			Code:    http.StatusInternalServerError,
		}
	}

	if !success {
		s.logger.Debug("No rows were affected when deleting product",
			zap.Int("product_id", productID))

		return false, &response.ErrorResponse{
			Status:  "not_found",
			Message: fmt.Sprintf("Product with ID %d not found", productID),
			Code:    http.StatusNotFound,
		}
	}

	s.logger.Debug("Product permanently deleted successfully",
		zap.Int("product_id", productID))

	return true, nil
}

func (s *productService) RestoreAllProducts() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all trashed products")

	success, err := s.productRepository.RestoreAllProducts()

	if err != nil {
		s.logger.Error("Failed to restore all trashed products",
			zap.Error(err))

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all products",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("All trashed products restored successfully",
		zap.Bool("success", success))

	return success, nil
}

func (s *productService) DeleteAllProductsPermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all trashed products")

	success, err := s.productRepository.DeleteAllProductPermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all trashed products",
			zap.Error(err))

		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all products",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("All trashed products permanently deleted successfully",
		zap.Bool("success", success))

	return success, nil
}
