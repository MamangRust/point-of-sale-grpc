package requests

import "github.com/go-playground/validator/v10"

type CreateCategoryRequest struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	SlugCategory *string `json:"slug_category"`
}

type UpdateCategoryRequest struct {
	CategoryID   *int    `json:"category_id"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	SlugCategory *string `json:"slug_category"`
}

func (r *CreateCategoryRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (r *UpdateCategoryRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}
