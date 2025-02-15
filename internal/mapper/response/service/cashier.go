package response_service

import (
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/response"
)

type cashierResponseMapper struct {
}

func NewCashierResponseMapper() *cashierResponseMapper {
	return &cashierResponseMapper{}
}

func (s *cashierResponseMapper) ToCashierResponse(cashier *record.CashierRecord) *response.CashierResponse {
	return &response.CashierResponse{
		ID:         cashier.ID,
		MerchantID: cashier.MerchantID,
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt,
		UpdatedAt:  cashier.UpdatedAt,
	}
}

func (s *cashierResponseMapper) ToCashiersResponse(cashiers []*record.CashierRecord) []*response.CashierResponse {
	var responses []*response.CashierResponse

	for _, cashier := range cashiers {
		responses = append(responses, s.ToCashierResponse(cashier))
	}

	return responses
}

func (s *cashierResponseMapper) ToCashierResponseDeleteAt(cashier *record.CashierRecord) *response.CashierResponseDeleteAt {
	return &response.CashierResponseDeleteAt{
		ID:         cashier.ID,
		MerchantID: cashier.MerchantID,
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt,
		UpdatedAt:  cashier.UpdatedAt,
		DeletedAt:  *cashier.DeletedAt,
	}
}

func (s *cashierResponseMapper) ToCashiersResponseDeleteAt(cashiers []*record.CashierRecord) []*response.CashierResponseDeleteAt {
	var responses []*response.CashierResponseDeleteAt

	for _, cashier := range cashiers {
		responses = append(responses, s.ToCashierResponseDeleteAt(cashier))
	}

	return responses
}
