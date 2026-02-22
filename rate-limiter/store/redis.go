package store

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(c *redis.Client) *RedisStore {
	return &RedisStore{
		client: c,
	}
}

func (r *RedisStore) Execute(ctx context.Context, scriptName, scriptBody string, keys []string, args ...any) (any, error) {
	return r.client.Eval(ctx, scriptBody, keys, args...).Result()
}
