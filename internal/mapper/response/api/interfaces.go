package response_api

import (
	"pointofsale/internal/domain/response"
	"pointofsale/internal/pb"
)

type AuthResponseMapper interface {
	ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin
	ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister
	ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken
	ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe
}

type RoleResponseMapper interface {
	ToApiResponseRoleAll(pbResponse *pb.ApiResponseRoleAll) *response.ApiResponseRoleAll
	ToApiResponseRoleDelete(pbResponse *pb.ApiResponseRoleDelete) *response.ApiResponseRoleDelete
	ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole
	ToApiResponsesRole(pbResponse *pb.ApiResponsesRole) *response.ApiResponsesRole
	ToApiResponsePaginationRole(pbResponse *pb.ApiResponsePaginationRole) *response.ApiResponsePaginationRole
	ToApiResponsePaginationRoleDeleteAt(pbResponse *pb.ApiResponsePaginationRoleDeleteAt) *response.ApiResponsePaginationRoleDeleteAt
}

type UserResponseMapper interface {
	ToApiResponseUserDeleteAt(pbResponse *pb.ApiResponseUserDeleteAt) *response.ApiResponseUserDeleteAt
	ToApiResponseUser(pbResponse *pb.ApiResponseUser) *response.ApiResponseUser
	ToApiResponsesUser(pbResponse *pb.ApiResponsesUser) *response.ApiResponsesUser

	ToApiResponseUserDelete(pbResponse *pb.ApiResponseUserDelete) *response.ApiResponseUserDelete
	ToApiResponseUserAll(pbResponse *pb.ApiResponseUserAll) *response.ApiResponseUserAll
	ToApiResponsePaginationUserDeleteAt(pbResponse *pb.ApiResponsePaginationUserDeleteAt) *response.ApiResponsePaginationUserDeleteAt
	ToApiResponsePaginationUser(pbResponse *pb.ApiResponsePaginationUser) *response.ApiResponsePaginationUser
}

type CategoryResponseMapper interface {
	ToApiResponseCategoryMonthlyTotalPrice(pbResponse *pb.ApiResponseCategoryMonthlyTotalPrice) *response.ApiResponseCategoryMonthlyTotalPrice
	ToApiResponseCategoryYearlyTotalPrice(pbResponse *pb.ApiResponseCategoryYearlyTotalPrice) *response.ApiResponseCategoryYearlyTotalPrice

	ToApiResponseCategoryMonthlyPrice(pbResponse *pb.ApiResponseCategoryMonthPrice) *response.ApiResponseCategoryMonthPrice
	ToApiResponseCategoryYearlyPrice(pbResponse *pb.ApiResponseCategoryYearPrice) *response.ApiResponseCategoryYearPrice

	ToApiResponseCategory(pbResponse *pb.ApiResponseCategory) *response.ApiResponseCategory
	ToApiResponseCategoryDeleteAt(pbResponse *pb.ApiResponseCategoryDeleteAt) *response.ApiResponseCategoryDeleteAt
	ToApiResponsesCategory(pbResponse *pb.ApiResponsesCategory) *response.ApiResponsesCategory
	ToApiResponseCategoryDelete(pbResponse *pb.ApiResponseCategoryDelete) *response.ApiResponseCategoryDelete
	ToApiResponseCategoryAll(pbResponse *pb.ApiResponseCategoryAll) *response.ApiResponseCategoryAll
	ToApiResponsePaginationCategoryDeleteAt(pbResponse *pb.ApiResponsePaginationCategoryDeleteAt) *response.ApiResponsePaginationCategoryDeleteAt
	ToApiResponsePaginationCategory(pbResponse *pb.ApiResponsePaginationCategory) *response.ApiResponsePaginationCategory
}

type CashierResponseMapper interface {
	ToApiResponseMonthlyTotalSales(pbResponse *pb.ApiResponseCashierMonthlyTotalSales) *response.ApiResponseCashierMonthlyTotalSales
	ToApiResponseYearlyTotalSales(pbResponse *pb.ApiResponseCashierYearlyTotalSales) *response.ApiResponseCashierYearlyTotalSales

	ToApiResponseCashierMonthlySale(pbResponse *pb.ApiResponseCashierMonthSales) *response.ApiResponseCashierMonthSales
	ToApiResponseCashierYearlySale(pbResponse *pb.ApiResponseCashierYearSales) *response.ApiResponseCashierYearSales

	ToApiResponseCashier(pbResponse *pb.ApiResponseCashier) *response.ApiResponseCashier
	ToApiResponsesCashier(pbResponse *pb.ApiResponsesCashier) *response.ApiResponsesCashier
	ToApiResponseCashierAll(pbResponse *pb.ApiResponseCashierAll) *response.ApiResponseCashierAll
	ToApiResponseCashierDelete(pbResponse *pb.ApiResponseCashierDelete) *response.ApiResponseCashierDelete
	ToApiResponseCashierDeleteAt(pbResponse *pb.ApiResponseCashierDeleteAt) *response.ApiResponseCashierDeleteAt
	ToApiResponsePaginationCashierDeleteAt(pbResponse *pb.ApiResponsePaginationCashierDeleteAt) *response.ApiResponsePaginationCashierDeleteAt
	ToApiResponsePaginationCashier(pbResponse *pb.ApiResponsePaginationCashier) *response.ApiResponsePaginationCashier
}

