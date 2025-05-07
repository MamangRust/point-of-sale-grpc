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
	FindAll(req *requests.FindAllRoles) ([]*response.RoleResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllRoles) ([]*response.RoleResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllRoles) ([]*response.RoleResponseDeleteAt, *int, *response.ErrorResponse)
	FindById(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	FindByUserId(id int) ([]*response.RoleResponse, *response.ErrorResponse)
	Create(request *requests.CreateRoleRequest) (*response.RoleResponse, *response.ErrorResponse)
	Update(request *requests.UpdateRoleRequest) (*response.RoleResponse, *response.ErrorResponse)
	Trashed(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	Restore(role_id int) (*response.RoleResponse, *response.ErrorResponse)
	DeletePermanent(role_id int) (bool, *response.ErrorResponse)

	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type UserService interface {
	FindAll(req *requests.FindAllUsers) ([]*response.UserResponse, *int, *response.ErrorResponse)
	FindByID(id int) (*response.UserResponse, *response.ErrorResponse)
	FindByActive(req *requests.FindAllUsers) ([]*response.UserResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllUsers) ([]*response.UserResponseDeleteAt, *int, *response.ErrorResponse)
	Create(request *requests.CreateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Update(request *requests.UpdateUserRequest) (*response.UserResponse, *response.ErrorResponse)
	Trashed(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	Restore(user_id int) (*response.UserResponseDeleteAt, *response.ErrorResponse)
	DeletePermanent(user_id int) (bool, *response.ErrorResponse)

	RestoreAll() (bool, *response.ErrorResponse)
	DeleteAllPermanent() (bool, *response.ErrorResponse)
}

type CategoryService interface {
	FindMonthlyTotalPrice(req *requests.MonthTotalPrice) ([]*response.CategoriesMonthlyTotalPriceResponse, *response.ErrorResponse)
	FindYearlyTotalPrice(year int) ([]*response.CategoriesYearlyTotalPriceResponse, *response.ErrorResponse)
	FindMonthlyTotalPriceById(req *requests.MonthTotalPriceCategory) ([]*response.CategoriesMonthlyTotalPriceResponse, *response.ErrorResponse)
	FindYearlyTotalPriceById(req *requests.YearTotalPriceCategory) ([]*response.CategoriesYearlyTotalPriceResponse, *response.ErrorResponse)
	FindMonthlyTotalPriceByMerchant(req *requests.MonthTotalPriceMerchant) ([]*response.CategoriesMonthlyTotalPriceResponse, *response.ErrorResponse)
	FindYearlyTotalPriceByMerchant(req *requests.YearTotalPriceMerchant) ([]*response.CategoriesYearlyTotalPriceResponse, *response.ErrorResponse)

	FindMonthPrice(year int) ([]*response.CategoryMonthPriceResponse, *response.ErrorResponse)
	FindYearPrice(year int) ([]*response.CategoryYearPriceResponse, *response.ErrorResponse)
	FindMonthPriceByMerchant(req *requests.MonthPriceMerchant) ([]*response.CategoryMonthPriceResponse, *response.ErrorResponse)
	FindYearPriceByMerchant(req *requests.YearPriceMerchant) ([]*response.CategoryYearPriceResponse, *response.ErrorResponse)
	FindMonthPriceById(req *requests.MonthPriceId) ([]*response.CategoryMonthPriceResponse, *response.ErrorResponse)
	FindYearPriceById(req *requests.YearPriceId) ([]*response.CategoryYearPriceResponse, *response.ErrorResponse)

	FindAll(req *requests.FindAllCategory) ([]*response.CategoryResponse, *int, *response.ErrorResponse)
	FindById(category_id int) (*response.CategoryResponse, *response.ErrorResponse)
	FindByActive(req *requests.FindAllCategory) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllCategory) ([]*response.CategoryResponseDeleteAt, *int, *response.ErrorResponse)
	CreateCategory(req *requests.CreateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	UpdateCategory(req *requests.UpdateCategoryRequest) (*response.CategoryResponse, *response.ErrorResponse)
	TrashedCategory(category_id int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	RestoreCategory(categoryID int) (*response.CategoryResponseDeleteAt, *response.ErrorResponse)
	DeleteCategoryPermanent(categoryID int) (bool, *response.ErrorResponse)
	RestoreAllCategories() (bool, *response.ErrorResponse)
	DeleteAllCategoriesPermanent() (bool, *response.ErrorResponse)
}

type CashierService interface {
	FindMonthlyTotalSales(req *requests.MonthTotalSales) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse)
	FindYearlyTotalSales(year int) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse)
	FindMonthlyTotalSalesById(req *requests.MonthTotalSalesCashier) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse)
	FindYearlyTotalSalesById(req *requests.YearTotalSalesCashier) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse)
	FindMonthlyTotalSalesByMerchant(req *requests.MonthTotalSalesMerchant) ([]*response.CashierResponseMonthTotalSales, *response.ErrorResponse)
	FindYearlyTotalSalesByMerchant(req *requests.YearTotalSalesMerchant) ([]*response.CashierResponseYearTotalSales, *response.ErrorResponse)

	FindMonthlySales(year int) ([]*response.CashierResponseMonthSales, *response.ErrorResponse)
	FindYearlySales(year int) ([]*response.CashierResponseYearSales, *response.ErrorResponse)
	FindMonthlyCashierByMerchant(req *requests.MonthCashierMerchant) ([]*response.CashierResponseMonthSales, *response.ErrorResponse)
	FindYearlyCashierByMerchant(req *requests.YearCashierMerchant) ([]*response.CashierResponseYearSales, *response.ErrorResponse)
	FindMonthlyCashierById(req *requests.MonthCashierId) ([]*response.CashierResponseMonthSales, *response.ErrorResponse)
	FindYearlyCashierById(req *requests.YearCashierId) ([]*response.CashierResponseYearSales, *response.ErrorResponse)

	FindAll(req *requests.FindAllCashiers) ([]*response.CashierResponse, *int, *response.ErrorResponse)
	FindById(cashierID int) (*response.CashierResponse, *response.ErrorResponse)
	FindByActive(req *requests.FindAllCashiers) ([]*response.CashierResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllCashiers) ([]*response.CashierResponseDeleteAt, *int, *response.ErrorResponse)
	FindByMerchant(req *requests.FindAllCashierMerchant) ([]*response.CashierResponse, *int, *response.ErrorResponse)
	CreateCashier(req *requests.CreateCashierRequest) (*response.CashierResponse, *response.ErrorResponse)
	UpdateCashier(req *requests.UpdateCashierRequest) (*response.CashierResponse, *response.ErrorResponse)
	TrashedCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse)
	RestoreCashier(cashierID int) (*response.CashierResponseDeleteAt, *response.ErrorResponse)
	DeleteCashierPermanent(cashierID int) (bool, *response.ErrorResponse)
	RestoreAllCashier() (bool, *response.ErrorResponse)
	DeleteAllCashierPermanent() (bool, *response.ErrorResponse)
}

