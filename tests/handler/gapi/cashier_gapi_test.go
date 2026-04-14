package gapi_test

import (
	"context"
	cashier_cache "pointofsale/internal/cache/cashier"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/handler/gapi"
	"pointofsale/internal/pb"
	"pointofsale/internal/repository"
	"pointofsale/internal/service"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/tests"
	"net"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CashierGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	client      pb.CashierServiceClient
	conn        *grpc.ClientConn
}

func (s *CashierGapiTestSuite) SetupSuite() {
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
	log, _ := logger.NewLogger("test-gapi", lp)
	obs, _ := observability.NewObservability("test-gapi", log)
	
	cacheMetrics, _ := observability.NewCacheMetrics("test-gapi")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)
	
	cashierCacheSrv := cashier_cache.NewCashierMencache(cacheStore)

	cashierService := service.NewCashierService(service.CashierServiceDeps{
		CashierRepo:   repos.Cashier,
		MerchantRepo:  repos.Merchant,
		UserRepo:      repos.User,
		Logger:        log,
		Observability: obs,
		Cache:         cashierCacheSrv,
	})

	cashierGapi := gapi.NewCashierHandleGrpc(cashierService)
	server := grpc.NewServer()
	pb.RegisterCashierServiceServer(server, cashierGapi)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)

	go func() {
		_ = server.Serve(lis)
	}()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.client = pb.NewCashierServiceClient(conn)
}

func (s *CashierGapiTestSuite) TearDownSuite() {
	if s.conn != nil {
		s.conn.Close()
	}
	if s.grpcServer != nil {
		s.grpcServer.Stop()
	}
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

func (s *CashierGapiTestSuite) TestCashierGapiLifecycle() {
	ctx := context.Background()

    // 0. Setup User and Merchant via Repository
    queries := db.New(s.dbPool)
    userRepo := repository.NewUserRepository(queries)
    merchantRepo := repository.NewMerchantRepository(queries)

    user, _ := userRepo.CreateUser(ctx, &requests.CreateUserRequest{
        FirstName: "Cashier",
        LastName:  "Gapi",
        Email:     "cashier.gapi@example.com",
        Password:  "password123",
    })
    userID := int(user.UserID)

    merchant, _ := merchantRepo.CreateMerchant(ctx, &requests.CreateMerchantRequest{
        UserID: userID,
        Name: "Gapi Merchant",
        Status: "active",
    })
    merchantID := int(merchant.MerchantID)

	// 1. Create
	createRes, err := s.client.CreateCashier(ctx, &pb.CreateCashierRequest{
		Name:       "gRPC Cashier",
        MerchantId: int32(merchantID),
        UserId:     int32(userID),
	})
	s.NoError(err)
	s.Equal("success", createRes.Status)
	cashierID := createRes.Data.Id

	// 2. Find By ID
	findRes, err := s.client.FindById(ctx, &pb.FindByIdCashierRequest{
		Id: cashierID,
	})
	s.NoError(err)
	s.Equal("success", findRes.Status)
	s.Equal("gRPC Cashier", findRes.Data.Name)

	// 3. Update
	updateRes, err := s.client.UpdateCashier(ctx, &pb.UpdateCashierRequest{
		CashierId: cashierID,
		Name:      "gRPC Cashier Updated",
	})
	s.NoError(err)
	s.Equal("success", updateRes.Status)

	// 4. Trash
	trashRes, err := s.client.TrashedCashier(ctx, &pb.FindByIdCashierRequest{
		Id: cashierID,
	})
	s.NoError(err)
	s.Equal("success", trashRes.Status)

	// 5. Restore
	restoreRes, err := s.client.RestoreCashier(ctx, &pb.FindByIdCashierRequest{
		Id: cashierID,
	})
	s.NoError(err)
	s.Equal("success", restoreRes.Status)

	// 6. Delete Permanent
	// Trash again
	_, _ = s.client.TrashedCashier(ctx, &pb.FindByIdCashierRequest{
		Id: cashierID,
	})

	deleteRes, err := s.client.DeleteCashierPermanent(ctx, &pb.FindByIdCashierRequest{
		Id: cashierID,
	})
	s.NoError(err)
	s.Equal("success", deleteRes.Status)
}

func TestCashierGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CashierGapiTestSuite))
}
