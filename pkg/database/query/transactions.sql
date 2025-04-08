-- name: GetTransactions :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%' OR payment_status ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Transactions with Pagination and Total Count
-- name: GetTransactionsActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%' OR payment_status ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- Get Trashed Transactions with Pagination and Total Count
-- name: GetTransactionsTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%' OR payment_status ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetTransactionByMerchant :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM transactions
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR payment_method ILIKE '%' || $1 || '%' OR payment_status ILIKE '%' || $1 || '%')
  AND ($2::INT IS NULL OR merchant_id = $2)
ORDER BY created_at DESC
LIMIT $3 OFFSET $4;



-- name: GetMonthlyAmountTransactionSuccess :many
WITH monthly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        EXTRACT(MONTH FROM t.created_at)::integer AS month,
        COUNT(*) AS total_success,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND (
            (t.created_at >= $1::timestamp AND t.created_at <= $2::timestamp)
            OR (t.created_at >= $3::timestamp AND t.created_at <= $4::timestamp)
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at),
        EXTRACT(MONTH FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        TO_CHAR(TO_DATE(month::text, 'MM'), 'Mon') AS month,
        total_success,
        total_amount
    FROM
        monthly_data

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $1::timestamp)::text AS year,
        TO_CHAR($1::timestamp, 'Mon') AS month,
        0 AS total_success,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $1::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $1::timestamp)::integer
    )

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $3::timestamp)::text AS year,
        TO_CHAR($3::timestamp, 'Mon') AS month,
        0 AS total_success,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $3::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $3::timestamp)::integer
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC,
    TO_DATE(month, 'Mon') DESC;


-- name: GetYearlyAmountTransactionSuccess :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        COUNT(*) AS total_success,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND (
            EXTRACT(YEAR FROM t.created_at) = $1::integer
            OR EXTRACT(YEAR FROM t.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        total_success::integer,
        total_amount::integer
    FROM
        yearly_data

    UNION ALL

    SELECT
        $1::text AS year,
        0::integer AS total_success,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer
    )

    UNION ALL

    SELECT
        ($1::integer - 1)::text AS year,
        0::integer AS total_success,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer - 1
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC;

-- name: GetMonthlyAmountTransactionFailed :many
WITH monthly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        EXTRACT(MONTH FROM t.created_at)::integer AS month,
        COUNT(*) AS total_failed,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'failed'
        AND (
            (t.created_at >= $1::timestamp AND t.created_at <= $2::timestamp)
            OR (t.created_at >= $3::timestamp AND t.created_at <= $4::timestamp)
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at),
        EXTRACT(MONTH FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        TO_CHAR(TO_DATE(month::text, 'MM'), 'Mon') AS month,
        total_failed,
        total_amount
    FROM
        monthly_data

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $1::timestamp)::text AS year,
        TO_CHAR($1::timestamp, 'Mon') AS month,
        0 AS total_failed,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $1::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $1::timestamp)::integer
    )

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $3::timestamp)::text AS year,
        TO_CHAR($3::timestamp, 'Mon') AS month,
        0 AS total_failed,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $3::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $3::timestamp)::integer
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC,
    TO_DATE(month, 'Mon') DESC;