type MerchantService interface {
	FindAll(req *requests.FindAllMerchants) ([]*response.MerchantResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllMerchants) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllMerchants) ([]*response.MerchantResponseDeleteAt, *int, *response.ErrorResponse)
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
	FindAllOrderItems(req *requests.FindAllOrderItems) ([]*response.OrderItemResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllOrderItems) ([]*response.OrderItemResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllOrderItems) ([]*response.OrderItemResponseDeleteAt, *int, *response.ErrorResponse)
	FindOrderItemByOrder(orderID int) ([]*response.OrderItemResponse, *response.ErrorResponse)
}

type OrderService interface {
	FindMonthlyTotalRevenue(req *requests.MonthTotalRevenue) ([]*response.OrderMonthlyTotalRevenueResponse, *response.ErrorResponse)
	FindYearlyTotalRevenue(year int) ([]*response.OrderYearlyTotalRevenueResponse, *response.ErrorResponse)
	FindMonthlyTotalRevenueById(req *requests.MonthTotalRevenueOrder) ([]*response.OrderMonthlyTotalRevenueResponse, *response.ErrorResponse)
	FindYearlyTotalRevenueById(req *requests.YearTotalRevenueOrder) ([]*response.OrderYearlyTotalRevenueResponse, *response.ErrorResponse)
	FindMonthlyTotalRevenueByMerchant(req *requests.MonthTotalRevenueMerchant) ([]*response.OrderMonthlyTotalRevenueResponse, *response.ErrorResponse)
	FindYearlyTotalRevenueByMerchant(req *requests.YearTotalRevenueMerchant) ([]*response.OrderYearlyTotalRevenueResponse, *response.ErrorResponse)

	FindMonthlyOrder(year int) ([]*response.OrderMonthlyResponse, *response.ErrorResponse)
	FindYearlyOrder(year int) ([]*response.OrderYearlyResponse, *response.ErrorResponse)
	FindMonthlyOrderByMerchant(req *requests.MonthOrderMerchant) ([]*response.OrderMonthlyResponse, *response.ErrorResponse)
	FindYearlyOrderByMerchant(req *requests.YearOrderMerchant) ([]*response.OrderYearlyResponse, *response.ErrorResponse)

	FindAll(req *requests.FindAllOrders) ([]*response.OrderResponse, *int, *response.ErrorResponse)
	FindById(order_id int) (*response.OrderResponse, *response.ErrorResponse)
	FindByActive(req *requests.FindAllOrders) ([]*response.OrderResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllOrders) ([]*response.OrderResponseDeleteAt, *int, *response.ErrorResponse)
	FindByMerchant(req *requests.FindAllOrderMerchant) ([]*response.OrderResponse, *int, *response.ErrorResponse)
	CreateOrder(req *requests.CreateOrderRequest) (*response.OrderResponse, *response.ErrorResponse)
	UpdateOrder(req *requests.UpdateOrderRequest) (*response.OrderResponse, *response.ErrorResponse)
	TrashedOrder(order_id int) (*response.OrderResponseDeleteAt, *response.ErrorResponse)
	RestoreOrder(order_id int) (*response.OrderResponseDeleteAt, *response.ErrorResponse)
	DeleteOrderPermanent(order_id int) (bool, *response.ErrorResponse)
	RestoreAllOrder() (bool, *response.ErrorResponse)
	DeleteAllOrderPermanent() (bool, *response.ErrorResponse)
}

