package gateway

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type LockRedisGateway struct {
	client *redis.Client
}

func NewLockRedisGateway(client *redis.Client) *LockRedisGateway {
	return &LockRedisGateway{client}
}

func (g *LockRedisGateway) Lock(ctx context.Context, key string, ttl time.Duration) error {
	key = fmt.Sprintf("lock:%s", key)
	if _, err := g.client.SetNX(ctx, key, true, ttl).Result(); err != nil {
		return err
	}
	return nil
}

func (g *LockRedisGateway) IsLocked(ctx context.Context, key string) (bool, error) {
	key = fmt.Sprintf("lock:%s", key)
	_, err := g.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
