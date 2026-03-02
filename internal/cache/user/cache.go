package user_cache

import "pointofsale/internal/cache"

type UserMencache interface {
	UserQueryCache
	UserCommandCache
}

type userMencache struct {
	UserQueryCache
	UserCommandCache
}

func NewUserMencache(cacheStore *cache.CacheStore) UserMencache {
	return &userMencache{
		UserQueryCache:   NewUserQueryCache(cacheStore),
		UserCommandCache: NewUserCommandCache(cacheStore),
	}
}
