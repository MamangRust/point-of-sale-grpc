package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	protomapper "pointofsale/internal/mapper/proto"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors_custom"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type transactionHandleGrpc struct {
	pb.UnimplementedTransactionServiceServer
	transactionService service.TransactionService
	mapping            protomapper.TransactionProtoMapper
}

func NewTransactionHandleGrpc(
	transactionService service.TransactionService,
	mapping protomapper.TransactionProtoMapper,
) *transactionHandleGrpc {
	return &transactionHandleGrpc{
		transactionService: transactionService,
		mapping:            mapping,
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

	transaction, totalRecords, err := s.transactionService.FindAllTransactions(search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationTransaction(paginationMeta, "success", "Successfully fetched transaction", transaction)
	return so, nil
}

func (s *transactionHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindAllTransactionMerchantRequest) (*pb.ApiResponsePaginationTransaction, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	merchant_id := int(request.GetMerchantId())

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	transaction, totalRecords, err := s.transactionService.FindByMerchant(merchant_id, search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(*totalRecords),
	}

	so := s.mapping.ToProtoResponsePaginationTransaction(paginationMeta, "success", "Successfully fetched transaction", transaction)
	return so, nil
}

func (s *transactionHandleGrpc) FindMonthStatusSuccess(ctx context.Context, request *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindMonthlyAmountSuccess(year, month)
	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthAmountSuccess("success", "Monthly success data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindYearStatusSuccess(ctx context.Context, request *pb.FindYearlyTransactionStatus) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	year := int(request.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindYearlyAmountSuccess(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearAmountSuccess("success", "Yearly success data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindMonthStatusFailed(ctx context.Context, request *pb.FindMonthlyTransactionStatus) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindMonthlyAmountFailed(year, month)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthAmountFailed("success", "Monthly failed data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindYearStatusFailed(ctx context.Context, request *pb.FindYearlyTransactionStatus) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	year := int(request.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindYearlyAmountFailed(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearAmountFailed("success", "Yearly failed data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindMonthStatusSuccessByMerchant(ctx context.Context, request *pb.FindMonthlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionMonthAmountSuccess, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindMonthlyAmountSuccessByMerchant(
		year, month, id,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthAmountSuccess("success", "Merchant monthly success data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindYearStatusSuccessByMerchant(ctx context.Context, request *pb.FindYearlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionYearAmountSuccess, error) {
	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindYearlyAmountSuccessByMerchant(
		year, id,
	)
	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  "operation_failed",
			Message: "Could not retrieve merchant yearly success data",
		})
	}

	return s.mapping.ToProtoResponseYearAmountSuccess("success", "Merchant yearly success data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindMonthStatusFailedByMerchant(ctx context.Context, request *pb.FindMonthlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionMonthAmountFailed, error) {
	year := int(request.GetYear())
	month := int(request.GetMonth())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if month <= 0 || month >= 12 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_month",
				Message: "Month must be between 1 and 12",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindMonthlyAmountFailedByMerchant(
		year, month, id,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthAmountFailed("success", "Merchant monthly failed data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindYearStatusFailedByMerchant(ctx context.Context, request *pb.FindYearlyTransactionStatusByMerchant) (*pb.ApiResponseTransactionYearAmountFailed, error) {
	year := int(request.GetYear())
	id := int(request.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	res, err := s.transactionService.FindYearlyAmountFailedByMerchant(
		year, id,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearAmountFailed("success", "Merchant yearly failed data retrieved successfully", res), nil
}

func (s *transactionHandleGrpc) FindMonthMethod(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.transactionService.FindMonthlyMethod(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthMethod("success", "Monthly payment methods retrieved successfully", methods), nil
}

func (s *transactionHandleGrpc) FindYearMethod(ctx context.Context, req *pb.FindYearTransaction) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.transactionService.FindYearlyMethod(year)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearMethod("success", "Yearly payment methods retrieved successfully", methods), nil
}

func (s *transactionHandleGrpc) FindMonthMethodByMerchant(ctx context.Context, req *pb.FindYearTransactionByMerchant) (*pb.ApiResponseTransactionMonthPaymentMethod, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.transactionService.FindMonthlyMethodByMerchant(
		year, id,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseMonthMethod("success", "Merchant monthly payment methods retrieved successfully", methods), nil
}

func (s *transactionHandleGrpc) FindYearMethodByMerchant(ctx context.Context, req *pb.FindYearTransactionByMerchant) (*pb.ApiResponseTransactionYearPaymentmethod, error) {
	year := int(req.GetYear())
	id := int(req.GetMerchantId())

	if year <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid year parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	if id <= 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id parameter",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	methods, err := s.transactionService.FindYearlyMethodByMerchant(
		year, id,
	)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	return s.mapping.ToProtoResponseYearMethod("success", "Merchant yearly payment methods retrieved successfully", methods), nil
}

func (s *transactionHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetId())

	if request.GetId() == 0 {
		return nil, status.Errorf(
			codes.Code(codes.InvalidArgument),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "invalid_input",
				Message: "Invalid id",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	transaction, err := s.transactionService.FindById(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransaction("success", "Successfully fetched transaction", transaction)

	return so, nil

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

	transaction, totalRecords, err := s.transactionService.FindByActive(search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(0),
	}
	so := s.mapping.ToProtoResponsePaginationTransactionDeleteAt(paginationMeta, "success", "Successfully fetched active transaction", transaction)

	return so, nil
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

	transaction, totalRecords, err := s.transactionService.FindByTrashed(search, page, pageSize)

	if err != nil {
		return nil, status.Errorf(codes.Code(err.Code), "%v", &pb.ErrorResponse{
			Status:  err.Status,
			Message: err.Message,
		})
	}

	totalPages := int(math.Ceil(float64(*totalRecords) / float64(pageSize)))

	paginationMeta := &pb.PaginationMeta{
		CurrentPage:  int32(page),
		PageSize:     int32(pageSize),
		TotalPages:   int32(totalPages),
		TotalRecords: int32(0),
	}

	so := s.mapping.ToProtoResponsePaginationTransactionDeleteAt(paginationMeta, "success", "Successfully fetched trashed transaction", transaction)

	return so, nil
}

func (s *transactionHandleGrpc) Create(ctx context.Context, request *pb.CreateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	req := &requests.CreateTransactionRequest{
		OrderID:       int(request.GetOrderId()),
		CashierID:     int(request.GetCashierId()),
		PaymentMethod: request.GetPaymentMethod(),
		Amount:        int(request.GetAmount()),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Unable to create new transaction. Please check your input.",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	transaction, err := s.transactionService.CreateTransaction(req)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransaction("success", "Successfully created transaction", transaction)
	return so, nil
}

func (s *transactionHandleGrpc) Update(ctx context.Context, request *pb.UpdateTransactionRequest) (*pb.ApiResponseTransaction, error) {
	id := int(request.GetTransactionId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Transaction ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	req := &requests.UpdateTransactionRequest{
		TransactionID: &id,
		OrderID:       int(request.GetOrderId()),
		CashierID:     int(request.GetCashierId()),
		PaymentMethod: request.GetPaymentMethod(),
		Amount:        int(request.GetAmount()),
	}

	if err := req.Validate(); err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Unable to process transaction update. Please review your data.",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	transaction, err := s.transactionService.UpdateTransaction(req)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransaction("success", "Successfully updated transaction", transaction)
	return so, nil
}

func (s *transactionHandleGrpc) TrashedTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Transaction ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	transaction, err := s.transactionService.TrashedTransaction(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransactionDeleteAt("success", "Successfully trashed transaction", transaction)

	return so, nil
}

func (s *transactionHandleGrpc) RestoreTransaction(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Transaction ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	transaction, err := s.transactionService.RestoreTransaction(id)

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransactionDeleteAt("success", "Successfully restored transaction", transaction)

	return so, nil
}

func (s *transactionHandleGrpc) DeleteTransactionPermanent(ctx context.Context, request *pb.FindByIdTransactionRequest) (*pb.ApiResponseTransactionDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  "validation_error",
				Message: "Transaction ID parameter cannot be empty and must be a positive number",
				Code:    int32(codes.InvalidArgument),
			}),
		)
	}

	_, err := s.transactionService.DeleteTransactionPermanently(int(request.GetId()))

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransactionDelete("success", "Successfully deleted Transaction permanently")

	return so, nil
}

func (s *transactionHandleGrpc) RestoreAllTransaction(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.RestoreAllTransactions()

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransactionAll("success", "Successfully restore all Transaction")

	return so, nil
}

func (s *transactionHandleGrpc) DeleteAllTransactionPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseTransactionAll, error) {
	_, err := s.transactionService.DeleteAllTransactionPermanent()

	if err != nil {
		return nil, status.Errorf(
			codes.Code(err.Code),
			errors_custom.GrpcErrorToJson(&pb.ErrorResponse{
				Status:  err.Status,
				Message: err.Message,
				Code:    int32(err.Code),
			}),
		)
	}

	so := s.mapping.ToProtoResponseTransactionAll("success", "Successfully delete Transaction permanen")

	return so, nil
}
