package response_service

import (
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/response"
)

type UserResponseMapper interface {
	ToUserResponse(user *record.UserRecord) *response.UserResponse
	ToUsersResponse(users []*record.UserRecord) []*response.UserResponse

	ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt
	ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt
}

type RoleResponseMapper interface {
	ToRoleResponse(role *record.RoleRecord) *response.RoleResponse
	ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse

	ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt
	ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt
}

type RefreshTokenResponseMapper interface {
	ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse
	ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse
}

type CategoryResponseMapper interface {
	ToCategoryResponse(category *record.CategoriesRecord) *response.CategoryResponse
	ToCategorysResponse(categories []*record.CategoriesRecord) []*response.CategoryResponse
	ToCategoryResponseDeleteAt(category *record.CategoriesRecord) *response.CategoryResponseDeleteAt
	ToCategorysResponseDeleteAt(categories []*record.CategoriesRecord) []*response.CategoryResponseDeleteAt
}

type CashierResponseMapper interface {
	ToCashierResponse(cashier *record.CashierRecord) *response.CashierResponse
	ToCashiersResponse(cashiers []*record.CashierRecord) []*response.CashierResponse
	ToCashierResponseDeleteAt(cashier *record.CashierRecord) *response.CashierResponseDeleteAt
	ToCashiersResponseDeleteAt(cashiers []*record.CashierRecord) []*response.CashierResponseDeleteAt
}

type MerchantResponseMapper interface {
	ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse
	ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse
	ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt
	ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt
}

type OrderResponseMapper interface {
	ToOrderResponse(order *record.OrderRecord) *response.OrderResponse
	ToOrdersResponse(orders []*record.OrderRecord) []*response.OrderResponse
	ToOrderResponseDeleteAt(order *record.OrderRecord) *response.OrderResponseDeleteAt
	ToOrdersResponseDeleteAt(orders []*record.OrderRecord) []*response.OrderResponseDeleteAt
}

type OrderItemResponseMapper interface {
	ToOrderItemResponse(order *record.OrderItemRecord) *response.OrderItemResponse
	ToOrderItemsResponse(orders []*record.OrderItemRecord) []*response.OrderItemResponse
	ToOrderItemResponseDeleteAt(order *record.OrderItemRecord) *response.OrderItemResponseDeleteAt
	ToOrderItemsResponseDeleteAt(orders []*record.OrderItemRecord) []*response.OrderItemResponseDeleteAt
}

type ProductResponseMapper interface {
	ToProductResponse(product *record.ProductRecord) *response.ProductResponse
	ToProductsResponse(products []*record.ProductRecord) []*response.ProductResponse
	ToProductResponseDeleteAt(product *record.ProductRecord) *response.ProductResponseDeleteAt
	ToProductsResponseDeleteAt(products []*record.ProductRecord) []*response.ProductResponseDeleteAt
}

type TransactionResponseMapper interface {
	ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse
	ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt
	ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
}
