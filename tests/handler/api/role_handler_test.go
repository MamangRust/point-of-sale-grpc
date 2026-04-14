package api_test

import (
	"bytes"
	api_role_cache "pointofsale/internal/cache/api/role"
	role_cache "pointofsale/internal/cache/role"
	"pointofsale/internal/cache"
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

type RoleApiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	echo        *echo.Echo
	client      pb.RoleServiceClient
	conn        *grpc.ClientConn
}

func (s *RoleApiTestSuite) SetupSuite() {
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
	
	roleCacheSrv := role_cache.NewRoleMencache(cacheStore)
	roleCacheApi := api_role_cache.NewRoleMencache(cacheStore)

	roleService := service.NewRoleService(service.RoleServiceDeps{
		RoleRepo:      repos.Role,
		Logger:        log,
		Observability: obs,
		Cache:         roleCacheSrv,
	})

	roleGapi := gapi.NewRoleHandleGrpc(roleService)
	server := grpc.NewServer()
	pb.RegisterRoleServiceServer(server, roleGapi)
	s.grpcServer = server

	lis, err := net.Listen("tcp", ":0")
	s.Require().NoError(err)

	go func() {
		_ = server.Serve(lis)
	}()

	conn, err := grpc.NewClient(lis.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err)
	s.conn = conn
	s.client = pb.NewRoleServiceClient(conn)

	s.echo = echo.New()
	mapping := response_api.NewRoleResponseMapper()
	apiHandler := errors.NewApiHandler(obs, log)

	api.NewHandlerRole(s.echo, s.client, log, mapping, apiHandler, roleCacheApi)
}

func (s *RoleApiTestSuite) TearDownSuite() {
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

func (s *RoleApiTestSuite) TestRoleApiLifecycle() {
	// 1. Create
	createReq := map[string]string{"name": "API Role"}
	body, _ := json.Marshal(createReq)
	req := httptest.NewRequest(http.MethodPost, "/api/role", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, rec.Code)
	var createRes map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &createRes)
	s.NoError(err)
	s.Equal("success", createRes["status"])
	
	data := createRes["data"].(map[string]interface{})
	roleID := int(data["id"].(float64))

	// 2. Find By ID
	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/role/%d", roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 3. Update
	updateReq := map[string]string{"name": "API Role Updated"}
	body, _ = json.Marshal(updateReq)
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/role/update/%d", roleID), bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 4. Trash
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/role/trashed/%d", roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 5. Restore
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/role/restore/%d", roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)

	// 6. Delete Permanent
	// Trash again
	req = httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/role/trashed/%d", roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)

	req = httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/role/permanent/%d", roleID), nil)
	rec = httptest.NewRecorder()
	s.echo.ServeHTTP(rec, req)
	s.Equal(http.StatusOK, rec.Code)
}

func TestRoleApiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleApiTestSuite))
}
