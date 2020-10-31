package infra

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/typical-go/typical-rest-server/pkg/cachekit"
)

// NewCacheStore return new instaence of cache store
// @ctor
func NewCacheStore(client *redis.Client) *cachekit.Store {
	return &cachekit.Store{
		Client:        client,
		DefaultMaxAge: 30 * time.Second,
		Prefix:        "cache_",
	}
}