-- Get Orders with Pagination and Total Count
-- name: GetOrders :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Orders with Pagination and Total Count
-- name: GetOrdersActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- Get Trashed Orders with Pagination and Total Count
-- name: GetOrdersTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Orders with Pagination and Total Count where merchant_id
-- name: GetOrdersByMerchant :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM orders
WHERE 
    deleted_at IS NULL
    AND ($1::TEXT IS NULL OR order_id::TEXT ILIKE '%' || $1 || '%' OR total_price::TEXT ILIKE '%' || $1 || '%')
    AND ($4::UUID IS NULL OR merchant_id = $4)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;




-- name: GetMonthlyTotalRevenue :many
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
    am.month DESC;


-- name: GetYearlyTotalRevenue :many
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
    ay.year DESC;




-- name: GetMonthlyTotalRevenueById :many
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
    am.month DESC;


-- name: GetYearlyTotalRevenueById :many
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
    ay.year DESC;



-- name: GetMonthlyTotalRevenueByMerchant :many
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
    am.month DESC;


-- name: GetYearlyTotalRevenueByMerchant :many
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
    ay.year DESC;




-- name: GetMonthlyOrder :many
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
    mo.activity_month;



-- name: GetYearlyOrder :many
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
    year;


-- name: GetMonthlyOrderByMerchant :many
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
    mo.activity_month;



-- name: GetYearlyOrderByMerchant :many
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
    year;


-- name: CreateOrder :one
INSERT INTO orders (merchant_id, cashier_id, total_price)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetOrderByID :one
SELECT *
FROM orders
WHERE order_id = $1
  AND deleted_at IS NULL;



-- name: UpdateOrder :one
UPDATE orders
SET total_price = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE order_id = $1
  AND deleted_at IS NULL
  RETURNING *;  

-- Trash Order
-- name: TrashedOrder :one
UPDATE orders
SET
    deleted_at = current_timestamp
WHERE
    order_id = $1
    AND deleted_at IS NULL
    RETURNING *;    


-- Restore Trashed Order
-- name: RestoreOrder :one
UPDATE orders
SET
    deleted_at = NULL
WHERE
    order_id = $1
    AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Order Permanently
-- name: DeleteOrderPermanently :exec
DELETE FROM orders WHERE order_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Order
-- name: RestoreAllOrders :exec
UPDATE orders
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Order Permanently
-- name: DeleteAllPermanentOrders :exec
DELETE FROM orders
WHERE
    deleted_at IS NOT NULL;
