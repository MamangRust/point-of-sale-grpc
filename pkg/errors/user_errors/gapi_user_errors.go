package user_errors

import (
	"pointofsale/pkg/errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcUserNotFound  = errors.NewGrpcError("User not found", int(codes.NotFound))
	ErrGrpcUserInvalidId = errors.NewGrpcError("Invalid User ID", int(codes.NotFound))

	ErrGrpcFailedFindAll     = errors.NewGrpcError("Failed to fetch Users", int(codes.Internal))
	ErrGrpcFailedFindActive  = errors.NewGrpcError("Failed to fetch active Users", int(codes.Internal))
	ErrGrpcFailedFindTrashed = errors.NewGrpcError("Failed to fetch trashed Users", int(codes.Internal))

	ErrGrpcFailedCreateUser   = errors.NewGrpcError("Failed to create User", int(codes.Internal))
	ErrGrpcFailedUpdateUser   = errors.NewGrpcError("Failed to update User", int(codes.Internal))
	ErrGrpcValidateCreateUser = errors.NewGrpcError("validation failed: invalid create User request", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateUser = errors.NewGrpcError("validation failed: invalid update User request", int(codes.InvalidArgument))

	ErrGrpcFailedTrashedUser     = errors.NewGrpcError("Failed to move User to trash", int(codes.Internal))
	ErrGrpcFailedRestoreUser     = errors.NewGrpcError("Failed to restore User", int(codes.Internal))
	ErrGrpcFailedDeletePermanent = errors.NewGrpcError("Failed to delete User permanently", int(codes.Internal))

	ErrGrpcFailedRestoreAll = errors.NewGrpcError("Failed to restore all Users", int(codes.Internal))
	ErrGrpcFailedDeleteAll  = errors.NewGrpcError("Failed to delete all Users permanently", int(codes.Internal))
)
