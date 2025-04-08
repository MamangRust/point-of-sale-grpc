package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
	"time"
)

type transactionRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.TransactionRecordMapping
}

func NewTransactionRepository(db *db.Queries, ctx context.Context, mapping recordmapper.TransactionRecordMapping) *transactionRepository {
	return &transactionRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *transactionRepository) FindAllTransactions(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactions(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionsRecordPagination(res), totalCount, nil
}

func (r *transactionRepository) FindByActive(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionsActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionsRecordActivePagination(res), totalCount, nil
}

func (r *transactionRepository) FindByTrashed(search string, page, pageSize int) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionsTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionsTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionsRecordTrashedPagination(res), totalCount, nil
}

func (r *transactionRepository) FindByMerchant(
	merchant_id int,
	search string,
	page,
	pageSize int,
) ([]*record.TransactionRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetTransactionByMerchantParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetTransactionByMerchant(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find transactions: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToTransactionMerchantsRecordPagination(res), totalCount, nil
}

func (r *transactionRepository) GetMonthlyAmountSuccess(year int, month int) ([]*record.TransactionMonthlyAmountSuccessRecord, error) {
	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionSuccess(r.ctx, db.GetMonthlyAmountTransactionSuccessParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly successful transaction amounts: %w", err)
	}

	return r.mapping.ToTransactionMonthlyAmountSuccess(res), nil
}

func (r *transactionRepository) GetYearlyAmountSuccess(year int) ([]*record.TransactionYearlyAmountSuccessRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionSuccess(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly successful transaction amounts: %w", err)
	}

	return r.mapping.ToTransactionYearlyAmountSuccess(res), nil
}

func (r *transactionRepository) GetMonthlyAmountFailed(year int, month int) ([]*record.TransactionMonthlyAmountFailedRecord, error) {
	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionFailed(r.ctx, db.GetMonthlyAmountTransactionFailedParams{
		Column1: currentDate,
		Column2: lastDayCurrentMonth,
		Column3: prevDate,
		Column4: lastDayPrevMonth,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly failed transaction amounts: %w", err)
	}

	return r.mapping.ToTransactionMonthlyAmountFailed(res), nil
}

func (r *transactionRepository) GetYearlyAmountFailed(year int) ([]*record.TransactionYearlyAmountFailedRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionFailed(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly failed transaction amounts: %w", err)
	}

	return r.mapping.ToTransactionYearlyAmountFailed(res), nil
}

func (r *transactionRepository) GetMonthlyAmountSuccessByMerchant(year int, month int, merchantID int) ([]*record.TransactionMonthlyAmountSuccessRecord, error) {
	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionSuccessByMerchant(r.ctx, db.GetMonthlyAmountTransactionSuccessByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: int32(merchantID),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly successful transaction amounts for merchant %d: %w", merchantID, err)
	}

	return r.mapping.ToTransactionMonthlyAmountSuccessByMerchant(res), nil
}

func (r *transactionRepository) GetYearlyAmountSuccessByMerchant(year int, merchantID int) ([]*record.TransactionYearlyAmountSuccessRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionSuccessByMerchant(r.ctx, db.GetYearlyAmountTransactionSuccessByMerchantParams{
		Column1:    int32(year),
		MerchantID: int32(merchantID),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly successful transaction amounts for merchant %d: %w", merchantID, err)
	}

	return r.mapping.ToTransactionYearlyAmountSuccessByMerchant(res), nil
}

func (r *transactionRepository) GetMonthlyAmountFailedByMerchant(year int, month int, merchantID int) ([]*record.TransactionMonthlyAmountFailedRecord, error) {
	currentDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	prevDate := currentDate.AddDate(0, -1, 0)

	lastDayCurrentMonth := currentDate.AddDate(0, 1, -1)
	lastDayPrevMonth := prevDate.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyAmountTransactionFailedByMerchant(r.ctx, db.GetMonthlyAmountTransactionFailedByMerchantParams{
		Column1:    currentDate,
		Column2:    lastDayCurrentMonth,
		Column3:    prevDate,
		Column4:    lastDayPrevMonth,
		MerchantID: int32(merchantID),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly failed transaction amounts for merchant %d: %w", merchantID, err)
	}

	return r.mapping.ToTransactionMonthlyAmountFailedByMerchant(res), nil
}

func (r *transactionRepository) GetYearlyAmountFailedByMerchant(year int, merchantID int) ([]*record.TransactionYearlyAmountFailedRecord, error) {
	res, err := r.db.GetYearlyAmountTransactionFailedByMerchant(r.ctx, db.GetYearlyAmountTransactionFailedByMerchantParams{
		Column1:    int32(year),
		MerchantID: int32(merchantID),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly failed transaction amounts for merchant %d: %w", merchantID, err)
	}

	return r.mapping.ToTransactionYearlyAmountFailedByMerchant(res), nil
}

func (r *transactionRepository) GetMonthlyTransactionMethod(year int) ([]*record.TransactionMonthlyMethodRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err := r.db.GetMonthlyTransactionMethods(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly Transactions: %w", err)
	}

	return r.mapping.ToTransactionMonthlyMethod(res), nil
}

func (r *transactionRepository) GetYearlyTransactionMethod(year int) ([]*record.TransactionYearlyMethodRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyTransactionMethods(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly Transactions: %w", err)
	}

	return r.mapping.ToTransactionYearlyMethod(res), nil
}

func (r *transactionRepository) GetMonthlyTransactionMethodByMerchant(year int, merchant_id int) ([]*record.TransactionMonthlyMethodRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err := r.db.GetMonthlyTransactionMethodsByMerchant(r.ctx, db.GetMonthlyTransactionMethodsByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly Transactions: %w", err)
	}

	return r.mapping.ToTransactionMonthlyByMerchantMethod(res), nil
}

func (r *transactionRepository) GetYearlyTransactionMethodByMerchant(year int, merchant_id int) ([]*record.TransactionYearlyMethodRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyTransactionMethodsByMerchant(r.ctx, db.GetYearlyTransactionMethodsByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly Transactions: %w", err)
	}

	return r.mapping.ToTransactionYearlyMethodByMerchant(res), nil
}

func (r *transactionRepository) FindById(transaction_id int) (*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionByID(r.ctx, int32(transaction_id))

	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)

		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) FindByOrderId(order_id int) (*record.TransactionRecord, error) {
	res, err := r.db.GetTransactionByOrderID(r.ctx, int32(order_id))

	if err != nil {
		fmt.Printf("Error fetching user: %v\n", err)

		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) CreateTransaction(request *requests.CreateTransactionRequest) (*record.TransactionRecord, error) {
	req := db.CreateTransactionsParams{
		OrderID:       int32(request.OrderID),
		MerchantID:    int32(request.MerchantID),
		PaymentMethod: request.PaymentMethod,
		Amount:        int32(request.Amount),
		ChangeAmount: sql.NullInt32{
			Int32: int32(*request.ChangeAmount),
			Valid: true,
		},
		PaymentStatus: *request.PaymentStatus,
	}

	transaction, err := r.db.CreateTransactions(r.ctx, req)
	if err != nil {
		return nil, errors.New("failed to create transaction")
	}

	return r.mapping.ToTransactionRecord(transaction), nil
}

func (r *transactionRepository) UpdateTransaction(request *requests.UpdateTransactionRequest) (*record.TransactionRecord, error) {
	req := db.UpdateTransactionParams{
		TransactionID: int32(*request.TransactionID),
		MerchantID:    int32(request.MerchantID),
		PaymentMethod: request.PaymentMethod,
		Amount:        int32(request.Amount),
		ChangeAmount: sql.NullInt32{
			Int32: int32(*request.ChangeAmount),
			Valid: true,
		},
		OrderID:       int32(request.OrderID),
		PaymentStatus: *request.PaymentStatus,
	}

	res, err := r.db.UpdateTransaction(r.ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update transaction: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) TrashTransaction(transaction_id int) (*record.TransactionRecord, error) {
	res, err := r.db.TrashTransaction(r.ctx, int32(transaction_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash user: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) RestoreTransaction(transaction_id int) (*record.TransactionRecord, error) {
	res, err := r.db.RestoreTransaction(r.ctx, int32(transaction_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore topup: %w", err)
	}

	return r.mapping.ToTransactionRecord(res), nil
}

func (r *transactionRepository) DeleteTransactionPermanently(transaction_id int) (bool, error) {
	err := r.db.DeleteTransactionPermanently(r.ctx, int32(transaction_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete transactions: %w", err)
	}

	return true, nil
}

func (r *transactionRepository) RestoreAllTransactions() (bool, error) {
	err := r.db.RestoreAllTransactions(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all transactions: %w", err)
	}
	return true, nil
}

func (r *transactionRepository) DeleteAllTransactionPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentTransactions(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all transactions permanently: %w", err)
	}
	return true, nil
}
