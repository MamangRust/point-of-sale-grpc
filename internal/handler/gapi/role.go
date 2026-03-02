package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/role_errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

type roleHandleGrpc struct {
	pb.UnimplementedRoleServiceServer
	roleService service.RoleService
}

func NewRoleHandleGrpc(role service.RoleService) *roleHandleGrpc {
	return &roleHandleGrpc{
		roleService: role,
	}
}

func (s *roleHandleGrpc) FindAllRole(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRole, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllRoles{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.roleService.FindAll(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var roleResponses []*pb.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationRole{
		Status:     "success",
		Message:    "Successfully fetched roles",
		Data:       roleResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *roleHandleGrpc) FindByIdRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRole, error) {
	id := int(req.GetRoleId())

	if id == 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	role, err := s.roleService.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully fetched role",
		Data: &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *roleHandleGrpc) FindByUserId(ctx context.Context, req *pb.FindByIdUserRoleRequest) (*pb.ApiResponsesRole, error) {
	id := int(req.GetUserId())

	if id == 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	roles, err := s.roleService.FindByUserId(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var roleResponses []*pb.RoleResponse
	for _, role := range roles {
		roleResponses = append(roleResponses, &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsesRole{
		Status:  "success",
		Message: "Successfully fetched roles by user ID",
		Data:    roleResponses,
	}, nil
}

func (s *roleHandleGrpc) FindByActive(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllRoles{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.roleService.FindByActiveRole(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var roleResponses []*pb.RoleResponseDeleteAt
	for _, role := range roles {
		var deletedAt string
		if role.DeletedAt.Valid {
			deletedAt = role.DeletedAt.Time.String()
		}

		roleResponses = append(roleResponses, &pb.RoleResponseDeleteAt{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
			DeletedAt: deletedAt,
		})
	}

	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active roles",
		Data:       roleResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *roleHandleGrpc) FindByTrashed(ctx context.Context, req *pb.FindAllRoleRequest) (*pb.ApiResponsePaginationRoleDeleteAt, error) {
	page := int(req.GetPage())
	pageSize := int(req.GetPageSize())
	search := req.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllRoles{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	roles, totalRecords, err := s.roleService.FindByTrashedRole(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var roleResponses []*pb.RoleResponseDeleteAt
	for _, role := range roles {
		var deletedAt string
		if role.DeletedAt.Valid {
			deletedAt = role.DeletedAt.Time.String()
		}

		roleResponses = append(roleResponses, &pb.RoleResponseDeleteAt{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
			DeletedAt: deletedAt,
		})
	}

	return &pb.ApiResponsePaginationRoleDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed roles",
		Data:       roleResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *roleHandleGrpc) CreateRole(ctx context.Context, req *pb.CreateRoleRequest) (*pb.ApiResponseRole, error) {
	name := req.GetName()

	request := &requests.CreateRoleRequest{
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, role_errors.ErrGrpcValidateCreateRole
	}

	role, err := s.roleService.CreateRole(ctx, &requests.CreateRoleRequest{
		Name: name,
	})
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully created role",
		Data: &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *roleHandleGrpc) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.ApiResponseRole, error) {
	id := int(req.GetId())

	if id == 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	name := req.GetName()

	request := &requests.UpdateRoleRequest{
		ID:   &id,
		Name: name,
	}

	if err := request.Validate(); err != nil {
		return nil, role_errors.ErrGrpcValidateUpdateRole
	}

	role, err := s.roleService.UpdateRole(ctx, request)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRole{
		Status:  "success",
		Message: "Successfully updated role",
		Data: &pb.RoleResponse{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *roleHandleGrpc) TrashedRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDeleteAt, error) {
	id := int(req.GetRoleId())

	if id == 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	role, err := s.roleService.TrashedRole(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleDeleteAt{
		Status:  "success",
		Message: "Successfully trashed role",
		Data: &pb.RoleResponseDeleteAt{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
			DeletedAt: role.DeletedAt.Time.String(),
		},
	}, nil
}

func (s *roleHandleGrpc) RestoreRole(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDeleteAt, error) {
	id := int(req.GetRoleId())

	if id == 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	role, err := s.roleService.RestoreRole(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleDeleteAt{
		Status:  "success",
		Message: "Successfully trashed role",
		Data: &pb.RoleResponseDeleteAt{
			Id:        int32(role.RoleID),
			Name:      role.RoleName,
			CreatedAt: role.CreatedAt.Time.String(),
			UpdatedAt: role.UpdatedAt.Time.String(),
			DeletedAt: role.DeletedAt.Time.String(),
		},
	}, nil
}

func (s *roleHandleGrpc) DeleteRolePermanent(ctx context.Context, req *pb.FindByIdRoleRequest) (*pb.ApiResponseRoleDelete, error) {
	id := int(req.GetRoleId())

	if id == 0 {
		return nil, role_errors.ErrGrpcRoleInvalidId
	}

	_, err := s.roleService.DeleteRolePermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleDelete{
		Status:  "success",
		Message: "Successfully deleted role permanently",
	}, nil
}

func (s *roleHandleGrpc) RestoreAllRole(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseRoleAll, error) {
	_, err := s.roleService.RestoreAllRole(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleAll{
		Status:  "success",
		Message: "Successfully restored all roles",
	}, nil
}

func (s *roleHandleGrpc) DeleteAllRolePermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseRoleAll, error) {
	_, err := s.roleService.DeleteAllRolePermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseRoleAll{
		Status:  "success",
		Message: "Successfully deleted all roles permanently",
	}, nil
}
