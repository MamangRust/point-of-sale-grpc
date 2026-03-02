package transaction_cache

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
)

type TransactionStatsCache interface {
	GetCachedMonthAmountSuccessCached(ctx context.Context, req *requests.MonthAmountTransaction) (*response.ApiResponsesTransactionMonthSuccess, bool)
	SetCachedMonthAmountSuccessCached(ctx context.Context, req *requests.MonthAmountTransaction, res *response.ApiResponsesTransactionMonthSuccess)

	GetCachedYearAmountSuccessCached(ctx context.Context, year int) (*response.ApiResponsesTransactionYearSuccess, bool)
	SetCachedYearAmountSuccessCached(ctx context.Context, year int, res *response.ApiResponsesTransactionYearSuccess)

	GetCachedMonthAmountFailedCached(ctx context.Context, req *requests.MonthAmountTransaction) (*response.ApiResponsesTransactionMonthFailed, bool)
	SetCachedMonthAmountFailedCached(ctx context.Context, req *requests.MonthAmountTransaction, res *response.ApiResponsesTransactionMonthFailed)

	GetCachedYearAmountFailedCached(ctx context.Context, year int) (*response.ApiResponsesTransactionYearFailed, bool)
	SetCachedYearAmountFailedCached(ctx context.Context, year int, res *response.ApiResponsesTransactionYearFailed)

	GetCachedMonthMethodSuccessCached(ctx context.Context, req *requests.MonthMethodTransaction) (*response.ApiResponsesTransactionMonthMethod, bool)
	SetCachedMonthMethodSuccessCached(ctx context.Context, req *requests.MonthMethodTransaction, res *response.ApiResponsesTransactionMonthMethod)

	GetCachedYearMethodSuccessCached(ctx context.Context, year int) (*response.ApiResponsesTransactionYearMethod, bool)
	SetCachedYearMethodSuccessCached(ctx context.Context, year int, res *response.ApiResponsesTransactionYearMethod)

	GetCachedMonthMethodFailedCached(ctx context.Context, req *requests.MonthMethodTransaction) (*response.ApiResponsesTransactionMonthMethod, bool)
	SetCachedMonthMethodFailedCached(ctx context.Context, req *requests.MonthMethodTransaction, res *response.ApiResponsesTransactionMonthMethod)

	GetCachedYearMethodFailedCached(ctx context.Context, year int) (*response.ApiResponsesTransactionYearMethod, bool)
	SetCachedYearMethodFailedCached(ctx context.Context, year int, res *response.ApiResponsesTransactionYearMethod)
}

type TransactionStatsByMerchantCache interface {
	GetCachedMonthAmountSuccessByMerchantCached(ctx context.Context, req *requests.MonthAmountTransactionMerchant) (*response.ApiResponsesTransactionMonthSuccess, bool)
	SetCachedMonthAmountSuccessByMerchantCached(ctx context.Context, req *requests.MonthAmountTransactionMerchant, res *response.ApiResponsesTransactionMonthSuccess)

	GetCachedYearAmountSuccessByMerchantCached(ctx context.Context, req *requests.YearAmountTransactionMerchant) (*response.ApiResponsesTransactionYearSuccess, bool)
	SetCachedYearAmountSuccessByMerchantCached(ctx context.Context, req *requests.YearAmountTransactionMerchant, res *response.ApiResponsesTransactionYearSuccess)

	GetCachedMonthAmountFailedByMerchantCached(ctx context.Context, req *requests.MonthAmountTransactionMerchant) (*response.ApiResponsesTransactionMonthFailed, bool)
	SetCachedMonthAmountFailedByMerchantCached(ctx context.Context, req *requests.MonthAmountTransactionMerchant, res *response.ApiResponsesTransactionMonthFailed)

	GetCachedYearAmountFailedByMerchantCached(ctx context.Context, req *requests.YearAmountTransactionMerchant) (*response.ApiResponsesTransactionYearFailed, bool)
	SetCachedYearAmountFailedByMerchantCached(ctx context.Context, req *requests.YearAmountTransactionMerchant, res *response.ApiResponsesTransactionYearFailed)

	GetCachedMonthMethodSuccessByMerchantCached(ctx context.Context, req *requests.MonthMethodTransactionMerchant) (*response.ApiResponsesTransactionMonthMethod, bool)
	SetCachedMonthMethodSuccessByMerchantCached(ctx context.Context, req *requests.MonthMethodTransactionMerchant, res *response.ApiResponsesTransactionMonthMethod)

	GetCachedYearMethodSuccessByMerchantCached(ctx context.Context, req *requests.YearMethodTransactionMerchant) (*response.ApiResponsesTransactionYearMethod, bool)
	SetCachedYearMethodSuccessByMerchantCached(ctx context.Context, req *requests.YearMethodTransactionMerchant, res *response.ApiResponsesTransactionYearMethod)

	GetCachedMonthMethodFailedByMerchantCached(ctx context.Context, req *requests.MonthMethodTransactionMerchant) (*response.ApiResponsesTransactionMonthMethod, bool)
	SetCachedMonthMethodFailedByMerchantCached(ctx context.Context, req *requests.MonthMethodTransactionMerchant, res *response.ApiResponsesTransactionMonthMethod)

	GetCachedYearMethodFailedByMerchantCached(ctx context.Context, req *requests.YearMethodTransactionMerchant) (*response.ApiResponsesTransactionYearMethod, bool)
	SetCachedYearMethodFailedByMerchantCached(ctx context.Context, req *requests.YearMethodTransactionMerchant, res *response.ApiResponsesTransactionYearMethod)
}

type TransactionQueryCache interface {
	GetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransaction) (*response.ApiResponsePaginationTransaction, bool)
	SetCachedTransactionsCache(ctx context.Context, req *requests.FindAllTransaction, res *response.ApiResponsePaginationTransaction)

	GetCachedTransactionByMerchant(ctx context.Context, req *requests.FindAllTransactionByMerchant) (*response.ApiResponsePaginationTransaction, bool)
	SetCachedTransactionByMerchant(ctx context.Context, req *requests.FindAllTransactionByMerchant, res *response.ApiResponsePaginationTransaction)

	GetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransaction) (*response.ApiResponsePaginationTransactionDeleteAt, bool)
	SetCachedTransactionActiveCache(ctx context.Context, req *requests.FindAllTransaction, res *response.ApiResponsePaginationTransactionDeleteAt)

	GetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransaction) (*response.ApiResponsePaginationTransactionDeleteAt, bool)
	SetCachedTransactionTrashedCache(ctx context.Context, req *requests.FindAllTransaction, res *response.ApiResponsePaginationTransactionDeleteAt)

	GetCachedTransactionCache(ctx context.Context, id int) (*response.ApiResponseTransaction, bool)
	SetCachedTransactionCache(ctx context.Context, res *response.ApiResponseTransaction)

	GetCachedTransactionByOrderId(ctx context.Context, orderID int) (*response.ApiResponseTransaction, bool)
	SetCachedTransactionByOrderId(ctx context.Context, orderID int, res *response.ApiResponseTransaction)
}

type TransactionCommandCache interface {
	DeleteTransactionCache(ctx context.Context, transactionID int)
}
