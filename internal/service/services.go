package service

import (
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/auth"
	"pointofsale/pkg/hash"
	"pointofsale/pkg/logger"
)

type Service struct {
	Auth AuthService
	User UserService
	Role RoleService
}

type Deps struct {
	Repositories *repository.Repositories
	Token        auth.TokenManager
	Hash         hash.HashPassword
	Logger       logger.LoggerInterface
	Mapper       response_service.ResponseServiceMapper
}

func NewService(deps Deps) *Service {
	return &Service{
		Auth: NewAuthService(deps.Repositories.User, deps.Repositories.RefreshToken, deps.Repositories.Role, deps.Repositories.UserRole, deps.Hash, deps.Token, deps.Logger, deps.Mapper.UserResponseMapper),
		User: NewUserService(deps.Repositories.User, deps.Logger, deps.Mapper.UserResponseMapper, deps.Hash),
		Role: NewRoleService(deps.Repositories.Role, deps.Logger, deps.Mapper.RoleResponseMapper),
	}
}
