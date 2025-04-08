-- Get Categories with Pagination and Total Count
-- name: GetCategories :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR slug_category ILIKE '%' || $1 || '%') --
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Categories with Pagination and Total Count
-- name: GetCategoriesActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR slug_category ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- Get Trashed Categories with Pagination and Total Count
-- name: GetCategoriesTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM categories
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%' OR slug_category ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;



-- name: GetMonthlyTotalPrice :many
WITH monthly_totals AS (
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
    COALESCE(mt.total_revenue, 0) AS total_revenue
FROM 
    all_months am
LEFT JOIN 
    monthly_totals mt ON am.year = mt.year AND am.month = mt.month
ORDER BY 
    am.year::INT DESC,
    am.month DESC;



-- name: GetYearlyTotalPrice :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
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
    a.year::text AS year,
    COALESCE(yd.total_revenue, 0) AS total_revenue
FROM 
    all_years a
LEFT JOIN 
    yearly_data yd ON a.year = yd.year
ORDER BY 
    a.year DESC;



-- name: GetMonthlyTotalPriceByMerchant :many
WITH monthly_totals AS (
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
    COALESCE(mt.total_revenue, 0) AS total_revenue
FROM 
    all_months am
LEFT JOIN 
    monthly_totals mt ON am.year = mt.year AND am.month = mt.month
ORDER BY 
    am.year::INT DESC,
    am.month DESC;



-- name: GetYearlyTotalPriceByMerchant :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
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
    SELECT $1 AS year
    UNION
    SELECT $1 - 1 AS year
)
SELECT 
    a.year::text AS year,
    COALESCE(yd.total_revenue, 0) AS total_revenue
FROM 
    all_years a
LEFT JOIN 
    yearly_data yd ON a.year = yd.year
ORDER BY 
    a.year DESC;



-- name: GetMonthlyTotalPriceById :many
WITH monthly_totals AS (
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
    COALESCE(mt.total_revenue, 0) AS total_revenue
FROM 
    all_months am
LEFT JOIN 
    monthly_totals mt ON am.year = mt.year AND am.month = mt.month
ORDER BY 
    am.year::INT DESC,
    am.month DESC;



-- name: GetYearlyTotalPriceById :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM o.created_at)::integer AS year,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND (
            EXTRACT(YEAR FROM o.created_at) = $1::integer
            OR EXTRACT(YEAR FROM o.created_at) = $1::integer - 1
        )
        AND o.order_id = $2
    GROUP BY
        EXTRACT(YEAR FROM o.created_at)
),
all_years AS (
    SELECT $1 AS year
    UNION
    SELECT $1 - 1 AS year
)
SELECT 
    a.year::text AS year,
    COALESCE(yd.total_revenue, 0) AS total_revenue
FROM 
    all_years a
LEFT JOIN 
    yearly_data yd ON a.year = yd.year
ORDER BY 
    a.year DESC;




-- name: GetMonthlyCategory :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
monthly_category_stats AS (
    SELECT
        c.category_id,
        c.name AS category_name,
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(DISTINCT o.order_id) AS order_count,
        SUM(oi.quantity) AS items_sold,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
    GROUP BY
        c.category_id, c.name, activity_month
)
SELECT
    TO_CHAR(mcs.activity_month, 'Mon') AS month,
    mcs.category_id,
    mcs.category_name,
    mcs.order_count,
    mcs.items_sold,
    mcs.total_revenue
FROM
    monthly_category_stats mcs
ORDER BY
    mcs.activity_month, mcs.total_revenue DESC;