-- name: GetYearlyAmountTransactionFailed :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        COUNT(*) AS total_failed,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'failed'
        AND (
            EXTRACT(YEAR FROM t.created_at) = $1::integer
            OR EXTRACT(YEAR FROM t.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        total_failed::integer,
        total_amount::integer
    FROM
        yearly_data

    UNION ALL

    SELECT
        $1::text AS year,
        0::integer AS total_failed,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer
    )

    UNION ALL

    SELECT
        ($1::integer - 1)::text AS year,
        0::integer AS total_failed,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer - 1
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC;



-- name: GetMonthlyTransactionMethods :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
payment_methods AS (
    SELECT DISTINCT payment_method
    FROM transactions
    WHERE deleted_at IS NULL
),
monthly_transactions AS (
    SELECT
        date_trunc('month', t.created_at) AS activity_month,
        t.payment_method,
        COUNT(t.transaction_id) AS total_transactions,
        SUM(t.amount)::NUMERIC AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
    GROUP BY
        activity_month, t.payment_method
)
SELECT
    TO_CHAR(mt.activity_month, 'Mon') AS month,
    mt.payment_method,
    mt.total_transactions,
    mt.total_amount
FROM
    monthly_transactions mt
ORDER BY
    mt.activity_month,
    mt.payment_method;



-- name: GetYearlyTransactionMethods :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::text AS year,
        t.payment_method,
        COUNT(t.transaction_id) AS total_transactions,
        SUM(t.amount)::NUMERIC AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND EXTRACT(YEAR FROM t.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
    GROUP BY
        EXTRACT(YEAR FROM t.created_at),
        t.payment_method
)
SELECT
    year,
    payment_method,
    total_transactions,
    total_amount
FROM
    last_five_years
ORDER BY
    year,
    payment_method;



-- name: GetMonthlyAmountTransactionSuccessByMerchant :many
WITH monthly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        EXTRACT(MONTH FROM t.created_at)::integer AS month,
        COUNT(*) AS total_success,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND t.merchant_id = $5
        AND (
            (t.created_at >= $1::timestamp AND t.created_at <= $2::timestamp)
            OR (t.created_at >= $3::timestamp AND t.created_at <= $4::timestamp)
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at),
        EXTRACT(MONTH FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        TO_CHAR(TO_DATE(month::text, 'MM'), 'Mon') AS month,
        total_success,
        total_amount
    FROM
        monthly_data

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $1::timestamp)::text AS year,
        TO_CHAR($1::timestamp, 'Mon') AS month,
        0 AS total_success,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $1::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $1::timestamp)::integer
    )

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $3::timestamp)::text AS year,
        TO_CHAR($3::timestamp, 'Mon') AS month,
        0 AS total_success,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $3::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $3::timestamp)::integer
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC,
    TO_DATE(month, 'Mon') DESC;


-- name: GetYearlyAmountTransactionSuccessByMerchant :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        COUNT(*) AS total_success,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND t.merchant_id = $2
        AND (
            EXTRACT(YEAR FROM t.created_at) = $1::integer
            OR EXTRACT(YEAR FROM t.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        total_success::integer,
        total_amount::integer
    FROM
        yearly_data

    UNION ALL

    SELECT
        $1::text AS year,
        0::integer AS total_success,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer
    )

    UNION ALL

    SELECT
        ($1::integer - 1)::text AS year,
        0::integer AS total_success,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer - 1
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC;



-- name: GetMonthlyAmountTransactionFailedByMerchant :many
WITH monthly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        EXTRACT(MONTH FROM t.created_at)::integer AS month,
        COUNT(*) AS total_failed,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'failed'
        AND t.merchant_id = $5
        AND (
            (t.created_at >= $1::timestamp AND t.created_at <= $2::timestamp)
            OR (t.created_at >= $3::timestamp AND t.created_at <= $4::timestamp)
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at),
        EXTRACT(MONTH FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        TO_CHAR(TO_DATE(month::text, 'MM'), 'Mon') AS month,
        total_failed,
        total_amount
    FROM
        monthly_data

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $1::timestamp)::text AS year,
        TO_CHAR($1::timestamp, 'Mon') AS month,
        0 AS total_failed,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $1::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $1::timestamp)::integer
    )

    UNION ALL

    SELECT
        EXTRACT(YEAR FROM $3::timestamp)::text AS year,
        TO_CHAR($3::timestamp, 'Mon') AS month,
        0 AS total_failed,
        0 AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM monthly_data
        WHERE year = EXTRACT(YEAR FROM $3::timestamp)::integer
        AND month = EXTRACT(MONTH FROM $3::timestamp)::integer
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC,
    TO_DATE(month, 'Mon') DESC;

