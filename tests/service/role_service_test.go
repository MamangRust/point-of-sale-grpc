package service_test

import (
	"context"
	"pointofsale/internal/cache"
	role_cache "pointofsale/internal/cache/role"
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

type RoleServiceTestSuite struct {
	suite.Suite
	ts          *tests.TestSuite
	dbPool      *pgxpool.Pool
	redisClient *redis.Client
	roleService service.RoleService
}

func (s *RoleServiceTestSuite) SetupSuite() {
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
	cacheRole := role_cache.NewRoleMencache(cacheStore)

	s.roleService = service.NewRoleService(service.RoleServiceDeps{
		RoleRepo:      repos.Role,
		Logger:        log,
		Observability: obs,
		Cache:         cacheRole,
	})
}

func (s *RoleServiceTestSuite) TearDownSuite() {
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

func (s *RoleServiceTestSuite) TestRoleLifecycle() {
	ctx := context.Background()

	// 1. Create Role
	req := &requests.CreateRoleRequest{
		Name: "Test Role",
	}
	role, err := s.roleService.CreateRole(ctx, req)
	s.NoError(err)
	s.NotNil(role)
	s.Equal(req.Name, role.RoleName)

	roleID := int(role.RoleID)

	// 2. Find By ID
	found, err := s.roleService.FindById(ctx, roleID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(role.RoleName, found.RoleName)

	// 3. Update Role
	updateReq := &requests.UpdateRoleRequest{
		ID:   &roleID,
		Name: "Updated Role",
	}
	updated, err := s.roleService.UpdateRole(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(updateReq.Name, updated.RoleName)

	// 4. Trash and Restore
	_, err = s.roleService.TrashedRole(ctx, roleID)
	s.NoError(err)

	_, err = s.roleService.RestoreRole(ctx, roleID)
	s.NoError(err)

	// 5. Delete Permanent
	success, err := s.roleService.DeleteRolePermanent(ctx, roleID)
	s.NoError(err)
	s.True(success)
}

func TestRoleServiceSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleServiceTestSuite))
}
