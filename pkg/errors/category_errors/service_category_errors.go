package category_errors

import (
	"net/http"
	"pointofsale/internal/domain/response"
)

var (
	ErrFailedFindMonthlyTotalPrice           = response.NewErrorResponse("Failed to find monthly total price", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalPrice            = response.NewErrorResponse("Failed to find yearly total price", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalPriceById       = response.NewErrorResponse("Failed to find monthly total price by ID", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalPriceById        = response.NewErrorResponse("Failed to find yearly total price by ID", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalPriceByMerchant = response.NewErrorResponse("Failed to find monthly total price by merchant", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalPriceByMerchant  = response.NewErrorResponse("Failed to find yearly total price by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthPrice           = response.NewErrorResponse("Failed to find month price", http.StatusInternalServerError)
	ErrFailedFindYearPrice            = response.NewErrorResponse("Failed to find year price", http.StatusInternalServerError)
	ErrFailedFindMonthPriceByMerchant = response.NewErrorResponse("Failed to find month price by merchant", http.StatusInternalServerError)
	ErrFailedFindYearPriceByMerchant  = response.NewErrorResponse("Failed to find year price by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthPriceById       = response.NewErrorResponse("Failed to find month price by ID", http.StatusInternalServerError)
	ErrFailedFindYearPriceById        = response.NewErrorResponse("Failed to find year price by ID", http.StatusInternalServerError)

	ErrFailedFindAllCategories  = response.NewErrorResponse("Failed to find all categories", http.StatusInternalServerError)
	ErrFailedFindCategoryById   = response.NewErrorResponse("Failed to find category by ID", http.StatusInternalServerError)
	ErrFailedFindCategoryByName = response.NewErrorResponse("Failed to find category by name", http.StatusInternalServerError)

	ErrFailedFindCategoryByActive  = response.NewErrorResponse("Failed to find active categories", http.StatusInternalServerError)
	ErrFailedFindCategoryByTrashed = response.NewErrorResponse("Failed to find trashed categories", http.StatusInternalServerError)

	ErrFailedCreateCategory               = response.NewErrorResponse("Failed to create category", http.StatusInternalServerError)
	ErrFailedUpdateCategory               = response.NewErrorResponse("Failed to update category", http.StatusInternalServerError)
	ErrFailedTrashedCategory              = response.NewErrorResponse("Failed to trash category", http.StatusInternalServerError)
	ErrFailedRestoreCategory              = response.NewErrorResponse("Failed to restore category", http.StatusInternalServerError)
	ErrFailedDeleteCategoryPermanent      = response.NewErrorResponse("Failed to permanently delete category", http.StatusInternalServerError)
	ErrFailedRestoreAllCategories         = response.NewErrorResponse("Failed to restore all categories", http.StatusInternalServerError)
	ErrFailedDeleteAllCategoriesPermanent = response.NewErrorResponse("Failed to permanently delete all categories", http.StatusInternalServerError)
)
