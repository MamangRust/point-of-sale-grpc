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
