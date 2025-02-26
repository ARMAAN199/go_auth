package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore interface {
	SetRefreshTokenWithExpiry(context.Context, string, string, time.Duration) error
	GetRefreshToken(context.Context, string) (string, error)
	DeleteRefreshToken(context.Context, string) error
	SetBlackListedAccessToken(context.Context, string) error
	IsAccessTokenBlackListed(context.Context, string) (bool, error)
}

type dbRedisStore struct {
	client *redis.Client
}

func NewDBRedisStore(client *redis.Client) RedisStore {
	store := &dbRedisStore{
		client: client,
	}
	return store
}
