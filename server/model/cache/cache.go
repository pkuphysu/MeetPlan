package cache

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/singleflight"

	"github.com/allegro/bigcache/v3"
	"github.com/eko/gocache/lib/v4/cache"
	bigcachestore "github.com/eko/gocache/store/bigcache/v4"
)

var (
	store     *bigcachestore.BigcacheStore
	longStore *bigcachestore.BigcacheStore
)

func init() {
	cacheClient, err := bigcache.New(context.TODO(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		panic(err)
	}
	store = bigcachestore.NewBigcache(cacheClient)

	longStoreConfig := bigcache.DefaultConfig(24 * time.Hour)
	longStoreConfig.CleanWindow = 3 * time.Hour
	longCacheClient, err := bigcache.New(context.TODO(), longStoreConfig)
	if err != nil {
		panic(err)
	}
	longStore = bigcachestore.NewBigcache(longCacheClient)
}

func New[T any](long ...bool) *cache.Cache[T] {
	if len(long) > 0 && long[0] {
		return cache.New[T](longStore)
	}
	return cache.New[T](store)
}

func NewLoader[T any](f func(ctx context.Context, key string) (T, error), sfg *singleflight.Group, long ...bool) *cache.LoadableCache[T] {
	loadFunc := func(ctx context.Context, key any) (T, error) {
		cacheKey, ok := key.(string)
		if !ok {
			return *new(T), errors.New("invalid key type")
		}
		v, err, _ := sfg.Do(cacheKey, func() (interface{}, error) {
			return f(ctx, cacheKey)
		})
		if err != nil {
			return *new(T), err
		}
		return v.(T), nil
	}
	if len(long) > 0 && long[0] {
		return cache.NewLoadable[T](loadFunc, cache.New[T](longStore))
	}
	return cache.NewLoadable[T](loadFunc, cache.New[T](store))
}
