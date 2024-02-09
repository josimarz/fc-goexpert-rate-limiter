package gateway

import (
	"context"
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type TokenRedisGateway struct {
	client *redis.Client
}

func NewTokenRedisGateway(client *redis.Client) *TokenRedisGateway {
	return &TokenRedisGateway{client}
}

func (g *TokenRedisGateway) GetLimit(ctx context.Context, token string) (int, error) {
	key := fmt.Sprintf("limit:%s", token)
	result, err := g.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return 0, nil
	}
	if result == "" {
		return 0, nil
	}
	limit, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return limit, nil
}
