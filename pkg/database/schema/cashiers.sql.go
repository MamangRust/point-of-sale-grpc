// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: cashiers.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCashier = `-- name: CreateCashier :one
INSERT INTO cashiers (merchant_id, user_id, name)
VALUES ($1, $2, $3) RETURNING cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at
`

type CreateCashierParams struct {
	MerchantID int32  `json:"merchant_id"`
	UserID     int32  `json:"user_id"`
	Name       string `json:"name"`
}

func (q *Queries) CreateCashier(ctx context.Context, arg CreateCashierParams) (*Cashier, error) {
	row := q.db.QueryRowContext(ctx, createCashier, arg.MerchantID, arg.UserID, arg.Name)
	var i Cashier
	err := row.Scan(
		&i.CashierID,
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteAllPermanentCashiers = `-- name: DeleteAllPermanentCashiers :exec
DELETE FROM cashiers
WHERE
    deleted_at IS NOT NULL
`

// Delete All Trashed Cashier Permanently
func (q *Queries) DeleteAllPermanentCashiers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllPermanentCashiers)
	return err
}

const deleteCashierPermanently = `-- name: DeleteCashierPermanently :exec
DELETE FROM cashiers WHERE cashier_id = $1 AND deleted_at IS NOT NULL
`

// Delete Cashier Permanently
func (q *Queries) DeleteCashierPermanently(ctx context.Context, cashierID int32) error {
	_, err := q.db.ExecContext(ctx, deleteCashierPermanently, cashierID)
	return err
}

const getCashierByID = `-- name: GetCashierByID :one
SELECT cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at
FROM cashiers
WHERE cashier_id = $1
  AND deleted_at IS NULL
