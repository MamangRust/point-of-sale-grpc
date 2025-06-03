package repository

import (
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
)

type UserRepository interface {
	FindAllUsers(req *requests.FindAllUsers) ([]*record.UserRecord, *int, error)
	FindById(user_id int) (*record.UserRecord, error)
	FindByEmail(email string) (*record.UserRecord, error)
	FindByActive(req *requests.FindAllUsers) ([]*record.UserRecord, *int, error)
	FindByTrashed(req *requests.FindAllUsers) ([]*record.UserRecord, *int, error)
	CreateUser(request *requests.CreateUserRequest) (*record.UserRecord, error)
	UpdateUser(request *requests.UpdateUserRequest) (*record.UserRecord, error)
	TrashedUser(user_id int) (*record.UserRecord, error)
	RestoreUser(user_id int) (*record.UserRecord, error)
	DeleteUserPermanent(user_id int) (bool, error)
	RestoreAllUser() (bool, error)
	DeleteAllUserPermanent() (bool, error)
}

type RoleRepository interface {
	FindAllRoles(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error)
	FindById(role_id int) (*record.RoleRecord, error)
	FindByName(name string) (*record.RoleRecord, error)
	FindByUserId(user_id int) ([]*record.RoleRecord, error)
	FindByActiveRole(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error)
	FindByTrashedRole(req *requests.FindAllRoles) ([]*record.RoleRecord, *int, error)
	CreateRole(request *requests.CreateRoleRequest) (*record.RoleRecord, error)
	UpdateRole(request *requests.UpdateRoleRequest) (*record.RoleRecord, error)
	TrashedRole(role_id int) (*record.RoleRecord, error)

	RestoreRole(role_id int) (*record.RoleRecord, error)
	DeleteRolePermanent(role_id int) (bool, error)
	RestoreAllRole() (bool, error)
	DeleteAllRolePermanent() (bool, error)
}
type RefreshTokenRepository interface {
	FindByToken(token string) (*record.RefreshTokenRecord, error)
	FindByUserId(user_id int) (*record.RefreshTokenRecord, error)
	CreateRefreshToken(req *requests.CreateRefreshToken) (*record.RefreshTokenRecord, error)
	UpdateRefreshToken(req *requests.UpdateRefreshToken) (*record.RefreshTokenRecord, error)
	DeleteRefreshToken(token string) error
	DeleteRefreshTokenByUserId(user_id int) error
}

type UserRoleRepository interface {
	AssignRoleToUser(req *requests.CreateUserRoleRequest) (*record.UserRoleRecord, error)
	RemoveRoleFromUser(req *requests.RemoveUserRoleRequest) error
}

type CategoryRepository interface {
	GetMonthlyTotalPrice(req *requests.MonthTotalPrice) ([]*record.CategoriesMonthlyTotalPriceRecord, error)
	GetYearlyTotalPrices(year int) ([]*record.CategoriesYearlyTotalPriceRecord, error)
	GetMonthlyTotalPriceById(req *requests.MonthTotalPriceCategory) ([]*record.CategoriesMonthlyTotalPriceRecord, error)
	GetYearlyTotalPricesById(req *requests.YearTotalPriceCategory) ([]*record.CategoriesYearlyTotalPriceRecord, error)
	GetMonthlyTotalPriceByMerchant(req *requests.MonthTotalPriceMerchant) ([]*record.CategoriesMonthlyTotalPriceRecord, error)
	GetYearlyTotalPricesByMerchant(req *requests.YearTotalPriceMerchant) ([]*record.CategoriesYearlyTotalPriceRecord, error)

	GetMonthPrice(year int) ([]*record.CategoriesMonthPriceRecord, error)
	GetYearPrice(year int) ([]*record.CategoriesYearPriceRecord, error)
	GetMonthPriceByMerchant(req *requests.MonthPriceMerchant) ([]*record.CategoriesMonthPriceRecord, error)
	GetYearPriceByMerchant(req *requests.YearPriceMerchant) ([]*record.CategoriesYearPriceRecord, error)
	GetMonthPriceById(req *requests.MonthPriceId) ([]*record.CategoriesMonthPriceRecord, error)
	GetYearPriceById(req *requests.YearPriceId) ([]*record.CategoriesYearPriceRecord, error)

	FindAllCategory(req *requests.FindAllCategory) ([]*record.CategoriesRecord, *int, error)
	FindById(category_id int) (*record.CategoriesRecord, error)
	FindByNameAndId(req *requests.CategoryNameAndId) (*record.CategoriesRecord, error)
	FindByName(name string) (*record.CategoriesRecord, error)
	FindByActive(req *requests.FindAllCategory) ([]*record.CategoriesRecord, *int, error)
	FindByTrashed(req *requests.FindAllCategory) ([]*record.CategoriesRecord, *int, error)
	CreateCategory(request *requests.CreateCategoryRequest) (*record.CategoriesRecord, error)
	UpdateCategory(request *requests.UpdateCategoryRequest) (*record.CategoriesRecord, error)
	TrashedCategory(category_id int) (*record.CategoriesRecord, error)
	RestoreCategory(category_id int) (*record.CategoriesRecord, error)
	DeleteCategoryPermanently(Category_id int) (bool, error)
	RestoreAllCategories() (bool, error)
	DeleteAllPermanentCategories() (bool, error)
}

