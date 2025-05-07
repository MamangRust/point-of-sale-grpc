package cashier_errors

import (
	"pointofsale/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcFailedInvalidId         = response.NewGrpcError("error", "Invalid ID", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidMerchantId = response.NewGrpcError("error", "Invalid merchant ID", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidYear       = response.NewGrpcError("error", "Invalid year", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidMonth      = response.NewGrpcError("error", "Invalid month", int(codes.InvalidArgument))

	ErrGrpcFailedFindMonthlyTotalSales           = response.NewGrpcError("error", "Failed to find monthly total sales", int(codes.Internal))
	ErrGrpcFailedFindYearlyTotalSales            = response.NewGrpcError("error", "Failed to find yearly total sales", int(codes.Internal))
	ErrGrpcFailedFindMonthlyTotalSalesById       = response.NewGrpcError("error", "Failed to find monthly total sales by ID", int(codes.Internal))
	ErrGrpcFailedFindYearlyTotalSalesById        = response.NewGrpcError("error", "Failed to find yearly total sales by ID", int(codes.Internal))
	ErrGrpcFailedFindMonthlyTotalSalesByMerchant = response.NewGrpcError("error", "Failed to find monthly total sales by merchant", int(codes.Internal))
	ErrGrpcFailedFindYearlyTotalSalesByMerchant  = response.NewGrpcError("error", "Failed to find yearly total sales by merchant", int(codes.Internal))

	ErrGrpcFailedFindAllCashier  = response.NewGrpcError("error", "Failed to find all cashiers", int(codes.Internal))
	ErrGrpcFailedFindCashierById = response.NewGrpcError("error", "Failed to find cashier by ID", int(codes.Internal))

	ErrGrpcFailedFindMonthSales           = response.NewGrpcError("error", "Failed to find monthly cashier sales", int(codes.Internal))
	ErrGrpcFailedFindYearSales            = response.NewGrpcError("error", "Failed to find yearly cashier sales", int(codes.Internal))
	ErrGrpcFailedFindMonthSalesByMerchant = response.NewGrpcError("error", "Failed to find monthly cashier sales by merchant", int(codes.Internal))
	ErrGrpcFailedFindYearSalesByMerchant  = response.NewGrpcError("error", "Failed to find yearly cashier sales by merchant", int(codes.Internal))
	ErrGrpcFailedFindMonthSalesById       = response.NewGrpcError("error", "Failed to find monthly cashier sales by ID", int(codes.Internal))
	ErrGrpcFailedFindYearSalesById        = response.NewGrpcError("error", "Failed to find yearly cashier sales by ID", int(codes.Internal))

	ErrGrpcFailedFindCashierByActive   = response.NewGrpcError("error", "Failed to find active cashiers", int(codes.Internal))
	ErrGrpcFailedFindCashierByTrashed  = response.NewGrpcError("error", "Failed to find trashed cashiers", int(codes.Internal))
	ErrGrpcFailedFindCashierByMerchant = response.NewGrpcError("error", "Failed to find cashiers by merchant", int(codes.Internal))

	ErrGrpcFailedCreateCashier = response.NewGrpcError("error", "Failed to create cashier", int(codes.Internal))
	ErrGrpcFailedUpdateCashier = response.NewGrpcError("error", "Failed to update cashier", int(codes.Internal))

	ErrGrpcFailedTrashedCashier         = response.NewGrpcError("error", "Failed to trash cashier", int(codes.Internal))
	ErrGrpcFailedRestoreCashier         = response.NewGrpcError("error", "Failed to restore cashier", int(codes.Internal))
	ErrGrpcFailedDeleteCashierPermanent = response.NewGrpcError("error", "Failed to permanently delete cashier", int(codes.Internal))

	ErrGrpcFailedRestoreAllCashier         = response.NewGrpcError("error", "Failed to restore all cashiers", int(codes.Internal))
	ErrGrpcFailedDeleteAllCashierPermanent = response.NewGrpcError("error", "Failed to permanently delete all cashiers", int(codes.Internal))

	ErrGrpcValidateCreateCashier = response.NewGrpcError("error", "validation failed: invalid create cashier request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateCashier = response.NewGrpcError("error", "validation failed: invalid update cashier request", int(codes.InvalidArgument))
)
