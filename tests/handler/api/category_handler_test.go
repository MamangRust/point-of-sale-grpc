package api_test

import (
	"bytes"
	"pointofsale/internal/cache"
	api_category_cache "pointofsale/internal/cache/api/category"
	category_cache "pointofsale/internal/cache/category"
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
	"pointofsale/internal/domain/requests"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CategoryApiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	echo        *echo.Echo
	client      pb.CategoryServiceClient
	conn        *grpc.ClientConn
}

func (s *CategoryApiTestSuite) SetupSuite() {
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
	
	// Service layer cache
	catCacheSrv := category_cache.NewCategoryMencache(cacheStore)
	// API layer cache
	catCacheApi := api_category_cache.NewCategoryMencache(cacheStore)

	categoryService := service.NewCategoryService(service.CategoryServiceDeps{
		CategoryRepo:  repos.Category,
		Logger:             log,
		Observability:      obs,
		Cache:              catCacheSrv,
	})

	// Start gRPC Server
	categoryGapi := gapi.NewCategoryHandleGrpc(categoryService)
	server := grpc.NewServer()
	pb.RegisterCategoryServiceServer(server, categoryGapi)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)

	go func() {
		_ = server.Serve(lis)
	}()

	// gRPC Client for the API Handler
	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.client = pb.NewCategoryServiceClient(conn)

	// Echo Setup
	s.echo = echo.New()
	mapping := response_api.NewCategoryResponseMapper()
	apiHandler := errors.NewApiHandler(obs, log)

	api.NewHandlerCategory(s.echo, s.client, log, mapping, apiHandler, catCacheApi)
}

func (s *CategoryApiTestSuite) TearDownSuite() {
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

func (s *CategoryApiTestSuite) TestCategoryApiLifecycle() {
	// 1. Create
	createReq := requests.CreateCategoryRequest{
		Name:        "Electronics API",
		Description: "Via Echo",
	}
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest(http.MethodPost, "/api/category/create", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusCreated, rec.Code)
	var createRes map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &createRes)
	s.NoError(err)
	s.Equal("success", createRes["status"])
	
	data := createRes["data"].(map[string]interface{})
	categoryID := int(data["id"].(float64))

	// 2. Find By ID
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/category/%d", categoryID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 3. Update
	updateReq := requests.UpdateCategoryRequest{
		CategoryID:  &categoryID,
		Name:        "Electronics API Updated",
		Description: "Via Echo Updated",
	}
	body, _ = json.Marshal(updateReq)
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/category/update/%d", categoryID), bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 4. Trash
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/category/trashed/%d", categoryID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 5. Restore
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/category/restore/%d", categoryID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 6. Delete Permanent
	// Trash again
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/category/trashed/%d", categoryID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/category/permanent/%d", categoryID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestCategoryApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(CategoryApiTestSuite))
}
