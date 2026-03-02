package api

import (
	"pointofsale/internal/cache"
	auth_cache "pointofsale/internal/cache/api/auth"
	cashier_cache "pointofsale/internal/cache/api/cashier"
	category_cache "pointofsale/internal/cache/api/category"
	merchant_cache "pointofsale/internal/cache/api/merchant"
	order_cache "pointofsale/internal/cache/api/order"
	orderitem_cache "pointofsale/internal/cache/api/order_item"
	product_cache "pointofsale/internal/cache/api/product"
	role_cache "pointofsale/internal/cache/api/role"
	transaction_cache "pointofsale/internal/cache/api/transaction"
	user_cache "pointofsale/internal/cache/api/user"
	response_api "pointofsale/internal/mapper"
	"pointofsale/internal/pb"
	"pointofsale/pkg/auth"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/pkg/upload_image"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Deps struct {
	Conn        *grpc.ClientConn
	Token       auth.TokenManager
	E           *echo.Echo
	Logger      logger.LoggerInterface
	Mapping     *response_api.ResponseApiMapper
	ImageUpload upload_image.ImageUploads
	Cache       *cache.CacheStore
}

func NewHandler(deps Deps) {
	observability, _ := observability.NewObservability("client", deps.Logger)

	apiHandler := errors.NewApiHandler(observability, deps.Logger)

	auth_cache := auth_cache.NewMencache(deps.Cache)
	user_cache := user_cache.NewUserMencache(deps.Cache)
	role_cache := role_cache.NewRoleMencache(deps.Cache)
	category_cache := category_cache.NewCategoryMencache(deps.Cache)
	merchant_cache := merchant_cache.NewMerchantMencache(deps.Cache)
	cashier_cache := cashier_cache.NewCashierMencache(deps.Cache)
	order_item_cache := orderitem_cache.NewOrderItemCache(deps.Cache)
	order_cache := order_cache.NewOrderMencache(deps.Cache)
	product_cache := product_cache.NewProductMencache(deps.Cache)
	transaction_cache := transaction_cache.NewTransactionMencache(deps.Cache)

	clientAuth := pb.NewAuthServiceClient(deps.Conn)
	clientRole := pb.NewRoleServiceClient(deps.Conn)
	clientUser := pb.NewUserServiceClient(deps.Conn)
	clientCategory := pb.NewCategoryServiceClient(deps.Conn)
	clientCashier := pb.NewCashierServiceClient(deps.Conn)
	clientMerchant := pb.NewMerchantServiceClient(deps.Conn)
	clientOrderItem := pb.NewOrderItemServiceClient(deps.Conn)
	clientOrder := pb.NewOrderServiceClient(deps.Conn)
	clientProduct := pb.NewProductServiceClient(deps.Conn)
	clientTransaction := pb.NewTransactionServiceClient(deps.Conn)

	NewHandlerAuth(deps.E, clientAuth, deps.Logger, deps.Mapping.AuthResponseMapper, apiHandler, auth_cache)
	NewHandlerRole(deps.E, clientRole, deps.Logger, deps.Mapping.RoleResponseMapper, apiHandler, role_cache)
	NewHandlerUser(deps.E, clientUser, deps.Logger, deps.Mapping.UserResponseMapper, apiHandler, user_cache)
	NewHandlerCategory(deps.E, clientCategory, deps.Logger, deps.Mapping.CategoryResponseMapper, apiHandler, category_cache)
	NewHandlerCashier(deps.E, clientCashier, deps.Logger, deps.Mapping.CashierResponseMapper, apiHandler, cashier_cache)
	NewHandlerMerchant(deps.E, clientMerchant, deps.Logger, deps.Mapping.MerchantResponseMapper, apiHandler, merchant_cache)
	NewHandlerOrderItem(deps.E, clientOrderItem, deps.Logger, deps.Mapping.OrderItemResponseMapper, apiHandler, order_item_cache)
	NewHandlerOrder(deps.E, clientOrder, deps.Logger, deps.Mapping.OrderResponseMapper, apiHandler, order_cache)
	NewHandlerProduct(deps.E, clientProduct, deps.Logger, deps.Mapping.ProductResponseMapper, deps.ImageUpload, apiHandler, product_cache)
	NewHandlerTransaction(deps.E, clientTransaction, deps.Logger, deps.Mapping.TransactionResponseMapper, apiHandler, transaction_cache)
}
