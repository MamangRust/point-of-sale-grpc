package category_errors

import (
	"pointofsale/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcFailedInvalidId    = response.NewGrpcError("error", "Invalid ID", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidYear  = response.NewGrpcError("error", "Invalid year", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidMonth = response.NewGrpcError("error", "Invalid month", int(codes.InvalidArgument))

	ErrGrpcFailedFindMonthlyTotalPrices           = response.NewGrpcError("error", "Failed to find monthly total prices", int(codes.Internal))
	ErrGrpcFailedFindYearlyTotalPrices            = response.NewGrpcError("error", "Failed to find yearly total prices", int(codes.Internal))
	ErrGrpcFailedFindMonthlyTotalPricesById       = response.NewGrpcError("error", "Failed to find monthly total prices by ID", int(codes.Internal))
	ErrGrpcFailedFindYearlyTotalPricesById        = response.NewGrpcError("error", "Failed to find yearly total prices by ID", int(codes.Internal))
	ErrGrpcFailedFindMonthlyTotalPricesByMerchant = response.NewGrpcError("error", "Failed to find monthly total prices by merchant", int(codes.Internal))
	ErrGrpcFailedFindYearlyTotalPricesByMerchant  = response.NewGrpcError("error", "Failed to find yearly total prices by merchant", int(codes.Internal))

	ErrGrpcFailedFindMonthPrice           = response.NewGrpcError("error", "Failed to find monthly category price", int(codes.Internal))
	ErrGrpcFailedFindYearPrice            = response.NewGrpcError("error", "Failed to find yearly category price", int(codes.Internal))
	ErrGrpcFailedFindMonthPriceByMerchant = response.NewGrpcError("error", "Failed to find monthly category price by merchant", int(codes.Internal))
	ErrGrpcFailedFindYearPriceByMerchant  = response.NewGrpcError("error", "Failed to find yearly category price by merchant", int(codes.Internal))
	ErrGrpcFailedFindMonthPriceById       = response.NewGrpcError("error", "Failed to find monthly category price by ID", int(codes.Internal))
	ErrGrpcFailedFindYearPriceById        = response.NewGrpcError("error", "Failed to find yearly category price by ID", int(codes.Internal))

	ErrGrpcFailedFindCategoryByActive  = response.NewGrpcError("error", "Failed to find active categories", int(codes.Internal))
	ErrGrpcFailedFindCategoryByTrashed = response.NewGrpcError("error", "Failed to find trashed categories", int(codes.Internal))

	ErrGrpcFailedFindAllCategory  = response.NewGrpcError("error", "Failed to find all categories", int(codes.Internal))
	ErrGrpcFailedFindCategoryById = response.NewGrpcError("error", "Failed to find category by ID", int(codes.Internal))

	ErrGrpcFailedCreateCategory          = response.NewGrpcError("error", "Failed to create category", int(codes.Internal))
	ErrGrpcFailedUpdateCategory          = response.NewGrpcError("error", "Failed to update category", int(codes.Internal))
	ErrGrpcFailedTrashedCategory         = response.NewGrpcError("error", "Failed to trash category", int(codes.Internal))
	ErrGrpcFailedRestoreCategory         = response.NewGrpcError("error", "Failed to restore category", int(codes.Internal))
	ErrGrpcFailedDeleteCategoryPermanent = response.NewGrpcError("error", "Failed to permanently delete category", int(codes.Internal))

	ErrGrpcFailedRestoreAllCategory         = response.NewGrpcError("error", "Failed to restore all categories", int(codes.Internal))
	ErrGrpcFailedDeleteAllCategoryPermanent = response.NewGrpcError("error", "Failed to permanently delete all categories", int(codes.Internal))

	ErrGrpcValidateCreateCategory = response.NewGrpcError("error", "validation failed: invalid create category request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateCategory = response.NewGrpcError("error", "validation failed: invalid update category request", int(codes.InvalidArgument))
)
