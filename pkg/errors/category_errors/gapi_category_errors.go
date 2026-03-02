package category_errors

import (
	"pointofsale/pkg/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcFailedInvalidId    = errors.NewGrpcError("Invalid ID", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidYear  = errors.NewGrpcError("Invalid year", int(codes.InvalidArgument))
	ErrGrpcFailedInvalidMonth = errors.NewGrpcError("Invalid month", int(codes.InvalidArgument))

	ErrGrpcValidateCreateCategory = errors.NewGrpcError("validation failed: invalid create category request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateCategory = errors.NewGrpcError("validation failed: invalid update category request", int(codes.InvalidArgument))
)
