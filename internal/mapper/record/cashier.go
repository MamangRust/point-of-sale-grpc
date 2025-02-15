package recordmapper

import (
	"pointofsale/internal/domain/record"
	db "pointofsale/pkg/database/schema"
)

type cashierRecordMapper struct {
}

func NewCashierRecordMapper() *cashierRecordMapper {
	return &cashierRecordMapper{}
}

func (s *cashierRecordMapper) ToCashierRecord(cashier *db.Cashier) *record.CashierRecord {
	var deletedAt *string
	if cashier.DeletedAt.Valid {
		deletedAtStr := cashier.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.CashierRecord{
		ID:         int(cashier.CashierID),
		MerchantID: int(cashier.MerchantID),
		UserID:     int(cashier.UserID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  deletedAt,
	}
}

func (s *cashierRecordMapper) ToCashierRecordPagination(cashier *db.GetCashiersRow) *record.CashierRecord {
	var deletedAt *string
	if cashier.DeletedAt.Valid {
		deletedAtStr := cashier.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.CashierRecord{
		ID:         int(cashier.CashierID),
		MerchantID: int(cashier.MerchantID),
		UserID:     int(cashier.UserID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  deletedAt,
	}
}

func (s *cashierRecordMapper) ToCashiersRecordPagination(cashiers []*db.GetCashiersRow) []*record.CashierRecord {
	var result []*record.CashierRecord

	for _, cashier := range cashiers {
		result = append(result, s.ToCashierRecordPagination(cashier))
	}

	return result
}

func (s *cashierRecordMapper) ToCashierRecordActivePagination(cashier *db.GetCashiersActiveRow) *record.CashierRecord {
	var deletedAt *string
	if cashier.DeletedAt.Valid {
		deletedAtStr := cashier.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.CashierRecord{
		ID:         int(cashier.CashierID),
		MerchantID: int(cashier.MerchantID),
		UserID:     int(cashier.UserID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  deletedAt,
	}
}

func (s *cashierRecordMapper) ToCashiersRecordActivePagination(cashiers []*db.GetCashiersActiveRow) []*record.CashierRecord {
	var result []*record.CashierRecord

	for _, cashier := range cashiers {
		result = append(result, s.ToCashierRecordActivePagination(cashier))
	}

	return result
}

func (s *cashierRecordMapper) ToCashierRecordTrashedPagination(cashier *db.GetCashiersTrashedRow) *record.CashierRecord {
	var deletedAt *string
	if cashier.DeletedAt.Valid {
		deletedAtStr := cashier.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.CashierRecord{
		ID:         int(cashier.CashierID),
		MerchantID: int(cashier.MerchantID),
		UserID:     int(cashier.UserID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  deletedAt,
	}
}

func (s *cashierRecordMapper) ToCashiersRecordTrashedPagination(cashiers []*db.GetCashiersTrashedRow) []*record.CashierRecord {
	var result []*record.CashierRecord

	for _, cashier := range cashiers {
		result = append(result, s.ToCashierRecordTrashedPagination(cashier))
	}

	return result
}

func (s *cashierRecordMapper) ToCashierMerchantRecordPagination(cashier *db.GetCashiersByMerchantRow) *record.CashierRecord {
	var deletedAt *string
	if cashier.DeletedAt.Valid {
		deletedAtStr := cashier.DeletedAt.Time.Format("2006-01-02 15:04:05.000")
		deletedAt = &deletedAtStr
	}

	return &record.CashierRecord{
		ID:         int(cashier.CashierID),
		MerchantID: int(cashier.MerchantID),
		UserID:     int(cashier.UserID),
		Name:       cashier.Name,
		CreatedAt:  cashier.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:  cashier.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:  deletedAt,
	}
}

func (s *cashierRecordMapper) ToCashiersMerchantRecordPagination(cashiers []*db.GetCashiersByMerchantRow) []*record.CashierRecord {
	var result []*record.CashierRecord

	for _, cashier := range cashiers {
		result = append(result, s.ToCashierMerchantRecordPagination(cashier))
	}

	return result
}
