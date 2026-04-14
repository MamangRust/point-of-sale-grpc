package service_test

import (
	"context"
	"pointofsale/internal/cache"
	cashier_cache "pointofsale/internal/cache/cashier"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/repository"
	"pointofsale/internal/service"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/tests"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

type CashierServiceTestSuite struct {
	suite.Suite
	ts             *tests.TestSuite
	dbPool         *pgxpool.Pool
	redisClient    *redis.Client
	cashierService service.CashierService
	userService    service.UserService
	merchantService service.MerchantService
}

func (s *CashierServiceTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := pgxpool.New(s.ts.Ctx, s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	opts, err := redis.ParseURL(s.ts.RedisURL)
	s.Require().NoError(err)
	s.redisClient = redis.NewClient(opts)

	queries := db.New(pool)
	repos := repository.NewRepositories(queries)

	logger.ResetInstance()
	lp := sdklog.NewLoggerProvider()
	log, _ := logger.NewLogger("test", lp)
	obs, _ := observability.NewObservability("test", log)
	cacheMetrics, _ := observability.NewCacheMetrics("test")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)
	cacheCashier := cashier_cache.NewCashierMencache(cacheStore)

	s.cashierService = service.NewCashierService(service.CashierServiceDeps{
		CashierRepo:   repos.Cashier,
		MerchantRepo:  repos.Merchant,
		UserRepo:      repos.User,
		Logger:        log,
		Observability: obs,
		Cache:         cacheCashier,
	})
    // Also need user and merchant services to setup data
    // For simplicity, we use repositories directly if services are complex, 
    // but here we can just use repos.
}

func (s *CashierServiceTestSuite) TearDownSuite() {
	if s.redisClient != nil {
		s.redisClient.Close()
	}
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *CashierServiceTestSuite) TestCashierLifecycle() {
	ctx := context.Background()

	// 0. Setup User and Merchant (using Repositories to avoid service complexity)
    queries := db.New(s.dbPool)
    userRepo := repository.NewUserRepository(queries)
    merchantRepo := repository.NewMerchantRepository(queries)

	userReq := &requests.CreateUserRequest{
		FirstName: "Cashier",
		LastName:  "Service",
		Email:     "cashier.service@example.com",
		Password:  "password123",
	}
	user, err := userRepo.CreateUser(ctx, userReq)
	s.NoError(err)
	userID := int(user.UserID)

	merchantReq := &requests.CreateMerchantRequest{
		UserID:       userID,
		Name:         "Service Merchant",
		Description:  "Merchant for service testing",
		Address:      "Jakarta, Indonesia",
		ContactEmail: "merchant.srv@example.com",
		ContactPhone: "08123456789",
		Status:       "active",
	}
	merchant, err := merchantRepo.CreateMerchant(ctx, merchantReq)
	s.NoError(err)
	merchantID := int(merchant.MerchantID)

	// 1. Create Cashier
	req := &requests.CreateCashierRequest{
		MerchantID: merchantID,
		UserID:     userID,
		Name:       "Service Cashier",
	}
	cashier, err := s.cashierService.CreateCashier(ctx, req)
	s.NoError(err)
	s.NotNil(cashier)
	s.Equal(req.Name, cashier.Name)

	cashierID := int(cashier.CashierID)

	// 2. Find By ID
	found, err := s.cashierService.FindById(ctx, cashierID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(cashier.Name, found.Name)

	// 3. Update Cashier
	updateReq := &requests.UpdateCashierRequest{
		CashierID: &cashierID,
		Name:      "Updated Cashier",
	}
	updated, err := s.cashierService.UpdateCashier(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(updateReq.Name, updated.Name)

	// 4. Trash and Restore
	_, err = s.cashierService.TrashedCashier(ctx, cashierID)
	s.NoError(err)

	_, err = s.cashierService.RestoreAllCashier(ctx) // Test bulk restore for variety
	s.NoError(err)

	// 5. Delete Permanent
    // Trash again
    _, err = s.cashierService.TrashedCashier(ctx, cashierID)
    s.NoError(err)

	success, err := s.cashierService.DeleteCashierPermanent(ctx, cashierID)
	s.NoError(err)
	s.True(success)
}

func TestCashierServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CashierServiceTestSuite))
}
