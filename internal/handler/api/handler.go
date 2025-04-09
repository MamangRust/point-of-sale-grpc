package api

import (
	response_api "pointofsale/internal/mapper/response/api"
	"pointofsale/internal/pb"
	"pointofsale/pkg/auth"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/upload_image"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Deps struct {
	Conn        *grpc.ClientConn
	Token       auth.TokenManager
	E           *echo.Echo
	Logger      logger.LoggerInterface
	Mapping     response_api.ResponseApiMapper
	ImageUpload upload_image.ImageUploads
}

func NewHandler(deps Deps) {
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

	NewHandlerAuth(deps.E, clientAuth, deps.Logger, deps.Mapping.AuthResponseMapper)
	NewHandlerRole(deps.E, clientRole, deps.Logger, deps.Mapping.RoleResponseMapper)
	NewHandlerUser(deps.E, clientUser, deps.Logger, deps.Mapping.UserResponseMapper)
	NewHandlerCategory(deps.E, clientCategory, deps.Logger, deps.Mapping.CategoryResponseMapper)
	NewHandlerCashier(deps.E, clientCashier, deps.Logger, deps.Mapping.CashierResponseMapper)
	NewHandlerMerchant(deps.E, clientMerchant, deps.Logger, deps.Mapping.MerchantResponseMapper)
	NewHandlerOrderItem(deps.E, clientOrderItem, deps.Logger, deps.Mapping.OrderItemResponseMapper)
	NewHandlerOrder(deps.E, clientOrder, deps.Logger, deps.Mapping.OrderResponseMapper)
	NewHandlerProduct(deps.E, clientProduct, deps.Logger, deps.Mapping.ProductResponseMapper, deps.ImageUpload)
	NewHandlerTransaction(deps.E, clientTransaction, deps.Logger, deps.Mapping.TransactionResponseMapper)
}
