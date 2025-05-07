package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"pointofsale/internal/domain/record"
	"pointofsale/internal/domain/requests"
	recordmapper "pointofsale/internal/mapper/record"
	db "pointofsale/pkg/database/schema"
	"time"
)

type refreshTokenRepository struct {
	db      *db.Queries
	ctx     context.Context
	mapping recordmapper.RefreshTokenRecordMapping
}

func NewRefreshTokenRepository(db *db.Queries, ctx context.Context, mapping recordmapper.RefreshTokenRecordMapping) *refreshTokenRepository {
	return &refreshTokenRepository{
		db:      db,
		ctx:     ctx,
		mapping: mapping,
	}
}

func (r *refreshTokenRepository) FindByToken(token string) (*record.RefreshTokenRecord, error) {
	res, err := r.db.FindRefreshTokenByToken(r.ctx, token)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("refresh token not found or expired: invalid token %s", token)
		}
		return nil, fmt.Errorf("failed to find refresh token: %w", err)
	}

	return r.mapping.ToRefreshTokenRecord(res), nil
}

func (r *refreshTokenRepository) FindByUserId(user_id int) (*record.RefreshTokenRecord, error) {
	res, err := r.db.FindRefreshTokenByUserId(r.ctx, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no refresh token found for user ID %d: %w", user_id, err)
		}
		return nil, fmt.Errorf("failed to retrieve refresh token for user ID %d: %w", user_id, err)
	}

	return r.mapping.ToRefreshTokenRecord(res), nil
}

func (r *refreshTokenRepository) CreateRefreshToken(req *requests.CreateRefreshToken) (*record.RefreshTokenRecord, error) {
	layout := "2006-01-02 15:04:05"
	expirationTime, err := time.Parse(layout, req.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expiration date: %w", err)
	}

	res, err := r.db.CreateRefreshToken(r.ctx, db.CreateRefreshTokenParams{
		UserID:     int32(req.UserId),
		Token:      req.Token,
		Expiration: expirationTime,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create refresh token: invalid or incomplete token data")
	}

	return r.mapping.ToRefreshTokenRecord(res), nil
}

func (r *refreshTokenRepository) UpdateRefreshToken(req *requests.UpdateRefreshToken) (*record.RefreshTokenRecord, error) {
	layout := "2006-01-02 15:04:05"
	expirationTime, err := time.Parse(layout, req.ExpiresAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse expiration date: %w", err)
	}

	res, err := r.db.UpdateRefreshTokenByUserId(r.ctx, db.UpdateRefreshTokenByUserIdParams{
		UserID:     int32(req.UserId),
		Token:      req.Token,
		Expiration: expirationTime,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("refresh token not found for user ID %d: %w", req.UserId, err)
		}
		return nil, fmt.Errorf("failed to update refresh token: %w", err)
	}

	return r.mapping.ToRefreshTokenRecord(res), nil
}

func (r *refreshTokenRepository) DeleteRefreshToken(token string) error {
	err := r.db.DeleteRefreshToken(r.ctx, token)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("refresh token with token %s not found or already expired: %w", token, err)
		}
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}

	return nil
}

func (r *refreshTokenRepository) DeleteRefreshTokenByUserId(user_id int) error {
	err := r.db.DeleteRefreshTokenByUserId(r.ctx, int32(user_id))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("no refresh tokens found for user ID %d: %w", user_id, err)
		}
		return fmt.Errorf("failed to delete refresh tokens for user ID %d: %w", user_id, err)
	}

	return nil
}
