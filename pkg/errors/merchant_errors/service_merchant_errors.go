package merchant_errors

import (
	"net/http"
	"pointofsale/internal/domain/response"
)

var (
	ErrFailedFindAllMerchants            = response.NewErrorResponse("Failed to find all merchants", http.StatusInternalServerError)
	ErrFailedFindMerchantsByActive       = response.NewErrorResponse("Failed to find active merchants", http.StatusInternalServerError)
	ErrFailedFindMerchantsByTrashed      = response.NewErrorResponse("Failed to find trashed merchants", http.StatusInternalServerError)
	ErrFailedFindMerchantById            = response.NewErrorResponse("Failed to find merchant by ID", http.StatusInternalServerError)
	ErrFailedCreateMerchant              = response.NewErrorResponse("Failed to create merchant", http.StatusInternalServerError)
	ErrFailedUpdateMerchant              = response.NewErrorResponse("Failed to update merchant", http.StatusInternalServerError)
	ErrFailedTrashMerchant               = response.NewErrorResponse("Failed to trash merchant", http.StatusInternalServerError)
	ErrFailedRestoreMerchant             = response.NewErrorResponse("Failed to restore merchant", http.StatusInternalServerError)
	ErrFailedDeleteMerchantPermanent     = response.NewErrorResponse("Failed to permanently delete merchant", http.StatusInternalServerError)
	ErrFailedRestoreAllMerchants         = response.NewErrorResponse("Failed to restore all merchants", http.StatusInternalServerError)
	ErrFailedDeleteAllMerchantsPermanent = response.NewErrorResponse("Failed to permanently delete all merchants", http.StatusInternalServerError)
)
