package role_cache

import "pointofsale/internal/cache"

type RoleMencache interface {
	RoleQueryCache
	RoleCommandCache
}

type roleMencache struct {
	RoleQueryCache
	RoleCommandCache
}

func NewRoleMencache(store *cache.CacheStore) RoleMencache {
	return &roleMencache{
		RoleQueryCache:   NewRoleQueryCache(store),
		RoleCommandCache: NewRoleCommandCache(store),
	}
}
