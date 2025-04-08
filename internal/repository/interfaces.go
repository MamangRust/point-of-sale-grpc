package repository

import (
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
)

type UserRepository interface {
	FindAllUsers(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindById(user_id int) (*record.UserRecord, error)
	FindByEmail(email string) (*record.UserRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.UserRecord, int, error)
	CreateUser(request *requests.CreateUserRequest) (*record.UserRecord, error)
	UpdateUser(request *requests.UpdateUserRequest) (*record.UserRecord, error)
	TrashedUser(user_id int) (*record.UserRecord, error)
	RestoreUser(user_id int) (*record.UserRecord, error)
	DeleteUserPermanent(user_id int) (bool, error)
	RestoreAllUser() (bool, error)
	DeleteAllUserPermanent() (bool, error)
}

type RoleRepository interface {
	FindAllRoles(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	FindById(role_id int) (*record.RoleRecord, error)
	FindByName(name string) (*record.RoleRecord, error)
	FindByUserId(user_id int) ([]*record.RoleRecord, error)
	FindByActiveRole(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
	FindByTrashedRole(page int, pageSize int, search string) ([]*record.RoleRecord, int, error)
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
	GetMonthlyTotalPrice(year int, month int) ([]*record.CategoriesMonthlyTotalPriceRecord, error)
	GetYearlyTotalPrices(year int) ([]*record.CategoriesYearlyTotalPriceRecord, error)
	GetMonthlyTotalPriceById(year int, month int, category_id int) ([]*record.CategoriesMonthlyTotalPriceRecord, error)
	GetYearlyTotalPricesById(year int, category_id int) ([]*record.CategoriesYearlyTotalPriceRecord, error)
	GetMonthlyTotalPriceByMerchant(year int, month int, merchant_id int) ([]*record.CategoriesMonthlyTotalPriceRecord, error)
	GetYearlyTotalPricesByMerchant(year int, merchant_id int) ([]*record.CategoriesYearlyTotalPriceRecord, error)

	GetMonthPrice(year int) ([]*record.CategoriesMonthPriceRecord, error)
	GetYearPrice(year int) ([]*record.CategoriesYearPriceRecord, error)
	GetMonthPriceByMerchant(year int, merchant_id int) ([]*record.CategoriesMonthPriceRecord, error)
	GetYearPriceByMerchant(year int, merchant_id int) ([]*record.CategoriesYearPriceRecord, error)
	GetMonthPriceById(year int, category_id int) ([]*record.CategoriesMonthPriceRecord, error)
	GetYearPriceById(year int, category_id int) ([]*record.CategoriesYearPriceRecord, error)

	FindAllCategory(search string, page, pageSize int) ([]*record.CategoriesRecord, int, error)
	FindById(category_id int) (*record.CategoriesRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.CategoriesRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.CategoriesRecord, int, error)
	CreateCategory(request *requests.CreateCategoryRequest) (*record.CategoriesRecord, error)
	UpdateCategory(request *requests.UpdateCategoryRequest) (*record.CategoriesRecord, error)
	TrashedCategory(category_id int) (*record.CategoriesRecord, error)
	RestoreCategory(category_id int) (*record.CategoriesRecord, error)
	DeleteCategoryPermanently(Category_id int) (bool, error)
	RestoreAllCategories() (bool, error)
	DeleteAllPermanentCategories() (bool, error)
}

type CashierRepository interface {
	GetMonthlyTotalSales(year int, month int) ([]*record.CashierRecordMonthTotalSales, error)
	GetYearlyTotalSales(year int) ([]*record.CashierRecordYearTotalSales, error)
	GetMonthlyTotalSalesById(year int, month int, cashier_id int) ([]*record.CashierRecordMonthTotalSales, error)
	GetYearlyTotalSalesById(year int, cashier_id int) ([]*record.CashierRecordYearTotalSales, error)
	GetMonthlyTotalSalesByMerchant(year int, month int, merchant_id int) ([]*record.CashierRecordMonthTotalSales, error)
	GetYearlyTotalSalesByMerchant(year int, merchant_id int) ([]*record.CashierRecordYearTotalSales, error)

	GetMonthyCashier(year int) ([]*record.CashierRecordMonthSales, error)
	GetYearlyCashier(year int) ([]*record.CashierRecordYearSales, error)
	GetMonthlyCashierByMerchant(year int, merchant_id int) ([]*record.CashierRecordMonthSales, error)
	GetYearlyCashierByMerchant(year int, merchant_id int) ([]*record.CashierRecordYearSales, error)
	GetMonthlyCashierById(year int, cashier_id int) ([]*record.CashierRecordMonthSales, error)
	GetYearlyCashierById(year int, cashier_id int) ([]*record.CashierRecordYearSales, error)

	FindAllCashiers(search string, page, pageSize int) ([]*record.CashierRecord, int, error)
	FindById(cashier_id int) (*record.CashierRecord, error)
	FindByActive(search string, page, pageSize int) ([]*record.CashierRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.CashierRecord, int, error)
	FindByMerchant(merchant_id int, search string, page, pageSize int) ([]*record.CashierRecord, int, error)
	CreateCashier(request *requests.CreateCashierRequest) (*record.CashierRecord, error)
	UpdateCashier(request *requests.UpdateCashierRequest) (*record.CashierRecord, error)
	TrashedCashier(cashier_id int) (*record.CashierRecord, error)
	RestoreCashier(cashier_id int) (*record.CashierRecord, error)
	DeleteCashierPermanent(cashier_id int) (bool, error)
	RestoreAllCashier() (bool, error)
	DeleteAllCashierPermanent() (bool, error)
}

type MerchantRepository interface {
	FindAllMerchants(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.MerchantRecord, int, error)
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
	GetMonthlyTotalRevenue(year int, month int) ([]*record.OrderMonthlyTotalRevenueRecord, error)
	GetYearlyTotalRevenue(year int) ([]*record.OrderYearlyTotalRevenueRecord, error)
	GetMonthlyTotalRevenueById(year int, month int, order_id int) ([]*record.OrderMonthlyTotalRevenueRecord, error)
	GetYearlyTotalRevenueById(year int, order_id int) ([]*record.OrderYearlyTotalRevenueRecord, error)
	GetMonthlyTotalRevenueByMerchant(year int, month int, merchant_id int) ([]*record.OrderMonthlyTotalRevenueRecord, error)
	GetYearlyTotalRevenueByMerchant(year int, merchant_id int) ([]*record.OrderYearlyTotalRevenueRecord, error)

	GetMonthlyOrder(year int) ([]*record.OrderMonthlyRecord, error)
	GetYearlyOrder(year int) ([]*record.OrderYearlyRecord, error)
	GetMonthlyOrderByMerchant(year int, merchant_id int) ([]*record.OrderMonthlyRecord, error)
	GetYearlyOrderByMerchant(year int, merchant_id int) ([]*record.OrderYearlyRecord, error)

	FindAllOrders(search string, page, pageSize int) ([]*record.OrderRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.OrderRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.OrderRecord, int, error)
	FindByMerchant(merchant_id int, search string, page, pageSize int) ([]*record.OrderRecord, int, error)
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
	FindAllOrderItems(search string, page, pageSize int) ([]*record.OrderItemRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.OrderItemRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.OrderItemRecord, int, error)
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
	FindAllProducts(search string, page, pageSize int) ([]*record.ProductRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.ProductRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.ProductRecord, int, error)
	FindByMerchant(req *requests.ProductByMerchantRequest) ([]*record.ProductRecord, int, error)

	FindByCategory(category_name string, search string, page, pageSize int) ([]*record.ProductRecord, int, error)

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
	GetMonthlyAmountSuccess(year int, month int) ([]*record.TransactionMonthlyAmountSuccessRecord, error)
	GetYearlyAmountSuccess(year int) ([]*record.TransactionYearlyAmountSuccessRecord, error)
	GetMonthlyAmountFailed(year int, month int) ([]*record.TransactionMonthlyAmountFailedRecord, error)
	GetYearlyAmountFailed(year int) ([]*record.TransactionYearlyAmountFailedRecord, error)
	GetMonthlyAmountSuccessByMerchant(year int, month int, merchantID int) ([]*record.TransactionMonthlyAmountSuccessRecord, error)
	GetYearlyAmountSuccessByMerchant(year int, merchantID int) ([]*record.TransactionYearlyAmountSuccessRecord, error)
	GetMonthlyAmountFailedByMerchant(year int, month int, merchantID int) ([]*record.TransactionMonthlyAmountFailedRecord, error)
	GetYearlyAmountFailedByMerchant(year int, merchantID int) ([]*record.TransactionYearlyAmountFailedRecord, error)

	GetMonthlyTransactionMethod(year int) ([]*record.TransactionMonthlyMethodRecord, error)
	GetYearlyTransactionMethod(year int) ([]*record.TransactionYearlyMethodRecord, error)
	GetMonthlyTransactionMethodByMerchant(year int, merchant_id int) ([]*record.TransactionMonthlyMethodRecord, error)
	GetYearlyTransactionMethodByMerchant(year int, merchant_id int) ([]*record.TransactionYearlyMethodRecord, error)

	FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByActive(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByTrashed(search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
	FindByMerchant(merchant_id int, search string, page, pageSize int) ([]*record.TransactionRecord, int, error)
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
