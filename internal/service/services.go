package service

import (
	"pointofsale/internal/cache"
	auth_cache "pointofsale/internal/cache/auth"
	cashier_cache "pointofsale/internal/cache/cashier"
	category_cache "pointofsale/internal/cache/category"
	merchant_cache "pointofsale/internal/cache/merchant"
	order_cache "pointofsale/internal/cache/order"
	orderitem_cache "pointofsale/internal/cache/order_item"
	product_cache "pointofsale/internal/cache/product"
	role_cache "pointofsale/internal/cache/role"
	transaction_cache "pointofsale/internal/cache/transaction"
	user_cache "pointofsale/internal/cache/user"
	"pointofsale/internal/repository"
	"pointofsale/pkg/auth"
	"pointofsale/pkg/hash"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
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
	Cache        *cache.CacheStore
}

func NewService(deps Deps) *Service {
	observability, _ := observability.NewObservability("grpc-server", deps.Logger)

	auth_cache := auth_cache.NewMencache(deps.Cache)
	role_cache := role_cache.NewRoleMencache(deps.Cache)
	user_cache := user_cache.NewUserMencache(deps.Cache)
	category_cache := category_cache.NewCategoryMencache(deps.Cache)
	cashier_cache := cashier_cache.NewCashierMencache(deps.Cache)
	merchant_cache := merchant_cache.NewMerchantMencache(deps.Cache)
	order_cache := order_cache.NewOrderMencache(deps.Cache)
	order_item_cache := orderitem_cache.NewOrderItemCache(deps.Cache)
	product_cache := product_cache.NewProductMencache(deps.Cache)
	transaction_cache := transaction_cache.NewTransactionMencache(deps.Cache)

	return &Service{
		Auth: NewAuthService(AuthServiceDeps{
			UserRepo:         deps.Repositories.User,
			RefreshTokenRepo: deps.Repositories.RefreshToken,
			RoleRepo:         deps.Repositories.Role,
			UserRoleRepo:     deps.Repositories.UserRole,
			Hash:             deps.Hash,
			TokenManager:     deps.Token,
			Logger:           deps.Logger,
			Observability:    observability,
			Cache:            auth_cache,
		}),

		User: NewUserService(UserServiceDeps{
			UserRepo:      deps.Repositories.User,
			Hash:          deps.Hash,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         user_cache,
		}),

		Role: NewRoleService(RoleServiceDeps{
			RoleRepo:      deps.Repositories.Role,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         role_cache,
		}),

		Cashier: NewCashierService(CashierServiceDeps{
			MerchantRepo:  deps.Repositories.Merchant,
			UserRepo:      deps.Repositories.User,
			CashierRepo:   deps.Repositories.Cashier,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         cashier_cache,
		}),
		Category: NewCategoryService(CategoryServiceDeps{
			CategoryRepo:  deps.Repositories.Category,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         category_cache,
		}),

		Merchant: NewMerchantService(MerchantServiceDeps{
			MerchantRepo:  deps.Repositories.Merchant,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         merchant_cache,
		}),

		OrderItem: NewOrderItemService(OrderItemServiceDeps{
			OrderItemRepo: deps.Repositories.OrderItem,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         order_item_cache,
		}),

		Order: NewOrderService(OrderServiceDeps{
			OrderRepo:     deps.Repositories.Order,
			OrderItemRepo: deps.Repositories.OrderItem,
			ProductRepo:   deps.Repositories.Product,
			CashierRepo:   deps.Repositories.Cashier,
			MerchantRepo:  deps.Repositories.Merchant,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         order_cache,
		}),

		Product: NewProductService(ProductServiceDeps{
			CategoryRepo:  deps.Repositories.Category,
			MerchantRepo:  deps.Repositories.Merchant,
			ProductRepo:   deps.Repositories.Product,
			Logger:        deps.Logger,
			Observability: observability,
			Cache:         product_cache,
		}),

		Transaction: NewTransactionService(TransactionServiceDeps{
			CashierRepo:     deps.Repositories.Cashier,
			MerchantRepo:    deps.Repositories.Merchant,
			TransactionRepo: deps.Repositories.Transaction,
			OrderRepo:       deps.Repositories.Order,
			OrderItemRepo:   deps.Repositories.OrderItem,
			Logger:          deps.Logger,
			Observability:   observability,
			Cache:           transaction_cache,
		}),
	}
}
