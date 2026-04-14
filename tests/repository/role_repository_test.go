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

type RoleRepositoryTestSuite struct {
	suite.Suite
	ts     *tests.TestSuite
	dbPool *pgxpool.Pool
	repo   repository.RoleRepository
}

func (s *RoleRepositoryTestSuite) SetupSuite() {
	ts, err := tests.SetupTestSuite()
	s.Require().NoError(err)
	s.ts = ts

	pool, err := pgxpool.New(s.ts.Ctx, s.ts.DBURL)
	s.Require().NoError(err)
	s.dbPool = pool

	queries := db.New(pool)
	s.repo = repository.NewRoleRepository(queries)
}

func (s *RoleRepositoryTestSuite) TearDownSuite() {
	if s.dbPool != nil {
		s.dbPool.Close()
	}
	if s.ts != nil {
		s.ts.Teardown()
	}
}

func (s *RoleRepositoryTestSuite) TestRoleLifecycle() {
	ctx := context.Background()

	// 1. Create Role
	createReq := &requests.CreateRoleRequest{
		Name: "Admin",
	}

	role, err := s.repo.CreateRole(ctx, createReq)
	s.NoError(err)
	s.NotNil(role)
	s.Equal(createReq.Name, role.RoleName)

	roleID := int(role.RoleID)

	// 2. Find By ID
	found, err := s.repo.FindById(ctx, roleID)
	s.NoError(err)
	s.NotNil(found)
	s.Equal(role.RoleName, found.RoleName)

	// 3. Update Role
	updateReq := &requests.UpdateRoleRequest{
		ID:   &roleID,
		Name: "Super Admin",
	}

	updated, err := s.repo.UpdateRole(ctx, updateReq)
	s.NoError(err)
	s.NotNil(updated)
	s.Equal(updateReq.Name, updated.RoleName)

	// 4. Trash Role
	trashed, err := s.repo.TrashedRole(ctx, roleID)
	s.NoError(err)
	s.NotNil(trashed)
	s.True(trashed.DeletedAt.Valid)

	// 5. Find By Trashed
	trashedList, err := s.repo.FindByTrashedRole(ctx, &requests.FindAllRoles{
		Page:     1,
		PageSize: 10,
	})
	s.NoError(err)
	s.NotEmpty(trashedList)

	// 6. Restore Role
	restored, err := s.repo.RestoreRole(ctx, roleID)
	s.NoError(err)
	s.NotNil(restored)
	s.False(restored.DeletedAt.Valid)

	// 7. Delete Permanent
	// Re-trash first
	_, err = s.repo.TrashedRole(ctx, roleID)
	s.NoError(err)

	success, err := s.repo.DeleteRolePermanent(ctx, roleID)
	s.NoError(err)
	s.True(success)

	// 8. Verify it's gone
	_, err = s.repo.FindById(ctx, roleID)
	s.Error(err)
}

func TestRoleRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	suite.Run(t, new(RoleRepositoryTestSuite))
}
