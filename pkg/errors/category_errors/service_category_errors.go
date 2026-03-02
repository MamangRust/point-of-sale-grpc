package category_errors

import (
	"net/http"
	"pointofsale/pkg/errors"
)

var (
	ErrFailedFindMonthlyTotalPrice           = errors.NewErrorResponse("Failed to find monthly total price", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalPrice            = errors.NewErrorResponse("Failed to find yearly total price", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalPriceById       = errors.NewErrorResponse("Failed to find monthly total price by ID", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalPriceById        = errors.NewErrorResponse("Failed to find yearly total price by ID", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalPriceByMerchant = errors.NewErrorResponse("Failed to find monthly total price by merchant", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalPriceByMerchant  = errors.NewErrorResponse("Failed to find yearly total price by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthPrice           = errors.NewErrorResponse("Failed to find month price", http.StatusInternalServerError)
	ErrFailedFindYearPrice            = errors.NewErrorResponse("Failed to find year price", http.StatusInternalServerError)
	ErrFailedFindMonthPriceByMerchant = errors.NewErrorResponse("Failed to find month price by merchant", http.StatusInternalServerError)
	ErrFailedFindYearPriceByMerchant  = errors.NewErrorResponse("Failed to find year price by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthPriceById       = errors.NewErrorResponse("Failed to find month price by ID", http.StatusInternalServerError)
	ErrFailedFindYearPriceById        = errors.NewErrorResponse("Failed to find year price by ID", http.StatusInternalServerError)

	ErrFailedFindAllCategories  = errors.NewErrorResponse("Failed to find all categories", http.StatusInternalServerError)
	ErrFailedFindCategoryById   = errors.NewErrorResponse("Failed to find category by ID", http.StatusInternalServerError)
	ErrFailedFindCategoryByName = errors.NewErrorResponse("Failed to find category by name", http.StatusInternalServerError)

	ErrFailedFindCategoryByActive  = errors.NewErrorResponse("Failed to find active categories", http.StatusInternalServerError)
	ErrFailedFindCategoryByTrashed = errors.NewErrorResponse("Failed to find trashed categories", http.StatusInternalServerError)

	ErrFailedCreateCategory               = errors.NewErrorResponse("Failed to create category", http.StatusInternalServerError)
	ErrFailedUpdateCategory               = errors.NewErrorResponse("Failed to update category", http.StatusInternalServerError)
	ErrFailedTrashedCategory              = errors.NewErrorResponse("Failed to trash category", http.StatusInternalServerError)
	ErrFailedRestoreCategory              = errors.NewErrorResponse("Failed to restore category", http.StatusInternalServerError)
	ErrFailedDeleteCategoryPermanent      = errors.NewErrorResponse("Failed to permanently delete category", http.StatusInternalServerError)
	ErrFailedRestoreAllCategories         = errors.NewErrorResponse("Failed to restore all categories", http.StatusInternalServerError)
	ErrFailedDeleteAllCategoriesPermanent = errors.NewErrorResponse("Failed to permanently delete all categories", http.StatusInternalServerError)
)
