package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/cashier_errors"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type cashierHandleGrpc struct {
	pb.UnimplementedCashierServiceServer
	cashierService service.CashierService
}

func NewCashierHandleGrpc(
	cashierService service.CashierService,
) *cashierHandleGrpc {
	return &cashierHandleGrpc{
		cashierService: cashierService,
	}
}

func (s *cashierHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllCashierRequest) (*pb.ApiResponsePaginationCashier, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCashiers{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	cashiers, totalRecords, err := s.cashierService.FindAllCashiers(ctx, &reqService)
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

	var cashierResponses []*pb.CashierResponse
	for _, cashier := range cashiers {
		cashierResponses = append(cashierResponses, &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationCashier{
		Status:     "success",
		Message:    "Successfully fetched cashier",
		Data:       cashierResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *cashierHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashier, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	cashier, err := s.cashierService.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCashier{
		Status:  "success",
		Message: "Successfully fetched cashier",
		Data: &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *cashierHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllCashierRequest) (*pb.ApiResponsePaginationCashierDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCashiers{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	cashiers, totalRecords, err := s.cashierService.FindByActive(ctx, &reqService)
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

	var cashierResponses []*pb.CashierResponseDeleteAt
	for _, cashier := range cashiers {
		var deletedAt string
		if cashier.DeletedAt.Valid {
			deletedAt = cashier.DeletedAt.Time.String()
		}

		cashierResponses = append(cashierResponses, &pb.CashierResponseDeleteAt{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationCashierDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active cashier",
		Data:       cashierResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *cashierHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllCashierRequest) (*pb.ApiResponsePaginationCashierDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCashiers{
		Search:   search,
		Page:     page,
		PageSize: pageSize,
	}

	cashiers, totalRecords, err := s.cashierService.FindByTrashed(ctx, &reqService)
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

	var cashierResponses []*pb.CashierResponseDeleteAt
	for _, cashier := range cashiers {
		var deletedAt string
		if cashier.DeletedAt.Valid {
			deletedAt = cashier.DeletedAt.Time.String()
		}

		cashierResponses = append(cashierResponses, &pb.CashierResponseDeleteAt{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: deletedAt},
		})
	}

	return &pb.ApiResponsePaginationCashierDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed cashier",
		Data:       cashierResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *cashierHandleGrpc) FindByMerchant(ctx context.Context, request *pb.FindByMerchantCashierRequest) (*pb.ApiResponsePaginationCashier, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()
	merchant_id := int(request.GetMerchantId())

	if merchant_id <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllCashierMerchant{
		Search:     search,
		Page:       page,
		PageSize:   pageSize,
		MerchantID: merchant_id,
	}

	cashiers, totalRecords, err := s.cashierService.FindByMerchant(ctx, &reqService)
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

	var cashierResponses []*pb.CashierResponse
	for _, cashier := range cashiers {
		cashierResponses = append(cashierResponses, &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationCashier{
		Status:     "success",
		Message:    "Successfully fetched cashier",
		Data:       cashierResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *cashierHandleGrpc) FindMonthlyTotalSales(ctx context.Context, req *pb.FindYearMonthTotalSales) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMonth
	}

	reqService := requests.MonthTotalSales{
		Year:  year,
		Month: month,
	}

	sales, err := s.cashierService.FindMonthlyTotalSales(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseMonthTotalSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseMonthTotalSales{
			Year:       sale.Year,
			Month:      sale.Month,
			TotalSales: int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierMonthlyTotalSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindYearlyTotalSales(ctx context.Context, req *pb.FindYearTotalSales) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	sales, err := s.cashierService.FindYearlyTotalSales(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseYearTotalSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseYearTotalSales{
			Year:       sale.Year,
			TotalSales: int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierYearlyTotalSales{
		Status:  "success",
		Message: "Yearly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindMonthlyTotalSalesById(ctx context.Context, req *pb.FindYearMonthTotalSalesById) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	id := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMonth
	}

	if id <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthTotalSalesCashier{
		Year:      year,
		Month:     month,
		CashierID: id,
	}

	sales, err := s.cashierService.FindMonthlyTotalSalesById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseMonthTotalSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseMonthTotalSales{
			Year:       sale.Year,
			Month:      sale.Month,
			TotalSales: int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierMonthlyTotalSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindYearlyTotalSalesById(ctx context.Context, req *pb.FindYearTotalSalesById) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	year := int(req.GetYear())
	id := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if id <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearTotalSalesCashier{
		Year:      year,
		CashierID: id,
	}

	sales, err := s.cashierService.FindYearlyTotalSalesById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseYearTotalSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseYearTotalSales{
			Year:       sale.Year,
			TotalSales: int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierYearlyTotalSales{
		Status:  "success",
		Message: "Yearly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindMonthlyTotalSalesByMerchant(ctx context.Context, req *pb.FindYearMonthTotalSalesByMerchant) (*pb.ApiResponseCashierMonthlyTotalSales, error) {
	year := int(req.GetYear())
	month := int(req.GetMonth())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if month <= 0 || month >= 12 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMonth
	}

	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthTotalSalesMerchant{
		Year:       year,
		Month:      month,
		MerchantID: merchantId,
	}

	sales, err := s.cashierService.FindMonthlyTotalSalesByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseMonthTotalSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseMonthTotalSales{
			Year:       sale.Year,
			Month:      sale.Month,
			TotalSales: int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierMonthlyTotalSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindYearlyTotalSalesByMerchant(ctx context.Context, req *pb.FindYearTotalSalesByMerchant) (*pb.ApiResponseCashierYearlyTotalSales, error) {
	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.YearTotalSalesMerchant{
		Year:       year,
		MerchantID: merchantId,
	}

	sales, err := s.cashierService.FindYearlyTotalSalesByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseYearTotalSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseYearTotalSales{
			Year:       sale.Year,
			TotalSales: int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierYearlyTotalSales{
		Status:  "success",
		Message: "Yearly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindMonthSales(ctx context.Context, req *pb.FindYearCashier) (*pb.ApiResponseCashierMonthSales, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	sales, err := s.cashierService.FindMonthyCashier(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseMonthSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseMonthSales{
			Month:       sale.Month,
			CashierId:   int32(sale.CashierID),
			CashierName: sale.CashierName,
			OrderCount:  int32(sale.OrderCount),
			TotalSales:  int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierMonthSales{
		Status:  "success",
		Message: "Monthly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindYearSales(ctx context.Context, req *pb.FindYearCashier) (*pb.ApiResponseCashierYearSales, error) {
	year := int(req.GetYear())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	sales, err := s.cashierService.FindYearlyCashier(ctx, year)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseYearSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseYearSales{
			Year:        sale.Year,
			CashierId:   int32(sale.CashierID),
			CashierName: sale.CashierName,
			OrderCount:  int32(sale.OrderCount),
			TotalSales:  int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierYearSales{
		Status:  "success",
		Message: "Yearly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindMonthSalesByMerchant(ctx context.Context, req *pb.FindYearCashierByMerchant) (*pb.ApiResponseCashierMonthSales, error) {
	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.MonthCashierMerchant{
		Year:       year,
		MerchantID: merchantId,
	}

	sales, err := s.cashierService.FindMonthlyCashierByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseMonthSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseMonthSales{
			Month:       sale.Month,
			CashierId:   int32(sale.CashierID),
			CashierName: sale.CashierName,
			OrderCount:  int32(sale.OrderCount),
			TotalSales:  int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierMonthSales{
		Status:  "success",
		Message: "Merchant monthly revenue retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindYearSalesByMerchant(ctx context.Context, req *pb.FindYearCashierByMerchant) (*pb.ApiResponseCashierYearSales, error) {
	year := int(req.GetYear())
	merchantId := int(req.GetMerchantId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if merchantId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidMerchantId
	}

	reqService := requests.YearCashierMerchant{
		Year:       year,
		MerchantID: merchantId,
	}

	sales, err := s.cashierService.FindYearlyCashierByMerchant(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseYearSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseYearSales{
			Year:        sale.Year,
			CashierId:   int32(sale.CashierID),
			CashierName: sale.CashierName,
			OrderCount:  int32(sale.OrderCount),
			TotalSales:  int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierYearSales{
		Status:  "success",
		Message: "Merchant yearly payment methods retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindMonthSalesById(ctx context.Context, req *pb.FindYearCashierById) (*pb.ApiResponseCashierMonthSales, error) {
	year := int(req.GetYear())
	cashierId := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if cashierId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.MonthCashierId{
		Year:      year,
		CashierID: cashierId,
	}

	sales, err := s.cashierService.FindMonthlyCashierById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseMonthSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseMonthSales{
			Month:       sale.Month,
			CashierId:   int32(sale.CashierID),
			CashierName: sale.CashierName,
			OrderCount:  int32(sale.OrderCount),
			TotalSales:  int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierMonthSales{
		Status:  "success",
		Message: "Cashier monthly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) FindYearSalesById(ctx context.Context, req *pb.FindYearCashierById) (*pb.ApiResponseCashierYearSales, error) {
	year := int(req.GetYear())
	cashierId := int(req.GetCashierId())

	if year <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidYear
	}

	if cashierId <= 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	reqService := requests.YearCashierId{
		Year:      year,
		CashierID: cashierId,
	}

	sales, err := s.cashierService.FindYearlyCashierById(ctx, &reqService)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var salesResponses []*pb.CashierResponseYearSales
	for _, sale := range sales {
		salesResponses = append(salesResponses, &pb.CashierResponseYearSales{
			Year:        sale.Year,
			CashierId:   int32(sale.CashierID),
			CashierName: sale.CashierName,
			OrderCount:  int32(sale.OrderCount),
			TotalSales:  int32(sale.TotalSales),
		})
	}

	return &pb.ApiResponseCashierYearSales{
		Status:  "success",
		Message: "Cashier yearly sales retrieved successfully",
		Data:    salesResponses,
	}, nil
}

func (s *cashierHandleGrpc) CreateCashier(ctx context.Context, request *pb.CreateCashierRequest) (*pb.ApiResponseCashier, error) {
	req := &requests.CreateCashierRequest{
		Name:       request.GetName(),
		MerchantID: int(request.GetMerchantId()),
		UserID:     int(request.GetUserId()),
	}

	if err := req.Validate(); err != nil {
		return nil, cashier_errors.ErrGrpcValidateCreateCashier
	}

	cashier, err := s.cashierService.CreateCashier(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCashier{
		Status:  "success",
		Message: "Successfully created cashier",
		Data: &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *cashierHandleGrpc) UpdateCashier(ctx context.Context, request *pb.UpdateCashierRequest) (*pb.ApiResponseCashier, error) {
	id := int(request.GetCashierId())

	if id == 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	req := &requests.UpdateCashierRequest{
		CashierID: &id,
		Name:      request.GetName(),
	}

	if err := req.Validate(); err != nil {
		return nil, cashier_errors.ErrGrpcValidateUpdateCashier
	}

	cashier, err := s.cashierService.UpdateCashier(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCashier{
		Status:  "success",
		Message: "Successfully updated cashier",
		Data: &pb.CashierResponse{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *cashierHandleGrpc) TrashedCashier(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	cashier, err := s.cashierService.TrashedCashier(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if cashier.DeletedAt.Valid {
		deletedAt = cashier.DeletedAt.Time.String()
	}

	return &pb.ApiResponseCashierDeleteAt{
		Status:  "success",
		Message: "Successfully trashed cashier",
		Data: &pb.CashierResponseDeleteAt{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: deletedAt},
		},
	}, nil
}

func (s *cashierHandleGrpc) RestoreCashier(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	cashier, err := s.cashierService.RestoreCashier(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if cashier.DeletedAt.Valid {
		deletedAt = cashier.DeletedAt.Time.String()
	}

	return &pb.ApiResponseCashierDeleteAt{
		Status:  "success",
		Message: "Successfully restored cashier",
		Data: &pb.CashierResponseDeleteAt{
			Id:         int32(cashier.CashierID),
			MerchantId: int32(cashier.MerchantID),
			Name:       cashier.Name,
			CreatedAt:  cashier.CreatedAt.Time.String(),
			UpdatedAt:  cashier.UpdatedAt.Time.String(),
			DeletedAt:  &wrapperspb.StringValue{Value: deletedAt},
		},
	}, nil
}

func (s *cashierHandleGrpc) DeleteCashierPermanent(ctx context.Context, request *pb.FindByIdCashierRequest) (*pb.ApiResponseCashierDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, cashier_errors.ErrGrpcFailedInvalidId
	}

	_, err := s.cashierService.DeleteCashierPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCashierDelete{
		Status:  "success",
		Message: "Successfully deleted cashier permanently",
	}, nil
}

func (s *cashierHandleGrpc) RestoreAllCashier(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCashierAll, error) {
	_, err := s.cashierService.RestoreAllCashier(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCashierAll{
		Status:  "success",
		Message: "Successfully restore all cashier",
	}, nil
}

func (s *cashierHandleGrpc) DeleteAllCashierPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseCashierAll, error) {
	_, err := s.cashierService.DeleteAllCashierPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseCashierAll{
		Status:  "success",
		Message: "Successfully delete cashier permanen",
	}, nil
}