type CashierRepository interface {
	GetMonthlyTotalSales(req *requests.MonthTotalSales) ([]*record.CashierRecordMonthTotalSales, error)
	GetYearlyTotalSales(year int) ([]*record.CashierRecordYearTotalSales, error)
	GetMonthlyTotalSalesById(req *requests.MonthTotalSalesCashier) ([]*record.CashierRecordMonthTotalSales, error)
	GetYearlyTotalSalesById(req *requests.YearTotalSalesCashier) ([]*record.CashierRecordYearTotalSales, error)
	GetMonthlyTotalSalesByMerchant(req *requests.MonthTotalSalesMerchant) ([]*record.CashierRecordMonthTotalSales, error)
	GetYearlyTotalSalesByMerchant(req *requests.YearTotalSalesMerchant) ([]*record.CashierRecordYearTotalSales, error)

	GetMonthyCashier(year int) ([]*record.CashierRecordMonthSales, error)
	GetYearlyCashier(year int) ([]*record.CashierRecordYearSales, error)
	GetMonthlyCashierByMerchant(req *requests.MonthCashierMerchant) ([]*record.CashierRecordMonthSales, error)
	GetYearlyCashierByMerchant(req *requests.YearCashierMerchant) ([]*record.CashierRecordYearSales, error)
	GetMonthlyCashierById(req *requests.MonthCashierId) ([]*record.CashierRecordMonthSales, error)
	GetYearlyCashierById(req *requests.YearCashierId) ([]*record.CashierRecordYearSales, error)

	FindAllCashiers(req *requests.FindAllCashiers) ([]*record.CashierRecord, *int, error)
	FindById(cashier_id int) (*record.CashierRecord, error)
	FindByActive(req *requests.FindAllCashiers) ([]*record.CashierRecord, *int, error)
	FindByTrashed(req *requests.FindAllCashiers) ([]*record.CashierRecord, *int, error)
	FindByMerchant(req *requests.FindAllCashierMerchant) ([]*record.CashierRecord, *int, error)
	CreateCashier(request *requests.CreateCashierRequest) (*record.CashierRecord, error)
	UpdateCashier(request *requests.UpdateCashierRequest) (*record.CashierRecord, error)
	TrashedCashier(cashier_id int) (*record.CashierRecord, error)
	RestoreCashier(cashier_id int) (*record.CashierRecord, error)
	DeleteCashierPermanent(cashier_id int) (bool, error)
	RestoreAllCashier() (bool, error)
	DeleteAllCashierPermanent() (bool, error)
}

type MerchantRepository interface {
	FindAllMerchants(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error)
	FindByActive(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error)
	FindByTrashed(req *requests.FindAllMerchants) ([]*record.MerchantRecord, *int, error)
	FindById(user_id int) (*record.MerchantRecord, error)
	CreateMerchant(request *requests.CreateMerchantRequest) (*record.MerchantRecord, error)
	UpdateMerchant(request *requests.UpdateMerchantRequest) (*record.MerchantRecord, error)
	TrashedMerchant(merchant_id int) (*record.MerchantRecord, error)
	RestoreMerchant(merchant_id int) (*record.MerchantRecord, error)
	DeleteMerchantPermanent(Merchant_id int) (bool, error)
	RestoreAllMerchant() (bool, error)
	DeleteAllMerchantPermanent() (bool, error)
}

