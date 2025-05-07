package merchant_errors

import (
	"pointofsale/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = response.NewGrpcError("error", "invalid ID", int(codes.InvalidArgument))

	ErrGrpcFailedFindAllMerchant       = response.NewGrpcError("error", "Failed to find all merchants", int(codes.Internal))
	ErrGrpcFailedFindMerchantById      = response.NewGrpcError("error", "Failed to find merchant by ID", int(codes.Internal))
	ErrGrpcFailedFindMerchantByActive  = response.NewGrpcError("error", "Failed to find active merchants", int(codes.Internal))
	ErrGrpcFailedFindMerchantByTrashed = response.NewGrpcError("error", "Failed to find trashed merchants", int(codes.Internal))

	ErrGrpcFailedCreateMerchant          = response.NewGrpcError("error", "Failed to create merchant", int(codes.Internal))
	ErrGrpcFailedUpdateMerchant          = response.NewGrpcError("error", "Failed to update merchant", int(codes.Internal))
	ErrGrpcFailedTrashedMerchant         = response.NewGrpcError("error", "Failed to trash merchant", int(codes.Internal))
	ErrGrpcFailedRestoreMerchant         = response.NewGrpcError("error", "Failed to restore merchant", int(codes.Internal))
	ErrGrpcFailedDeleteMerchantPermanent = response.NewGrpcError("error", "Failed to permanently delete merchant", int(codes.Internal))

	ErrGrpcFailedRestoreAllMerchant         = response.NewGrpcError("error", "Failed to restore all merchants", int(codes.Internal))
	ErrGrpcFailedDeleteAllMerchantPermanent = response.NewGrpcError("error", "Failed to permanently delete all merchants", int(codes.Internal))

	ErrGrpcValidateCreateMerchant = response.NewGrpcError("error", "validation failed: invalid create merchant request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant = response.NewGrpcError("error", "validation failed: invalid update merchant request", int(codes.InvalidArgument))
)
