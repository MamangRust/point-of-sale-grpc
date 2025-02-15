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
