package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/transaction_errors"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type transactionHandleGrpc struct {
	pb.UnimplementedTransactionServiceServer
	transactionService service.TransactionService
}

func NewTransactionHandleGrpc(
	transactionService service.TransactionService,
) *transactionHandleGrpc {
	return &transactionHandleGrpc{
		transactionService: transactionService,
	}
}

func (s *transactionHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransaction, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransaction{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	transactions, totalRecords, err := s.transactionService.FindAllTransactions(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var transactionResponses []*pb.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, &pb.TransactionResponse{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationTransaction{
		Status:     "success",
		Message:    "Successfully fetched transactions",
		Data:       transactionResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *transactionHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindAllTransactionMerchantRequest) (*pb.ApiResponsePaginationTransaction, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	merchant_id := int(request.GetMerchantId())

	if merchant_id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransactionByMerchant{
		Search:     search,
		Page:       page,
		PageSize:   pageSize,
		MerchantID: merchant_id,
	}

	transactions, totalRecords, err := s.transactionService.FindByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var transactionResponses []*pb.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, &pb.TransactionResponse{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationTransaction{
		Status:     "success",
		Message:    "Successfully fetched transactions by merchant",
		Data:       transactionResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthStatusSuccess(ctx context.Context, request *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthAmountTransaction{
		Year:  year,
		Month: month,
	}

	monthlyAmounts, err := s.transactionService.FindMonthlyAmountSuccess(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyAmountResponses []*pb.TransactionMonthlyAmountSuccess
	for _, amount := range monthlyAmounts {
		monthlyAmountResponses = append(monthlyAmountResponses, &pb.TransactionMonthlyAmountSuccess{
			Year:         amount.Year,
			Month:        amount.Month,
			TotalSuccess: int32(amount.TotalSuccess),
			TotalAmount:  int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  "success",
		Message: "Successfully fetched monthly transaction amounts",
		Data:    monthlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearStatusSuccess(ctx context.Context, request *pb.FindYearlyTransactionStatus) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	year := int(request.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	yearlyAmounts, err := s.transactionService.FindYearlyAmountSuccess(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyAmountResponses []*pb.TransactionYearlyAmountSuccess
	for _, amount := range yearlyAmounts {
		yearlyAmountResponses = append(yearlyAmountResponses, &pb.TransactionYearlyAmountSuccess{
			Year:         amount.Year,
			TotalSuccess: int32(amount.TotalSuccess),
			TotalAmount:  int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  "success",
		Message: "Successfully fetched yearly transaction amounts",
		Data:    yearlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthStatusFailed(ctx context.Context, request *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthAmountTransaction{
		Year:  year,
		Month: month,
	}

	monthlyAmounts, err := s.transactionService.FindMonthlyAmountFailed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyAmountResponses []*pb.TransactionMonthlyAmountFailed
	for _, amount := range monthlyAmounts {
		monthlyAmountResponses = append(monthlyAmountResponses, &pb.TransactionMonthlyAmountFailed{
			Year:        amount.Year,
			Month:       amount.Month,
			TotalFailed: int32(amount.TotalFailed),
			TotalAmount: int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  "success",
		Message: "Successfully fetched monthly failed transaction amounts",
		Data:    monthlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearStatusFailed(ctx context.Context, request *pb.FindYearlyTransactionStatus) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	year := int(request.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	yearlyAmounts, err := s.transactionService.FindYearlyAmountFailed(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyAmountResponses []*pb.TransactionYearlyAmountFailed
	for _, amount := range yearlyAmounts {
		yearlyAmountResponses = append(yearlyAmountResponses, &pb.TransactionYearlyAmountFailed{
			Year:        amount.Year,
			TotalFailed: int32(amount.TotalFailed),
			TotalAmount: int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  "success",
		Message: "Successfully fetched yearly failed transaction amounts",
		Data:    yearlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthStatusSuccessByMerchant(ctx context.Context, request *pb.FindMonthlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.MonthAmountTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	monthlyAmounts, err := s.transactionService.FindMonthlyAmountSuccessByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyAmountResponses []*pb.TransactionMonthlyAmountSuccess
	for _, amount := range monthlyAmounts {
		monthlyAmountResponses = append(monthlyAmountResponses, &pb.TransactionMonthlyAmountSuccess{
			Year:         amount.Year,
			Month:        amount.Month,
			TotalSuccess: int32(amount.TotalSuccess),
			TotalAmount:  int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthAmountSuccess{
		Status:  "success",
		Message: "Successfully fetched merchant monthly transaction amounts",
		Data:    monthlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearStatusSuccessByMerchant(ctx context.Context, request *pb.FindYearlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearAmountTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	yearlyAmounts, err := s.transactionService.FindYearlyAmountSuccessByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyAmountResponses []*pb.TransactionYearlyAmountSuccess
	for _, amount := range yearlyAmounts {
		yearlyAmountResponses = append(yearlyAmountResponses, &pb.TransactionYearlyAmountSuccess{
			Year:         amount.Year,
			TotalSuccess: int32(amount.TotalSuccess),
			TotalAmount:  int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearAmountSuccess{
		Status:  "success",
		Message: "Successfully fetched merchant yearly transaction amounts",
		Data:    yearlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthStatusFailedByMerchant(ctx context.Context, request *pb.FindMonthlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.MonthAmountTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	monthlyAmounts, err := s.transactionService.FindMonthlyAmountFailedByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var monthlyAmountResponses []*pb.TransactionMonthlyAmountFailed
	for _, amount := range monthlyAmounts {
		monthlyAmountResponses = append(monthlyAmountResponses, &pb.TransactionMonthlyAmountFailed{
			Year:        amount.Year,
			Month:       amount.Month,
			TotalFailed: int32(amount.TotalFailed),
			TotalAmount: int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthAmountFailed{
		Status:  "success",
		Message: "Successfully fetched merchant monthly failed transaction amounts",
		Data:    monthlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearStatusFailedByMerchant(ctx context.Context, request *pb.FindYearlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearAmountTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	yearlyAmounts, err := s.transactionService.FindYearlyAmountFailedByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var yearlyAmountResponses []*pb.TransactionYearlyAmountFailed
	for _, amount := range yearlyAmounts {
		yearlyAmountResponses = append(yearlyAmountResponses, &pb.TransactionYearlyAmountFailed{
			Year:        amount.Year,
			TotalFailed: int32(amount.TotalFailed),
			TotalAmount: int32(amount.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearAmountFailed{
		Status:  "success",
		Message: "Successfully fetched merchant yearly failed transaction amounts",
		Data:    yearlyAmountResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthMethodSuccess(ctx context.Context, req *pb.MonthTransactionMethod) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthMethodTransaction{
		Year:  year,
		Month: month,
	}

	methods, err := s.transactionService.FindMonthlyMethodSuccess(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionMonthlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionMonthlyMethod{
			Month:             method.Month,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched monthly payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearMethodSuccess(ctx context.Context, req *pb.YearTransactionMethod) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	methods, err := s.transactionService.FindYearlyMethodSuccess(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionYearlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionYearlyMethod{
			Year:              method.Year,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "Successfully fetched yearly payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthMethodByMerchantSuccess(ctx context.Context, req *pb.MonthTransactionMethodByMerchant) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthMethodTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	methods, err := s.transactionService.FindMonthlyMethodByMerchantSuccess(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionMonthlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionMonthlyMethod{
			Month:             method.Month,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched merchant monthly payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearMethodByMerchantSuccess(ctx context.Context, req *pb.YearTransactionMethodByMerchant) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearMethodTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.transactionService.FindYearlyMethodByMerchantSuccess(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionYearlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionYearlyMethod{
			Year:              method.Year,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "Successfully fetched merchant yearly payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthMethodFailed(ctx context.Context, req *pb.MonthTransactionMethod) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	reqService := requests.MonthMethodTransaction{
		Year:  year,
		Month: month,
	}

	methods, err := s.transactionService.FindMonthlyMethodFailed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionMonthlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionMonthlyMethod{
			Month:             method.Month,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched monthly failed payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearMethodFailed(ctx context.Context, req *pb.YearTransactionMethod) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	methods, err := s.transactionService.FindYearlyMethodFailed(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionYearlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionYearlyMethod{
			Year:              method.Year,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "Successfully fetched yearly failed payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindMonthMethodByMerchantFailed(ctx context.Context, req *pb.MonthTransactionMethodByMerchant) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, transaction_errors.ErrGrpcInvalidMonth
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.MonthMethodTransactionMerchant{
		Year:       year,
		Month:      month,
		MerchantID: id,
	}

	methods, err := s.transactionService.FindMonthlyMethodByMerchantFailed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionMonthlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionMonthlyMethod{
			Month:             method.Month,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionMonthPaymentMethod{
		Status:  "success",
		Message: "Successfully fetched merchant monthly failed payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindYearMethodByMerchantFailed(ctx context.Context, req *pb.YearTransactionMethodByMerchant) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidYear
	}

	if id <= 0 {
		return nil, transaction_errors.ErrGrpcInvalidMerchantId
	}

	reqService := requests.YearMethodTransactionMerchant{
		Year:       year,
		MerchantID: id,
	}

	methods, err := s.transactionService.FindYearlyMethodByMerchantFailed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var methodResponses []*pb.TransactionYearlyMethod
	for _, method := range methods {
		methodResponses = append(methodResponses, &pb.TransactionYearlyMethod{
			Year:              method.Year,
			PaymentMethod:     method.PaymentMethod,
			TotalTransactions: int32(method.TotalTransactions),
			TotalAmount:       int32(method.TotalAmount),
		})
	}

	return &pb.ApiResponseTransactionYearPaymentmethod{
		Status:  "success",
		Message: "Successfully fetched merchant yearly failed payment methods",
		Data:    methodResponses,
	}, nil
}

func (s *transactionHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcInvalidID
	}

	transaction, err := s.transactionService.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully fetched transaction",
		Data: &pb.TransactionResponse{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *transactionHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransaction{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	transactions, totalRecords, err := s.transactionService.FindByActive(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var transactionResponses []*pb.TransactionResponseDeleteAt
	for _, transaction := range transactions {
		var deletedAt string
		if transaction.DeletedAt.Valid {
			deletedAt = transaction.DeletedAt.Time.String()
		}

		transactionResponses = append(transactionResponses, &pb.TransactionResponseDeleteAt{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active transactions",
		Data:       transactionResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *transactionHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllTransactionRequest) (*pb.ApiResponsePaginationTransactionDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllTransaction{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	transactions, totalRecords, err := s.transactionService.FindByTrashed(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	var transactionResponses []*pb.TransactionResponseDeleteAt
	for _, transaction := range transactions {
		var deletedAt string
		if transaction.DeletedAt.Valid {
			deletedAt = transaction.DeletedAt.Time.String()
		}

		transactionResponses = append(transactionResponses, &pb.TransactionResponseDeleteAt{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed transactions",
		Data:       transactionResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *transactionHandleGrpc) Create(ctx context.Context, request *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	req := &requests.CreateTransactionRequest{
		OrderID:       int(request.GetOrderId()),
		CashierID:     int(request.GetCashierId()),
		PaymentMethod: request.GetPaymentMethod(),
		Amount:        int(request.GetAmount()),
	}

	if err := req.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateCreateTransaction
	}

	transaction, err := s.transactionService.CreateTransaction(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data: &pb.TransactionResponse{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *transactionHandleGrpc) UpdateTransaction(ctx context.Context, request *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcInvalidID
	}

	req := &requests.UpdateTransactionRequest{
		TransactionID: &id,
		OrderID:       int(request.GetOrderId()),
		CashierID:     int(request.GetCashierId()),
		PaymentMethod: request.GetPaymentMethod(),
		Amount:        int(request.GetAmount()),
	}

	if err := req.Validate(); err != nil {
		return nil, transaction_errors.ErrGrpcValidateUpdateTransaction
	}

	transaction, err := s.transactionService.UpdateTransaction(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransaction{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data: &pb.TransactionResponse{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *transactionHandleGrpc) TrashedTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcInvalidID
	}

	transaction, err := s.transactionService.TrashedTransaction(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if transaction.DeletedAt.Valid {
		deletedAt = transaction.DeletedAt.Time.String()
	}

	return &pb.ApiResponseTransactionDeleteAt{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data: &pb.TransactionResponseDeleteAt{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		},
	}, nil
}

func (s *transactionHandleGrpc) RestoreTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcInvalidID
	}

	transaction, err := s.transactionService.RestoreTransaction(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if transaction.DeletedAt.Valid {
		deletedAt = transaction.DeletedAt.Time.String()
	}

	return &pb.ApiResponseTransactionDeleteAt{
		Status:  "success",
		Message: "Successfully restored transaction",
		Data: &pb.TransactionResponseDeleteAt{
			Id:            int32(transaction.TransactionID),
			OrderId:       int32(transaction.OrderID),
			MerchantId:    int32(transaction.MerchantID),
			PaymentMethod: transaction.PaymentMethod,
			Amount:        int32(transaction.Amount),
			ChangeAmount:  int32(*transaction.ChangeAmount),
			PaymentStatus: transaction.PaymentStatus,
			CreatedAt:     transaction.CreatedAt.Time.String(),
			UpdatedAt:     transaction.UpdatedAt.Time.String(),
			DeletedAt:     &wrapperspb.StringValue{Value: deletedAt},
		},
	}, nil
}

func (s *transactionHandleGrpc) DeleteTransactionPermanent(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, transaction_errors.ErrGrpcInvalidID
	}

	_, err := s.transactionService.DeleteTransactionPermanently(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionDelete{
		Status:  "success",
		Message: "Successfully deleted transaction permanently",
	}, nil
}

func (s *transactionHandleGrpc) RestoreAllTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.RestoreAllTransactions(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "success",
		Message: "Successfully restored all transactions",
	}, nil
}

func (s *transactionHandleGrpc) DeleteAllTransactionPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.DeleteAllTransactionPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseTransactionAll{
		Status:  "Success",
		Message: "Successfully deleted all transactions permanently",
	}, nil
}
