package repository

import (
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

func NewRepositories(db *db.Queries) *Repositories {
	return &Repositories{
		User:         NewUserRepository(db),
		Role:         NewRoleRepository(db),
		UserRole:     NewUserRoleRepository(db),
		Category:     NewCategoryRepository(db),
		RefreshToken: NewRefreshTokenRepository(db),
		Cashier:      NewCashierRepository(db),
		Product:      NewProductRepository(db),
		Merchant:     NewMerchantRepository(db),
		OrderItem:    NewOrderItemRepository(db),
		Order:        NewOrderRepository(db),
		Transaction:  NewTransactionRepository(db),
	}
}