type MerchantResponseMapper interface {
	ToApiResponseMerchant(pbResponse *pb.ApiResponseMerchant) *response.ApiResponseMerchant

	ToApiResponseMerchantDeleteAt(pbResponse *pb.ApiResponseMerchantDeleteAt) *response.ApiResponseMerchantDeleteAt
	ToApiResponsesMerchant(pbResponse *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant
	ToApiResponseMerchantDelete(pbResponse *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete
	ToApiResponseMerchantAll(pbResponse *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll
	ToApiResponsePaginationMerchantDeleteAt(pbResponse *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt
	ToApiResponsePaginationMerchant(pbResponse *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant
}

type OrderItemResponseMapper interface {
	ToApiResponseOrderItem(pbResponse *pb.ApiResponseOrderItem) *response.ApiResponseOrderItem
	ToApiResponsesOrderItem(pbResponse *pb.ApiResponsesOrderItem) *response.ApiResponsesOrderItem
	ToApiResponseOrderItemDelete(pbResponse *pb.ApiResponseOrderItemDelete) *response.ApiResponseOrderItemDelete
	ToApiResponseOrderItemAll(pbResponse *pb.ApiResponseOrderItemAll) *response.ApiResponseOrderItemAll
	ToApiResponsePaginationOrderItemDeleteAt(pbResponse *pb.ApiResponsePaginationOrderItemDeleteAt) *response.ApiResponsePaginationOrderItemDeleteAt
	ToApiResponsePaginationOrderItem(pbResponse *pb.ApiResponsePaginationOrderItem) *response.ApiResponsePaginationOrderItem
}

type OrderResponseMapper interface {
	ToApiResponseMonthlyTotalRevenue(pbResponse *pb.ApiResponseOrderMonthlyTotalRevenue) *response.ApiResponseOrderMonthlyTotalRevenue
	ToApiResponseYearlyTotalRevenue(pbResponse *pb.ApiResponseOrderYearlyTotalRevenue) *response.ApiResponseOrderYearlyTotalRevenue

	ToApiResponseMonthlyOrder(pbResponse *pb.ApiResponseOrderMonthly) *response.ApiResponseOrderMonthly
	ToApiResponseYearlyOrder(pbResponse *pb.ApiResponseOrderYearly) *response.ApiResponseOrderYearly

	ToApiResponseOrder(pbResponse *pb.ApiResponseOrder) *response.ApiResponseOrder
	ToApiResponseOrderDeleteAt(pbResponse *pb.ApiResponseOrderDeleteAt) *response.ApiResponseOrderDeleteAt
	ToApiResponsesOrder(pbResponse *pb.ApiResponsesOrder) *response.ApiResponsesOrder
	ToApiResponseOrderDelete(pbResponse *pb.ApiResponseOrderDelete) *response.ApiResponseOrderDelete
	ToApiResponseOrderAll(pbResponse *pb.ApiResponseOrderAll) *response.ApiResponseOrderAll
	ToApiResponsePaginationOrderDeleteAt(pbResponse *pb.ApiResponsePaginationOrderDeleteAt) *response.ApiResponsePaginationOrderDeleteAt
	ToApiResponsePaginationOrder(pbResponse *pb.ApiResponsePaginationOrder) *response.ApiResponsePaginationOrder
}

type ProductResponseMapper interface {
	ToApiResponseProduct(pbResponse *pb.ApiResponseProduct) *response.ApiResponseProduct
	ToApiResponsesProductDeleteAt(pbResponse *pb.ApiResponseProductDeleteAt) *response.ApiResponseProductDeleteAt
	ToApiResponsesProduct(pbResponse *pb.ApiResponsesProduct) *response.ApiResponsesProduct
	ToApiResponseProductDelete(pbResponse *pb.ApiResponseProductDelete) *response.ApiResponseProductDelete
	ToApiResponseProductAll(pbResponse *pb.ApiResponseProductAll) *response.ApiResponseProductAll
	ToApiResponsePaginationProductDeleteAt(pbResponse *pb.ApiResponsePaginationProductDeleteAt) *response.ApiResponsePaginationProductDeleteAt
	ToApiResponsePaginationProduct(pbResponse *pb.ApiResponsePaginationProduct) *response.ApiResponsePaginationProduct
}

type TransactionResponseMapper interface {
	ToApiResponseTransactionMonthAmountSuccess(pbResponse *pb.ApiResponseTransactionMonthAmountSuccess) *response.ApiResponsesTransactionMonthSuccess
	ToApiResponseTransactionMonthAmountFailed(pbResponse *pb.ApiResponseTransactionMonthAmountFailed) *response.ApiResponsesTransactionMonthFailed
	ToApiResponseTransactionYearAmountSuccess(pbResponse *pb.ApiResponseTransactionYearAmountSuccess) *response.ApiResponsesTransactionYearSuccess
	ToApiResponseTransactionYearAmountFailed(pbResponse *pb.ApiResponseTransactionYearAmountFailed) *response.ApiResponsesTransactionYearFailed
	ToApiResponseTransactionMonthMethod(pbResponse *pb.ApiResponseTransactionMonthPaymentMethod) *response.ApiResponsesTransactionMonthMethod
	ToApiResponseTransactionYearMethod(pbResponse *pb.ApiResponseTransactionYearPaymentmethod) *response.ApiResponsesTransactionYearMethod

	ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction
	ToApiResponseTransactionDeleteAt(pbResponse *pb.ApiResponseTransactionDeleteAt) *response.ApiResponseTransactionDeleteAt
	ToApiResponsesTransaction(pbResponse *pb.ApiResponsesTransaction) *response.ApiResponsesTransaction
	ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete
	ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll
	ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt
	ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction
}