-- name: GetYearlyCategory :many
WITH last_five_years AS (
    SELECT
        c.category_id,
        c.name AS category_name,
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(DISTINCT o.order_id) AS order_count,
        SUM(oi.quantity) AS items_sold,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue,
        COUNT(DISTINCT oi.product_id) AS unique_products_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
    GROUP BY
        c.category_id, c.name, EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    category_id,
    category_name,
    order_count,
    items_sold,
    total_revenue,
    unique_products_sold
FROM
    last_five_years
ORDER BY
    year, total_revenue DESC;


-- name: GetMonthlyCategoryByMerchant :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
monthly_category_stats AS (
    SELECT
        c.category_id,
        c.name AS category_name,
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(DISTINCT o.order_id) AS order_count,
        SUM(oi.quantity) AS items_sold,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
        AND o.merchant_id = $2
    GROUP BY
        c.category_id, c.name, activity_month
)
SELECT
    TO_CHAR(mcs.activity_month, 'Mon') AS month,
    mcs.category_id,
    mcs.category_name,
    mcs.order_count,
    mcs.items_sold,
    mcs.total_revenue
FROM
    monthly_category_stats mcs
ORDER BY
    mcs.activity_month, mcs.total_revenue DESC;



-- name: GetYearlyCategoryByMerchant :many
WITH last_five_years AS (
    SELECT
        c.category_id,
        c.name AS category_name,
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(DISTINCT o.order_id) AS order_count,
        SUM(oi.quantity) AS items_sold,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue,
        COUNT(DISTINCT oi.product_id) AS unique_products_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
        AND o.merchant_id = $2
    GROUP BY
        c.category_id, c.name, EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    category_id,
    category_name,
    order_count,
    items_sold,
    total_revenue,
    unique_products_sold
FROM
    last_five_years
ORDER BY
    year, total_revenue DESC;


-- name: GetMonthlyCategoryById :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
monthly_category_stats AS (
    SELECT
        c.category_id,
        c.name AS category_name,
        date_trunc('month', o.created_at) AS activity_month,
        COUNT(DISTINCT o.order_id) AS order_count,
        SUM(oi.quantity) AS items_sold,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND o.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
        AND c.category_id = $2
    GROUP BY
        c.category_id, c.name, activity_month
)
SELECT
    TO_CHAR(mcs.activity_month, 'Mon') AS month,
    mcs.category_id,
    mcs.category_name,
    mcs.order_count,
    mcs.items_sold,
    mcs.total_revenue
FROM
    monthly_category_stats mcs
ORDER BY
    mcs.activity_month, mcs.total_revenue DESC;



-- name: GetYearlyCategoryById :many
WITH last_five_years AS (
    SELECT
        c.category_id,
        c.name AS category_name,
        EXTRACT(YEAR FROM o.created_at)::text AS year,
        COUNT(DISTINCT o.order_id) AS order_count,
        SUM(oi.quantity) AS items_sold,
        COALESCE(SUM(o.total_price), 0)::INTEGER AS total_revenue,
        COUNT(DISTINCT oi.product_id) AS unique_products_sold
    FROM
        orders o
    JOIN
        order_items oi ON o.order_id = oi.order_id
    JOIN
        products p ON oi.product_id = p.product_id
    JOIN
        categories c ON p.category_id = c.category_id
    WHERE
        o.deleted_at IS NULL
        AND oi.deleted_at IS NULL
        AND p.deleted_at IS NULL
        AND c.deleted_at IS NULL
        AND EXTRACT(YEAR FROM o.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
        AND c.category_id = $2
    GROUP BY
        c.category_id, c.name, EXTRACT(YEAR FROM o.created_at)
)
SELECT
    year,
    category_id,
    category_name,
    order_count,
    items_sold,
    total_revenue,
    unique_products_sold
FROM
    last_five_years
ORDER BY
    year, total_revenue DESC;


-- name: CreateCategory :one
INSERT INTO categories (name, description, slug_category)
VALUES ($1, $2, $3)
  RETURNING *;

-- name: GetCategoryByID :one
SELECT *
FROM categories
WHERE category_id = $1
  AND deleted_at IS NULL;


-- name: UpdateCategory :one
UPDATE categories
SET name = $2,
    description = $3,
    slug_category = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE category_id = $1
  AND deleted_at IS NULL
  RETURNING *;

-- Trash Category
-- name: TrashCategory :one
UPDATE categories
SET
    deleted_at = current_timestamp
WHERE
    category_id = $1
    AND deleted_at IS NULL
    RETURNING *;    


-- Restore Trashed Category
-- name: RestoreCategory :one
UPDATE categories
SET
    deleted_at = NULL
WHERE
    category_id = $1
    AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Category Permanently
-- name: DeleteCategoryPermanently :exec
DELETE FROM categories WHERE category_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Category
-- name: RestoreAllCategories :exec
UPDATE categories
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Category Permanently
-- name: DeleteAllPermanentCategories :exec
DELETE FROM categories
WHERE
    deleted_at IS NOT NULL;