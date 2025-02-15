// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
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
    ) RETURNING user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Create User
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const deleteAllPermanentUsers = `-- name: DeleteAllPermanentUsers :exec
DELETE FROM users
WHERE
    deleted_at IS NOT NULL
`

// Delete All Trashed Users Permanently
func (q *Queries) DeleteAllPermanentUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteAllPermanentUsers)
	return err
}

const deleteUserPermanently = `-- name: DeleteUserPermanently :exec
DELETE FROM users WHERE user_id = $1 AND deleted_at IS NOT NULL
`

// Delete User Permanently
func (q *Queries) DeleteUserPermanently(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserPermanently, userID)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at FROM users WHERE email = $1 AND deleted_at IS NULL
`

// Get User by Email
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1 AND deleted_at IS NULL
`

// Get User by ID
func (q *Queries) GetUserByID(ctx context.Context, userID int32) (*User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const getUserTrashed = `-- name: GetUserTrashed :many
SELECT
    user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NOT NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetUserTrashedParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetUserTrashedRow struct {
	UserID     int32        `json:"user_id"`
	Firstname  string       `json:"firstname"`
	Lastname   string       `json:"lastname"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Trashed Users with Pagination and Total Count
func (q *Queries) GetUserTrashed(ctx context.Context, arg GetUserTrashedParams) ([]*GetUserTrashedRow, error) {
	rows, err := q.db.QueryContext(ctx, getUserTrashed, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetUserTrashedRow
	for rows.Next() {
		var i GetUserTrashedRow
		if err := rows.Scan(
			&i.UserID,
			&i.Firstname,
			&i.Lastname,
			&i.Email,
			&i.Password,
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

const getUsers = `-- name: GetUsers :many
SELECT
    user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetUsersParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetUsersRow struct {
	UserID     int32        `json:"user_id"`
	Firstname  string       `json:"firstname"`
	Lastname   string       `json:"lastname"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Users with Pagination and Total Count
func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]*GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.UserID,
			&i.Firstname,
			&i.Lastname,
			&i.Email,
			&i.Password,
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

const getUsersActive = `-- name: GetUsersActive :many
SELECT
    user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at,
    COUNT(*) OVER() AS total_count
FROM users
WHERE deleted_at IS NULL
  AND ($1::TEXT IS NULL OR firstname ILIKE '%' || $1 || '%' OR lastname ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type GetUsersActiveParams struct {
	Column1 string `json:"column_1"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
}

type GetUsersActiveRow struct {
	UserID     int32        `json:"user_id"`
	Firstname  string       `json:"firstname"`
	Lastname   string       `json:"lastname"`
	Email      string       `json:"email"`
	Password   string       `json:"password"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at"`
	TotalCount int64        `json:"total_count"`
}

// Get Active Users with Pagination and Total Count
func (q *Queries) GetUsersActive(ctx context.Context, arg GetUsersActiveParams) ([]*GetUsersActiveRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsersActive, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*GetUsersActiveRow
	for rows.Next() {
		var i GetUsersActiveRow
		if err := rows.Scan(
			&i.UserID,
			&i.Firstname,
			&i.Lastname,
			&i.Email,
			&i.Password,
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

const restoreAllUsers = `-- name: RestoreAllUsers :exec
UPDATE users
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL
`

// Restore All Trashed Users
func (q *Queries) RestoreAllUsers(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, restoreAllUsers)
	return err
}

const restoreUser = `-- name: RestoreUser :one
UPDATE users
SET
    deleted_at = NULL
WHERE
    user_id = $1
    AND deleted_at IS NOT NULL
    RETURNING user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at
`

// Restore Trashed User
func (q *Queries) RestoreUser(ctx context.Context, userID int32) (*User, error) {
	row := q.db.QueryRowContext(ctx, restoreUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const trashUser = `-- name: TrashUser :one
UPDATE users
SET
    deleted_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
    RETURNING user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at
`

// Trash User
func (q *Queries) TrashUser(ctx context.Context, userID int32) (*User, error) {
	row := q.db.QueryRowContext(ctx, trashUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    firstname = $2,
    lastname = $3,
    email = $4,
    password = $5,
    updated_at = current_timestamp
WHERE
    user_id = $1
    AND deleted_at IS NULL
    RETURNING user_id, firstname, lastname, email, password, created_at, updated_at, deleted_at
`

type UpdateUserParams struct {
	UserID    int32  `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Update User
func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UserID,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return &i, err
}
