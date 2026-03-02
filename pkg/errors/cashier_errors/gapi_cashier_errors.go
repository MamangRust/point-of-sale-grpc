package cashier_errors

import (
	"pointofsale/pkg/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcFailedInvalidId         = errors.NewGrpcError("Invalid ID", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidMerchantId = errors.NewGrpcError("Invalid merchant ID", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidYear       = errors.NewGrpcError("Invalid year", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidMonth      = errors.NewGrpcError("Invalid month", int(codes.InvalidArgument))

	ErrGrpcValidateCreateCashier = errors.NewGrpcError("validation failed: invalid create cashier request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateCashier = errors.NewGrpcError("validation failed: invalid update cashier request", int(codes.InvalidArgument))
)
