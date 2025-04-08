package recordmapper

import (
	"pointofsale/internal/domain/record"
	db "pointofsale/pkg/database/schema"
)

type transactionRecordMapper struct {
}

func NewTransactionRecordMapper() *transactionRecordMapper {
	return &transactionRecordMapper{}
}

func (s *transactionRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord {
	var deletedAt *string
	if transaction.DeletedAt.Valid {
		deletedAtStr := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		OrderID:       int(transaction.OrderID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int(transaction.Amount),
		MerchantID:    int(transaction.MerchantID),
		ChangeAmount:  int(transaction.ChangeAmount.Int32),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, transaction := range transactions {
		result = append(result, s.ToTransactionRecord(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionMerchantRecordPagination(transaction *db.GetTransactionByMerchantRow) *record.TransactionRecord {
	var deletedAt *string
	if transaction.DeletedAt.Valid {
		deletedAtStr := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		OrderID:       int(transaction.OrderID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int(transaction.Amount),
		MerchantID:    int(transaction.MerchantID),
		ChangeAmount:  int(transaction.ChangeAmount.Int32),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionMerchantsRecordPagination(products []*db.GetTransactionByMerchantRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, product := range products {
		result = append(result, s.ToTransactionMerchantRecordPagination(product))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordPagination(transaction *db.GetTransactionsRow) *record.TransactionRecord {
	var deletedAt *string
	if transaction.DeletedAt.Valid {
		deletedAtStr := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		OrderID:       int(transaction.OrderID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int(transaction.Amount),
		ChangeAmount:  int(transaction.ChangeAmount.Int32),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordPagination(products []*db.GetTransactionsRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, product := range products {
		result = append(result, s.ToTransactionRecordPagination(product))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordActivePagination(transaction *db.GetTransactionsActiveRow) *record.TransactionRecord {
	var deletedAt *string
	if transaction.DeletedAt.Valid {
		deletedAtStr := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		OrderID:       int(transaction.OrderID),
		MerchantID:    int(transaction.MerchantID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int(transaction.Amount),
		ChangeAmount:  int(transaction.ChangeAmount.Int32),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordActivePagination(transactions []*db.GetTransactionsActiveRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, transaction := range transactions {
		result = append(result, s.ToTransactionRecordActivePagination(transaction))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionRecordTrashedPagination(transaction *db.GetTransactionsTrashedRow) *record.TransactionRecord {
	var deletedAt *string
	if transaction.DeletedAt.Valid {
		deletedAtStr := transaction.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.TransactionRecord{
		ID:            int(transaction.TransactionID),
		OrderID:       int(transaction.OrderID),
		MerchantID:    int(transaction.MerchantID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int(transaction.Amount),
		ChangeAmount:  int(transaction.ChangeAmount.Int32),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:     transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:     deletedAt,
	}
}

func (s *transactionRecordMapper) ToTransactionsRecordTrashedPagination(products []*db.GetTransactionsTrashedRow) []*record.TransactionRecord {
	var result []*record.TransactionRecord

	for _, product := range products {
		result = append(result, s.ToTransactionRecordTrashedPagination(product))
	}

	return result
}

func (s *transactionRecordMapper) ToTransactionMonthAmountSuccess(row *db.GetMonthlyAmountTransactionSuccessRow) *record.TransactionMonthlyAmountSuccessRecord {
	return &record.TransactionMonthlyAmountSuccessRecord{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: int(row.TotalSuccess),
		TotalAmount:  int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionMonthlyAmountSuccess(rows []*db.GetMonthlyAmountTransactionSuccessRow) []*record.TransactionMonthlyAmountSuccessRecord {
	var transaction []*record.TransactionMonthlyAmountSuccessRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionMonthAmountSuccess(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionYearAmountSuccess(row *db.GetYearlyAmountTransactionSuccessRow) *record.TransactionYearlyAmountSuccessRecord {
	return &record.TransactionYearlyAmountSuccessRecord{
		Year:         row.Year,
		TotalSuccess: int(row.TotalSuccess),
		TotalAmount:  int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionYearlyAmountSuccess(rows []*db.GetYearlyAmountTransactionSuccessRow) []*record.TransactionYearlyAmountSuccessRecord {
	var transaction []*record.TransactionYearlyAmountSuccessRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionYearAmountSuccess(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionMonthAmountFailed(row *db.GetMonthlyAmountTransactionFailedRow) *record.TransactionMonthlyAmountFailedRecord {
	return &record.TransactionMonthlyAmountFailedRecord{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: int(row.TotalFailed),
		TotalAmount: int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionMonthlyAmountFailed(rows []*db.GetMonthlyAmountTransactionFailedRow) []*record.TransactionMonthlyAmountFailedRecord {
	var transaction []*record.TransactionMonthlyAmountFailedRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionMonthAmountFailed(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionYearAmountFailed(row *db.GetYearlyAmountTransactionFailedRow) *record.TransactionYearlyAmountFailedRecord {
	return &record.TransactionYearlyAmountFailedRecord{
		Year:        row.Year,
		TotalFailed: int(row.TotalFailed),
		TotalAmount: int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionYearlyAmountFailed(rows []*db.GetYearlyAmountTransactionFailedRow) []*record.TransactionYearlyAmountFailedRecord {
	var transaction []*record.TransactionYearlyAmountFailedRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionYearAmountFailed(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionMonthAmountSuccessByMerchant(row *db.GetMonthlyAmountTransactionSuccessByMerchantRow) *record.TransactionMonthlyAmountSuccessRecord {
	return &record.TransactionMonthlyAmountSuccessRecord{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: int(row.TotalSuccess),
		TotalAmount:  int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionMonthlyAmountSuccessByMerchant(rows []*db.GetMonthlyAmountTransactionSuccessByMerchantRow) []*record.TransactionMonthlyAmountSuccessRecord {
	var transaction []*record.TransactionMonthlyAmountSuccessRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionMonthAmountSuccessByMerchant(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionYearAmountSuccessByMerchant(row *db.GetYearlyAmountTransactionSuccessByMerchantRow) *record.TransactionYearlyAmountSuccessRecord {
	return &record.TransactionYearlyAmountSuccessRecord{
		Year:         row.Year,
		TotalSuccess: int(row.TotalSuccess),
		TotalAmount:  int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionYearlyAmountSuccessByMerchant(rows []*db.GetYearlyAmountTransactionSuccessByMerchantRow) []*record.TransactionYearlyAmountSuccessRecord {
	var transaction []*record.TransactionYearlyAmountSuccessRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionYearAmountSuccessByMerchant(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionMonthAmountFailedByMerchant(row *db.GetMonthlyAmountTransactionFailedByMerchantRow) *record.TransactionMonthlyAmountFailedRecord {
	return &record.TransactionMonthlyAmountFailedRecord{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: int(row.TotalFailed),
		TotalAmount: int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionMonthlyAmountFailedByMerchant(rows []*db.GetMonthlyAmountTransactionFailedByMerchantRow) []*record.TransactionMonthlyAmountFailedRecord {
	var transaction []*record.TransactionMonthlyAmountFailedRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionMonthAmountFailedByMerchant(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionYearAmountFailedByMerchant(row *db.GetYearlyAmountTransactionFailedByMerchantRow) *record.TransactionYearlyAmountFailedRecord {
	return &record.TransactionYearlyAmountFailedRecord{
		Year:        row.Year,
		TotalFailed: int(row.TotalFailed),
		TotalAmount: int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionYearlyAmountFailedByMerchant(rows []*db.GetYearlyAmountTransactionFailedByMerchantRow) []*record.TransactionYearlyAmountFailedRecord {
	var transaction []*record.TransactionYearlyAmountFailedRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionYearAmountFailedByMerchant(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionMonthMethod(row *db.GetMonthlyTransactionMethodsRow) *record.TransactionMonthlyMethodRecord {
	return &record.TransactionMonthlyMethodRecord{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int(row.TotalTransactions),
		TotalAmount:       int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionMonthlyMethod(rows []*db.GetMonthlyTransactionMethodsRow) []*record.TransactionMonthlyMethodRecord {
	var transaction []*record.TransactionMonthlyMethodRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionMonthMethod(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionYearMethod(row *db.GetYearlyTransactionMethodsRow) *record.TransactionYearlyMethodRecord {
	return &record.TransactionYearlyMethodRecord{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int(row.TotalTransactions),
		TotalAmount:       int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionYearlyMethod(rows []*db.GetYearlyTransactionMethodsRow) []*record.TransactionYearlyMethodRecord {
	var transaction []*record.TransactionYearlyMethodRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionYearMethod(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionMonthMethodByMerchant(row *db.GetMonthlyTransactionMethodsByMerchantRow) *record.TransactionMonthlyMethodRecord {
	return &record.TransactionMonthlyMethodRecord{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int(row.TotalTransactions),
		TotalAmount:       int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionMonthlyByMerchantMethod(rows []*db.GetMonthlyTransactionMethodsByMerchantRow) []*record.TransactionMonthlyMethodRecord {
	var transaction []*record.TransactionMonthlyMethodRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionMonthMethodByMerchant(row))
	}

	return transaction
}

func (s *transactionRecordMapper) ToTransactionYearMethodByMerchant(row *db.GetYearlyTransactionMethodsByMerchantRow) *record.TransactionYearlyMethodRecord {
	return &record.TransactionYearlyMethodRecord{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int(row.TotalTransactions),
		TotalAmount:       int(row.TotalAmount),
	}
}

func (s *transactionRecordMapper) ToTransactionYearlyMethodByMerchant(rows []*db.GetYearlyTransactionMethodsByMerchantRow) []*record.TransactionYearlyMethodRecord {
	var transaction []*record.TransactionYearlyMethodRecord

	for _, row := range rows {
		transaction = append(transaction, s.ToTransactionYearMethodByMerchant(row))
	}

	return transaction
}
