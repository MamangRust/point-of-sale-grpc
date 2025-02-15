package repository

import (
	"context"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
)

type Repositories struct {
	User         UserRepository
	Role         RoleRepository
	UserRole     UserRoleRepository
	Category     CategoryRepository
	RefreshToken RefreshTokenRepository
	Cashier      CashierRepository
	Product      ProductRepository
	Merchant     MerchantRepository
	OrderItem    OrderItemRepository
	Order        OrderRepository
	Transaction  TransactionRepository
}

type Deps struct {
	DB           *db.Queries
	Ctx          context.Context
	MapperRecord *recordmapper.RecordMapper
}

func NewRepositories(deps Deps) *Repositories {
	return &Repositories{
		User:         NewUserRepository(deps.DB, deps.Ctx, deps.MapperRecord.UserRecordMapper),
		Role:         NewRoleRepository(deps.DB, deps.Ctx, deps.MapperRecord.RoleRecordMapper),
		UserRole:     NewUserRoleRepository(deps.DB, deps.Ctx, deps.MapperRecord.UserRoleRecordMapper),
		Category:     NewCategoryRepository(deps.DB, deps.Ctx, deps.MapperRecord.CategoryRecordMapper),
		RefreshToken: NewRefreshTokenRepository(deps.DB, deps.Ctx, deps.MapperRecord.RefreshTokenRecordMapper),
		Cashier:      NewCashierRepository(deps.DB, deps.Ctx, deps.MapperRecord.CashierRecordMapper),
		Product:      NewProductRepository(deps.DB, deps.Ctx, deps.MapperRecord.ProductRecordMapper),
		Merchant:     NewMerchantRepository(deps.DB, deps.Ctx, deps.MapperRecord.MerchantRecordMapper),
		OrderItem:    NewOrderItemRepository(deps.DB, deps.Ctx, deps.MapperRecord.OrderItemRecordMapper),
		Order:        NewOrderRepository(deps.DB, deps.Ctx, deps.MapperRecord.OrderRecordMapper),
		Transaction:  NewTransactionRepository(deps.DB, deps.Ctx, deps.MapperRecord.TransactionRecordMapper),
	}
}
