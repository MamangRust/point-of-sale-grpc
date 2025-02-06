-- Get Users with Pagination and Total Count
-- name: GetUsers :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get User by ID
-- name: GetUserByID :one
SELECT * FROM users WHERE user_id = $1 AND deleted_at IS NULL;


-- Get User by Email
-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL;



-- Get Trashed By User ID
-- name: GetTrashedUserByID :one
SELECT *
FROM users
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL;


-- Get Active Users with Pagination and Total Count
-- name: GetUsersActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;



-- Get Trashed Users with Pagination and Total Count
-- name: GetUserTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;



-- Create User
-- name: CreateUser :one
INSERT INTO
    users (
        firstname,
        lastname,
        email,
        password,
        created_at,
        updated_at
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        current_timestamp,
        current_timestamp
    ) RETURNING *;



-- Update User
-- name: UpdateUser :exec
UPDATE users
SET
    firstname = $2,
    lastname = $3,
    email = $4,
    password = $5,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL;



-- Trash User
-- name: TrashUser :exec
UPDATE users
SET
    deleted_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL;

-- Restore Trashed User
-- name: RestoreUser :exec
UPDATE users
SET
    deleted_at = NULL
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL;


-- Delete User Permanently
-- name: DeleteUserPermanently :exec
DELETE FROM users WHERE user_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Users
-- name: RestoreAllUsers :exec
UPDATE users
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;


-- Delete All Trashed Users Permanently
-- name: DeleteAllPermanentUsers :exec
DELETE FROM users
WHERE
    deleted_at IS NOT NULL;