type OrderRepository interface {
	GetMonthlyTotalRevenue(req *requests.MonthTotalRevenue) ([]*record.OrderMonthlyTotalRevenueRecord, error)
	GetYearlyTotalRevenue(year int) ([]*record.OrderYearlyTotalRevenueRecord, error)
	GetMonthlyTotalRevenueById(req *requests.MonthTotalRevenueOrder) ([]*record.OrderMonthlyTotalRevenueRecord, error)
	GetYearlyTotalRevenueById(req *requests.YearTotalRevenueOrder) ([]*record.OrderYearlyTotalRevenueRecord, error)
	GetMonthlyTotalRevenueByMerchant(req *requests.MonthTotalRevenueMerchant) ([]*record.OrderMonthlyTotalRevenueRecord, error)
	GetYearlyTotalRevenueByMerchant(req *requests.YearTotalRevenueMerchant) ([]*record.OrderYearlyTotalRevenueRecord, error)

	GetMonthlyOrder(year int) ([]*record.OrderMonthlyRecord, error)
	GetYearlyOrder(year int) ([]*record.OrderYearlyRecord, error)
	GetMonthlyOrderByMerchant(req *requests.MonthOrderMerchant) ([]*record.OrderMonthlyRecord, error)
	GetYearlyOrderByMerchant(req *requests.YearOrderMerchant) ([]*record.OrderYearlyRecord, error)

	FindAllOrders(req *requests.FindAllOrders) ([]*record.OrderRecord, *int, error)
	FindByActive(req *requests.FindAllOrders) ([]*record.OrderRecord, *int, error)
	FindByTrashed(req *requests.FindAllOrders) ([]*record.OrderRecord, *int, error)
	FindByMerchant(req *requests.FindAllOrderMerchant) ([]*record.OrderRecord, *int, error)
	FindById(order_id int) (*record.OrderRecord, error)
	CreateOrder(request *requests.CreateOrderRecordRequest) (*record.OrderRecord, error)
	UpdateOrder(request *requests.UpdateOrderRecordRequest) (*record.OrderRecord, error)
	TrashedOrder(order_id int) (*record.OrderRecord, error)
	RestoreOrder(order_id int) (*record.OrderRecord, error)
	DeleteOrderPermanent(order_id int) (bool, error)
	RestoreAllOrder() (bool, error)
	DeleteAllOrderPermanent() (bool, error)
}

type OrderItemRepository interface {
	FindAllOrderItems(req *requests.FindAllOrderItems) ([]*record.OrderItemRecord, *int, error)
	FindByActive(req *requests.FindAllOrderItems) ([]*record.OrderItemRecord, *int, error)
	FindByTrashed(req *requests.FindAllOrderItems) ([]*record.OrderItemRecord, *int, error)
	FindOrderItemByOrder(order_id int) ([]*record.OrderItemRecord, error)
	CalculateTotalPrice(order_id int) (*int32, error)
	CreateOrderItem(req *requests.CreateOrderItemRecordRequest) (*record.OrderItemRecord, error)
	UpdateOrderItem(req *requests.UpdateOrderItemRecordRequest) (*record.OrderItemRecord, error)
	TrashedOrderItem(order_id int) (*record.OrderItemRecord, error)
	RestoreOrderItem(order_id int) (*record.OrderItemRecord, error)
	DeleteOrderItemPermanent(order_id int) (bool, error)
	RestoreAllOrderItem() (bool, error)
	DeleteAllOrderPermanent() (bool, error)
}

