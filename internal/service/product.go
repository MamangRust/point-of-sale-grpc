package service

import (
	"context"
	"os"
	product_cache "pointofsale/internal/cache/product"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/errorhandler"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors/category_errors"
	"pointofsale/pkg/errors/merchant_errors"
	"pointofsale/pkg/errors/product_errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/pkg/utils"

	"go.opentelemetry.io/otel/attribute"

	"go.uber.org/zap"
)

type productService struct {
	categoryRepository repository.CategoryRepository
	merchantRepository repository.MerchantRepository
	productRepository  repository.ProductRepository
	logger             logger.LoggerInterface
	observability      observability.TraceLoggerObservability
	cache              product_cache.ProductMencache
}

type ProductServiceDeps struct {
	CategoryRepo  repository.CategoryRepository
	MerchantRepo  repository.MerchantRepository
	ProductRepo   repository.ProductRepository
	Logger        logger.LoggerInterface
	Observability observability.TraceLoggerObservability
	cache         product_cache.ProductMencache
}

func NewProductService(deps ProductServiceDeps) *productService {
	return &productService{
		categoryRepository: deps.CategoryRepo,
		merchantRepository: deps.MerchantRepo,
		productRepository:  deps.ProductRepo,
		logger:             deps.Logger,
		observability:      deps.Observability,
		cache:              deps.cache,
	}
}

