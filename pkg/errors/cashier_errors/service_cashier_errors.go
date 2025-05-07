package cashier_errors

import (
	"net/http"
	"pointofsale/internal/domain/response"
)

var (
	ErrFailedFindMonthlyTotalSales           = response.NewErrorResponse("Failed to find monthly total sales", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalSales            = response.NewErrorResponse("Failed to find yearly total sales", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalSalesById       = response.NewErrorResponse("Failed to find monthly total sales by ID", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalSalesById        = response.NewErrorResponse("Failed to find yearly total sales by ID", http.StatusInternalServerError)
	ErrFailedFindMonthlyTotalSalesByMerchant = response.NewErrorResponse("Failed to find monthly total sales by merchant", http.StatusInternalServerError)
	ErrFailedFindYearlyTotalSalesByMerchant  = response.NewErrorResponse("Failed to find yearly total sales by merchant", http.StatusInternalServerError)

	ErrFailedFindMonthlySales             = response.NewErrorResponse("Failed to find monthly sales", http.StatusInternalServerError)
	ErrFailedFindYearlySales              = response.NewErrorResponse("Failed to find yearly sales", http.StatusInternalServerError)
	ErrFailedFindMonthlyCashierByMerchant = response.NewErrorResponse("Failed to find monthly cashier sales by merchant", http.StatusInternalServerError)
	ErrFailedFindYearlyCashierByMerchant  = response.NewErrorResponse("Failed to find yearly cashier sales by merchant", http.StatusInternalServerError)
	ErrFailedFindMonthlyCashierById       = response.NewErrorResponse("Failed to find monthly cashier sales by ID", http.StatusInternalServerError)
	ErrFailedFindYearlyCashierById        = response.NewErrorResponse("Failed to find yearly cashier sales by ID", http.StatusInternalServerError)

	ErrFailedFindAllCashiers       = response.NewErrorResponse("Failed to find all cashiers", http.StatusInternalServerError)
	ErrFailedFindCashierById       = response.NewErrorResponse("Failed to find cashier by ID", http.StatusInternalServerError)
	ErrFailedFindCashierByActive   = response.NewErrorResponse("Failed to find active cashiers", http.StatusInternalServerError)
	ErrFailedFindCashierByTrashed  = response.NewErrorResponse("Failed to find trashed cashiers", http.StatusInternalServerError)
	ErrFailedFindCashierByMerchant = response.NewErrorResponse("Failed to find cashiers by merchant", http.StatusInternalServerError)

	ErrFailedCreateCashier             = response.NewErrorResponse("Failed to create cashier", http.StatusInternalServerError)
	ErrFailedUpdateCashier             = response.NewErrorResponse("Failed to update cashier", http.StatusInternalServerError)
	ErrFailedTrashedCashier            = response.NewErrorResponse("Failed to trash cashier", http.StatusInternalServerError)
	ErrFailedRestoreCashier            = response.NewErrorResponse("Failed to restore cashier", http.StatusInternalServerError)
	ErrFailedDeleteCashierPermanent    = response.NewErrorResponse("Failed to permanently delete cashier", http.StatusInternalServerError)
	ErrFailedRestoreAllCashiers        = response.NewErrorResponse("Failed to restore all cashiers", http.StatusInternalServerError)
	ErrFailedDeleteAllCashierPermanent = response.NewErrorResponse("Failed to permanently delete all cashiers", http.StatusInternalServerError)
)
