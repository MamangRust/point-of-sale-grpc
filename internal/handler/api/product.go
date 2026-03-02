package api

import (
	"net/http"
	product_cache "pointofsale/internal/cache/api/product"
	"pointofsale/internal/domain/requests"
	response_api "pointofsale/internal/mapper"
	"pointofsale/internal/pb"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/upload_image"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productHandleApi struct {
	client       pb.ProductServiceClient
	logger       logger.LoggerInterface
	mapping      response_api.ProductResponseMapper
	upload_image upload_image.ImageUploads
	apiHandler   errors.ApiHandler
	cache        product_cache.ProductMencache
}

func NewHandlerProduct(
	router *echo.Echo,
	client pb.ProductServiceClient,
	logger logger.LoggerInterface,
	mapping response_api.ProductResponseMapper,
	upload_image upload_image.ImageUploads,
	apiHandler errors.ApiHandler,
	cache product_cache.ProductMencache,
) *productHandleApi {
	productHandler := &productHandleApi{
		client:       client,
		logger:       logger,
		mapping:      mapping,
		upload_image: upload_image,
		apiHandler:   apiHandler,
		cache:        cache,
	}

	routerProduct := router.Group("/api/product")

	routerProduct.GET("", productHandler.FindAllProduct)
	routerProduct.GET("/:id", productHandler.FindById)
	routerProduct.GET("/merchant/:merchant_id", productHandler.FindByMerchant)
	routerProduct.GET("/category/:category_name", productHandler.FindByCategory)

	routerProduct.GET("/active", productHandler.FindByActive)
	routerProduct.GET("/trashed", productHandler.FindByTrashed)

	routerProduct.POST("/create", apiHandler.Handle("create", productHandler.Create))
	routerProduct.POST("/update/:id", apiHandler.Handle("update", productHandler.Update))

	routerProduct.POST("/trashed/:id", apiHandler.Handle("trashed", productHandler.TrashedProduct))
	routerProduct.POST("/restore/:id", apiHandler.Handle("restore", productHandler.RestoreProduct))
	routerProduct.DELETE("/permanent/:id", apiHandler.Handle("delete", productHandler.DeleteProductPermanent))

	routerProduct.POST("/restore/all", apiHandler.Handle("restore-all", productHandler.RestoreAllProduct))
	routerProduct.POST("/permanent/all", apiHandler.Handle("delete-all", productHandler.DeleteAllProductPermanent))

	return productHandler
}

