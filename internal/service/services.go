package service

import (
	response_service "pointofsale/internal/mapper/response/service"
	"pointofsale/internal/repository"
	"pointofsale/pkg/auth"
	"pointofsale/pkg/hash"
	"pointofsale/pkg/logger"
)

type Service struct {
	Auth        AuthService
	User        UserService
	Role        RoleService
	Cashier     CashierService
	Category    CategoryService
	Merchant    MerchantService
	OrderItem   OrderItemService
	Order       OrderService
	Product     ProductService
	Transaction TransactionService
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
		Auth:        NewAuthService(deps.Repositories.User, deps.Repositories.RefreshToken, deps.Repositories.Role, deps.Repositories.UserRole, deps.Hash, deps.Token, deps.Logger, deps.Mapper.UserResponseMapper),
		User:        NewUserService(deps.Repositories.User, deps.Logger, deps.Mapper.UserResponseMapper, deps.Hash),
		Role:        NewRoleService(deps.Repositories.Role, deps.Logger, deps.Mapper.RoleResponseMapper),
		Cashier:     NewCashierService(deps.Repositories.Merchant, deps.Repositories.User, deps.Repositories.Cashier, deps.Logger, deps.Mapper.CashierResponseMapper),
		Category:    NewCategoryService(deps.Repositories.Category, deps.Logger, deps.Mapper.CategoryResponseMapper),
		Merchant:    NewMerchantService(deps.Repositories.Merchant, deps.Logger, deps.Mapper.MerchantResponseMapper),
		OrderItem:   NewOrderItemService(deps.Repositories.OrderItem, deps.Logger, deps.Mapper.OrderItemResponseMapper),
		Order:       NewOrderServiceMapper(deps.Repositories.Order, deps.Repositories.OrderItem, deps.Repositories.Cashier, deps.Repositories.Merchant, deps.Repositories.Product, deps.Logger, deps.Mapper.OrderResponseMapper),
		Product:     NewProductService(deps.Repositories.Category, deps.Repositories.Merchant, deps.Repositories.Product, deps.Logger, deps.Mapper.ProductResponseMapper),
		Transaction: NewTransactionService(deps.Repositories.Cashier, deps.Repositories.Merchant, deps.Repositories.Transaction, deps.Repositories.Order, deps.Repositories.OrderItem, deps.Logger, deps.Mapper.TransactionResponseMapper),
	}
}
