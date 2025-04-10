// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: orders.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (merchant_id, cashier_id, total_price)
VALUES ($1, $2, $3)
RETURNING order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at
`

type CreateOrderParams struct {
	MerchantID int32 `json:"merchant_id"`
	CashierID  int32 `json:"cashier_id"`
	TotalPrice int64 `json:"total_price"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (*Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder, arg.MerchantID, arg.CashierID, arg.TotalPrice)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.MerchantID,
		&i.CashierID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteAllPermanentOrders = `-- name: DeleteAllPermanentOrders :exec
DELETE FROM orders
WHERE
    deleted_at IS NOT NULL
`

// Delete All Trashed Order Permanently
func (q *Queries) DeleteAllPermanentOrders(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllPermanentOrders)
	return err
}

const deleteOrderPermanently = `-- name: DeleteOrderPermanently :exec
DELETE FROM orders WHERE order_id = $1 AND deleted_at IS NOT NULL
`

// Delete Order Permanently
func (q *Queries) DeleteOrderPermanently(ctx context.Context, orderID int32) error {
	_, err := q.db.ExecContext(ctx, deleteOrderPermanently, orderID)
	return err
}

const getMonthlyOrder = `-- name: GetMonthlyOrder :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
monthly_orders AS (
    SELECT
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price)::NUMERIC AS total_revenue,
        SUM(oi.quantity) AS total_items_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
    GROUP BY
        activity_month
)
SELECT
    TO_CHAR(mo.activity_month, 'Mon') AS month,
    mo.order_count,
    mo.total_revenue,
    mo.total_items_sold
FROM
    monthly_orders mo
ORDER BY
    mo.activity_month
`

type GetMonthlyOrderRow struct {
	Month          string  `json:"month"`
	OrderCount     int64   `json:"order_count"`
	TotalRevenue   float64 `json:"total_revenue"`
	TotalItemsSold int64   `json:"total_items_sold"`
}

