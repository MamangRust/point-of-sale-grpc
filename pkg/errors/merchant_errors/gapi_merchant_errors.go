package merchant_errors

import (
	"pointofsale/internal/domain/response"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = response.NewGrpcError("error", "invalid ID", int(codes.InvalidArgument))

	ErrGrpcValidateCreateMerchant = response.NewGrpcError("error", "validation failed: invalid create merchant request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant = response.NewGrpcError("error", "validation failed: invalid update merchant request", int(codes.InvalidArgument))
)
