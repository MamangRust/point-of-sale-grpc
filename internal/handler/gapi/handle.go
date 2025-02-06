package gapi

import (
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/service"
)

type Deps struct {
	Service service.Service
	Mapper  protomapper.ProtoMapper
}

type Handler struct {
	Auth AuthHandleGrpc
	Role RoleHandleGrpc
	User UserHandleGrpc
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		Auth: NewAuthHandleGrpc(deps.Service.Auth, deps.Mapper.AuthProtoMapper),
		Role: NewRoleHandleGrpc(deps.Service.Role, deps.Mapper.RoleProtoMapper),
		User: NewUserHandleGrpc(deps.Service.User, deps.Mapper.UserProtoMapper),
	}
}
