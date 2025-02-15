package seeder

import (
	"context"
	"database/sql"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/logger"

	"go.uber.org/zap"
)

type transactionSeeder struct {
	db     *db.Queries
	ctx    context.Context
	logger logger.LoggerInterface
}

func NewTransactionSeeder(db *db.Queries, ctx context.Context, logger logger.LoggerInterface) *transactionSeeder {
	return &transactionSeeder{
		db:     db,
		ctx:    ctx,
		logger: logger,
	}
}

func (r *transactionSeeder) Seed() error {
	orders, err := r.db.GetOrders(r.ctx, db.GetOrdersParams{
		Column1: "",
		Limit:   20,
		Offset:  0,
	})

	if err != nil {
		r.logger.Error("Failed to get transactions:", zap.Any("error", err))
		return err
	}

	merchants, err := r.db.GetMerchants(r.ctx, db.GetMerchantsParams{
		Column1: "",
		Limit:   20,
		Offset:  0,
	})

	if err != nil {
		r.logger.Error("Failed to get transactions:", zap.Any("error", err))
		return err
	}

	for i := 0; i < 10; i++ {
		var orderID int
		var merchantID int
		var paymentMethod string
		var amount, changeAmount float64
		var paymentStatus string

		if len(orders) > 0 {

			orderID = int(orders[i%len(orders)].OrderID)
		} else {
			orderID = i + 1
		}

		if len(merchants) > 0 {
			merchantID = int(merchants[i%len(merchants)].MerchantID)
		} else {
			merchantID = i + 1
		}

		paymentMethod = "Credit Card"
		amount = float64(100 + i)
		changeAmount = float64(5 + i)
		paymentStatus = "Completed"

		_, err := r.db.CreateTransactions(r.ctx, db.CreateTransactionsParams{
			OrderID:       int32(orderID),
			PaymentMethod: paymentMethod,
			Amount:        int32(amount),
			ChangeAmount: sql.NullInt32{
				Int32: int32(changeAmount),
				Valid: true,
			},
			PaymentStatus: paymentStatus,
			MerchantID:    int32(merchantID),
		})
		if err != nil {
			r.logger.Error("Failed to create transaction:", zap.Any("error", err))
			return err
		}
	}

	r.logger.Info("Successfully seeded 10 transactions.")
	return nil
}