type ProductRepository interface {
	FindAllProducts(req *requests.FindAllProducts) ([]*record.ProductRecord, *int, error)
	FindByActive(req *requests.FindAllProducts) ([]*record.ProductRecord, *int, error)
	FindByTrashed(req *requests.FindAllProducts) ([]*record.ProductRecord, *int, error)
	FindByMerchant(req *requests.ProductByMerchantRequest) ([]*record.ProductRecord, *int, error)
	FindByCategory(req *requests.ProductByCategoryRequest) ([]*record.ProductRecord, *int, error)
	FindById(user_id int) (*record.ProductRecord, error)
	FindByIdTrashed(id int) (*record.ProductRecord, error)
	CreateProduct(request *requests.CreateProductRequest) (*record.ProductRecord, error)
	UpdateProduct(request *requests.UpdateProductRequest) (*record.ProductRecord, error)
	UpdateProductCountStock(product_id int, stock int) (*record.ProductRecord, error)
	TrashedProduct(user_id int) (*record.ProductRecord, error)
	RestoreProduct(user_id int) (*record.ProductRecord, error)
	DeleteProductPermanent(user_id int) (bool, error)
	RestoreAllProducts() (bool, error)
	DeleteAllProductPermanent() (bool, error)
}

type TransactionRepository interface {
	GetMonthlyAmountSuccess(req *requests.MonthAmountTransaction) ([]*record.TransactionMonthlyAmountSuccessRecord, error)
	GetYearlyAmountSuccess(year int) ([]*record.TransactionYearlyAmountSuccessRecord, error)
	GetMonthlyAmountFailed(req *requests.MonthAmountTransaction) ([]*record.TransactionMonthlyAmountFailedRecord, error)
	GetYearlyAmountFailed(year int) ([]*record.TransactionYearlyAmountFailedRecord, error)

	GetMonthlyAmountSuccessByMerchant(req *requests.MonthAmountTransactionMerchant) ([]*record.TransactionMonthlyAmountSuccessRecord, error)
	GetYearlyAmountSuccessByMerchant(req *requests.YearAmountTransactionMerchant) ([]*record.TransactionYearlyAmountSuccessRecord, error)
	GetMonthlyAmountFailedByMerchant(req *requests.MonthAmountTransactionMerchant) ([]*record.TransactionMonthlyAmountFailedRecord, error)
	GetYearlyAmountFailedByMerchant(req *requests.YearAmountTransactionMerchant) ([]*record.TransactionYearlyAmountFailedRecord, error)

	GetMonthlyTransactionMethodSuccess(req *requests.MonthMethodTransaction) ([]*record.TransactionMonthlyMethodRecord, error)
	GetYearlyTransactionMethodSuccess(year int) ([]*record.TransactionYearlyMethodRecord, error)
	GetMonthlyTransactionMethodByMerchantSuccess(req *requests.MonthMethodTransactionMerchant) ([]*record.TransactionMonthlyMethodRecord, error)
	GetYearlyTransactionMethodByMerchantSuccess(req *requests.YearMethodTransactionMerchant) ([]*record.TransactionYearlyMethodRecord, error)

	GetMonthlyTransactionMethodFailed(req *requests.MonthMethodTransaction) ([]*record.TransactionMonthlyMethodRecord, error)
	GetYearlyTransactionMethodFailed(year int) ([]*record.TransactionYearlyMethodRecord, error)
	GetMonthlyTransactionMethodByMerchantFailed(req *requests.MonthMethodTransactionMerchant) ([]*record.TransactionMonthlyMethodRecord, error)
	GetYearlyTransactionMethodByMerchantFailed(req *requests.YearMethodTransactionMerchant) ([]*record.TransactionYearlyMethodRecord, error)

	FindAllTransactions(req *requests.FindAllTransaction) ([]*record.TransactionRecord, *int, error)
	FindByActive(req *requests.FindAllTransaction) ([]*record.TransactionRecord, *int, error)
	FindByTrashed(req *requests.FindAllTransaction) ([]*record.TransactionRecord, *int, error)
	FindByMerchant(req *requests.FindAllTransactionByMerchant) ([]*record.TransactionRecord, *int, error)
	FindById(transaction_id int) (*record.TransactionRecord, error)
	FindByOrderId(order_id int) (*record.TransactionRecord, error)
	CreateTransaction(request *requests.CreateTransactionRequest) (*record.TransactionRecord, error)
	UpdateTransaction(request *requests.UpdateTransactionRequest) (*record.TransactionRecord, error)
	TrashTransaction(transaction_id int) (*record.TransactionRecord, error)
	RestoreTransaction(transaction_id int) (*record.TransactionRecord, error)
	DeleteTransactionPermanently(transaction_id int) (bool, error)
	RestoreAllTransactions() (bool, error)
	DeleteAllTransactionPermanent() (bool, error)
}
