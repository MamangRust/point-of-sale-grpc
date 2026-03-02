package merchant_errors

import (
	"pointofsale/pkg/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcInvalidID = errors.NewGrpcError("invalid ID", int(codes.InvalidArgument))

	ErrGrpcValidateCreateMerchant = errors.NewGrpcError("validation failed: invalid create merchant request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant = errors.NewGrpcError("validation failed: invalid update merchant request", int(codes.InvalidArgument))
)
