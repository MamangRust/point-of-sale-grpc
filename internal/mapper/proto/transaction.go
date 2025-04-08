package protomapper

import (
	"pointofsale/internal/domain/response"
	"pointofsale/internal/pb"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type transactionProtoMapper struct{}

func NewTransactionProtoMapper() *transactionProtoMapper {
	return &transactionProtoMapper{}
}

func (t *transactionProtoMapper) ToProtoResponseTransaction(status string, message string, trans *response.TransactionResponse) *pb.ApiResponseTransaction {
	return &pb.ApiResponseTransaction{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTransaction(trans),
	}
}

func (t *transactionProtoMapper) ToProtoResponsesTransaction(status string, message string, transList []*response.TransactionResponse) *pb.ApiResponsesTransaction {
	return &pb.ApiResponsesTransaction{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransaction(transList),
	}
}

func (t *transactionProtoMapper) ToProtoResponseTransactionDeleteAt(status string, message string, trans *response.TransactionResponseDeleteAt) *pb.ApiResponseTransactionDeleteAt {
	return &pb.ApiResponseTransactionDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTransactionDeleteAt(trans),
	}
}

func (t *transactionProtoMapper) ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete {
	return &pb.ApiResponseTransactionDelete{
		Status:  status,
		Message: message,
	}
}

func (t *transactionProtoMapper) ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll {
	return &pb.ApiResponseTransactionAll{
		Status:  status,
		Message: message,
	}
}

func (t *transactionProtoMapper) ToProtoResponsePaginationTransactionDeleteAt(pagination *pb.PaginationMeta, status string, message string, transactions []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt {
	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTransactionDeleteAt(transactions),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *transactionProtoMapper) ToProtoResponsePaginationTransaction(pagination *pb.PaginationMeta, status string, message string, transactions []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction {
	return &pb.ApiResponsePaginationTransaction{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTransaction(transactions),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (t *transactionProtoMapper) ToProtoResponseMonthAmountSuccess(status string, message string, row []*response.TransactionMonthlyAmountSuccessResponse) *pb.ApiResponseTransactionMonthAmountSuccess {
	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthlyAmountSuccess(row),
	}
}

func (t *transactionProtoMapper) ToProtoResponseYearAmountSuccess(status string, message string, row []*response.TransactionYearlyAmountSuccessResponse) *pb.ApiResponseTransactionYearAmountSuccess {
	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearlyAmountSuccess(row),
	}
}

func (t *transactionProtoMapper) ToProtoResponseMonthAmountFailed(status string, message string, row []*response.TransactionMonthlyAmountFailedResponse) *pb.ApiResponseTransactionMonthAmountFailed {
	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthlyAmountFailed(row),
	}
}

func (t *transactionProtoMapper) ToProtoResponseYearAmountFailed(status string, message string, row []*response.TransactionYearlyAmountFailedResponse) *pb.ApiResponseTransactionYearAmountFailed {
	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearlyAmountFailed(row),
	}
}

func (t *transactionProtoMapper) ToProtoResponseMonthMethod(status string, message string, row []*response.TransactionMonthlyMethodResponse) *pb.ApiResponseTransactionMonthPaymentMethod {
	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionMonthlyMethod(row),
	}
}

func (t *transactionProtoMapper) ToProtoResponseYearMethod(status string, message string, row []*response.TransactionYearlyMethodResponse) *pb.ApiResponseTransactionYearPaymentmethod {
	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTransactionYearlyMethod(row),
	}
}

func (t *transactionProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Id:            int32(transaction.ID),
		OrderId:       int32(transaction.OrderID),
		MerchantId:    int32(transaction.MerchantID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int32(transaction.Amount),
		ChangeAmount:  int32(transaction.ChangeAmount),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}
}

func (t *transactionProtoMapper) mapResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse {
	var mappedTransactions []*pb.TransactionResponse

	for _, transaction := range transactions {
		mappedTransactions = append(mappedTransactions, t.mapResponseTransaction(transaction))
	}

	return mappedTransactions
}

func (t *transactionProtoMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if transaction.DeletedAt != nil {
		deletedAt = wrapperspb.String(*transaction.DeletedAt)
	}

	return &pb.TransactionResponseDeleteAt{
		Id:            int32(transaction.ID),
		OrderId:       int32(transaction.OrderID),
		MerchantId:    int32(transaction.MerchantID),
		PaymentMethod: transaction.PaymentMethod,
		Amount:        int32(transaction.Amount),
		ChangeAmount:  int32(transaction.ChangeAmount),
		PaymentStatus: transaction.PaymentStatus,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
		DeletedAt:     deletedAt,
	}
}