// @Security Bearer
// @Summary Find all products
// @Tags Product
// @Description Retrieve a list of all products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationProduct "List of products"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve product data"
// @Router /api/product [get]
func (h *productHandleApi) FindAllProduct(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllProducts{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedProducts(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllProductRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindAll(ctx, grpcReq)

	if err != nil {
		h.logger.Debug("Failed to retrieve product data", zap.Error(err))
		return h.handleGrpcError(err, "FindAllProduct")
	}

	so := h.mapping.ToApiResponsePaginationProduct(res)

	h.cache.SetCachedProducts(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find products by merchant
// @Tags Product
// @Description Retrieve a list of products filtered by merchant
// @Accept json
// @Produce json
// @Param merchant_id path int true "Merchant ID"
// @Param page query int false "Page number" minimum(1) default(1)
// @Param page_size query int false "Number of items per page" minimum(1) maximum(100) default(10)
// @Param search query string false "Search query"
// @Param category_id query int false "Category ID filter"
// @Param min_price query int false "Minimum price filter"
// @Param max_price query int false "Maximum price filter"
// @Success 200 {object} response.ApiResponsePaginationProduct "List of products"
// @Failure 400 {object} response.ErrorResponse "Invalid request parameters"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve product data"
// @Router /api/product/merchant/{merchant_id} [get]
func (h *productHandleApi) FindByMerchant(c echo.Context) error {
	merchantID, err := strconv.Atoi(c.Param("merchant_id"))
	if err != nil || merchantID <= 0 {
		h.logger.Debug("Invalid merchant ID", zap.Error(err))
		return errors.NewBadRequestError("Invalid merchant ID")
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 12
	}

	req := &requests.ProductByMerchantRequest{
		MerchantID: merchantID,
		Page:       page,
		PageSize:   pageSize,
		Search:     strings.TrimSpace(c.QueryParam("search")),
	}

	if c.QueryParam("category_id") != "" {
		if id, err := strconv.Atoi(c.QueryParam("category_id")); err == nil && id > 0 {
			req.CategoryID = id
		}
	}

	if minPriceStr := c.QueryParam("min_price"); minPriceStr != "" {
		if price, err := strconv.Atoi(minPriceStr); err == nil && price >= 0 {
			req.MinPrice = price
		}
	}

	if maxPriceStr := c.QueryParam("max_price"); maxPriceStr != "" {
		if price, err := strconv.Atoi(maxPriceStr); err == nil && price >= 0 {
			req.MaxPrice = price
		}
	}
	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedProductsByMerchant(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllProductMerchantRequest{
		MerchantId: int32(merchantID),
		Page:       int32(page),
		PageSize:   int32(pageSize),
		Search:     req.Search,
		CategoryId: int32(req.CategoryID),
		MinPrice:   int32(req.MinPrice),
		MaxPrice:   int32(req.MaxPrice),
	}

	res, err := h.client.FindByMerchant(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to retrieve product data by merchant",
			zap.Error(err),
			zap.Int("merchant_id", merchantID),
			zap.Any("request", grpcReq),
		)

		return h.handleGrpcError(err, "FindByMerchant")
	}

	so := h.mapping.ToApiResponsePaginationProduct(res)

	h.cache.SetCachedProductsByMerchant(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find products by category
// @Tags Product
// @Description Retrieve a list of products filtered by category
// @Accept json
// @Produce json
// @Param category_name query string true "Category Name"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationProduct "List of products"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve product data"
// @Router /api/product/category [get]
func (h *productHandleApi) FindByCategory(c echo.Context) error {
	categoryName := c.Param("category_name")
	if categoryName == "" {
		return errors.NewBadRequestError("category name is required")
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.ProductByCategoryRequest{
		CategoryName: categoryName,
		Page:         page,
		PageSize:     pageSize,
		Search:       search,
	}

	if cached, found := h.cache.GetCachedProductsByCategory(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllProductCategoryRequest{
		CategoryName: categoryName,
		Page:         int32(page),
		PageSize:     int32(pageSize),
		Search:       search,
	}

	res, err := h.client.FindByCategory(ctx, grpcReq)
	if err != nil {
		h.logger.Debug("Failed to retrieve product data by category", zap.Error(err))
		return h.handleGrpcError(err, "FindByCategory")
	}

	so := h.mapping.ToApiResponsePaginationProduct(res)

	h.cache.SetCachedProductsByCategory(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Find product by ID
// @Tags Product
// @Description Retrieve a product by ID
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.ApiResponseProduct "Product data"
// @Failure 400 {object} response.ErrorResponse "Invalid product ID"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve product data"
// @Router /api/product/{id} [get]
func (h *productHandleApi) FindById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid product ID", zap.Error(err))
		return errors.NewBadRequestError("Invalid product ID")
	}

	ctx := c.Request().Context()

	if cached, found := h.cache.GetCachedProduct(ctx, id); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindByIdProductRequest{
		Id: int32(id),
	}

	res, err := h.client.FindById(ctx, grpcReq)

	if err != nil {
		h.logger.Debug("Failed to retrieve product data", zap.Error(err))
		return h.handleGrpcError(err, "FindById")
	}

	so := h.mapping.ToApiResponseProduct(res)

	h.cache.SetCachedProduct(ctx, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve active products
// @Tags Product
// @Description Retrieve a list of active products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationProductDeleteAt "List of active products"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve product data"
// @Router /api/product/active [get]
func (h *productHandleApi) FindByActive(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllProducts{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedProductActive(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllProductRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByActive(ctx, grpcReq)

	if err != nil {
		h.logger.Debug("Failed to retrieve product data", zap.Error(err))
		return h.handleGrpcError(err, "FindByActive")
	}

	so := h.mapping.ToApiResponsePaginationProductDeleteAt(res)

	h.cache.SetCachedProductActive(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Retrieve trashed products
// @Tags Product
// @Description Retrieve a list of trashed products
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Number of items per page" default(10)
// @Param search query string false "Search query"
// @Success 200 {object} response.ApiResponsePaginationProductDeleteAt "List of trashed products"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve product data"
// @Router /api/product/trashed [get]
func (h *productHandleApi) FindByTrashed(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.QueryParam("page_size"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	search := c.QueryParam("search")

	ctx := c.Request().Context()

	req := &requests.FindAllProducts{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	if cached, found := h.cache.GetCachedProductTrashed(ctx, req); found {
		return c.JSON(http.StatusOK, cached)
	}

	grpcReq := &pb.FindAllProductRequest{
		Page:     int32(page),
		PageSize: int32(pageSize),
		Search:   search,
	}

	res, err := h.client.FindByTrashed(ctx, grpcReq)

	if err != nil {
		h.logger.Debug("Failed to retrieve product data", zap.Error(err))
		return h.handleGrpcError(err, "FindByTrashed")
	}

	so := h.mapping.ToApiResponsePaginationProductDeleteAt(res)

	h.cache.SetCachedProductTrashed(ctx, req, so)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// @Summary Create a new product
// @Tags Product
// @Description Create a new product with the provided details and an image file
// @Accept multipart/form-data
// @Produce json
// @Param merchant_id formData string true "Merchant ID"
// @Param category_id formData string true "Category ID"
// @Param name formData string true "Product name"
// @Param description formData string true "Product description"
// @Param price formData number true "Product price"
// @Param count_in_stock formData integer true "Product count in stock"
// @Param brand formData string true "Product brand"
// @Param weight formData number true "Product weight"
// @Param rating formData number true "Product rating"
// @Param slug_product formData string true "Product slug"
// @Param image formData file true "Product image file"
// @Param barcode formData string true "Product barcode"
// @Success 200 {object} response.ApiResponseProduct "Successfully created product"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create product"
// @Router /api/product/create [post]
func (h *productHandleApi) Create(c echo.Context) error {
	formData, err := h.parseProductForm(c, true)
	if err != nil {
		return h.handleGrpcError(err, "Create")
	}

	ctx := c.Request().Context()
	grpcReq := &pb.CreateProductRequest{
		MerchantId:   int32(formData.MerchantID),
		CategoryId:   int32(formData.CategoryID),
		Name:         formData.Name,
		Description:  formData.Description,
		Price:        int32(formData.Price),
		CountInStock: int32(formData.CountInStock),
		Brand:        formData.Brand,
		Weight:       int32(formData.Weight),
		ImageProduct: formData.ImagePath,
	}

	res, err := h.client.Create(ctx, grpcReq)
	if err != nil {
		if formData.ImagePath != "" {
			h.upload_image.CleanupImageOnFailure(formData.ImagePath)
		}

		h.logger.Error("Product creation failed",
			zap.Error(err),
			zap.Any("request", grpcReq),
		)

		return h.handleGrpcError(err, "Create")
	}

	so := h.mapping.ToApiResponseProduct(res)

	// Invalidate cache for related data
	h.cache.DeleteCachedProduct(ctx, int(res.Data.Id))

	return c.JSON(http.StatusCreated, so)
}

// @Security Bearer
// @Summary Update a product
// @Tags Product
// @Description Update a product with the provided details and an image file
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Product ID"
// @Param merchant_id formData string true "Merchant ID"
// @Param category_id formData string true "Category ID"
// @Param name formData string true "Product name"
// @Param description formData string true "Product description"
// @Param price formData number true "Product price"
// @Param count_in_stock formData integer true "Product count in stock"
// @Param brand formData string true "Product brand"
// @Param weight formData number true "Product weight"
// @Param rating formData number true "Product rating"
// @Param slug_product formData string true "Product slug"
// @Param image formData file true "Product image file"
// @Param barcode formData string true "Product barcode"
// @Success 200 {object} response.ApiResponseProduct "Successfully created product"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to create product"
// @Router /api/product/update/{id} [post]
func (h *productHandleApi) Update(c echo.Context) error {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.NewBadRequestError("Invalid product ID")
	}

	formData, err := h.parseProductForm(c, false)
	if err != nil {
		return h.handleGrpcError(err, "Update")
	}

	ctx := c.Request().Context()
	grpcReq := &pb.UpdateProductRequest{
		ProductId:    int32(productID),
		MerchantId:   int32(formData.MerchantID),
		CategoryId:   int32(formData.CategoryID),
		Name:         formData.Name,
		Description:  formData.Description,
		Price:        int32(formData.Price),
		CountInStock: int32(formData.CountInStock),
		Brand:        formData.Brand,
		Weight:       int32(formData.Weight),
		ImageProduct: formData.ImagePath,
	}

	res, err := h.client.Update(ctx, grpcReq)
	if err != nil {
		if formData.ImagePath != "" {
			h.upload_image.CleanupImageOnFailure(formData.ImagePath)
		}

		h.logger.Error("Product update failed",
			zap.Error(err),
			zap.Any("request", grpcReq),
		)

		return h.handleGrpcError(err, "Update")
	}

	so := h.mapping.ToApiResponseProduct(res)

	// Invalidate cache for related data
	h.cache.DeleteCachedProduct(ctx, productID)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// TrashedProduct retrieves a trashed product record by its ID.
// @Summary Retrieve a trashed product
// @Tags Product
// @Description Retrieve a trashed product record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.ApiResponseProductDeleteAt "Successfully retrieved trashed product"
// @Failure 400 {object} response.ErrorResponse "Invalid request body or validation error"
// @Failure 500 {object} response.ErrorResponse "Failed to retrieve trashed product"
// @Router /api/product/trashed/{id} [get]
func (h *productHandleApi) TrashedProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid product ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid product ID")
	}

	ctx := c.Request().Context()

	grpcReq := &pb.FindByIdProductRequest{
		Id: int32(id),
	}

	res, err := h.client.TrashedProduct(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to archive product", zap.Error(err))
		return h.handleGrpcError(err, "TrashedProduct")
	}

	so := h.mapping.ToApiResponsesProductDeleteAt(res)

	// Invalidate cache for related data
	h.cache.DeleteCachedProduct(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreProduct restores a product record from the trash by its ID.
// @Summary Restore a trashed product
// @Tags Product
// @Description Restore a trashed product record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.ApiResponseProductDeleteAt "Successfully restored product"
// @Failure 400 {object} response.ErrorResponse "Invalid product ID"
// @Failure 500 {object} response.ErrorResponse "Failed to restore product"
// @Router /api/product/restore/{id} [post]
func (h *productHandleApi) RestoreProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid product ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid product ID")
	}

	ctx := c.Request().Context()

	grpcReq := &pb.FindByIdProductRequest{
		Id: int32(id),
	}

	res, err := h.client.RestoreProduct(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to restore product", zap.Error(err))
		return h.handleGrpcError(err, "RestoreProduct")
	}

	so := h.mapping.ToApiResponsesProductDeleteAt(res)

	// Invalidate cache for related data
	h.cache.DeleteCachedProduct(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteProductPermanent permanently deletes a product record by its ID.
// @Summary Permanently delete a product
// @Tags Product
// @Description Permanently delete a product record by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.ApiResponseProductDelete "Successfully deleted product record permanently"
// @Failure 400 {object} response.ErrorResponse "Bad Request: Invalid ID"
// @Failure 500 {object} response.ErrorResponse "Failed to delete product:"
// @Router /api/product/delete/{id} [delete]
func (h *productHandleApi) DeleteProductPermanent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		h.logger.Debug("Invalid product ID format", zap.Error(err))
		return errors.NewBadRequestError("Invalid product ID")
	}

	ctx := c.Request().Context()

	grpcReq := &pb.FindByIdProductRequest{
		Id: int32(id),
	}

	res, err := h.client.DeleteProductPermanent(ctx, grpcReq)

	if err != nil {
		h.logger.Error("Failed to permanently delete product", zap.Error(err))
		return h.handleGrpcError(err, "DeleteProductPermanent")
	}

	so := h.mapping.ToApiResponseProductDelete(res)

	// Invalidate cache for related data
	h.cache.DeleteCachedProduct(ctx, id)

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// RestoreAllProduct restores all trashed product records.
// @Summary Restore all trashed products
// @Tags Product
// @Description Restore all trashed product records.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseProductAll "Successfully restored all products"
// @Failure 500 {object} response.ErrorResponse "Failed to restore all products"
// @Router /api/product/restore/all [post]
func (h *productHandleApi) RestoreAllProduct(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.RestoreAllProduct(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Bulk products restoration failed", zap.Error(err))
		return h.handleGrpcError(err, "RestoreAllProduct")
	}

	so := h.mapping.ToApiResponseProductAll(res)

	h.logger.Debug("Successfully restored all products")

	return c.JSON(http.StatusOK, so)
}

// @Security Bearer
// DeleteAllProductPermanent permanently deletes all product records.
// @Summary Permanently delete all products
// @Tags Product
// @Description Permanently delete all product records.
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponseProductAll "Successfully deleted all product records permanently"
// @Failure 500 {object} response.ErrorResponse "Failed to delete all products"
// @Router /api/product/delete/all [post]
func (h *productHandleApi) DeleteAllProductPermanent(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := h.client.DeleteAllProductPermanent(ctx, &emptypb.Empty{})

	if err != nil {
		h.logger.Error("Bulk products deletion failed", zap.Error(err))
		return h.handleGrpcError(err, "DeleteAllProductPermanent")
	}

	so := h.mapping.ToApiResponseProductAll(res)

	h.logger.Debug("Successfully deleted all products permanently")

	return c.JSON(http.StatusOK, so)
}

func (h *productHandleApi) parseProductForm(c echo.Context, requireImage bool) (requests.ProductFormData, error) {
	var formData requests.ProductFormData
	var err error

	if h.upload_image == nil {
		h.logger.Error("upload_image not initialized")
		return formData, errors.NewInternalError(nil).WithMessage("internal server error")
	}

	formData.MerchantID, err = strconv.Atoi(c.FormValue("merchant_id"))
	if err != nil || formData.MerchantID <= 0 {
		return formData, errors.NewBadRequestError("Please provide a valid merchant ID")
	}

	formData.CategoryID, err = strconv.Atoi(c.FormValue("category_id"))
	if err != nil || formData.CategoryID <= 0 {
		return formData, errors.NewBadRequestError("Please provide a valid category ID")
	}

	formData.Name = strings.TrimSpace(c.FormValue("name"))
	if formData.Name == "" {
		return formData, errors.NewBadRequestError("Product name is required")
	}

	formData.Description = strings.TrimSpace(c.FormValue("description"))
	formData.Brand = strings.TrimSpace(c.FormValue("brand"))

	formData.Price, err = strconv.Atoi(c.FormValue("price"))
	if err != nil || formData.Price <= 0 {
		return formData, errors.NewBadRequestError("Please provide a valid positive price")
	}

	formData.CountInStock, err = strconv.Atoi(c.FormValue("count_in_stock"))
	if err != nil || formData.CountInStock < 0 {
		return formData, errors.NewBadRequestError("Please provide a valid stock count (zero or positive)")
	}

	formData.Weight, err = strconv.Atoi(c.FormValue("weight"))
	if err != nil || formData.Weight <= 0 {
		return formData, errors.NewBadRequestError("Please provide a valid positive weight")
	}

	file, err := c.FormFile("image_product")
	if err != nil {
		if requireImage {
			h.logger.Debug("Image upload error", zap.Error(err))
			return formData, errors.NewBadRequestError("A product image is required")
		}

		return formData, nil
	}

	imagePath, err := h.upload_image.ProcessImageUpload(c, file)
	if err != nil {
		return formData, err
	}

	formData.ImagePath = imagePath
	return formData, nil
}

func (h *productHandleApi) handleGrpcError(err error, operation string) *errors.AppError {
	st, ok := status.FromError(err)
	if !ok {
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}

	switch st.Code() {
	case codes.NotFound:
		return errors.NewNotFoundError("Product").WithInternal(err)

	case codes.AlreadyExists:
		return errors.NewConflictError("Product already exists").WithInternal(err)

	case codes.InvalidArgument:
		return errors.NewBadRequestError(st.Message()).WithInternal(err)

	case codes.PermissionDenied:
		return errors.ErrForbidden.WithInternal(err)

	case codes.Unauthenticated:
		return errors.ErrUnauthorized.WithInternal(err)

	case codes.ResourceExhausted:
		return errors.ErrTooManyRequests.WithInternal(err)

	case codes.Unavailable:
		return errors.NewServiceUnavailableError("Product service").WithInternal(err)

	case codes.DeadlineExceeded:
		return errors.ErrTimeout.WithInternal(err)

	default:
		return errors.NewInternalError(err).WithMessage("Failed to " + operation)
	}
}
