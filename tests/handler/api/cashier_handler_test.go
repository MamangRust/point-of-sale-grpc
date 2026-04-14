package api_test

import (
	"bytes"
	"context"
	api_cashier_cache "pointofsale/internal/cache/api/cashier"
	cashier_cache "pointofsale/internal/cache/cashier"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/handler/api"
	"pointofsale/internal/handler/gapi"
	response_api "pointofsale/internal/mapper"
	"pointofsale/internal/pb"
	"pointofsale/internal/repository"
	"pointofsale/internal/service"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/errors"
	"pointofsale/pkg/logger"
	"pointofsale/pkg/observability"
	"pointofsale/tests"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CashierApiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	echo        *echo.Echo
	client      pb.CashierServiceClient
	conn        *grpc.ClientConn
}

func (s *CashierApiTestSuite) SetupSuite() {
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
	log, _ := logger.NewLogger("test-api", lp)
	obs, _ := observability.NewObservability("test-api", log)
	
	cacheMetrics, _ := observability.NewCacheMetrics("test-api")
	cacheStore := cache.NewCacheStore(s.redisClient, log, cacheMetrics)
	
	cashierCacheSrv := cashier_cache.NewCashierMencache(cacheStore)
	cashierCacheApi := api_cashier_cache.NewCashierMencache(cacheStore)

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

	s.echo = echo.New()
	mapping := response_api.NewCashierResponseMapper()
	apiHandler := errors.NewApiHandler(obs, log)

	api.NewHandlerCashier(s.echo, s.client, log, mapping, apiHandler, cashierCacheApi)
}

func (s *CashierApiTestSuite) TearDownSuite() {
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

func (s *CashierApiTestSuite) TestCashierApiLifecycle() {
    ctx := context.Background()
    // 0. Setup User and Merchant via Repository
    queries := db.New(s.dbPool)
    userRepo := repository.NewUserRepository(queries)
    merchantRepo := repository.NewMerchantRepository(queries)

    user, _ := userRepo.CreateUser(ctx, &requests.CreateUserRequest{
        FirstName: "Cashier",
        LastName:  "Api",
        Email:     "cashier.api@example.com",
        Password:  "password123",
    })
    userID := int(user.UserID)

    merchant, _ := merchantRepo.CreateMerchant(ctx, &requests.CreateMerchantRequest{
        UserID: userID,
        Name: "Api Merchant",
        Status: "active",
    })
    merchantID := int(merchant.MerchantID)

	// 1. Create
	createReq := map[string]interface{}{
        "name": "API Cashier",
        "merchant_id": merchantID,
        "user_id": userID,
    }
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest(http.MethodPost, "/api/cashier/create", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusCreated, rec.Code)
	var createRes map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &createRes)
	s.NoError(err)
	s.Equal("success", createRes["status"])
	
	data := createRes["data"].(map[string]interface{})
	cashierID := int(data["id"].(float64))

	// 2. Find By ID
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/cashier/%d", cashierID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 3. Update
	updateReq := map[string]string{"name": "API Cashier Updated"}
	body, _ = json.Marshal(updateReq)
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/cashier/update/%d", cashierID), bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 4. Trash
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/cashier/trashed/%d", cashierID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 5. Restore
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/cashier/restore/%d", cashierID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 6. Delete Permanent
	// Trash again
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/cashier/trashed/%d", cashierID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/cashier/permanent/%d", cashierID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestCashierApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CashierApiTestSuite))
}
