package service

import (
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type AuthService interface {
	Register(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Login(request *requests.AuthRequest) (*response.TokenResponse, *response.ErrorResponse)
	RefreshToken(token string) (*response.TokenResponse, *response.ErrorResponse)
	GetMe(token string) (*response.UserResponse, *response.ErrorResponse)
}

type RoleService interface {
	FindAll(page int, pageSize int, search string) ([]*response.RoleResponse, int, *response.ErrorResponse)
	FindByActiveRole(page int, pageSize int, search string) ([]*response.RoleResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashedRole(page int, pageSize int, search string) ([]*response.RoleResponseDeleteAt, int, *response.ErrorResponse)
	FindById(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	FindByUserId(id int) ([]*response.RoleResponse, *response.ErrorResponse)
	CreateRole(request *requests.CreateRoleRequest) (*response.RoleResponse, *response.ErrorResponse)
	UpdateRole(request *requests.UpdateRoleRequest) (*response.RoleResponse, *response.ErrorResponse)
	TrashedRole(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	RestoreRole(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	DeleteRolePermanent(role_id int) (bool, *response.ErrorResponse)

	RestoreAllRole() (bool, *response.ErrorResponse)
	DeleteAllRolePermanent() (bool, *response.ErrorResponse)
}

type UserService interface {
	FindAll(page int, pageSize int, search string) ([]*response.UserResponse, int, *response.ErrorResponse)
	FindByID(id int) (*response.UserResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.UserResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.UserResponseDeleteAt, int, *response.ErrorResponse)
	CreateUser(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	UpdateUser(request *requests.UpdateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	TrashedUser(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	RestoreUser(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	DeleteUserPermanent(user_id int) (bool, *response.ErrorResponse)

	RestoreAllUser() (bool, *response.ErrorResponse)
	DeleteAllUserPermanent() (bool, *response.ErrorResponse)
}

type CategoryService interface {
	FindAll(page int, pageSize int, search string) ([]*response.CategoryResponse, int, *response.ErrorResponse)
	FindById(category_id int) (*response.CategoryResponse, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.CategoryResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.CategoryResponseDeleteAt, int, *response.ErrorResponse)
	CreateCategory(req *requests.CreateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	UpdateCategory(req *requests.UpdateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	TrashedCategory(category_id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	RestoreCategory(categoryID int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	DeleteCategoryPermanent(categoryID int) (bool, *response.ErrorResponse)
	RestoreAllCategories() (bool, *response.ErrorResponse)
	DeleteAllCategoriesPermanent() (bool, *response.ErrorResponse)
}

type CashierService interface {
	FindAll(page int, pageSize int, search string) ([]*response.CashierResponse, int, *response.ErrorResponse)
	FindById(cashierID int) (*response.CashierResponse, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.CashierResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.CashierResponseDeleteAt, int, *response.ErrorResponse)
	FindByMerchant(merchantID int, search string, page, pageSize int) ([]*response.CashierResponse, int, *response.ErrorResponse)
	CreateCashier(req *requests.CreateCashierRequest) (*response.CashierResponse, *response.ErrorResponse)
	UpdateCashier(req *requests.UpdateCashierRequest) (*response.CashierResponse, *response.ErrorResponse)
	TrashedCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse)
	RestoreCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse)
	DeleteCashierPermanent(cashierID int) (bool, *response.ErrorResponse)
	RestoreAllCashier() (bool, *response.ErrorResponse)
	DeleteAllCashierPermanent() (bool, *response.ErrorResponse)
}

type MerchantService interface {
	FindAll(page, pageSize int, search string) ([]*response.MerchantResponse, int, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.MerchantResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.MerchantResponseDeleteAt, int, *response.ErrorResponse)
	FindById(merchantID int) (*response.MerchantResponse, *response.ErrorResponse)
	CreateMerchant(req *requests.CreateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse)
	UpdateMerchant(req *requests.UpdateMerchantRequest) (*response.MerchantResponse, *response.ErrorResponse)
	TrashedMerchant(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse)
	RestoreMerchant(merchantID int) (*response.MerchantResponseDeleteAt, *response.ErrorResponse)
	DeleteMerchantPermanent(merchantID int) (bool, *response.ErrorResponse)
	RestoreAllMerchant() (bool, *response.ErrorResponse)
	DeleteAllMerchantPermanent() (bool, *response.ErrorResponse)
}

type OrderItemService interface {
	FindAllOrderItems(search string, page, pageSize int) ([]*response.OrderItemResponse, int, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.OrderItemResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.OrderItemResponseDeleteAt, int, *response.ErrorResponse)
	FindOrderItemByOrder(orderID int) ([]*response.OrderItemResponse, *response.ErrorResponse)
}

type OrderService interface {
	FindAll(page int, pageSize int, search string) ([]*response.OrderResponse, int, *response.ErrorResponse)
	FindById(order_id int) (*response.OrderResponse, *response.ErrorResponse)
	FindByActive(page int, pageSize int, search string) ([]*response.OrderResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(page int, pageSize int, search string) ([]*response.OrderResponseDeleteAt, int, *response.ErrorResponse)
	CreateOrder(req *requests.CreateOrderRequest) (*response.OrderResponse, *response.ErrorResponse)
	UpdateOrder(req *requests.UpdateOrderRequest) (*response.OrderResponse, *response.ErrorResponse)
	TrashedOrder(order_id int) (*response.OrderResponseDeleteAt, *response.ErrorResponse)
	RestoreOrder(order_id int) (*response.OrderResponseDeleteAt, *response.ErrorResponse)
	DeleteOrderPermanent(order_id int) (bool, *response.ErrorResponse)
	RestoreAllOrder() (bool, *response.ErrorResponse)
	DeleteAllOrderPermanent() (bool, *response.ErrorResponse)
}

type ProductService interface {
	FindAll(page, pageSize int, search string) ([]*response.ProductResponse, int, *response.ErrorResponse)
	FindByMerchant(merchant_id int, page, pageSize int, search string) ([]*response.ProductResponse, int, *response.ErrorResponse)
	FindByCategory(category_name string, page, pageSize int, search string) ([]*response.ProductResponse, int, *response.ErrorResponse)
	FindById(productID int) (*response.ProductResponse, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.ProductResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.ProductResponseDeleteAt, int, *response.ErrorResponse)
	CreateProduct(req *requests.CreateProductRequest) (*response.ProductResponse, *response.ErrorResponse)
	UpdateProduct(req *requests.UpdateProductRequest) (*response.ProductResponse, *response.ErrorResponse)
	TrashProduct(productID int) (*response.ProductResponseDeleteAt, *response.ErrorResponse)
	RestoreProduct(productID int) (*response.ProductResponseDeleteAt, *response.ErrorResponse)
	DeleteProductPermanent(productID int) (bool, *response.ErrorResponse)
	RestoreAllProducts() (bool, *response.ErrorResponse)
	DeleteAllProductsPermanent() (bool, *response.ErrorResponse)
}

type TransactionService interface {
	FindAllTransactions(search string, page, pageSize int) ([]*response.TransactionResponse, int, *response.ErrorResponse)
	FindByMerchant(merchant_id int, search string, page, pageSize int) ([]*response.TransactionResponse, int, *response.ErrorResponse)
	FindByActive(search string, page, pageSize int) ([]*response.TransactionResponseDeleteAt, int, *response.ErrorResponse)
	FindByTrashed(search string, page, pageSize int) ([]*response.TransactionResponseDeleteAt, int, *response.ErrorResponse)
	FindById(transactionID int) (*response.TransactionResponse, *response.ErrorResponse)
	FindByOrderId(orderID int) (*response.TransactionResponse, *response.ErrorResponse)
	CreateTransaction(req *requests.CreateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	UpdateTransaction(req *requests.UpdateTransactionRequest) (*response.TransactionResponse, *response.ErrorResponse)
	TrashedTransaction(transaction_id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse)
	RestoreTransaction(transaction_id int) (*response.TransactionResponseDeleteAt, *response.ErrorResponse)
	DeleteTransactionPermanently(transactionID int) (bool, *response.ErrorResponse)
	RestoreAllTransactions() (bool, *response.ErrorResponse)
	DeleteAllTransactionPermanent() (bool, *response.ErrorResponse)
}
