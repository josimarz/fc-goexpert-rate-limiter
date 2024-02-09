package gateway

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RequestRedisGateway struct {
	client *redis.Client
}

func NewRequestRedisGateway(client *redis.Client) *RequestRedisGateway {
	return &RequestRedisGateway{client}
}

func (g *RequestRedisGateway) Save(ctx context.Context, key string) error {
	key = fmt.Sprintf("request:%s", key)
	if _, err := g.client.SetNX(ctx, key, 0, time.Second).Result(); err != nil {
		return err
	}
	if _, err := g.client.Incr(ctx, key).Result(); err != nil {
		return err
	}
	return nil
}

func (g *RequestRedisGateway) Count(ctx context.Context, key string) (int, error) {
	key = fmt.Sprintf("request:%s", key)
	result, err := g.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	value, err := strconv.Atoi(result)
	if err != nil {
		return 0, nil
	}
	return value, nil
}