type ProductService interface {
	FindAll(req *requests.FindAllProducts) ([]*response.ProductResponse, *int, *response.ErrorResponse)
	FindByMerchant(req *requests.ProductByMerchantRequest) ([]*response.ProductResponse, *int, *response.ErrorResponse)
	FindByCategory(req *requests.ProductByCategoryRequest) ([]*response.ProductResponse, *int, *response.ErrorResponse)
	FindById(productID int) (*response.ProductResponse, *response.ErrorResponse)
	FindByActive(req *requests.FindAllProducts) ([]*response.ProductResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllProducts) ([]*response.ProductResponseDeleteAt, *int, *response.ErrorResponse)
	CreateProduct(req *requests.CreateProductRequest) (*response.ProductResponse, *response.ErrorResponse)
	UpdateProduct(req *requests.UpdateProductRequest) (*response.ProductResponse, *response.ErrorResponse)
	TrashProduct(productID int) (*response.ProductResponseDeleteAt, *response.ErrorResponse)
	RestoreProduct(productID int) (*response.ProductResponseDeleteAt, *response.ErrorResponse)
	DeleteProductPermanent(productID int) (bool, *response.ErrorResponse)
	RestoreAllProducts() (bool, *response.ErrorResponse)
	DeleteAllProductsPermanent() (bool, *response.ErrorResponse)
}

type TransactionService interface {
	FindMonthlyAmountSuccess(req *requests.MonthAmountTransaction) ([]*response.TransactionMonthlyAmountSuccessResponse, *response.ErrorResponse)
	FindYearlyAmountSuccess(year int) ([]*response.TransactionYearlyAmountSuccessResponse, *response.ErrorResponse)
	FindMonthlyAmountFailed(req *requests.MonthAmountTransaction) ([]*response.TransactionMonthlyAmountFailedResponse, *response.ErrorResponse)
	FindYearlyAmountFailed(year int) ([]*response.TransactionYearlyAmountFailedResponse, *response.ErrorResponse)

	FindMonthlyAmountSuccessByMerchant(req *requests.MonthAmountTransactionMerchant) ([]*response.TransactionMonthlyAmountSuccessResponse, *response.ErrorResponse)
	FindYearlyAmountSuccessByMerchant(req *requests.YearAmountTransactionMerchant) ([]*response.TransactionYearlyAmountSuccessResponse, *response.ErrorResponse)
	FindMonthlyAmountFailedByMerchant(req *requests.MonthAmountTransactionMerchant) ([]*response.TransactionMonthlyAmountFailedResponse, *response.ErrorResponse)
	FindYearlyAmountFailedByMerchant(req *requests.YearAmountTransactionMerchant) ([]*response.TransactionYearlyAmountFailedResponse, *response.ErrorResponse)

	FindMonthlyMethod(year int) ([]*response.TransactionMonthlyMethodResponse, *response.ErrorResponse)
	FindYearlyMethod(year int) ([]*response.TransactionYearlyMethodResponse, *response.ErrorResponse)
	FindMonthlyMethodByMerchant(req *requests.MonthlyYearTransactionMethodMerchant) ([]*response.TransactionMonthlyMethodResponse, *response.ErrorResponse)
	FindYearlyMethodByMerchant(req *requests.MonthlyYearTransactionMethodMerchant) ([]*response.TransactionYearlyMethodResponse, *response.ErrorResponse)

	FindAllTransactions(req *requests.FindAllTransaction) ([]*response.TransactionResponse, *int, *response.ErrorResponse)
	FindByMerchant(req *requests.FindAllTransactionByMerchant) ([]*response.TransactionResponse, *int, *response.ErrorResponse)
	FindByActive(req *requests.FindAllTransaction) ([]*response.TransactionResponseDeleteAt, *int, *response.ErrorResponse)
	FindByTrashed(req *requests.FindAllTransaction) ([]*response.TransactionResponseDeleteAt, *int, *response.ErrorResponse)
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
