// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"time"
)

type Querier interface {
	AssignRoleToUser(ctx context.Context, arg AssignRoleToUserParams) (*UserRole, error)
	CalculateTotalPrice(ctx context.Context, orderID int32) (int32, error)
	CreateCashier(ctx context.Context, arg CreateCashierParams) (*Cashier, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (*Category, error)
	// Create Merchant
	CreateMerchant(ctx context.Context, arg CreateMerchantParams) (*Merchant, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (*Order, error)
	CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (*OrderItem, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (*Product, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (*RefreshToken, error)
	CreateRole(ctx context.Context, roleName string) (*Role, error)
	CreateTransactions(ctx context.Context, arg CreateTransactionsParams) (*Transaction, error)
	// Create User
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)
	// Delete All Trashed Cashier Permanently
	DeleteAllPermanentCashiers(ctx context.Context) error
	// Delete All Trashed Category Permanently
	DeleteAllPermanentCategories(ctx context.Context) error
	// Delete All Trashed Merchant Permanently
	DeleteAllPermanentMerchants(ctx context.Context) error
	// Delete All Trashed Order Permanently
	DeleteAllPermanentOrders(ctx context.Context) error
	// Delete All Trashed Order Item Permanently
	DeleteAllPermanentOrdersItem(ctx context.Context) error
	// Delete All Trashed Product Permanently
	DeleteAllPermanentProducts(ctx context.Context) error
	// Delete All Trashed Roles Permanently
	DeleteAllPermanentRoles(ctx context.Context) error
	// Delete All Trashed Transaction Permanently
	DeleteAllPermanentTransactions(ctx context.Context) error
	// Delete All Trashed Users Permanently
	DeleteAllPermanentUsers(ctx context.Context) error
	// Delete Cashier Permanently
	DeleteCashierPermanently(ctx context.Context, cashierID int32) error
	// Delete Category Permanently
	DeleteCategoryPermanently(ctx context.Context, categoryID int32) error
	// Delete Merchant Permanently
	DeleteMerchantPermanently(ctx context.Context, merchantID int32) error
	// Delete Order Item Permanently
	DeleteOrderItemPermanently(ctx context.Context, orderItemID int32) error
	// Delete Order Permanently
	DeleteOrderPermanently(ctx context.Context, orderID int32) error
	DeletePermanentRole(ctx context.Context, roleID int32) error
	// Delete Product Permanently
	DeleteProductPermanently(ctx context.Context, productID int32) error
	DeleteRefreshToken(ctx context.Context, token string) error
	DeleteRefreshTokenByUserId(ctx context.Context, userID int32) error
	// Delete Transaction Permanently
	DeleteTransactionPermanently(ctx context.Context, transactionID int32) error
	// Delete User Permanently
	DeleteUserPermanently(ctx context.Context, userID int32) error
	FindRefreshTokenByToken(ctx context.Context, token string) (*RefreshToken, error)
	FindRefreshTokenByUserId(ctx context.Context, userID int32) (*RefreshToken, error)
	// Get All Active Roles
	GetActiveRoles(ctx context.Context, arg GetActiveRolesParams) ([]*GetActiveRolesRow, error)
	GetCashierByID(ctx context.Context, cashierID int32) (*Cashier, error)
	GetCashiers(ctx context.Context, arg GetCashiersParams) ([]*GetCashiersRow, error)
	GetCashiersActive(ctx context.Context, arg GetCashiersActiveParams) ([]*GetCashiersActiveRow, error)
	GetCashiersByMerchant(ctx context.Context, arg GetCashiersByMerchantParams) ([]*GetCashiersByMerchantRow, error)
	GetCashiersTrashed(ctx context.Context, arg GetCashiersTrashedParams) ([]*GetCashiersTrashedRow, error)
	// Get Categories with Pagination and Total Count
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]*GetCategoriesRow, error)
	// Get Active Categories with Pagination and Total Count
	GetCategoriesActive(ctx context.Context, arg GetCategoriesActiveParams) ([]*GetCategoriesActiveRow, error)
	// Get Trashed Categories with Pagination and Total Count
	GetCategoriesTrashed(ctx context.Context, arg GetCategoriesTrashedParams) ([]*GetCategoriesTrashedRow, error)
	GetCategoryByID(ctx context.Context, categoryID int32) (*Category, error)
	GetMerchantByID(ctx context.Context, merchantID int32) (*Merchant, error)
	GetMerchants(ctx context.Context, arg GetMerchantsParams) ([]*GetMerchantsRow, error)
	// Get Active Merchants with Pagination and Total Count
	GetMerchantsActive(ctx context.Context, arg GetMerchantsActiveParams) ([]*GetMerchantsActiveRow, error)
	// Get Trashed Merchants with Pagination and Total Count
	GetMerchantsTrashed(ctx context.Context, arg GetMerchantsTrashedParams) ([]*GetMerchantsTrashedRow, error)
	GetMonthlyAmountTransactionFailed(ctx context.Context, arg GetMonthlyAmountTransactionFailedParams) ([]*GetMonthlyAmountTransactionFailedRow, error)
	GetMonthlyAmountTransactionFailedByMerchant(ctx context.Context, arg GetMonthlyAmountTransactionFailedByMerchantParams) ([]*GetMonthlyAmountTransactionFailedByMerchantRow, error)
	GetMonthlyAmountTransactionSuccess(ctx context.Context, arg GetMonthlyAmountTransactionSuccessParams) ([]*GetMonthlyAmountTransactionSuccessRow, error)
	GetMonthlyAmountTransactionSuccessByMerchant(ctx context.Context, arg GetMonthlyAmountTransactionSuccessByMerchantParams) ([]*GetMonthlyAmountTransactionSuccessByMerchantRow, error)
	GetMonthlyCashier(ctx context.Context, dollar_1 time.Time) ([]*GetMonthlyCashierRow, error)
	GetMonthlyCashierByCashierId(ctx context.Context, arg GetMonthlyCashierByCashierIdParams) ([]*GetMonthlyCashierByCashierIdRow, error)
	GetMonthlyCashierByMerchant(ctx context.Context, arg GetMonthlyCashierByMerchantParams) ([]*GetMonthlyCashierByMerchantRow, error)
	GetMonthlyCategory(ctx context.Context, dollar_1 time.Time) ([]*GetMonthlyCategoryRow, error)
	GetMonthlyCategoryById(ctx context.Context, arg GetMonthlyCategoryByIdParams) ([]*GetMonthlyCategoryByIdRow, error)
	GetMonthlyCategoryByMerchant(ctx context.Context, arg GetMonthlyCategoryByMerchantParams) ([]*GetMonthlyCategoryByMerchantRow, error)
	GetMonthlyOrder(ctx context.Context, dollar_1 time.Time) ([]*GetMonthlyOrderRow, error)
	GetMonthlyOrderByMerchant(ctx context.Context, arg GetMonthlyOrderByMerchantParams) ([]*GetMonthlyOrderByMerchantRow, error)
	GetMonthlyTotalPrice(ctx context.Context, arg GetMonthlyTotalPriceParams) ([]*GetMonthlyTotalPriceRow, error)
	GetMonthlyTotalPriceById(ctx context.Context, arg GetMonthlyTotalPriceByIdParams) ([]*GetMonthlyTotalPriceByIdRow, error)
	GetMonthlyTotalPriceByMerchant(ctx context.Context, arg GetMonthlyTotalPriceByMerchantParams) ([]*GetMonthlyTotalPriceByMerchantRow, error)
	GetMonthlyTotalRevenue(ctx context.Context, arg GetMonthlyTotalRevenueParams) ([]*GetMonthlyTotalRevenueRow, error)
	GetMonthlyTotalRevenueById(ctx context.Context, arg GetMonthlyTotalRevenueByIdParams) ([]*GetMonthlyTotalRevenueByIdRow, error)
	GetMonthlyTotalRevenueByMerchant(ctx context.Context, arg GetMonthlyTotalRevenueByMerchantParams) ([]*GetMonthlyTotalRevenueByMerchantRow, error)
	GetMonthlyTotalSalesById(ctx context.Context, arg GetMonthlyTotalSalesByIdParams) ([]*GetMonthlyTotalSalesByIdRow, error)
	GetMonthlyTotalSalesByMerchant(ctx context.Context, arg GetMonthlyTotalSalesByMerchantParams) ([]*GetMonthlyTotalSalesByMerchantRow, error)
	GetMonthlyTotalSalesCashier(ctx context.Context, arg GetMonthlyTotalSalesCashierParams) ([]*GetMonthlyTotalSalesCashierRow, error)
	GetMonthlyTransactionMethods(ctx context.Context, dollar_1 time.Time) ([]*GetMonthlyTransactionMethodsRow, error)
	GetMonthlyTransactionMethodsByMerchant(ctx context.Context, arg GetMonthlyTransactionMethodsByMerchantParams) ([]*GetMonthlyTransactionMethodsByMerchantRow, error)
	GetOrderByID(ctx context.Context, orderID int32) (*Order, error)
	// Get Order Items  with Pagination and Total Count
	GetOrderItems(ctx context.Context, arg GetOrderItemsParams) ([]*GetOrderItemsRow, error)
	// Get Active Order Item with Pagination and Total Count
	GetOrderItemsActive(ctx context.Context, arg GetOrderItemsActiveParams) ([]*GetOrderItemsActiveRow, error)
	GetOrderItemsByOrder(ctx context.Context, orderID int32) ([]*OrderItem, error)
	// Get Trashed Orders Items with Pagination and Total Count
	GetOrderItemsTrashed(ctx context.Context, arg GetOrderItemsTrashedParams) ([]*GetOrderItemsTrashedRow, error)
	// Get Orders with Pagination and Total Count
	GetOrders(ctx context.Context, arg GetOrdersParams) ([]*GetOrdersRow, error)
	// Get Active Orders with Pagination and Total Count
	GetOrdersActive(ctx context.Context, arg GetOrdersActiveParams) ([]*GetOrdersActiveRow, error)
	// Get Orders with Pagination and Total Count where merchant_id
	GetOrdersByMerchant(ctx context.Context, arg GetOrdersByMerchantParams) ([]*GetOrdersByMerchantRow, error)
	// Get Trashed Orders with Pagination and Total Count
	GetOrdersTrashed(ctx context.Context, arg GetOrdersTrashedParams) ([]*GetOrdersTrashedRow, error)
	GetProductByID(ctx context.Context, productID int32) (*Product, error)
	GetProductByIdTrashed(ctx context.Context, productID int32) (*Product, error)
	GetProducts(ctx context.Context, arg GetProductsParams) ([]*GetProductsRow, error)
	// Get Active Products with Pagination and Total Count
	GetProductsActive(ctx context.Context, arg GetProductsActiveParams) ([]*GetProductsActiveRow, error)
	// Get Products by Category Name with Filters
	GetProductsByCategoryName(ctx context.Context, arg GetProductsByCategoryNameParams) ([]*GetProductsByCategoryNameRow, error)
	GetProductsByMerchant(ctx context.Context, arg GetProductsByMerchantParams) ([]*GetProductsByMerchantRow, error)
	// Get Trashed Products with Pagination and Total Count
	GetProductsTrashed(ctx context.Context, arg GetProductsTrashedParams) ([]*GetProductsTrashedRow, error)
	GetRole(ctx context.Context, roleID int32) (*Role, error)
	GetRoleByName(ctx context.Context, roleName string) (*Role, error)
	GetRoles(ctx context.Context, arg GetRolesParams) ([]*GetRolesRow, error)
	GetTransactionByID(ctx context.Context, transactionID int32) (*Transaction, error)
	GetTransactionByMerchant(ctx context.Context, arg GetTransactionByMerchantParams) ([]*GetTransactionByMerchantRow, error)
	GetTransactionByOrderID(ctx context.Context, orderID int32) (*Transaction, error)
	GetTransactions(ctx context.Context, arg GetTransactionsParams) ([]*GetTransactionsRow, error)
	// Get Active Transactions with Pagination and Total Count
	GetTransactionsActive(ctx context.Context, arg GetTransactionsActiveParams) ([]*GetTransactionsActiveRow, error)
	// Get Trashed Transactions with Pagination and Total Count
	GetTransactionsTrashed(ctx context.Context, arg GetTransactionsTrashedParams) ([]*GetTransactionsTrashedRow, error)
	// Get All Trashed Roles
	GetTrashedRoles(ctx context.Context, arg GetTrashedRolesParams) ([]*GetTrashedRolesRow, error)
	GetTrashedUserRoles(ctx context.Context, userID int32) ([]*GetTrashedUserRolesRow, error)
	// Get User by Email
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	// Get User by ID
	GetUserByID(ctx context.Context, userID int32) (*User, error)
	GetUserRoles(ctx context.Context, userID int32) ([]*Role, error)
	// Get Trashed Users with Pagination and Total Count
	GetUserTrashed(ctx context.Context, arg GetUserTrashedParams) ([]*GetUserTrashedRow, error)
	// Get Users with Pagination and Total Count
	GetUsers(ctx context.Context, arg GetUsersParams) ([]*GetUsersRow, error)
	// Get Active Users with Pagination and Total Count
	GetUsersActive(ctx context.Context, arg GetUsersActiveParams) ([]*GetUsersActiveRow, error)
	GetYearlyAmountTransactionFailed(ctx context.Context, dollar_1 int32) ([]*GetYearlyAmountTransactionFailedRow, error)
	GetYearlyAmountTransactionFailedByMerchant(ctx context.Context, arg GetYearlyAmountTransactionFailedByMerchantParams) ([]*GetYearlyAmountTransactionFailedByMerchantRow, error)
	GetYearlyAmountTransactionSuccess(ctx context.Context, dollar_1 int32) ([]*GetYearlyAmountTransactionSuccessRow, error)
	GetYearlyAmountTransactionSuccessByMerchant(ctx context.Context, arg GetYearlyAmountTransactionSuccessByMerchantParams) ([]*GetYearlyAmountTransactionSuccessByMerchantRow, error)
	GetYearlyCashier(ctx context.Context, dollar_1 time.Time) ([]*GetYearlyCashierRow, error)
	GetYearlyCashierByCashierId(ctx context.Context, arg GetYearlyCashierByCashierIdParams) ([]*GetYearlyCashierByCashierIdRow, error)
	GetYearlyCashierByMerchant(ctx context.Context, arg GetYearlyCashierByMerchantParams) ([]*GetYearlyCashierByMerchantRow, error)
	GetYearlyCategory(ctx context.Context, dollar_1 time.Time) ([]*GetYearlyCategoryRow, error)
	GetYearlyCategoryById(ctx context.Context, arg GetYearlyCategoryByIdParams) ([]*GetYearlyCategoryByIdRow, error)
	GetYearlyCategoryByMerchant(ctx context.Context, arg GetYearlyCategoryByMerchantParams) ([]*GetYearlyCategoryByMerchantRow, error)
	GetYearlyOrder(ctx context.Context, dollar_1 time.Time) ([]*GetYearlyOrderRow, error)
	GetYearlyOrderByMerchant(ctx context.Context, arg GetYearlyOrderByMerchantParams) ([]*GetYearlyOrderByMerchantRow, error)
	GetYearlyTotalPrice(ctx context.Context, dollar_1 int32) ([]*GetYearlyTotalPriceRow, error)
	GetYearlyTotalPriceById(ctx context.Context, arg GetYearlyTotalPriceByIdParams) ([]*GetYearlyTotalPriceByIdRow, error)
	GetYearlyTotalPriceByMerchant(ctx context.Context, arg GetYearlyTotalPriceByMerchantParams) ([]*GetYearlyTotalPriceByMerchantRow, error)
	GetYearlyTotalRevenue(ctx context.Context, dollar_1 int32) ([]*GetYearlyTotalRevenueRow, error)
	GetYearlyTotalRevenueById(ctx context.Context, arg GetYearlyTotalRevenueByIdParams) ([]*GetYearlyTotalRevenueByIdRow, error)
	GetYearlyTotalRevenueByMerchant(ctx context.Context, arg GetYearlyTotalRevenueByMerchantParams) ([]*GetYearlyTotalRevenueByMerchantRow, error)
	GetYearlyTotalSalesById(ctx context.Context, arg GetYearlyTotalSalesByIdParams) ([]*GetYearlyTotalSalesByIdRow, error)
	GetYearlyTotalSalesByMerchant(ctx context.Context, arg GetYearlyTotalSalesByMerchantParams) ([]*GetYearlyTotalSalesByMerchantRow, error)
	GetYearlyTotalSalesCashier(ctx context.Context, dollar_1 int32) ([]*GetYearlyTotalSalesCashierRow, error)
	GetYearlyTransactionMethods(ctx context.Context, dollar_1 time.Time) ([]*GetYearlyTransactionMethodsRow, error)
	GetYearlyTransactionMethodsByMerchant(ctx context.Context, arg GetYearlyTransactionMethodsByMerchantParams) ([]*GetYearlyTransactionMethodsByMerchantRow, error)
	RemoveRoleFromUser(ctx context.Context, arg RemoveRoleFromUserParams) error
	// Restore All Trashed Cashier
	RestoreAllCashiers(ctx context.Context) error
	// Restore All Trashed Category
	RestoreAllCategories(ctx context.Context) error
	// Restore All Trashed Merchant
	RestoreAllMerchants(ctx context.Context) error
	// Restore All Trashed Order
	RestoreAllOrders(ctx context.Context) error
	// Restore All Trashed Order Item
	RestoreAllOrdersItem(ctx context.Context) error
	// Restore All Trashed Product
	RestoreAllProducts(ctx context.Context) error
	// Restore All Trashed Roles
	RestoreAllRoles(ctx context.Context) error
	// Restore All Trashed Transaction
	RestoreAllTransactions(ctx context.Context) error
	// Restore All Trashed Users
	RestoreAllUsers(ctx context.Context) error
	// Restore Trashed Cashier
	RestoreCashier(ctx context.Context, cashierID int32) (*Cashier, error)
	// Restore Trashed Category
	RestoreCategory(ctx context.Context, categoryID int32) (*Category, error)
	// Restore Trashed Merchant
	RestoreMerchant(ctx context.Context, merchantID int32) (*Merchant, error)
	// Restore Trashed Order
	RestoreOrder(ctx context.Context, orderID int32) (*Order, error)
	// Restore Trashed Order Item
	RestoreOrderItem(ctx context.Context, orderItemID int32) (*OrderItem, error)
	// Restore Trashed Product
	RestoreProduct(ctx context.Context, productID int32) (*Product, error)
	RestoreRole(ctx context.Context, roleID int32) error
	// Restore Trashed Transaction
	RestoreTransaction(ctx context.Context, transactionID int32) (*Transaction, error)
	// Restore Trashed User
	RestoreUser(ctx context.Context, userID int32) (*User, error)
	RestoreUserRole(ctx context.Context, userRoleID int32) error
	// Trash Cashier
	TrashCashier(ctx context.Context, cashierID int32) (*Cashier, error)
	// Trash Category
	TrashCategory(ctx context.Context, categoryID int32) (*Category, error)
	// Trash Merchant
	TrashMerchant(ctx context.Context, merchantID int32) (*Merchant, error)
	// Correct query to trash a specific order item
	TrashOrderItem(ctx context.Context, orderItemID int32) (*OrderItem, error)
	// Trash Product
	TrashProduct(ctx context.Context, productID int32) (*Product, error)
	TrashRole(ctx context.Context, roleID int32) error
	// Trash Transaction
	TrashTransaction(ctx context.Context, transactionID int32) (*Transaction, error)
	// Trash User
	TrashUser(ctx context.Context, userID int32) (*User, error)
	TrashUserRole(ctx context.Context, userRoleID int32) error
	// Trash Order
	TrashedOrder(ctx context.Context, orderID int32) (*Order, error)
	UpdateCashier(ctx context.Context, arg UpdateCashierParams) (*Cashier, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (*Category, error)
	// Update Merchant
	UpdateMerchant(ctx context.Context, arg UpdateMerchantParams) (*Merchant, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (*Order, error)
	UpdateOrderItem(ctx context.Context, arg UpdateOrderItemParams) (*OrderItem, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (*Product, error)
	UpdateProductCountStock(ctx context.Context, arg UpdateProductCountStockParams) (*Product, error)
	UpdateRefreshTokenByUserId(ctx context.Context, arg UpdateRefreshTokenByUserIdParams) error
	UpdateRole(ctx context.Context, arg UpdateRoleParams) (*Role, error)
	UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (*Transaction, error)
	// Update User
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error)
}

var _ Querier = (*Queries)(nil)
