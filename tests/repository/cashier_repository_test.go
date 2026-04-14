package repository_test

import (
	"context"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/repository"
	db "pointofsale/pkg/database/schema"
	"pointofsale/tests"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/suite"
)

type CashierRepositoryTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	repo        repository.CashierRepository
	userRepo    repository.UserRepository
	merchantRepo repository.MerchantRepository
}

func (s *CashierRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := pgxpool.New(s.ts.Ctx, s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := db.New(pool)
	s.repo = repository.NewCashierRepository(queries)
	s.userRepo = repository.NewUserRepository(queries)
	s.merchantRepo = repository.NewMerchantRepository(queries)
}

func (s *CashierRepositoryTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *CashierRepositoryTestSuite) TestCashierLifecycle() {
	ctx := context.Background()

	// 0. Create User and Merchant first
	userReq := &requests.CreateUserRequest{
		FirstName: "Cashier",
		LastName:  "User",
		Email:     "cashier@example.com",
		Password:  "password123",
	}
	user, err := s.userRepo.CreateUser(ctx, userReq)
	s.NoError(err)
	s.NotNil(user)
	userID := int(user.UserID)

	merchantReq := &requests.CreateMerchantRequest{
		UserID:       userID,
		Name:         "Cashier Merchant",
		Description:  "Merchant for cashier testing",
		Address:      "Jakarta, Indonesia",
		ContactEmail: "merchant@example.com",
		ContactPhone: "08123456789",
		Status:       "active",
	}
	merchant, err := s.merchantRepo.CreateMerchant(ctx, merchantReq)
	s.NoError(err)
	s.NotNil(merchant)
	merchantID := int(merchant.MerchantID)

	// 1. Create Cashier
	createReq := &requests.CreateCashierRequest{
		MerchantID: merchantID,
		UserID:     userID,
		Name:       "Morning Cashier",
	}

	cashier, err := s.repo.CreateCashier(ctx, createReq)
	s.NoError(err)
	s.NotNil(cashier)
	s.Equal(createReq.Name, cashier.Name)

	cashierID := int(cashier.CashierID)

	// 2. Find By ID
	found, err := s.repo.FindById(ctx, cashierID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(cashier.Name, found.Name)

	// 3. Update Cashier
	updateReq := &requests.UpdateCashierRequest{
		CashierID: &cashierID,
		Name:      "Evening Cashier",
	}

	updated, err := s.repo.UpdateCashier(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(updateReq.Name, updated.Name)

	// 4. Find All
	all, err := s.repo.FindAllCashiers(ctx, &requests.FindAllCashiers{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotEmpty(all)

	// 5. Trash Cashier
	trashed, err := s.repo.TrashedCashier(ctx, cashierID)
	s.NoError(err)
	s.NotNil(trashed)
	s.True(trashed.DeletedAt.Valid)

	// 6. Find By Trashed
	trashedList, err := s.repo.FindByTrashed(ctx, &requests.FindAllCashiers{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotEmpty(trashedList)

	// 7. Restore Cashier
	restored, err := s.repo.RestoreCashier(ctx, cashierID)
	s.NoError(err)
	s.NotNil(restored)
	s.False(restored.DeletedAt.Valid)

	// 8. Delete Permanent
	// Trash again first
	_, err = s.repo.TrashedCashier(ctx, cashierID)
	s.NoError(err)

	success, err := s.repo.DeleteCashierPermanent(ctx, cashierID)
	s.NoError(err)
	s.True(success)

	// 9. Verify it's gone
	_, err = s.repo.FindById(ctx, cashierID)
	s.Error(err)
}

func TestCashierRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CashierRepositoryTestSuite))
}
