package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
	"time"
)

type orderRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.OrderRecordMapping
}

func NewOrderRepository(db *db.Queries, ctx context.Context, mapping recordmapper.OrderRecordMapping) *orderRepository {
	return &orderRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *orderRepository) FindAllOrders(search string, page, pageSize int) ([]*record.OrderRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetOrdersParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetOrders(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToOrdersRecordPagination(res), totalCount, nil
}

func (r *orderRepository) FindByActive(search string, page, pageSize int) ([]*record.OrderRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetOrdersActiveParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetOrdersActive(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToOrdersRecordActivePagination(res), totalCount, nil
}

func (r *orderRepository) FindByTrashed(search string, page, pageSize int) ([]*record.OrderRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetOrdersTrashedParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetOrdersTrashed(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToOrdersRecordTrashedPagination(res), totalCount, nil
}

func (r *orderRepository) GetMonthlyTotalRevenue(year int, month int) ([]*record.OrderMonthlyTotalRevenueRecord, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalRevenue(r.ctx, db.GetMonthlyTotalRevenueParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get month total revenue order: %w", err)
	}

	so := r.mapping.ToOrderMonthlyTotalRevenues(res)

	return so, nil
}

func (r *orderRepository) GetYearlyTotalRevenue(year int) ([]*record.OrderYearlyTotalRevenueRecord, error) {
	res, err := r.db.GetYearlyTotalRevenue(r.ctx, int32(year))

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total revenue order: %w", err)
	}

	so := r.mapping.ToOrderYearlyTotalRevenues(res)

	return so, nil
}

func (r *orderRepository) GetMonthlyTotalRevenueById(year int, month int, order_id int) ([]*record.OrderMonthlyTotalRevenueRecord, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalRevenueById(r.ctx, db.GetMonthlyTotalRevenueByIdParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		OrderID:     int32(order_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get month total revenue order: %w", err)
	}

	so := r.mapping.ToOrderMonthlyTotalRevenuesById(res)

	return so, nil
}

func (r *orderRepository) GetYearlyTotalRevenueById(year int, order_id int) ([]*record.OrderYearlyTotalRevenueRecord, error) {
	res, err := r.db.GetYearlyTotalRevenueById(r.ctx, db.GetYearlyTotalRevenueByIdParams{
		Column1: int32(year),
		OrderID: int32(order_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total revenue order: %w", err)
	}

	so := r.mapping.ToOrderYearlyTotalRevenuesById(res)

	return so, nil
}

func (r *orderRepository) GetMonthlyTotalRevenueByMerchant(year int, month int, merchant_id int) ([]*record.OrderMonthlyTotalRevenueRecord, error) {
	currentMonthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := currentMonthStart.AddDate(0, 1, -1)
	prevMonthStart := currentMonthStart.AddDate(0, -1, 0)
	prevMonthEnd := prevMonthStart.AddDate(0, 1, -1)

	res, err := r.db.GetMonthlyTotalRevenueByMerchant(r.ctx, db.GetMonthlyTotalRevenueByMerchantParams{
		Extract:     currentMonthStart,
		CreatedAt:   sql.NullTime{Time: currentMonthEnd, Valid: true},
		CreatedAt_2: sql.NullTime{Time: prevMonthStart, Valid: true},
		CreatedAt_3: sql.NullTime{Time: prevMonthEnd, Valid: true},
		MerchantID:  int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get month total revenue order: %w", err)
	}

	so := r.mapping.ToOrderMonthlyTotalRevenuesByMerchant(res)

	return so, nil
}

func (r *orderRepository) GetYearlyTotalRevenueByMerchant(year int, merchant_id int) ([]*record.OrderYearlyTotalRevenueRecord, error) {
	res, err := r.db.GetYearlyTotalRevenueByMerchant(r.ctx, db.GetYearlyTotalRevenueByMerchantParams{
		Column1:    int32(year),
		MerchantID: int32(merchant_id),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to get yearly total revenue order: %w", err)
	}

	so := r.mapping.ToOrderYearlyTotalRevenuesByMerchant(res)

	return so, nil
}

func (r *orderRepository) GetMonthlyOrder(year int) ([]*record.OrderMonthlyRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	res, err := r.db.GetMonthlyOrder(r.ctx, yearStart)

	if err != nil {
		return nil, fmt.Errorf("failed to get monthly orders: %w", err)
	}

	return r.mapping.ToOrderMonthlyPrices(res), nil
}

func (r *orderRepository) GetYearlyOrder(year int) ([]*record.OrderYearlyRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyOrder(r.ctx, yearStart)
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly orders: %w", err)
	}

	return r.mapping.ToOrderYearlyPrices(res), nil
}

func (r *orderRepository) GetMonthlyOrderByMerchant(year int, merchant_id int) ([]*record.OrderMonthlyRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetMonthlyOrderByMerchant(r.ctx, db.GetMonthlyOrderByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get monthly orders for merchant %d: %w", merchant_id, err)
	}

	return r.mapping.ToOrderMonthlyPricesByMerchant(res), nil
}

func (r *orderRepository) GetYearlyOrderByMerchant(year int, merchant_id int) ([]*record.OrderYearlyRecord, error) {
	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)

	res, err := r.db.GetYearlyOrderByMerchant(r.ctx, db.GetYearlyOrderByMerchantParams{
		Column1:    yearStart,
		MerchantID: int32(merchant_id),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get yearly orders for merchant %d: %w", merchant_id, err)
	}

	return r.mapping.ToOrderYearlyPricesByMerchant(res), nil
}

func (r *orderRepository) FindByMerchant(merchant_id int, search string, page, pageSize int) ([]*record.OrderRecord, int, error) {
	offset := (page - 1) * pageSize

	req := db.GetOrdersByMerchantParams{
		Column1: search,
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	}

	res, err := r.db.GetOrdersByMerchant(r.ctx, req)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to find users: %w", err)
	}

	var totalCount int
	if len(res) > 0 {
		totalCount = int(res[0].TotalCount)
	} else {
		totalCount = 0
	}

	return r.mapping.ToOrdersRecordByMerchantPagination(res), totalCount, nil
}

func (r *orderRepository) FindById(user_id int) (*record.OrderRecord, error) {
	res, err := r.db.GetOrderByID(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}

	return r.mapping.ToOrderRecord(res), nil
}

func (r *orderRepository) CreateOrder(request *requests.CreateOrderRecordRequest) (*record.OrderRecord, error) {
	req := db.CreateOrderParams{
		MerchantID: int32(request.MerchantID),
		CashierID:  int32(request.CashierID),
		TotalPrice: int64(request.TotalPrice),
	}

	user, err := r.db.CreateOrder(r.ctx, req)

	if err != nil {
		return nil, errors.New("failed create order")
	}

	return r.mapping.ToOrderRecord(user), nil
}

func (r *orderRepository) UpdateOrder(request *requests.UpdateOrderRecordRequest) (*record.OrderRecord, error) {
	req := db.UpdateOrderParams{
		OrderID:    int32(request.OrderID),
		TotalPrice: int64(request.TotalPrice),
	}

	res, err := r.db.UpdateOrder(r.ctx, req)

	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return r.mapping.ToOrderRecord(res), nil
}

func (r *orderRepository) TrashedOrder(user_id int) (*record.OrderRecord, error) {
	res, err := r.db.TrashedOrder(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to trash order: %w", err)
	}

	return r.mapping.ToOrderRecord(res), nil
}

func (r *orderRepository) RestoreOrder(user_id int) (*record.OrderRecord, error) {
	res, err := r.db.RestoreOrder(r.ctx, int32(user_id))

	if err != nil {
		return nil, fmt.Errorf("failed to restore order: %w", err)
	}

	return r.mapping.ToOrderRecord(res), nil
}

func (r *orderRepository) DeleteOrderPermanent(user_id int) (bool, error) {
	err := r.db.DeleteOrderPermanently(r.ctx, int32(user_id))

	if err != nil {
		return false, fmt.Errorf("failed to delete order: %w", err)
	}

	return true, nil
}

func (r *orderRepository) RestoreAllOrder() (bool, error) {
	err := r.db.RestoreAllOrders(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to restore all orders: %w", err)
	}
	return true, nil
}

func (r *orderRepository) DeleteAllOrderPermanent() (bool, error) {
	err := r.db.DeleteAllPermanentOrders(r.ctx)

	if err != nil {
		return false, fmt.Errorf("failed to delete all orders permanently: %w", err)
	}
	return true, nil
}
