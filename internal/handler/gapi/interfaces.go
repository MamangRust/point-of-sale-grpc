package gapi

import (
	"context"
	"pointofsale/internal/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthHandleGrpc interface {
	pb.AuthServiceServer
	LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.ApiResponseLogin, error)
	RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.ApiResponseRegister, error)
}

type RoleHandleGrpc interface {
	pb.RoleServiceServer

	FindAllRole(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRole, error)
	FindByIdRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error)
	FindByUserId(ctx context.Context, req *pb.FindByIdUserRoleRequest) (*pb.ApiResponsesRole, error)
	FindByActive(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error)
	FindByTrashed(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error)
	CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.ApiResponseRole, error)
	UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.ApiResponseRole, error)
	TrashedRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error)
	RestoreRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error)
	DeleteRolePermanent(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDelete, error)
	RestoreAllRole(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseRoleAll, error)
	DeleteAllRolePermanent(ctx context.Context, req *emptypb.Empty) (*pb.ApiResponseRoleAll, error)
}

type UserHandleGrpc interface {
	pb.UserServiceServer

	FindAll(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUser, error)
	FindById(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error)
	FindByActive(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUserDeleteAt, error)
	FindByTrashed(ctx context.Context, request *pb.FindAllUserRequest) (*pb.ApiResponsePaginationUserDeleteAt, error)
	Create(ctx context.Context, request *pb.CreateUserRequest) (*pb.ApiResponseUser, error)
	Update(ctx context.Context, request *pb.UpdateUserRequest) (*pb.ApiResponseUser, error)
	TrashedUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error)
	RestoreUser(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUser, error)
	DeleteUserPermanent(ctx context.Context, request *pb.FindByIdUserRequest) (*pb.ApiResponseUserDelete, error)

	RestoreAllUser(context.Context, *emptypb.Empty) (*pb.ApiResponseUserAll, error)
	DeleteAllUserPermanent(context.Context, *emptypb.Empty) (*pb.ApiResponseUserAll, error)
}
