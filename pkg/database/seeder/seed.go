package seeder

import (
	"context"
	"fmt"
	db "pointofsale/pkg/database/schema"
	"pointofsale/pkg/hash"
	"pointofsale/pkg/logger"
	"time"
)

type Deps struct {
	Db     *db.Queries
	Ctx    context.Context
	Logger logger.LoggerInterface
	Hash   hash.HashPassword
}

type Seeder struct {
	User     *userSeeder
	Role     *roleSeeder
	UserRole *userRoleSeeder
}

func NewSeeder(deps Deps) *Seeder {
	return &Seeder{
		User:     NewUserSeeder(deps.Db, deps.Hash, deps.Ctx, deps.Logger),
		Role:     NewRoleSeeder(deps.Db, deps.Ctx, deps.Logger),
		UserRole: NewUserRoleSeeder(deps.Db, deps.Ctx, deps.Logger),
	}
}

func (s *Seeder) Run() error {
	if err := s.seedWithDelay("users", s.User.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("roles", s.Role.Seed); err != nil {
		return err
	}

	if err := s.seedWithDelay("user_role", s.UserRole.Seed); err != nil {
		return nil
	}

	return nil
}

func (s *Seeder) seedWithDelay(entityName string, seedFunc func() error) error {
	if err := seedFunc(); err != nil {
		return fmt.Errorf("failed to seed %s: %w", entityName, err)
	}

	time.Sleep(30 * time.Second)
	return nil
}