func (s *productService) FindAllProducts(ctx context.Context, req *requests.FindAllProducts) ([]*db.GetProductsRow, *int, error) {
	const method = "FindAllProducts"

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

	if data, total, found := s.cache.GetCachedProducts(ctx, req); found {
		logSuccess("Successfully retrieved all product records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	products, err := s.productRepository.FindAllProducts(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetProductsRow](
			s.logger,
			product_errors.ErrFailedFindAllProducts,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(products) > 0 {
		totalCount = int(products[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedProducts(ctx, req, products, &totalCount)

	logSuccess("Successfully fetched products",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return products, &totalCount, nil
}

func (s *productService) FindByMerchant(ctx context.Context, req *requests.ProductByMerchantRequest) ([]*db.GetProductsByMerchantRow, *int, error) {
	const method = "FindByMerchant"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search
	merchantId := req.MerchantID
	minPrice := req.MinPrice
	maxPrice := req.MaxPrice

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	if minPrice <= 0 {
		minPrice = 0
	}

	if maxPrice <= 0 {
		maxPrice = 0
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search),
		attribute.Int("merchant_id", merchantId),
		attribute.Int("minPrice", minPrice),
		attribute.Int("maxPrice", maxPrice))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedProductsByMerchant(ctx, req); found {
		logSuccess("Successfully retrieved merchant product records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.Int("merchant_id", merchantId))
		return data, total, nil
	}

	products, err := s.productRepository.FindByMerchant(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetProductsByMerchantRow](
			s.logger,
			product_errors.ErrFailedFindProductsByMerchant,
			method,
			span,
			zap.Int("merchant_id", merchantId),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(products) > 0 {
		totalCount = int(products[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedProductsByMerchant(ctx, req, products, &totalCount)

	logSuccess("Successfully fetched products",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return products, &totalCount, nil
}

func (s *productService) FindByCategory(ctx context.Context, req *requests.ProductByCategoryRequest) ([]*db.GetProductsByCategoryNameRow, *int, error) {
	const method = "FindByCategory"

	page := req.Page
	pageSize := req.PageSize
	search := req.Search
	category_name := req.CategoryName
	minPrice := req.MinPrice
	maxPrice := req.MaxPrice

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	if minPrice <= 0 {
		minPrice = 0
	}

	if maxPrice <= 0 {
		maxPrice = 0
	}

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("page", page),
		attribute.Int("pageSize", pageSize),
		attribute.String("search", search),
		attribute.String("category_name", category_name),
		attribute.Int("minPrice", minPrice),
		attribute.Int("maxPrice", maxPrice))

	defer func() {
		end(status)
	}()

	if data, total, found := s.cache.GetCachedProductsByCategory(ctx, req); found {
		logSuccess("Successfully retrieved category product records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize),
			zap.String("category_name", category_name))
		return data, total, nil
	}

	products, err := s.productRepository.FindByCategory(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetProductsByCategoryNameRow](
			s.logger,
			product_errors.ErrFailedFindProductsByCategory,
			method,
			span,
			zap.String("category_name", category_name),
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(products) > 0 {
		totalCount = int(products[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedProductsByCategory(ctx, req, products, &totalCount)

	logSuccess("Successfully fetched products",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return products, &totalCount, nil
}

func (s *productService) FindById(ctx context.Context, productID int) (*db.GetProductByIDRow, error) {
	const method = "FindById"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("product_id", productID))

	defer func() {
		end(status)
	}()

	if data, found := s.cache.GetCachedProduct(ctx, productID); found {
		logSuccess("Successfully retrieved product record from cache",
			zap.Int("product_id", productID))
		return data, nil
	}

	product, err := s.productRepository.FindById(ctx, productID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.GetProductByIDRow](
			s.logger,
			product_errors.ErrFailedFindProductById,
			method,
			span,
			zap.Int("product_id", productID))
	}

	s.cache.SetCachedProduct(ctx, product)

	logSuccess("Successfully fetched product",
		zap.Int("product_id", productID))

	return product, nil
}

func (s *productService) FindByActive(ctx context.Context, req *requests.FindAllProducts) ([]*db.GetProductsActiveRow, *int, error) {
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

	if data, total, found := s.cache.GetCachedProductActive(ctx, req); found {
		logSuccess("Successfully retrieved active product records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	products, err := s.productRepository.FindByActive(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetProductsActiveRow](
			s.logger,
			product_errors.ErrFailedFindProductsByActive,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(products) > 0 {
		totalCount = int(products[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedProductActive(ctx, req, products, &totalCount)

	logSuccess("Successfully fetched active products",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return products, &totalCount, nil
}

func (s *productService) FindByTrashed(ctx context.Context, req *requests.FindAllProducts) ([]*db.GetProductsTrashedRow, *int, error) {
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

	if data, total, found := s.cache.GetCachedProductTrashed(ctx, req); found {
		logSuccess("Successfully retrieved trashed product records from cache",
			zap.Int("totalRecords", *total),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
		return data, total, nil
	}

	products, err := s.productRepository.FindByTrashed(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandlerErrorPagination[[]*db.GetProductsTrashedRow](
			s.logger,
			product_errors.ErrFailedFindProductsByTrashed,
			method,
			span,
			zap.String("search", search),
			zap.Int("page", page),
			zap.Int("pageSize", pageSize))
	}

	var totalCount int

	if len(products) > 0 {
		totalCount = int(products[0].TotalCount)
	} else {
		totalCount = 0
	}

	s.cache.SetCachedProductTrashed(ctx, req, products, &totalCount)

	logSuccess("Successfully fetched trashed products",
		zap.Int("totalRecords", totalCount),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	return products, &totalCount, nil
}

func (s *productService) CreateProduct(ctx context.Context, req *requests.CreateProductRequest) (*db.CreateProductRow, error) {
	const method = "CreateProduct"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.String("name", req.Name),
		attribute.Int("categoryID", req.CategoryID),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	_, err := s.categoryRepository.FindById(ctx, req.CategoryID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateProductRow](
			s.logger,
			category_errors.ErrFailedFindCategoryById,
			method,
			span,
			zap.Int("categoryID", req.CategoryID))
	}

	_, err = s.merchantRepository.FindById(ctx, req.MerchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateProductRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchantID", req.MerchantID))
	}

	barcode := utils.GenerateBarcode(req.Name)
	slug := utils.GenerateSlug(req.Name)
	req.Barcode = &barcode
	req.SlugProduct = &slug

	product, err := s.productRepository.CreateProduct(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.CreateProductRow](
			s.logger,
			product_errors.ErrFailedCreateProduct,
			method,
			span)
	}

	s.cache.DeleteCachedProduct(ctx, int(product.ProductID))

	logSuccess("Successfully created product",
		zap.Int("productID", int(product.ProductID)),
		zap.String("name", req.Name))

	return product, nil
}

func (s *productService) UpdateProduct(ctx context.Context, req *requests.UpdateProductRequest) (*db.UpdateProductRow, error) {
	const method = "UpdateProduct"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("productID", *req.ProductID),
		attribute.Int("categoryID", req.CategoryID),
		attribute.Int("merchantID", req.MerchantID))

	defer func() {
		end(status)
	}()

	_, err := s.categoryRepository.FindById(ctx, req.CategoryID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateProductRow](
			s.logger,
			category_errors.ErrFailedFindCategoryById,
			method,
			span,
			zap.Int("categoryID", req.CategoryID))
	}

	_, err = s.merchantRepository.FindById(ctx, req.MerchantID)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateProductRow](
			s.logger,
			merchant_errors.ErrFailedFindMerchantById,
			method,
			span,
			zap.Int("merchantID", req.MerchantID))
	}

	barcode := utils.GenerateBarcode(req.Name)
	slug := utils.GenerateSlug(req.Name)
	req.Barcode = &barcode
	req.SlugProduct = &slug

	product, err := s.productRepository.UpdateProduct(ctx, req)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.UpdateProductRow](
			s.logger,
			product_errors.ErrFailedUpdateProduct,
			method,
			span)
	}

	s.cache.DeleteCachedProduct(ctx, int(product.ProductID))

	logSuccess("Successfully updated product",
		zap.Int("productID", int(product.ProductID)),
		zap.String("name", req.Name))

	return product, nil
}

func (s *productService) TrashedProduct(ctx context.Context, product_id int) (*db.Product, error) {
	const method = "TrashedProduct"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("product_id", product_id))

	defer func() {
		end(status)
	}()

	product, err := s.productRepository.TrashedProduct(ctx, product_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Product](
			s.logger,
			product_errors.ErrFailedTrashProduct,
			method,
			span,
			zap.Int("product_id", product_id))
	}

	s.cache.DeleteCachedProduct(ctx, product_id)

	logSuccess("Successfully trashed product",
		zap.Int("product_id", product_id))

	return product, nil
}

func (s *productService) RestoreProduct(ctx context.Context, product_id int) (*db.Product, error) {
	const method = "RestoreProduct"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("product_id", product_id))

	defer func() {
		end(status)
	}()

	product, err := s.productRepository.RestoreProduct(ctx, product_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[*db.Product](
			s.logger,
			product_errors.ErrFailedRestoreProduct,
			method,
			span,
			zap.Int("product_id", product_id))
	}

	s.cache.DeleteCachedProduct(ctx, product_id)

	logSuccess("Successfully restored product",
		zap.Int("product_id", product_id))

	return product, nil
}

func (s *productService) DeleteProductPermanent(ctx context.Context, product_id int) (bool, error) {
	const method = "DeleteProductPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method,
		attribute.Int("product_id", product_id))

	defer func() {
		end(status)
	}()

	res, err := s.productRepository.FindByIdTrashed(ctx, product_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			product_errors.ErrFailedFindProductByTrashed,
			method,
			span,
			zap.Int("product_id", product_id))
	}

	if res.ImageProduct == nil {
		err := os.Remove(*res.ImageProduct)
		if err != nil {
			if os.IsNotExist(err) {
				s.logger.Debug("Product image file not found, continuing with product deletion",
					zap.String("image_path", *res.ImageProduct))
			} else {
				status = "error"
				return errorhandler.HandleError[bool](
					s.logger,
					product_errors.ErrFailedDeleteImageProduct,
					method,
					span,
					zap.String("image_path", *res.ImageProduct))
			}
		} else {
			s.logger.Debug("Successfully deleted product image",
				zap.String("image_path", *res.ImageProduct))
		}
	}

	success, err := s.productRepository.DeleteProductPermanent(ctx, product_id)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			product_errors.ErrFailedDeleteProductPermanent,
			method,
			span,
			zap.Int("product_id", product_id))
	}

	if !success {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			product_errors.ErrFailedDeletingNotFoundProduct,
			method,
			span,
			zap.Int("product_id", product_id))
	}

	s.cache.DeleteCachedProduct(ctx, product_id)

	logSuccess("Successfully permanently deleted product",
		zap.Int("product_id", product_id))

	return true, nil
}

func (s *productService) RestoreAllProducts(ctx context.Context) (bool, error) {
	const method = "RestoreAllProducts"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.productRepository.RestoreAllProducts(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			product_errors.ErrFailedRestoreAllProducts,
			method,
			span)
	}

	s.logger.Debug("All product caches should be invalidated after restore all operation")

	logSuccess("Successfully restored all trashed products",
		zap.Bool("success", success))

	return success, nil
}

func (s *productService) DeleteAllProductPermanent(ctx context.Context) (bool, error) {
	const method = "DeleteAllProductPermanent"

	ctx, span, end, status, logSuccess := s.observability.StartTracingAndLogging(ctx, method)

	defer func() {
		end(status)
	}()

	success, err := s.productRepository.DeleteAllProductPermanent(ctx)
	if err != nil {
		status = "error"
		return errorhandler.HandleError[bool](
			s.logger,
			product_errors.ErrFailedDeleteAllProductsPermanent,
			method,
			span)
	}

	s.logger.Debug("All product caches should be invalidated after delete all operation")

	logSuccess("Successfully permanently deleted all trashed products",
		zap.Bool("success", success))

	return success, nil
}
