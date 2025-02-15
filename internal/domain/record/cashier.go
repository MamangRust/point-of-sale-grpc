package record

type CashierRecord struct {
	ID         int     `json:"id"`
	MerchantID int     `json:"merchant_id"`
	UserID     int     `json:"user_id"`
	Name       string  `json:"name"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}
