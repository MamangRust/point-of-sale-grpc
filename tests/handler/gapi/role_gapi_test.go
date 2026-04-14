package gapi_test

import (
	"context"
	role_cache "pointofsale/internal/cache/role"
	"pointofsale/internal/cache"
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

type RoleGapiTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	grpcServer  *grpc.Server
	client      pb.RoleServiceClient
	conn        *grpc.ClientConn
}

func (s *RoleGapiTestSuite) SetupSuite() {
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
	
	roleCacheSrv := role_cache.NewRoleMencache(cacheStore)

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
}

func (s *RoleGapiTestSuite) TearDownSuite() {
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

func (s *RoleGapiTestSuite) TestRoleGapiLifecycle() {
	ctx := context.Background()

	// 1. Create
	createRes, err := s.client.CreateRole(ctx, &pb.CreateRoleRequest{
		Name: "gRPC Role",
	})
	s.NoError(err)
	s.Equal("success", createRes.Status)
	roleID := createRes.Data.Id

	// 2. Find By ID
	findRes, err := s.client.FindByIdRole(ctx, &pb.FindByIdRoleRequest{
		RoleId: roleID,
	})
	s.NoError(err)
	s.Equal("success", findRes.Status)
	s.Equal("gRPC Role", findRes.Data.Name)

	// 3. Update
	updateRes, err := s.client.UpdateRole(ctx, &pb.UpdateRoleRequest{
		Id:   roleID,
		Name: "gRPC Role Updated",
	})
	s.NoError(err)
	s.Equal("success", updateRes.Status)

	// 4. Trash
	trashRes, err := s.client.TrashedRole(ctx, &pb.FindByIdRoleRequest{
		RoleId: roleID,
	})
	s.NoError(err)
	s.Equal("success", trashRes.Status)

	// 5. Restore
	restoreRes, err := s.client.RestoreRole(ctx, &pb.FindByIdRoleRequest{
		RoleId: roleID,
	})
	s.NoError(err)
	s.Equal("success", restoreRes.Status)

	// 6. Delete Permanent
	// Trash again
	_, _ = s.client.TrashedRole(ctx, &pb.FindByIdRoleRequest{
		RoleId: roleID,
	})

	deleteRes, err := s.client.DeleteRolePermanent(ctx, &pb.FindByIdRoleRequest{
		RoleId: roleID,
	})
	s.NoError(err)
	s.Equal("success", deleteRes.Status)
}

func TestRoleGapiSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleGapiTestSuite))
}
