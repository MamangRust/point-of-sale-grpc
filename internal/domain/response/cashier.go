package response

type CashierResponse struct {
	ID         int    `json:"id"`
	MerchantID int    `json:"merchant_id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type CashierResponseDeleteAt struct {
	ID         int    `json:"id"`
	MerchantID int    `json:"merchant_id"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"deleted_at"`
}

type ApiResponseCashier struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    *CashierResponse `json:"data"`
}

type ApiResponsesCashier struct {
	Status  string             `json:"status"`
	Message string             `json:"message"`
	Data    []*CashierResponse `json:"data"`
}

type ApiResponseCashierDeleteAt struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message"`
	Data    *CashierResponseDeleteAt `json:"data"`
}

type ApiResponseCashierDelete struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponseCashierAll struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ApiResponsePaginationCashierDeleteAt struct {
	Status     string                     `json:"status"`
	Message    string                     `json:"message"`
	Data       []*CashierResponseDeleteAt `json:"data"`
	Pagination PaginationMeta             `json:"pagination"`
}

type ApiResponsePaginationCashier struct {
	Status     string             `json:"status"`
	Message    string             `json:"message"`
	Data       []*CashierResponse `json:"data"`
	Pagination PaginationMeta     `json:"pagination"`
}
