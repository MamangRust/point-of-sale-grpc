package cashier_errors

import (
	"net/http"
	"pointofsale/pkg/errors"
)

var (
	ErrFailedFindMonthlyTotalSales           = errors.NewErrorResponse("Failed to find monthly total sales", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalSales            = errors.NewErrorResponse("Failed to find yearly total sales", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalSalesById       = errors.NewErrorResponse("Failed to find monthly total sales by ID", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalSalesById        = errors.NewErrorResponse("Failed to find yearly total sales by ID", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalSalesByMerchant = errors.NewErrorResponse("Failed to find monthly total sales by merchant", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalSalesByMerchant  = errors.NewErrorResponse("Failed to find yearly total sales by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthlySales             = errors.NewErrorResponse("Failed to find monthly sales", http.StatusInternalServerError)
	ErrFailedFindYearlySales              = errors.NewErrorResponse("Failed to find yearly sales", http.StatusInternalServerError)
	ErrFailedFindMonthlyCashierByMerchant = errors.NewErrorResponse("Failed to find monthly cashier sales by merchant", http.StatusInternalServerError)
	ErrFailedFindYearlyCashierByMerchant  = errors.NewErrorResponse("Failed to find yearly cashier sales by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthlyCashierById       = errors.NewErrorResponse("Failed to find monthly cashier sales by ID", http.StatusInternalServerError)
	ErrFailedFindYearlyCashierById        = errors.NewErrorResponse("Failed to find yearly cashier sales by ID", http.StatusInternalServerError)

	ErrFailedFindAllCashiers       = errors.NewErrorResponse("Failed to find all cashiers", http.StatusInternalServerError)
	ErrFailedFindCashierById       = errors.NewErrorResponse("Failed to find cashier by ID", http.StatusInternalServerError)
	ErrFailedFindCashierByActive   = errors.NewErrorResponse("Failed to find active cashiers", http.StatusInternalServerError)
	ErrFailedFindCashierByTrashed  = errors.NewErrorResponse("Failed to find trashed cashiers", http.StatusInternalServerError)
	ErrFailedFindCashierByMerchant = errors.NewErrorResponse("Failed to find cashiers by merchant", http.StatusInternalServerError)

	ErrFailedCreateCashier             = errors.NewErrorResponse("Failed to create cashier", http.StatusInternalServerError)
	ErrFailedUpdateCashier             = errors.NewErrorResponse("Failed to update cashier", http.StatusInternalServerError)
	ErrFailedTrashedCashier            = errors.NewErrorResponse("Failed to trash cashier", http.StatusInternalServerError)
	ErrFailedRestoreCashier            = errors.NewErrorResponse("Failed to restore cashier", http.StatusInternalServerError)
	ErrFailedDeleteCashierPermanent    = errors.NewErrorResponse("Failed to permanently delete cashier", http.StatusInternalServerError)
	ErrFailedRestoreAllCashiers        = errors.NewErrorResponse("Failed to restore all cashiers", http.StatusInternalServerError)
	ErrFailedDeleteAllCashierPermanent = errors.NewErrorResponse("Failed to permanently delete all cashiers", http.StatusInternalServerError)
)
