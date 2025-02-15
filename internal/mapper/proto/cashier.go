package protomapper

import (
	"pointofsale/internal/domain/response"
	"pointofsale/internal/pb"
)

type cashierProtoMapper struct {
}

func NewCashierProtoMapper() *cashierProtoMapper {
	return &cashierProtoMapper{}
}

func (c *cashierProtoMapper) ToProtoResponseCashier(status string, message string, pbResponse *response.CashierResponse) *pb.ApiResponseCashier {
	return &pb.ApiResponseCashier{
		Status:  status,
		Message: message,
		Data:    c.mapResponseCashier(pbResponse),
	}
}

func (c *cashierProtoMapper) ToProtoResponsesCashier(status string, message string, pbResponse []*response.CashierResponse) *pb.ApiResponsesCashier {
	return &pb.ApiResponsesCashier{
		Status:  status,
		Message: message,
		Data:    c.mapResponsesCashier(pbResponse),
	}
}

func (c *cashierProtoMapper) ToProtoResponseCashierDeleteAt(status string, message string, pbResponse *response.CashierResponseDeleteAt) *pb.ApiResponseCashierDeleteAt {
	return &pb.ApiResponseCashierDeleteAt{
		Status:  status,
		Message: message,
		Data:    c.mapResponseCashierDeleteAt(pbResponse),
	}
}

func (c *cashierProtoMapper) ToProtoResponseCashierDelete(status string, message string) *pb.ApiResponseCashierDelete {
	return &pb.ApiResponseCashierDelete{
		Status:  status,
		Message: message,
	}
}

func (u *cashierProtoMapper) ToProtoResponseCashierAll(status string, message string) *pb.ApiResponseCashierAll {
	return &pb.ApiResponseCashierAll{
		Status:  status,
		Message: message,
	}
}

func (u *cashierProtoMapper) ToProtoResponsePaginationCashierDeleteAt(pagination *pb.PaginationMeta, status string, message string, users []*response.CashierResponseDeleteAt) *pb.ApiResponsePaginationCashierDeleteAt {
	return &pb.ApiResponsePaginationCashierDeleteAt{
		Status:     status,
		Message:    message,
		Data:       u.mapResponsesCashierDeleteAt(users),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (u *cashierProtoMapper) ToProtoResponsePaginationCashier(pagination *pb.PaginationMeta, status string, message string, users []*response.CashierResponse) *pb.ApiResponsePaginationCashier {
	return &pb.ApiResponsePaginationCashier{
		Status:     status,
		Message:    message,
		Data:       u.mapResponsesCashier(users),
		Pagination: mapPaginationMeta(pagination),
	}
}

func (c *cashierProtoMapper) mapResponseCashier(cashier *response.CashierResponse) *pb.CashierResponse {
	return &pb.CashierResponse{
		Id:         int32(cashier.ID),
		MerchantId: int32(cashier.MerchantID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt,
		UpdatedAt:  cashier.UpdatedAt,
	}
}

func (c *cashierProtoMapper) mapResponsesCashier(cashiers []*response.CashierResponse) []*pb.CashierResponse {
	var mappedCashiers []*pb.CashierResponse

	for _, cashier := range cashiers {
		mappedCashiers = append(mappedCashiers, c.mapResponseCashier(cashier))
	}

	return mappedCashiers
}

func (c *cashierProtoMapper) mapResponseCashierDeleteAt(cashier *response.CashierResponseDeleteAt) *pb.CashierResponseDeleteAt {
	return &pb.CashierResponseDeleteAt{
		Id:         int32(cashier.ID),
		MerchantId: int32(cashier.MerchantID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt,
		UpdatedAt:  cashier.UpdatedAt,
		DeletedAt:  cashier.DeletedAt,
	}
}

func (c *cashierProtoMapper) mapResponsesCashierDeleteAt(cashiers []*response.CashierResponseDeleteAt) []*pb.CashierResponseDeleteAt {
	var mappedCashiers []*pb.CashierResponseDeleteAt

	for _, cashier := range cashiers {
		mappedCashiers = append(mappedCashiers, c.mapResponseCashierDeleteAt(cashier))
	}

	return mappedCashiers
}
