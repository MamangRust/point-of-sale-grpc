package requests

import "github.com/go-playground/validator/v10"

type CreateCashierRequest struct {
	MerchantID int    `json:"merchant_id" validate:"required"`
	UserID     int    `json:"user_id" validate:"required"`
	Name       string `json:"name" validate:"required"`
}

type UpdateCashierRequest struct {
	CashierID int    `json:"cashier_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

func (r *CreateCashierRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (r *UpdateCashierRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}
