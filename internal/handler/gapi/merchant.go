package gapi

import (
	"context"
	"math"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/pb"
	"pointofsale/internal/service"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/errors/merchant_errors"

	"google.golang.org/protobuf/types/known/emptypb"
)

type merchantHandleGrpc struct {
	pb.UnimplementedMerchantServiceServer
	merchantService service.MerchantService
}

func NewMerchantHandleGrpc(
	merchantService service.MerchantService,
) *merchantHandleGrpc {
	return &merchantHandleGrpc{
		merchantService: merchantService,
	}
}

func (s *merchantHandleGrpc) FindAll(ctx context.Context, request *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchant, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	merchants, totalRecords, err := s.merchantService.FindAllMerchants(ctx, &reqService)
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

	var merchantResponses []*pb.MerchantResponse
	for _, merchant := range merchants {
		merchantResponses = append(merchantResponses, &pb.MerchantResponse{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
		})
	}

	return &pb.ApiResponsePaginationMerchant{
		Status:     "success",
		Message:    "Successfully fetched merchants",
		Data:       merchantResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *merchantHandleGrpc) FindById(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchant, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, merchant_errors.ErrGrpcInvalidID
	}

	merchant, err := s.merchantService.FindById(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully fetched merchant",
		Data: &pb.MerchantResponse{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *merchantHandleGrpc) FindByActive(ctx context.Context, request *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	merchants, totalRecords, err := s.merchantService.FindByActive(ctx, &reqService)
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

	var merchantResponses []*pb.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		var deletedAt string
		if merchant.DeletedAt.Valid {
			deletedAt = merchant.DeletedAt.Time.String()
		}

		merchantResponses = append(merchantResponses, &pb.MerchantResponseDeleteAt{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		})
	}

	return &pb.ApiResponsePaginationMerchantDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched active merchants",
		Data:       merchantResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *merchantHandleGrpc) FindByTrashed(ctx context.Context, request *pb.FindAllMerchantRequest) (*pb.ApiResponsePaginationMerchantDeleteAt, error) {
	page := int(request.GetPage())
	pageSize := int(request.GetPageSize())
	search := request.GetSearch()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	reqService := requests.FindAllMerchants{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
	}

	merchants, totalRecords, err := s.merchantService.FindByTrashed(ctx, &reqService)
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

	var merchantResponses []*pb.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		var deletedAt string
		if merchant.DeletedAt.Valid {
			deletedAt = merchant.DeletedAt.Time.String()
		}

		merchantResponses = append(merchantResponses, &pb.MerchantResponseDeleteAt{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		})
	}

	return &pb.ApiResponsePaginationMerchantDeleteAt{
		Status:     "success",
		Message:    "Successfully fetched trashed merchants",
		Data:       merchantResponses,
		Pagination: paginationMeta,
	}, nil
}

func (s *merchantHandleGrpc) Create(ctx context.Context, request *pb.CreateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	req := &requests.CreateMerchantRequest{
		UserID:       int(request.GetUserId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Address:      request.GetAddress(),
		ContactEmail: request.GetContactEmail(),
		ContactPhone: request.GetContactPhone(),
		Status:       request.GetStatus(),
	}

	if err := req.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateCreateMerchant
	}

	merchant, err := s.merchantService.CreateMerchant(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully created merchant",
		Data: &pb.MerchantResponse{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *merchantHandleGrpc) Update(ctx context.Context, request *pb.UpdateMerchantRequest) (*pb.ApiResponseMerchant, error) {
	id := int(request.GetMerchantId())

	if id == 0 {
		return nil, merchant_errors.ErrGrpcInvalidID
	}

	req := &requests.UpdateMerchantRequest{
		MerchantID:   &id,
		UserID:       int(request.GetUserId()),
		Name:         request.GetName(),
		Description:  request.GetDescription(),
		Address:      request.GetAddress(),
		ContactEmail: request.GetContactEmail(),
		ContactPhone: request.GetContactPhone(),
		Status:       request.GetStatus(),
	}

	if err := req.Validate(); err != nil {
		return nil, merchant_errors.ErrGrpcValidateUpdateMerchant
	}

	merchant, err := s.merchantService.UpdateMerchant(ctx, req)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseMerchant{
		Status:  "success",
		Message: "Successfully updated merchant",
		Data: &pb.MerchantResponse{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
		},
	}, nil
}

func (s *merchantHandleGrpc) TrashedMerchant(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchantDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, merchant_errors.ErrGrpcInvalidID
	}

	merchant, err := s.merchantService.TrashedMerchant(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if merchant.DeletedAt.Valid {
		deletedAt = merchant.DeletedAt.Time.String()
	}

	return &pb.ApiResponseMerchantDeleteAt{
		Status:  "success",
		Message: "Successfully trashed merchant",
		Data: &pb.MerchantResponseDeleteAt{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		},
	}, nil
}

func (s *merchantHandleGrpc) RestoreMerchant(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchantDeleteAt, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, merchant_errors.ErrGrpcInvalidID
	}

	merchant, err := s.merchantService.RestoreMerchant(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	var deletedAt string
	if merchant.DeletedAt.Valid {
		deletedAt = merchant.DeletedAt.Time.String()
	}

	return &pb.ApiResponseMerchantDeleteAt{
		Status:  "success",
		Message: "Successfully restored merchant",
		Data: &pb.MerchantResponseDeleteAt{
			Id:           int32(merchant.MerchantID),
			UserId:       int32(merchant.UserID),
			Name:         merchant.Name,
			Description:  *merchant.Description,
			Address:      *merchant.Address,
			ContactEmail: *merchant.ContactEmail,
			ContactPhone: *merchant.ContactPhone,
			Status:       merchant.Status,
			CreatedAt:    merchant.CreatedAt.Time.String(),
			UpdatedAt:    merchant.UpdatedAt.Time.String(),
			DeletedAt:    deletedAt,
		},
	}, nil
}

func (s *merchantHandleGrpc) DeleteMerchantPermanent(ctx context.Context, request *pb.FindByIdMerchantRequest) (*pb.ApiResponseMerchantDelete, error) {
	id := int(request.GetId())

	if id == 0 {
		return nil, merchant_errors.ErrGrpcInvalidID
	}

	_, err := s.merchantService.DeleteMerchantPermanent(ctx, id)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseMerchantDelete{
		Status:  "success",
		Message: "Successfully deleted merchant permanently",
	}, nil
}

func (s *merchantHandleGrpc) RestoreAllMerchant(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseMerchantAll, error) {
	_, err := s.merchantService.RestoreAllMerchant(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseMerchantAll{
		Status:  "success",
		Message: "Successfully restore all merchants",
	}, nil
}

func (s *merchantHandleGrpc) DeleteAllMerchantPermanent(ctx context.Context, _ *emptypb.Empty) (*pb.ApiResponseMerchantAll, error) {
	_, err := s.merchantService.DeleteAllMerchantPermanent(ctx)
	if err != nil {
		return nil, errors.ToGrpcError(err)
	}

	return &pb.ApiResponseMerchantAll{
		Status:  "success",
		Message: "Successfully delete merchant permanent",
	}, nil
}