`

func (q *Queries) GetCashierByID(ctx context.Context, cashierID int32) (*Cashier, error) {
	row := q.db.QueryRowContext(ctx, getCashierByID, cashierID)
	var i Cashier
	err := row.Scan(
		&i.CashierID,
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getCashiers = `-- name: GetCashiers :many
SELECT
    cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetCashiersParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetCashiersRow struct {
	CashierID  int32        `json:"cashier_id"`
	MerchantID int32        `json:"merchant_id"`
	UserID     int32        `json:"user_id"`
	Name       string       `json:"name"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetCashiers(ctx context.Context, arg GetCashiersParams) ([]*GetCashiersRow, error) {
	rows, err := q.db.QueryContext(ctx, getCashiers, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCashiersRow
	for rows.Next() {
		var i GetCashiersRow
		if err := rows.Scan(
			&i.CashierID,
			&i.MerchantID,
			&i.UserID,
			&i.Name,
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

const getCashiersActive = `-- name: GetCashiersActive :many
SELECT
    cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetCashiersActiveParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetCashiersActiveRow struct {
	CashierID  int32        `json:"cashier_id"`
	MerchantID int32        `json:"merchant_id"`
	UserID     int32        `json:"user_id"`
	Name       string       `json:"name"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetCashiersActive(ctx context.Context, arg GetCashiersActiveParams) ([]*GetCashiersActiveRow, error) {
	rows, err := q.db.QueryContext(ctx, getCashiersActive, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCashiersActiveRow
	for rows.Next() {
		var i GetCashiersActiveRow
		if err := rows.Scan(
			&i.CashierID,
			&i.MerchantID,
			&i.UserID,
			&i.Name,
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

const getCashiersByMerchant = `-- name: GetCashiersByMerchant :many
SELECT
    cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE merchant_id = $1
  AND deleted_at IS NULL
  AND ($2::TEXT IS NULL OR name ILIKE '%' || $2 || '%')
ORDER BY created_at DESC
LIMIT $3 OFFSET $4
`

type GetCashiersByMerchantParams struct {
	MerchantID int32  `json:"merchant_id"`
	Column2    string `json:"column_2"`
	Limit      int32  `json:"limit"`
	Offset     int32  `json:"offset"`
}

type GetCashiersByMerchantRow struct {
	CashierID  int32        `json:"cashier_id"`
	MerchantID int32        `json:"merchant_id"`
	UserID     int32        `json:"user_id"`
	Name       string       `json:"name"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetCashiersByMerchant(ctx context.Context, arg GetCashiersByMerchantParams) ([]*GetCashiersByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getCashiersByMerchant,
		arg.MerchantID,
		arg.Column2,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCashiersByMerchantRow
	for rows.Next() {
		var i GetCashiersByMerchantRow
		if err := rows.Scan(
			&i.CashierID,
			&i.MerchantID,
			&i.UserID,
			&i.Name,
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

const getCashiersTrashed = `-- name: GetCashiersTrashed :many
SELECT
    cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetCashiersTrashedParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetCashiersTrashedRow struct {
	CashierID  int32        `json:"cashier_id"`
	MerchantID int32        `json:"merchant_id"`
	UserID     int32        `json:"user_id"`
	Name       string       `json:"name"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

func (q *Queries) GetCashiersTrashed(ctx context.Context, arg GetCashiersTrashedParams) ([]*GetCashiersTrashedRow, error) {
	rows, err := q.db.QueryContext(ctx, getCashiersTrashed, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetCashiersTrashedRow
	for rows.Next() {
		var i GetCashiersTrashedRow
		if err := rows.Scan(
			&i.CashierID,
			&i.MerchantID,
			&i.UserID,
			&i.Name,
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

const getMonthlyCashier = `-- name: GetMonthlyCashier :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
cashier_activity AS (
    SELECT
        c.cashier_id,
        c.name AS cashier_name,
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price)::NUMERIC AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
    GROUP BY
        c.cashier_id, c.name, activity_month
)
SELECT
    ca.cashier_id,
    ca.cashier_name,
    TO_CHAR(ca.activity_month, 'Mon') AS month,
    ca.order_count,
    ca.total_sales
FROM
    cashier_activity ca
ORDER BY
    ca.activity_month, ca.cashier_id
`

type GetMonthlyCashierRow struct {
	CashierID   int32   `json:"cashier_id"`
	CashierName string  `json:"cashier_name"`
	Month       string  `json:"month"`
	OrderCount  int64   `json:"order_count"`
	TotalSales  float64 `json:"total_sales"`
}

func (q *Queries) GetMonthlyCashier(ctx context.Context, dollar_1 time.Time) ([]*GetMonthlyCashierRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyCashier, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyCashierRow
	for rows.Next() {
		var i GetMonthlyCashierRow
		if err := rows.Scan(
			&i.CashierID,
			&i.CashierName,
			&i.Month,
			&i.OrderCount,
			&i.TotalSales,
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

const getMonthlyCashierByCashierId = `-- name: GetMonthlyCashierByCashierId :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
cashier_activity AS (
    SELECT
        c.cashier_id,
        c.name AS cashier_name,
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price) AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND c.cashier_id = $2
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
    GROUP BY
        c.cashier_id, c.name, activity_month
)
SELECT
    ca.cashier_id,
    ca.cashier_name,
    TO_CHAR(ca.activity_month, 'Mon') AS month,
    ca.order_count,
    ca.total_sales
FROM
    cashier_activity ca
ORDER BY
    ca.activity_month, ca.cashier_id
`

type GetMonthlyCashierByCashierIdParams struct {
	Column1   time.Time `json:"column_1"`
	CashierID int32     `json:"cashier_id"`
}

type GetMonthlyCashierByCashierIdRow struct {
	CashierID   int32  `json:"cashier_id"`
	CashierName string `json:"cashier_name"`
	Month       string `json:"month"`
	OrderCount  int64  `json:"order_count"`
	TotalSales  int64  `json:"total_sales"`
}

func (q *Queries) GetMonthlyCashierByCashierId(ctx context.Context, arg GetMonthlyCashierByCashierIdParams) ([]*GetMonthlyCashierByCashierIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyCashierByCashierId, arg.Column1, arg.CashierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyCashierByCashierIdRow
	for rows.Next() {
		var i GetMonthlyCashierByCashierIdRow
		if err := rows.Scan(
			&i.CashierID,
			&i.CashierName,
			&i.Month,
			&i.OrderCount,
			&i.TotalSales,
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

const getMonthlyCashierByMerchant = `-- name: GetMonthlyCashierByMerchant :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
cashier_activity AS (
    SELECT
        c.cashier_id,
        c.name AS cashier_name,
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price) AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND c.merchant_id = $2
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
    GROUP BY
        c.cashier_id, c.name, activity_month
)
SELECT
    ca.cashier_id,
    ca.cashier_name,
    TO_CHAR(ca.activity_month, 'Mon') AS month,
    ca.order_count,
    ca.total_sales
FROM
    cashier_activity ca
ORDER BY
    ca.activity_month, ca.cashier_id
`

type GetMonthlyCashierByMerchantParams struct {
	Column1    time.Time `json:"column_1"`
	MerchantID int32     `json:"merchant_id"`
}

type GetMonthlyCashierByMerchantRow struct {
	CashierID   int32  `json:"cashier_id"`
	CashierName string `json:"cashier_name"`
	Month       string `json:"month"`
	OrderCount  int64  `json:"order_count"`
	TotalSales  int64  `json:"total_sales"`
}

func (q *Queries) GetMonthlyCashierByMerchant(ctx context.Context, arg GetMonthlyCashierByMerchantParams) ([]*GetMonthlyCashierByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyCashierByMerchant, arg.Column1, arg.MerchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyCashierByMerchantRow
	for rows.Next() {
		var i GetMonthlyCashierByMerchantRow
		if err := rows.Scan(
			&i.CashierID,
			&i.CashierName,
			&i.Month,
			&i.OrderCount,
			&i.TotalSales,
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

const getMonthlyTotalSalesById = `-- name: GetMonthlyTotalSalesById :many
WITH monthly_totals AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::TEXT AS year,
        EXTRACT(MONTH FROM o.created_at)::integer AS month,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND (
            (o.created_at >= $1 AND o.created_at <= $2)  
            OR (o.created_at >= $3 AND o.created_at <= $4)  
        )
        AND c.cashier_id = $5
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
    COALESCE(mt.total_sales, 0) AS total_sales
FROM 
    all_months am
LEFT JOIN 
    monthly_totals mt ON am.year = mt.year AND am.month = mt.month
ORDER BY 
    am.year::INT DESC,
    am.month DESC
`

type GetMonthlyTotalSalesByIdParams struct {
	Extract     time.Time    `json:"extract"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CreatedAt_2 sql.NullTime `json:"created_at_2"`
	CreatedAt_3 sql.NullTime `json:"created_at_3"`
	CashierID   int32        `json:"cashier_id"`
}

type GetMonthlyTotalSalesByIdRow struct {
	Year       string `json:"year"`
	Month      string `json:"month"`
	TotalSales int32  `json:"total_sales"`
}

func (q *Queries) GetMonthlyTotalSalesById(ctx context.Context, arg GetMonthlyTotalSalesByIdParams) ([]*GetMonthlyTotalSalesByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyTotalSalesById,
		arg.Extract,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.CreatedAt_3,
		arg.CashierID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyTotalSalesByIdRow
	for rows.Next() {
		var i GetMonthlyTotalSalesByIdRow
		if err := rows.Scan(&i.Year, &i.Month, &i.TotalSales); err != nil {
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

const getMonthlyTotalSalesByMerchant = `-- name: GetMonthlyTotalSalesByMerchant :many
WITH monthly_totals AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::TEXT AS year,
        EXTRACT(MONTH FROM o.created_at)::integer AS month,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
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
    COALESCE(mt.total_sales, 0) AS total_sales
FROM 
    all_months am
LEFT JOIN 
    monthly_totals mt ON am.year = mt.year AND am.month = mt.month
ORDER BY 
    am.year::INT DESC,
    am.month DESC
`

type GetMonthlyTotalSalesByMerchantParams struct {
	Extract     time.Time    `json:"extract"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CreatedAt_2 sql.NullTime `json:"created_at_2"`
	CreatedAt_3 sql.NullTime `json:"created_at_3"`
	MerchantID  int32        `json:"merchant_id"`
}

type GetMonthlyTotalSalesByMerchantRow struct {
	Year       string `json:"year"`
	Month      string `json:"month"`
	TotalSales int32  `json:"total_sales"`
}

func (q *Queries) GetMonthlyTotalSalesByMerchant(ctx context.Context, arg GetMonthlyTotalSalesByMerchantParams) ([]*GetMonthlyTotalSalesByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyTotalSalesByMerchant,
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
	var items []*GetMonthlyTotalSalesByMerchantRow
	for rows.Next() {
		var i GetMonthlyTotalSalesByMerchantRow
		if err := rows.Scan(&i.Year, &i.Month, &i.TotalSales); err != nil {
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

const getMonthlyTotalSalesCashier = `-- name: GetMonthlyTotalSalesCashier :many
WITH monthly_totals AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::TEXT AS year,
        EXTRACT(MONTH FROM o.created_at)::integer AS month,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
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
    COALESCE(mt.total_sales, 0) AS total_sales
FROM 
    all_months am
LEFT JOIN 
    monthly_totals mt ON am.year = mt.year AND am.month = mt.month
ORDER BY 
    am.year::INT DESC,
    am.month DESC
`

type GetMonthlyTotalSalesCashierParams struct {
	Extract     time.Time    `json:"extract"`
	CreatedAt   sql.NullTime `json:"created_at"`
	CreatedAt_2 sql.NullTime `json:"created_at_2"`
	CreatedAt_3 sql.NullTime `json:"created_at_3"`
}

type GetMonthlyTotalSalesCashierRow struct {
	Year       string `json:"year"`
	Month      string `json:"month"`
	TotalSales int32  `json:"total_sales"`
}

func (q *Queries) GetMonthlyTotalSalesCashier(ctx context.Context, arg GetMonthlyTotalSalesCashierParams) ([]*GetMonthlyTotalSalesCashierRow, error) {
	rows, err := q.db.QueryContext(ctx, getMonthlyTotalSalesCashier,
		arg.Extract,
		arg.CreatedAt,
		arg.CreatedAt_2,
		arg.CreatedAt_3,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetMonthlyTotalSalesCashierRow
	for rows.Next() {
		var i GetMonthlyTotalSalesCashierRow
		if err := rows.Scan(&i.Year, &i.Month, &i.TotalSales); err != nil {
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

const getYearlyCashier = `-- name: GetYearlyCashier :many
WITH last_five_years AS (
    SELECT
        c.cashier_id,
        c.name AS cashier_name,
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price) AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
    GROUP BY
        c.cashier_id, c.name, EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    cashier_id,
    cashier_name,
    order_count,
    total_sales
FROM
    last_five_years
ORDER BY
    year, cashier_id
`

type GetYearlyCashierRow struct {
	Year        string `json:"year"`
	CashierID   int32  `json:"cashier_id"`
	CashierName string `json:"cashier_name"`
	OrderCount  int64  `json:"order_count"`
	TotalSales  int64  `json:"total_sales"`
}

func (q *Queries) GetYearlyCashier(ctx context.Context, dollar_1 time.Time) ([]*GetYearlyCashierRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyCashier, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyCashierRow
	for rows.Next() {
		var i GetYearlyCashierRow
		if err := rows.Scan(
			&i.Year,
			&i.CashierID,
			&i.CashierName,
			&i.OrderCount,
			&i.TotalSales,
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

const getYearlyCashierByCashierId = `-- name: GetYearlyCashierByCashierId :many
WITH last_five_years AS (
    SELECT
        c.cashier_id,
        c.name AS cashier_name,
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price) AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND c.cashier_id = $2
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
    GROUP BY
        c.cashier_id, c.name, EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    cashier_id,
    cashier_name,
    order_count,
    total_sales
FROM
    last_five_years
ORDER BY
    year, cashier_id
`

type GetYearlyCashierByCashierIdParams struct {
	Column1   time.Time `json:"column_1"`
	CashierID int32     `json:"cashier_id"`
}

type GetYearlyCashierByCashierIdRow struct {
	Year        string `json:"year"`
	CashierID   int32  `json:"cashier_id"`
	CashierName string `json:"cashier_name"`
	OrderCount  int64  `json:"order_count"`
	TotalSales  int64  `json:"total_sales"`
}

func (q *Queries) GetYearlyCashierByCashierId(ctx context.Context, arg GetYearlyCashierByCashierIdParams) ([]*GetYearlyCashierByCashierIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyCashierByCashierId, arg.Column1, arg.CashierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyCashierByCashierIdRow
	for rows.Next() {
		var i GetYearlyCashierByCashierIdRow
		if err := rows.Scan(
			&i.Year,
			&i.CashierID,
			&i.CashierName,
			&i.OrderCount,
			&i.TotalSales,
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

const getYearlyCashierByMerchant = `-- name: GetYearlyCashierByMerchant :many
WITH last_five_years AS (
    SELECT
        c.cashier_id,
        c.name AS cashier_name,
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(o.order_id) AS order_count,
        SUM(o.total_price) AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND c.merchant_id = $2
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
    GROUP BY
        c.cashier_id, c.name, EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    cashier_id,
    cashier_name,
    order_count,
    total_sales
FROM
    last_five_years
ORDER BY
    year, cashier_id
`

type GetYearlyCashierByMerchantParams struct {
	Column1    time.Time `json:"column_1"`
	MerchantID int32     `json:"merchant_id"`
}

type GetYearlyCashierByMerchantRow struct {
	Year        string `json:"year"`
	CashierID   int32  `json:"cashier_id"`
	CashierName string `json:"cashier_name"`
	OrderCount  int64  `json:"order_count"`
	TotalSales  int64  `json:"total_sales"`
}

func (q *Queries) GetYearlyCashierByMerchant(ctx context.Context, arg GetYearlyCashierByMerchantParams) ([]*GetYearlyCashierByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyCashierByMerchant, arg.Column1, arg.MerchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyCashierByMerchantRow
	for rows.Next() {
		var i GetYearlyCashierByMerchantRow
		if err := rows.Scan(
			&i.Year,
			&i.CashierID,
			&i.CashierName,
			&i.OrderCount,
			&i.TotalSales,
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

const getYearlyTotalSalesById = `-- name: GetYearlyTotalSalesById :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
        AND c.cashier_id = $2
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1::integer AS year
    UNION
    SELECT $1::integer - 1 AS year
)
SELECT 
    a.year::text AS year,
    COALESCE(yd.total_sales, 0) AS total_sales
FROM 
    all_years a
LEFT JOIN 
    yearly_data yd ON a.year = yd.year
ORDER BY 
    a.year DESC
`

type GetYearlyTotalSalesByIdParams struct {
	Column1   int32 `json:"column_1"`
	CashierID int32 `json:"cashier_id"`
}

type GetYearlyTotalSalesByIdRow struct {
	Year       string `json:"year"`
	TotalSales int32  `json:"total_sales"`
}

func (q *Queries) GetYearlyTotalSalesById(ctx context.Context, arg GetYearlyTotalSalesByIdParams) ([]*GetYearlyTotalSalesByIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyTotalSalesById, arg.Column1, arg.CashierID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyTotalSalesByIdRow
	for rows.Next() {
		var i GetYearlyTotalSalesByIdRow
		if err := rows.Scan(&i.Year, &i.TotalSales); err != nil {
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

const getYearlyTotalSalesByMerchant = `-- name: GetYearlyTotalSalesByMerchant :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
        AND o.merchant_id = $2
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1::integer AS year
    UNION
    SELECT $1::integer - 1 AS year
)
SELECT 
    a.year::text AS year,
    COALESCE(yd.total_sales, 0) AS total_sales
FROM 
    all_years a
LEFT JOIN 
    yearly_data yd ON a.year = yd.year
ORDER BY 
    a.year DESC
`

type GetYearlyTotalSalesByMerchantParams struct {
	Column1    int32 `json:"column_1"`
	MerchantID int32 `json:"merchant_id"`
}

type GetYearlyTotalSalesByMerchantRow struct {
	Year       string `json:"year"`
	TotalSales int32  `json:"total_sales"`
}

func (q *Queries) GetYearlyTotalSalesByMerchant(ctx context.Context, arg GetYearlyTotalSalesByMerchantParams) ([]*GetYearlyTotalSalesByMerchantRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyTotalSalesByMerchant, arg.Column1, arg.MerchantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyTotalSalesByMerchantRow
	for rows.Next() {
		var i GetYearlyTotalSalesByMerchantRow
		if err := rows.Scan(&i.Year, &i.TotalSales); err != nil {
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

const getYearlyTotalSalesCashier = `-- name: GetYearlyTotalSalesCashier :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_sales
    FROM
        orders o
    JOIN
        cashiers c ON o.cashier_id = c.cashier_id
    WHERE
        o.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1::integer AS year
    UNION
    SELECT $1::integer - 1 AS year
)
SELECT 
    a.year::text AS year,
    COALESCE(yd.total_sales, 0) AS total_sales
FROM 
    all_years a
LEFT JOIN 
    yearly_data yd ON a.year = yd.year
ORDER BY 
    a.year DESC
`

type GetYearlyTotalSalesCashierRow struct {
	Year       string `json:"year"`
	TotalSales int32  `json:"total_sales"`
}

func (q *Queries) GetYearlyTotalSalesCashier(ctx context.Context, dollar_1 int32) ([]*GetYearlyTotalSalesCashierRow, error) {
	rows, err := q.db.QueryContext(ctx, getYearlyTotalSalesCashier, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetYearlyTotalSalesCashierRow
	for rows.Next() {
		var i GetYearlyTotalSalesCashierRow
		if err := rows.Scan(&i.Year, &i.TotalSales); err != nil {
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

const restoreAllCashiers = `-- name: RestoreAllCashiers :exec
UPDATE cashiers
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL
`

// Restore All Trashed Cashier
func (q *Queries) RestoreAllCashiers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, restoreAllCashiers)
	return err
}

const restoreCashier = `-- name: RestoreCashier :one
UPDATE cashiers
SET
    deleted_at = NULL
WHERE
    cashier_id = $1
    AND deleted_at IS NOT NULL
  RETURNING cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at
`

// Restore Trashed Cashier
func (q *Queries) RestoreCashier(ctx context.Context, cashierID int32) (*Cashier, error) {
	row := q.db.QueryRowContext(ctx, restoreCashier, cashierID)
	var i Cashier
	err := row.Scan(
		&i.CashierID,
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const trashCashier = `-- name: TrashCashier :one
UPDATE cashiers
SET
    deleted_at = current_timestamp
WHERE
    cashier_id = $1
    AND deleted_at IS NULL
    RETURNING cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at
`

// Trash Cashier
func (q *Queries) TrashCashier(ctx context.Context, cashierID int32) (*Cashier, error) {
	row := q.db.QueryRowContext(ctx, trashCashier, cashierID)
	var i Cashier
	err := row.Scan(
		&i.CashierID,
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updateCashier = `-- name: UpdateCashier :one
UPDATE cashiers
SET name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE cashier_id = $1
  AND deleted_at IS NULL
  RETURNING cashier_id, merchant_id, user_id, name, created_at, updated_at, deleted_at
`

type UpdateCashierParams struct {
	CashierID int32  `json:"cashier_id"`
	Name      string `json:"name"`
}

func (q *Queries) UpdateCashier(ctx context.Context, arg UpdateCashierParams) (*Cashier, error) {
	row := q.db.QueryRowContext(ctx, updateCashier, arg.CashierID, arg.Name)
	var i Cashier
	err := row.Scan(
		&i.CashierID,
		&i.MerchantID,
		&i.UserID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
