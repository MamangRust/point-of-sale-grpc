package user_cache

import (
	"context"
	"fmt"
	"pointofsale/internal/cache"
	"pointofsale/internal/domain/requests"
	"pointofsale/internal/domain/response"
	"time"
)

const (
	userAllCacheKey     = "user:all:page:%d:pageSize:%d:search:%s"
	userByIdCacheKey    = "user:id:%d"
	userActiveCacheKey  = "user:active:page:%d:pageSize:%d:search:%s"
	userTrashedCacheKey = "user:trashed:page:%d:pageSize:%d:search:%s"

	ttlDefault = 5 * time.Minute
)

type userQueryCache struct {
	store *cache.CacheStore
}

func NewUserQueryCache(store *cache.CacheStore) *userQueryCache {
	return &userQueryCache{store: store}
}

func (s *userQueryCache) GetCachedUsersCache(ctx context.Context, req *requests.FindAllUsers) (*response.ApiResponsePaginationUser, bool) {
	key := fmt.Sprintf(userAllCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationUser](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *userQueryCache) SetCachedUsersCache(ctx context.Context, req *requests.FindAllUsers, res *response.ApiResponsePaginationUser) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(userAllCacheKey, req.Page, req.PageSize, req.Search)

	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *userQueryCache) GetCachedUserActiveCache(ctx context.Context, req *requests.FindAllUsers) (*response.ApiResponsePaginationUserDeleteAt, bool) {
	key := fmt.Sprintf(userActiveCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationUserDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *userQueryCache) SetCachedUserActiveCache(ctx context.Context, req *requests.FindAllUsers, res *response.ApiResponsePaginationUserDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(userActiveCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *userQueryCache) GetCachedUserTrashedCache(ctx context.Context, req *requests.FindAllUsers) (*response.ApiResponsePaginationUserDeleteAt, bool) {
	key := fmt.Sprintf(userTrashedCacheKey, req.Page, req.PageSize, req.Search)

	result, found := cache.GetFromCache[*response.ApiResponsePaginationUserDeleteAt](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *userQueryCache) SetCachedUserTrashedCache(ctx context.Context, req *requests.FindAllUsers, res *response.ApiResponsePaginationUserDeleteAt) {
	if res == nil {
		return
	}

	key := fmt.Sprintf(userTrashedCacheKey, req.Page, req.PageSize, req.Search)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}

func (s *userQueryCache) GetCachedUserCache(ctx context.Context, id int) (*response.ApiResponseUser, bool) {
	key := fmt.Sprintf(userByIdCacheKey, id)

	result, found := cache.GetFromCache[*response.ApiResponseUser](ctx, s.store, key)

	if !found || result == nil {
		return nil, false
	}

	return result, true
}

func (s *userQueryCache) SetCachedUserCache(ctx context.Context, res *response.ApiResponseUser) {
	if res == nil || res.Data == nil {
		return
	}

	key := fmt.Sprintf(userByIdCacheKey, res.Data.ID)
	cache.SetToCache(ctx, s.store, key, res, ttlDefault)
}
