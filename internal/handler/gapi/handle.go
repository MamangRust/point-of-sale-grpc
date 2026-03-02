package gapi

import (
	"pointofsale/internal/service"
)

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

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Auth:        NewAuthHandleGrpc(service.Auth),
		Role:        NewRoleHandleGrpc(service.Role),
		User:        NewUserHandleGrpc(service.User),
		Category:    NewCategoryHandleGrpc(service.Category),
		Cashier:     NewCashierHandleGrpc(service.Cashier),
		Merchant:    NewMerchantHandleGrpc(service.Merchant),
		OrderItem:   NewOrderItemHandleGrpc(service.OrderItem),
		Order:       NewOrderHandleGrpc(service.Order),
		Product:     NewProductHandleGrpc(service.Product),
		Transaction: NewTransactionHandleGrpc(service.Transaction),
	}
}
