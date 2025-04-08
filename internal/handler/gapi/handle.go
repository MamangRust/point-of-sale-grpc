package gapi

import (
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/service"
)

type Deps struct {
	Service *service.Service
	Mapper  *protomapper.ProtoMapper
}

type Handler struct {
	Auth        AuthHandleGrpc
	Role        RoleHandleGrpc
	User        UserHandleGrpc
	Category    CategoryHandleGrpc
	Cashier     CashierHandleGrpc
	Merchant    MerchantHandleGrpc
	OrderItem   OrderItemHandleGrpc
	Order       OrderHandleGrpc
	Product     ProductHandleGrpc
	Transaction TransactionHandleGrpc
}

func NewHandler(deps Deps) *Handler {
	return &Handler{
		Auth:        NewAuthHandleGrpc(deps.Service.Auth, deps.Mapper.AuthProtoMapper),
		Role:        NewRoleHandleGrpc(deps.Service.Role, deps.Mapper.RoleProtoMapper),
		User:        NewUserHandleGrpc(deps.Service.User, deps.Mapper.UserProtoMapper),
		Category:    NewCategoryHandleGrpc(deps.Service.Category, deps.Mapper.CategoryProtoMapper),
		Cashier:     NewCashierHandleGrpc(deps.Service.Cashier, deps.Mapper.CashierProtoMapper),
		Merchant:    NewMerchantHandleGrpc(deps.Service.Merchant, deps.Mapper.MerchantProtoMapper),
		OrderItem:   NewOrderItemHandleGrpc(deps.Service.OrderItem, deps.Mapper.OrderItemProtoMapper),
		Order:       NewOrderHandleGrpc(deps.Service.Order, deps.Mapper.OrderProtoMapper),
		Product:     NewProductHandleGrpc(deps.Service.Product, deps.Mapper.ProductProtoMapper),
		Transaction: NewTransactionHandleGrpc(deps.Service.Transaction, deps.Mapper.TransactionProtoMapper),
	}
}