func (t *transactionProtoMapper) mapResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt {
	var mappedTransactions []*pb.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		mappedTransactions = append(mappedTransactions, t.mapResponseTransactionDeleteAt(transaction))
	}

	return mappedTransactions
}

func (s *transactionProtoMapper) mapResponseTransactionMonthAmountSuccess(row *response.TransactionMonthlyAmountSuccessResponse) *pb.TransactionMonthlyAmountSuccess {
	return &pb.TransactionMonthlyAmountSuccess{
		Year:         row.Year,
		Month:        row.Month,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionMonthlyAmountSuccess(rows []*response.TransactionMonthlyAmountSuccessResponse) []*pb.TransactionMonthlyAmountSuccess {
	var transaction []*pb.TransactionMonthlyAmountSuccess

	for _, row := range rows {
		transaction = append(transaction, s.mapResponseTransactionMonthAmountSuccess(row))
	}

	return transaction
}

func (s *transactionProtoMapper) mapResponseTransactionYearAmountSuccess(row *response.TransactionYearlyAmountSuccessResponse) *pb.TransactionYearlyAmountSuccess {
	return &pb.TransactionYearlyAmountSuccess{
		Year:         row.Year,
		TotalSuccess: int32(row.TotalSuccess),
		TotalAmount:  int32(row.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionYearlyAmountSuccess(rows []*response.TransactionYearlyAmountSuccessResponse) []*pb.TransactionYearlyAmountSuccess {
	var transaction []*pb.TransactionYearlyAmountSuccess

	for _, row := range rows {
		transaction = append(transaction, s.mapResponseTransactionYearAmountSuccess(row))
	}

	return transaction
}

func (s *transactionProtoMapper) mapResponseTransactionMonthAmountFailed(row *response.TransactionMonthlyAmountFailedResponse) *pb.TransactionMonthlyAmountFailed {
	return &pb.TransactionMonthlyAmountFailed{
		Year:        row.Year,
		Month:       row.Month,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionMonthlyAmountFailed(rows []*response.TransactionMonthlyAmountFailedResponse) []*pb.TransactionMonthlyAmountFailed {
	var transaction []*pb.TransactionMonthlyAmountFailed

	for _, row := range rows {
		transaction = append(transaction, s.mapResponseTransactionMonthAmountFailed(row))
	}

	return transaction
}

func (s *transactionProtoMapper) mapResponseTransactionYearAmountFailed(row *response.TransactionYearlyAmountFailedResponse) *pb.TransactionYearlyAmountFailed {
	return &pb.TransactionYearlyAmountFailed{
		Year:        row.Year,
		TotalFailed: int32(row.TotalFailed),
		TotalAmount: int32(row.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionYearlyAmountFailed(rows []*response.TransactionYearlyAmountFailedResponse) []*pb.TransactionYearlyAmountFailed {
	var transaction []*pb.TransactionYearlyAmountFailed

	for _, row := range rows {
		transaction = append(transaction, s.mapResponseTransactionYearAmountFailed(row))
	}

	return transaction
}

func (s *transactionProtoMapper) mapResponseTransactionMonthMethod(row *response.TransactionMonthlyMethodResponse) *pb.TransactionMonthlyMethod {
	return &pb.TransactionMonthlyMethod{
		Month:             row.Month,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionMonthlyMethod(rows []*response.TransactionMonthlyMethodResponse) []*pb.TransactionMonthlyMethod {
	var transaction []*pb.TransactionMonthlyMethod

	for _, row := range rows {
		transaction = append(transaction, s.mapResponseTransactionMonthMethod(row))
	}

	return transaction
}

func (s *transactionProtoMapper) mapResponseTransactionYearMethod(row *response.TransactionYearlyMethodResponse) *pb.TransactionYearlyMethod {
	return &pb.TransactionYearlyMethod{
		Year:              row.Year,
		PaymentMethod:     row.PaymentMethod,
		TotalTransactions: int32(row.TotalTransactions),
		TotalAmount:       int32(row.TotalAmount),
	}
}

func (s *transactionProtoMapper) mapResponsesTransactionYearlyMethod(rows []*response.TransactionYearlyMethodResponse) []*pb.TransactionYearlyMethod {
	var transaction []*pb.TransactionYearlyMethod

	for _, row := range rows {
		transaction = append(transaction, s.mapResponseTransactionYearMethod(row))
	}

	return transaction
}
