-- name: GetCashiers :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetCashiersActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetCashiersTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetCashiersByMerchant :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM cashiers
WHERE merchant_id = $1
  AND deleted_at IS NULL
  AND ($2::TEXT IS NULL OR name ILIKE '%' || $2 || '%')
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;



-- name: GetMonthlyTotalSalesCashier :many
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
    am.month DESC;




-- name: GetYearlyTotalSalesCashier :many
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
    a.year DESC;




-- name: GetMonthlyTotalSalesByMerchant :many
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
    am.month DESC;



-- name: GetYearlyTotalSalesByMerchant :many
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
    a.year DESC;




-- name: GetMonthlyTotalSalesById :many
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
    am.month DESC;



-- name: GetYearlyTotalSalesById :many
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
    a.year DESC;




-- name: GetMonthlyCashier :many
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
    ca.activity_month, ca.cashier_id;



-- name: GetYearlyCashier :many
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
    year, cashier_id;




-- name: GetMonthlyCashierByCashierId :many
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
    ca.activity_month, ca.cashier_id;



-- name: GetYearlyCashierByCashierId :many
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
    year, cashier_id;


-- name: GetMonthlyCashierByMerchant :many
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
    ca.activity_month, ca.cashier_id;



-- name: GetYearlyCashierByMerchant :many
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
    year, cashier_id;



-- name: CreateCashier :one
INSERT INTO cashiers (merchant_id, user_id, name)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCashierByID :one
SELECT *
FROM cashiers
WHERE cashier_id = $1
  AND deleted_at IS NULL;


-- name: UpdateCashier :one
UPDATE cashiers
SET name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE cashier_id = $1
  AND deleted_at IS NULL
  RETURNING *;
  

-- Trash Cashier
-- name: TrashCashier :one
UPDATE cashiers
SET
    deleted_at = current_timestamp
WHERE
    cashier_id = $1
    AND deleted_at IS NULL
    RETURNING *;    


-- Restore Trashed Cashier
-- name: RestoreCashier :one
UPDATE cashiers
SET
    deleted_at = NULL
WHERE
    cashier_id = $1
    AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Cashier Permanently
-- name: DeleteCashierPermanently :exec
DELETE FROM cashiers WHERE cashier_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Cashier
-- name: RestoreAllCashiers :exec
UPDATE cashiers
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Cashier Permanently
-- name: DeleteAllPermanentCashiers :exec
DELETE FROM cashiers
WHERE
    deleted_at IS NOT NULL;