func (q *Queries) GetMonthlyOrder(ctx context.Context, dollar_1 time.Time) ([]*GetMonthlyOrderRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyOrder, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyOrderRow
	for rows.Next() {
		var i GetMonthlyOrderRow
		if err := rows.Scan(
			&i.Month,
			&i.OrderCount,
			&i.TotalRevenue,
			&i.TotalItemsSold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyOrderByMerchant = `-- name: GetMonthlyOrderByMerchant :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
monthly_orders AS (
    SELECT
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price)::NUMERIC AS total_revenue,
        SUM(oi.quantity) AS total_items_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
        AND o.merchant_id = $2
    GROUP BY
        activity_month
)
SELECT
    TO_CHAR(mo.activity_month, 'Mon') AS month,
    mo.order_count,
    mo.total_revenue,
    mo.total_items_sold
FROM
    monthly_orders mo
ORDER BY
    mo.activity_month
`

type GetMonthlyOrderByMerchantParams struct {
	Column1    time.Time `json:"column_1"`
	MerchantID int32     `json:"merchant_id"`
}

type GetMonthlyOrderByMerchantRow struct {
	Month          string  `json:"month"`
	OrderCount     int64   `json:"order_count"`
	TotalRevenue   float64 `json:"total_revenue"`
	TotalItemsSold int64   `json:"total_items_sold"`
}

func (q *Queries) GetMonthlyOrderByMerchant(ctx context.Context, arg GetMonthlyOrderByMerchantParams) ([]*GetMonthlyOrderByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyOrderByMerchant, arg.Column1, arg.MerchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyOrderByMerchantRow
	for rows.Next() {
		var i GetMonthlyOrderByMerchantRow
		if err := rows.Scan(
			&i.Month,
			&i.OrderCount,
			&i.TotalRevenue,
			&i.TotalItemsSold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyTotalRevenue = `-- name: GetMonthlyTotalRevenue :many
WITH monthly_revenue AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::TEXT AS year,
        EXTRACT(MONTH FROM o.created_at)::integer AS month,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND (
            (o.created_at >= $1 AND o.created_at <= $2)  
            OR (o.created_at >= $3 AND o.created_at <= $4) 
        )
    GROUP BY
        EXTRACT(YEAR FROM o.created_at),
        EXTRACT(MONTH FROM o.created_at)
),
all_months AS (
    SELECT 
        EXTRACT(YEAR FROM $1)::TEXT AS year,
        EXTRACT(MONTH FROM $1)::integer AS month,
        TO_CHAR($1, 'FMMonth') AS month_name
    
    UNION
    
    SELECT 
        EXTRACT(YEAR FROM $3)::TEXT AS year,
        EXTRACT(MONTH FROM $3)::integer AS month,
        TO_CHAR($3, 'FMMonth') AS month_name
)
SELECT 
    COALESCE(am.year, EXTRACT(YEAR FROM $1)::TEXT) AS year,
    COALESCE(am.month_name, TO_CHAR($1, 'FMMonth')) AS month,
    COALESCE(mr.total_revenue, 0) AS total_revenue
FROM 
    all_months am
LEFT JOIN 
    monthly_revenue mr ON am.year = mr.year 
                      AND am.month = mr.month
ORDER BY 
    am.year DESC,
    am.month DESC
`

type GetMonthlyTotalRevenueParams struct {
	Extract     time.Time    `json:"extract"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CreatedAt_2 sql.NullTime `json:"created_at_2"`
	CreatedAt_3 sql.NullTime `json:"created_at_3"`
}

type GetMonthlyTotalRevenueRow struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalRevenue int32  `json:"total_revenue"`
}

func (q *Queries) GetMonthlyTotalRevenue(ctx context.Context, arg GetMonthlyTotalRevenueParams) ([]*GetMonthlyTotalRevenueRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyTotalRevenue,
		arg.Extract,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.CreatedAt_3,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyTotalRevenueRow
	for rows.Next() {
		var i GetMonthlyTotalRevenueRow
		if err := rows.Scan(&i.Year, &i.Month, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyTotalRevenueById = `-- name: GetMonthlyTotalRevenueById :many
WITH monthly_revenue AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::TEXT AS year,
        EXTRACT(MONTH FROM o.created_at)::integer AS month,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND (
            (o.created_at >= $1 AND o.created_at <= $2)  
            OR (o.created_at >= $3 AND o.created_at <= $4) 
        )
        AND o.order_id = $5
    GROUP BY
        EXTRACT(YEAR FROM o.created_at),
        EXTRACT(MONTH FROM o.created_at)
),
all_months AS (
    SELECT 
        EXTRACT(YEAR FROM $1)::TEXT AS year,
        EXTRACT(MONTH FROM $1)::integer AS month,
        TO_CHAR($1, 'FMMonth') AS month_name
    
    UNION
    
    SELECT 
        EXTRACT(YEAR FROM $3)::TEXT AS year,
        EXTRACT(MONTH FROM $3)::integer AS month,
        TO_CHAR($3, 'FMMonth') AS month_name
)
SELECT 
    COALESCE(am.year, EXTRACT(YEAR FROM $1)::TEXT) AS year,
    COALESCE(am.month_name, TO_CHAR($1, 'FMMonth')) AS month,
    COALESCE(mr.total_revenue, 0) AS total_revenue
FROM 
    all_months am
LEFT JOIN 
    monthly_revenue mr ON am.year = mr.year 
                      AND am.month = mr.month
ORDER BY 
    am.year DESC,
    am.month DESC
`

type GetMonthlyTotalRevenueByIdParams struct {
	Extract     time.Time    `json:"extract"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CreatedAt_2 sql.NullTime `json:"created_at_2"`
	CreatedAt_3 sql.NullTime `json:"created_at_3"`
	OrderID     int32        `json:"order_id"`
}

type GetMonthlyTotalRevenueByIdRow struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalRevenue int32  `json:"total_revenue"`
}

func (q *Queries) GetMonthlyTotalRevenueById(ctx context.Context, arg GetMonthlyTotalRevenueByIdParams) ([]*GetMonthlyTotalRevenueByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyTotalRevenueById,
		arg.Extract,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.CreatedAt_3,
		arg.OrderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyTotalRevenueByIdRow
	for rows.Next() {
		var i GetMonthlyTotalRevenueByIdRow
		if err := rows.Scan(&i.Year, &i.Month, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMonthlyTotalRevenueByMerchant = `-- name: GetMonthlyTotalRevenueByMerchant :many
WITH monthly_revenue AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::TEXT AS year,
        EXTRACT(MONTH FROM o.created_at)::integer AS month,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND (
            (o.created_at >= $1 AND o.created_at <= $2)  
            OR (o.created_at >= $3 AND o.created_at <= $4) 
        )
        AND o.merchant_id = $5
    GROUP BY
        EXTRACT(YEAR FROM o.created_at),
        EXTRACT(MONTH FROM o.created_at)
),
all_months AS (
    SELECT 
        EXTRACT(YEAR FROM $1)::TEXT AS year,
        EXTRACT(MONTH FROM $1)::integer AS month,
        TO_CHAR($1, 'FMMonth') AS month_name
    
    UNION
    
    SELECT 
        EXTRACT(YEAR FROM $3)::TEXT AS year,
        EXTRACT(MONTH FROM $3)::integer AS month,
        TO_CHAR($3, 'FMMonth') AS month_name
)
SELECT 
    COALESCE(am.year, EXTRACT(YEAR FROM $1)::TEXT) AS year,
    COALESCE(am.month_name, TO_CHAR($1, 'FMMonth')) AS month,
    COALESCE(mr.total_revenue, 0) AS total_revenue
FROM 
    all_months am
LEFT JOIN 
    monthly_revenue mr ON am.year = mr.year 
                      AND am.month = mr.month
ORDER BY 
    am.year DESC,
    am.month DESC
`

type GetMonthlyTotalRevenueByMerchantParams struct {
	Extract     time.Time    `json:"extract"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CreatedAt_2 sql.NullTime `json:"created_at_2"`
	CreatedAt_3 sql.NullTime `json:"created_at_3"`
	MerchantID  int32        `json:"merchant_id"`
}

type GetMonthlyTotalRevenueByMerchantRow struct {
	Year         string `json:"year"`
	Month        string `json:"month"`
	TotalRevenue int32  `json:"total_revenue"`
}

func (q *Queries) GetMonthlyTotalRevenueByMerchant(ctx context.Context, arg GetMonthlyTotalRevenueByMerchantParams) ([]*GetMonthlyTotalRevenueByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyTotalRevenueByMerchant,
		arg.Extract,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.CreatedAt_3,
		arg.MerchantID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyTotalRevenueByMerchantRow
	for rows.Next() {
		var i GetMonthlyTotalRevenueByMerchantRow
		if err := rows.Scan(&i.Year, &i.Month, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderByID = `-- name: GetOrderByID :one
SELECT order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at
FROM orders
WHERE order_id = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetOrderByID(ctx context.Context, orderID int32) (*Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByID, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.MerchantID,
		&i.CashierID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getOrders = `-- name: GetOrders :many
SELECT
    order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetOrdersParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetOrdersRow struct {
	OrderID    int32        `json:"order_id"`
	MerchantID int32        `json:"merchant_id"`
	CashierID  int32        `json:"cashier_id"`
	TotalPrice int64        `json:"total_price"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Orders with Pagination and Total Count
func (q *Queries) GetOrders(ctx context.Context, arg GetOrdersParams) ([]*GetOrdersRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrders, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetOrdersRow
	for rows.Next() {
		var i GetOrdersRow
		if err := rows.Scan(
			&i.OrderID,
			&i.MerchantID,
			&i.CashierID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrdersActive = `-- name: GetOrdersActive :many
SELECT
    order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetOrdersActiveParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetOrdersActiveRow struct {
	OrderID    int32        `json:"order_id"`
	MerchantID int32        `json:"merchant_id"`
	CashierID  int32        `json:"cashier_id"`
	TotalPrice int64        `json:"total_price"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Active Orders with Pagination and Total Count
func (q *Queries) GetOrdersActive(ctx context.Context, arg GetOrdersActiveParams) ([]*GetOrdersActiveRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersActive, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetOrdersActiveRow
	for rows.Next() {
		var i GetOrdersActiveRow
		if err := rows.Scan(
			&i.OrderID,
			&i.MerchantID,
			&i.CashierID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrdersByMerchant = `-- name: GetOrdersByMerchant :many
SELECT
    order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE 
    deleted_at IS NULL
    AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
    AND ($4::UUID IS NULL OR merchant_id = $4)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetOrdersByMerchantParams struct {
	Column1 string    `json:"column_1"`
	Limit   int32     `json:"limit"`
	Offset  int32     `json:"offset"`
	Column4 uuid.UUID `json:"column_4"`
}

type GetOrdersByMerchantRow struct {
	OrderID    int32        `json:"order_id"`
	MerchantID int32        `json:"merchant_id"`
	CashierID  int32        `json:"cashier_id"`
	TotalPrice int64        `json:"total_price"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Orders with Pagination and Total Count where merchant_id
func (q *Queries) GetOrdersByMerchant(ctx context.Context, arg GetOrdersByMerchantParams) ([]*GetOrdersByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersByMerchant,
		arg.Column1,
		arg.Limit,
		arg.Offset,
		arg.Column4,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetOrdersByMerchantRow
	for rows.Next() {
		var i GetOrdersByMerchantRow
		if err := rows.Scan(
			&i.OrderID,
			&i.MerchantID,
			&i.CashierID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrdersTrashed = `-- name: GetOrdersTrashed :many
SELECT
    order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetOrdersTrashedParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetOrdersTrashedRow struct {
	OrderID    int32        `json:"order_id"`
	MerchantID int32        `json:"merchant_id"`
	CashierID  int32        `json:"cashier_id"`
	TotalPrice int64        `json:"total_price"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Trashed Orders with Pagination and Total Count
func (q *Queries) GetOrdersTrashed(ctx context.Context, arg GetOrdersTrashedParams) ([]*GetOrdersTrashedRow, error) {
	rows, err := q.db.QueryContext(ctx, getOrdersTrashed, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetOrdersTrashedRow
	for rows.Next() {
		var i GetOrdersTrashedRow
		if err := rows.Scan(
			&i.OrderID,
			&i.MerchantID,
			&i.CashierID,
			&i.TotalPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyOrder = `-- name: GetYearlyOrder :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price)::NUMERIC AS total_revenue,
        SUM(oi.quantity) AS total_items_sold,
        COUNT(DISTINCT o.cashier_id) AS active_cashiers,
        COUNT(DISTINCT oi.product_id) AS unique_products_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    order_count,
    total_revenue,
    total_items_sold,
    active_cashiers,
    unique_products_sold
FROM
    last_five_years
ORDER BY
    year
`

type GetYearlyOrderRow struct {
	Year               string  `json:"year"`
	OrderCount         int64   `json:"order_count"`
	TotalRevenue       float64 `json:"total_revenue"`
	TotalItemsSold     int64   `json:"total_items_sold"`
	ActiveCashiers     int64   `json:"active_cashiers"`
	UniqueProductsSold int64   `json:"unique_products_sold"`
}

func (q *Queries) GetYearlyOrder(ctx context.Context, dollar_1 time.Time) ([]*GetYearlyOrderRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyOrder, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyOrderRow
	for rows.Next() {
		var i GetYearlyOrderRow
		if err := rows.Scan(
			&i.Year,
			&i.OrderCount,
			&i.TotalRevenue,
			&i.TotalItemsSold,
			&i.ActiveCashiers,
			&i.UniqueProductsSold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyOrderByMerchant = `-- name: GetYearlyOrderByMerchant :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price)::NUMERIC AS total_revenue,
        SUM(oi.quantity) AS total_items_sold,
        COUNT(DISTINCT o.cashier_id) AS active_cashiers,
        COUNT(DISTINCT oi.product_id) AS unique_products_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
        AND o.merchant_id = $2
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    order_count,
    total_revenue,
    total_items_sold,
    active_cashiers,
    unique_products_sold
FROM
    last_five_years
ORDER BY
    year
`

type GetYearlyOrderByMerchantParams struct {
	Column1    time.Time `json:"column_1"`
	MerchantID int32     `json:"merchant_id"`
}

type GetYearlyOrderByMerchantRow struct {
	Year               string  `json:"year"`
	OrderCount         int64   `json:"order_count"`
	TotalRevenue       float64 `json:"total_revenue"`
	TotalItemsSold     int64   `json:"total_items_sold"`
	ActiveCashiers     int64   `json:"active_cashiers"`
	UniqueProductsSold int64   `json:"unique_products_sold"`
}

func (q *Queries) GetYearlyOrderByMerchant(ctx context.Context, arg GetYearlyOrderByMerchantParams) ([]*GetYearlyOrderByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyOrderByMerchant, arg.Column1, arg.MerchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyOrderByMerchantRow
	for rows.Next() {
		var i GetYearlyOrderByMerchantRow
		if err := rows.Scan(
			&i.Year,
			&i.OrderCount,
			&i.TotalRevenue,
			&i.TotalItemsSold,
			&i.ActiveCashiers,
			&i.UniqueProductsSold,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyTotalRevenue = `-- name: GetYearlyTotalRevenue :many
WITH yearly_revenue AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1 AS year
    UNION
    SELECT $1 - 1 AS year
)
SELECT 
    ay.year::text AS year,
    COALESCE(yr.total_revenue, 0) AS total_revenue
FROM 
    all_years ay
LEFT JOIN 
    yearly_revenue yr ON ay.year = yr.year
ORDER BY 
    ay.year DESC
`

type GetYearlyTotalRevenueRow struct {
	Year         string `json:"year"`
	TotalRevenue int32  `json:"total_revenue"`
}

func (q *Queries) GetYearlyTotalRevenue(ctx context.Context, dollar_1 int32) ([]*GetYearlyTotalRevenueRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyTotalRevenue, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyTotalRevenueRow
	for rows.Next() {
		var i GetYearlyTotalRevenueRow
		if err := rows.Scan(&i.Year, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyTotalRevenueById = `-- name: GetYearlyTotalRevenueById :many
WITH yearly_revenue AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
        AND o.order_id =  $2
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1 AS year
    UNION
    SELECT $1 - 1 AS year
)
SELECT 
    ay.year::text AS year,
    COALESCE(yr.total_revenue, 0) AS total_revenue
FROM 
    all_years ay
LEFT JOIN 
    yearly_revenue yr ON ay.year = yr.year
ORDER BY 
    ay.year DESC
`

type GetYearlyTotalRevenueByIdParams struct {
	Column1 int32 `json:"column_1"`
	OrderID int32 `json:"order_id"`
}

type GetYearlyTotalRevenueByIdRow struct {
	Year         string `json:"year"`
	TotalRevenue int32  `json:"total_revenue"`
}

func (q *Queries) GetYearlyTotalRevenueById(ctx context.Context, arg GetYearlyTotalRevenueByIdParams) ([]*GetYearlyTotalRevenueByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyTotalRevenueById, arg.Column1, arg.OrderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyTotalRevenueByIdRow
	for rows.Next() {
		var i GetYearlyTotalRevenueByIdRow
		if err := rows.Scan(&i.Year, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getYearlyTotalRevenueByMerchant = `-- name: GetYearlyTotalRevenueByMerchant :many
WITH yearly_revenue AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
        AND o.merchant_id = $2
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1 AS year
    UNION
    SELECT $1 - 1 AS year
)
SELECT 
    ay.year::text AS year,
    COALESCE(yr.total_revenue, 0) AS total_revenue
FROM 
    all_years ay
LEFT JOIN 
    yearly_revenue yr ON ay.year = yr.year
ORDER BY 
    ay.year DESC
`

type GetYearlyTotalRevenueByMerchantParams struct {
	Column1    int32 `json:"column_1"`
	MerchantID int32 `json:"merchant_id"`
}

type GetYearlyTotalRevenueByMerchantRow struct {
	Year         string `json:"year"`
	TotalRevenue int32  `json:"total_revenue"`
}

func (q *Queries) GetYearlyTotalRevenueByMerchant(ctx context.Context, arg GetYearlyTotalRevenueByMerchantParams) ([]*GetYearlyTotalRevenueByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyTotalRevenueByMerchant, arg.Column1, arg.MerchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyTotalRevenueByMerchantRow
	for rows.Next() {
		var i GetYearlyTotalRevenueByMerchantRow
		if err := rows.Scan(&i.Year, &i.TotalRevenue); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const restoreAllOrders = `-- name: RestoreAllOrders :exec
UPDATE orders
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL
`

// Restore All Trashed Order
func (q *Queries) RestoreAllOrders(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, restoreAllOrders)
	return err
}

const restoreOrder = `-- name: RestoreOrder :one
UPDATE orders
SET
    deleted_at = NULL
WHERE
    order_id = $1
    AND deleted_at IS NOT NULL
  RETURNING order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at
`

// Restore Trashed Order
func (q *Queries) RestoreOrder(ctx context.Context, orderID int32) (*Order, error) {
	row := q.db.QueryRowContext(ctx, restoreOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.MerchantID,
		&i.CashierID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const trashedOrder = `-- name: TrashedOrder :one
UPDATE orders
SET
    deleted_at = current_timestamp
WHERE
    order_id = $1
    AND deleted_at IS NULL
    RETURNING order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at
`

// Trash Order
func (q *Queries) TrashedOrder(ctx context.Context, orderID int32) (*Order, error) {
	row := q.db.QueryRowContext(ctx, trashedOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.MerchantID,
		&i.CashierID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updateOrder = `-- name: UpdateOrder :one
UPDATE orders
SET total_price = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE order_id = $1
  AND deleted_at IS NULL
  RETURNING order_id, merchant_id, cashier_id, total_price, created_at, updated_at, deleted_at
`

type UpdateOrderParams struct {
	OrderID    int32 `json:"order_id"`
	TotalPrice int64 `json:"total_price"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (*Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder, arg.OrderID, arg.TotalPrice)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.MerchantID,
		&i.CashierID,
		&i.TotalPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
