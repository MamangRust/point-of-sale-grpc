package service

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type roleService struct {
	roleRepository repository.RoleRepository
	logger         logger.LoggerInterface
	mapping        response_service.RoleResponseMapper
}

func NewRoleService(roleRepository repository.RoleRepository, logger logger.LoggerInterface, mapping response_service.RoleResponseMapper) *roleService {
	return &roleService{
		roleRepository: roleRepository,
		logger:         logger,
		mapping:        mapping,
	}
}

func (s *roleService) FindAll(page int, pageSize int, search string) ([]*response.RoleResponse, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching role",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.roleRepository.FindAllRoles(page, pageSize, search)

	if err != nil {
		s.logger.Error("Failed to retrieve role list from database",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.String("search", search))

		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve role list",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched role",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToRolesResponse(roles)

	return so, &totalRecords, nil
}

func (s *roleService) FindById(id int) (*response.RoleResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching role by ID", zap.Int("id", id))

	role, err := s.roleRepository.FindById(id)
	if err != nil {
		s.logger.Error("Failed to retrieve role details",
			zap.Int("role_id", id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Role with ID %d not found", id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve role details",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched role", zap.Int("id", id))

	so := s.mapping.ToRoleResponse(role)

	return so, nil
}

func (s *roleService) FindByUserId(id int) ([]*response.RoleResponse, *response.ErrorResponse) {
	s.logger.Debug("Fetching role by user ID", zap.Int("id", id))

	role, err := s.roleRepository.FindByUserId(id)
	if err != nil {
		s.logger.Error("Failed to retrieve role by user ID",
			zap.Int("user_id", id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Role for user ID %d not found", id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve user role",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched role by user ID", zap.Int("id", id))

	so := s.mapping.ToRolesResponse(role)

	return so, nil
}

func (s *roleService) FindByActiveRole(page int, pageSize int, search string) ([]*response.RoleResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching active role",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.roleRepository.FindByActiveRole(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to retrieve active roles from database",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve active roles",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched active role",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToRolesResponseDeleteAt(roles)

	return so, &totalRecords, nil
}

func (s *roleService) FindByTrashedRole(page int, pageSize int, search string) ([]*response.RoleResponseDeleteAt, *int, *response.ErrorResponse) {
	s.logger.Debug("Fetching trashed role",
		zap.Int("page", page),
		zap.Int("pageSize", pageSize),
		zap.String("search", search))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	roles, totalRecords, err := s.roleRepository.FindByTrashedRole(page, pageSize, search)
	if err != nil {
		s.logger.Error("Failed to retrieve trashed roles from database",
			zap.Error(err),
			zap.Int("page", page),
			zap.Int("page_size", pageSize),
			zap.String("search", search))
		return nil, nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to retrieve trashed roles",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully fetched trashed role",
		zap.Int("totalRecords", totalRecords),
		zap.Int("page", page),
		zap.Int("pageSize", pageSize))

	so := s.mapping.ToRolesResponseDeleteAt(roles)

	return so, &totalRecords, nil
}

func (s *roleService) CreateRole(request *requests.CreateRoleRequest) (*response.RoleResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting CreateRole process",
		zap.String("roleName", request.Name),
	)

	role, err := s.roleRepository.CreateRole(request)
	if err != nil {
		s.logger.Error("Failed to create new role record",
			zap.String("role_name", request.Name),
			zap.Error(err))
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to create new role",
			Code:    http.StatusInternalServerError,
		}
	}

	so := s.mapping.ToRoleResponse(role)

	s.logger.Debug("CreateRole process completed",
		zap.String("roleName", request.Name),
		zap.Int("roleID", role.ID),
	)

	return so, nil
}

func (s *roleService) UpdateRole(request *requests.UpdateRoleRequest) (*response.RoleResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting UpdateRole process",
		zap.Int("roleID", *request.ID),
		zap.String("newRoleName", request.Name),
	)

	role, err := s.roleRepository.UpdateRole(request)
	if err != nil {
		s.logger.Error("Failed to update role record",
			zap.Int("role_id", *request.ID),
			zap.String("new_name", request.Name),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Role with ID %d not found", request.ID),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to update role",
			Code:    http.StatusInternalServerError,
		}
	}

	so := s.mapping.ToRoleResponse(role)

	s.logger.Debug("UpdateRole process completed",
		zap.Int("roleID", *request.ID),
		zap.String("newRoleName", request.Name),
	)

	return so, nil
}

func (s *roleService) TrashedRole(id int) (*response.RoleResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting TrashedRole process",
		zap.Int("roleID", id),
	)

	role, err := s.roleRepository.TrashedRole(id)
	if err != nil {
		s.logger.Error("Failed to move role to trash",
			zap.Int("role_id", id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Role with ID %d not found", id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to move role to trash",
			Code:    http.StatusInternalServerError,
		}
	}

	so := s.mapping.ToRoleResponse(role)

	s.logger.Debug("TrashedRole process completed",
		zap.Int("roleID", id),
	)

	return so, nil
}

func (s *roleService) RestoreRole(id int) (*response.RoleResponse, *response.ErrorResponse) {
	s.logger.Debug("Starting RestoreRole process",
		zap.Int("roleID", id),
	)

	role, err := s.roleRepository.RestoreRole(id)

	if err != nil {
		s.logger.Error("Failed to restore role from trash",
			zap.Int("role_id", id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Role with ID %d not found in trash", id),
				Code:    http.StatusNotFound,
			}
		}
		return nil, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore role from trash",
			Code:    http.StatusInternalServerError,
		}
	}

	so := s.mapping.ToRoleResponse(role)

	s.logger.Debug("RestoreRole process completed",
		zap.Int("roleID", id),
	)

	return so, nil
}

func (s *roleService) DeleteRolePermanent(id int) (bool, *response.ErrorResponse) {
	s.logger.Debug("Starting DeleteRolePermanent process",
		zap.Int("roleID", id),
	)

	_, err := s.roleRepository.DeleteRolePermanent(id)
	if err != nil {
		s.logger.Error("Failed to permanently delete role",
			zap.Int("role_id", id),
			zap.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return false, &response.ErrorResponse{
				Status:  "not_found",
				Message: fmt.Sprintf("Role with ID %d not found", id),
				Code:    http.StatusNotFound,
			}
		}
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete role",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("DeleteRolePermanent process completed",
		zap.Int("roleID", id),
	)

	return true, nil
}

func (s *roleService) RestoreAllRole() (bool, *response.ErrorResponse) {
	s.logger.Debug("Restoring all roles")

	_, err := s.roleRepository.RestoreAllRole()
	if err != nil {
		s.logger.Error("Failed to restore all trashed roles",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to restore all roles",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully restored all roles")
	return true, nil
}

func (s *roleService) DeleteAllRolePermanent() (bool, *response.ErrorResponse) {
	s.logger.Debug("Permanently deleting all roles")

	_, err := s.roleRepository.DeleteAllRolePermanent()

	if err != nil {
		s.logger.Error("Failed to permanently delete all trashed roles",
			zap.Error(err))
		return false, &response.ErrorResponse{
			Status:  "error",
			Message: "Failed to permanently delete all roles",
			Code:    http.StatusInternalServerError,
		}
	}

	s.logger.Debug("Successfully deleted all roles permanently")
	return true, nil
}
