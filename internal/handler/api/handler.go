package api

import (
	response_api "pointofsale/internal/mapper/response/api"
	"pointofsale/internal/pb"
	"pointofsale/pkg/auth"
	"pointofsale/pkg/logger"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

type Deps struct {
	Conn    *grpc.ClientConn
	Token   auth.TokenManager
	E       *echo.Echo
	Logger  logger.LoggerInterface
	Mapping response_api.ResponseApiMapper
}

func NewHandler(deps Deps) {

	clientAuth := pb.NewAuthServiceClient(deps.Conn)
	clientRole := pb.NewRoleServiceClient(deps.Conn)
	clientUser := pb.NewUserServiceClient(deps.Conn)

	NewHandlerAuth(clientAuth, deps.E, deps.Logger, deps.Mapping.AuthResponseMapper)
	NewHandlerRole(clientRole, deps.E, deps.Logger, deps.Mapping.RoleResponseMapper)
	NewHandlerUser(clientUser, deps.E, deps.Logger, deps.Mapping.UserResponseMapper)
}