-- name: GetYearlyAmountTransactionFailedByMerchant :many
WITH yearly_data AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::integer AS year,
        COUNT(*) AS total_failed,
        COALESCE(SUM(t.amount), 0)::integer AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'failed'
        AND t.merchant_id = $2
        AND (
            EXTRACT(YEAR FROM t.created_at) = $1::integer
            OR EXTRACT(YEAR FROM t.created_at) = $1::integer - 1
        )
    GROUP BY
        EXTRACT(YEAR FROM t.created_at)
), formatted_data AS (
    SELECT
        year::text,
        total_failed::integer,
        total_amount::integer
    FROM
        yearly_data

    UNION ALL

    SELECT
        $1::text AS year,
        0::integer AS total_failed,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer
    )

    UNION ALL

    SELECT
        ($1::integer - 1)::text AS year,
        0::integer AS total_failed,
        0::integer AS total_amount
    WHERE NOT EXISTS (
        SELECT 1
        FROM yearly_data
        WHERE year = $1::integer - 1
    )
)
SELECT * FROM formatted_data
ORDER BY
    year DESC;



-- name: GetMonthlyTransactionMethodsByMerchant :many
WITH date_range AS (
    SELECT 
        date_trunc('month', $1::timestamp) AS start_date,
        date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
),
payment_methods AS (
    SELECT DISTINCT payment_method
    FROM transactions
    WHERE deleted_at IS NULL
),
monthly_transactions AS (
    SELECT
        date_trunc('month', t.created_at) AS activity_month,
        t.payment_method,
        COUNT(t.transaction_id) AS total_transactions,
        SUM(t.amount)::NUMERIC AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                             AND (SELECT end_date FROM date_range)
        AND t.merchant_id = $2
    GROUP BY
        activity_month, t.payment_method
)
SELECT
    TO_CHAR(mt.activity_month, 'Mon') AS month,
    mt.payment_method,
    mt.total_transactions,
    mt.total_amount
FROM
    monthly_transactions mt
ORDER BY
    mt.activity_month,
    mt.payment_method;



-- name: GetYearlyTransactionMethodsByMerchant :many
WITH last_five_years AS (
    SELECT
        EXTRACT(YEAR FROM t.created_at)::text AS year,
        t.payment_method,
        COUNT(t.transaction_id) AS total_transactions,
        SUM(t.amount)::NUMERIC AS total_amount
    FROM
        transactions t
    WHERE
        t.deleted_at IS NULL
        AND t.payment_status = 'success'
        AND EXTRACT(YEAR FROM t.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
        AND t.merchant_id = $2
    GROUP BY
        EXTRACT(YEAR FROM t.created_at),
        t.payment_method
)
SELECT
    year,
    payment_method,
    total_transactions,
    total_amount
FROM
    last_five_years
ORDER BY
    year,
    payment_method;


-- name: CreateTransactions :one
INSERT INTO transactions (order_id, merchant_id, payment_method, amount, change_amount, payment_status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: GetTransactionByOrderID :one
SELECT *
FROM transactions
WHERE order_id = $1
  AND deleted_at IS NULL;

-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE transaction_id = $1
  AND deleted_at IS NULL;


-- name: UpdateTransaction :one
UPDATE transactions
SET merchant_id = $2,
    payment_method = $3,
    amount = $4,
    change_amount = $5,
    payment_status = $6,
    order_id = $7,
    updated_at = CURRENT_TIMESTAMP
WHERE transaction_id = $1
  AND deleted_at IS NULL
RETURNING *;

-- Trash Transaction
-- name: TrashTransaction :one
UPDATE transactions
SET
    deleted_at = current_timestamp
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
    RETURNING *;    


-- Restore Trashed Transaction
-- name: RestoreTransaction :one
UPDATE transactions
SET
    deleted_at = NULL
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Transaction Permanently
-- name: DeleteTransactionPermanently :exec
DELETE FROM transactions WHERE transaction_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Transaction
-- name: RestoreAllTransactions :exec
UPDATE transactions
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Transaction Permanently
-- name: DeleteAllPermanentTransactions :exec
DELETE FROM transactions
WHERE
    deleted_at IS NOT NULL;



