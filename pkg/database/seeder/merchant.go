package seeder

import (
	"context"
	"fmt"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type merchantSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewMerchantSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *merchantSeeder {
	return &merchantSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *merchantSeeder) Seed() error {
	users, err := r.db.GetUsers(r.ctx, db.GetUsersParams{
		Column1: "",
		Limit:   int32(20),
		Offset:  0,
	})
	if err != nil {
		r.logger.Error("Failed to fetch merchants:", zap.Any("error", err))
		return err
	}

	for i := 1; i <= 10; i++ {
		userID := users[i%len(users)].UserID

		desc := fmt.Sprintf("Deskripsi untuk Toko %d", i)
		addr := fmt.Sprintf("Jl. Toko %d", i)
		email := fmt.Sprintf("toko%d@example.com", i)
		phone := fmt.Sprintf("0812345678%d", i)

		merchant := db.CreateMerchantParams{
			UserID:       userID,
			Name:         fmt.Sprintf("Toko %d", i),
			Description:  &desc,
			Address:      &addr,
			ContactEmail: &email,
			ContactPhone: &phone,
			Status:       "active",
		}

		_, err = r.db.CreateMerchant(r.ctx, merchant)
		if err != nil {
			r.logger.Error("Failed to create merchant:", zap.Any("error", err))
			return err
		}
	}

	r.logger.Info("Merchant seeding completed successfully.")
	return nil
}
