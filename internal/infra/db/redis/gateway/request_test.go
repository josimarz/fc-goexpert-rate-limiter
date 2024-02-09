package gateway

import (
	"context"
	"testing"
)

var (
	requestGateway *RequestRedisGateway
)

func TestNewRequestRedisGateway(t *testing.T) {
	requestGateway = NewRequestRedisGateway(redisClient)
	if requestGateway == nil {
		t.Errorf("NewRequestRedisGateway(%v) = %v, should be defined", redisClient, nil)
	}
}

func TestRequestGateway(t *testing.T) {
	ctx := context.Background()
	key := "127.0.0.1"
	t.Run("Save", func(t *testing.T) {
		if err := requestGateway.Save(ctx, key); err != nil {
			t.Errorf("Save(%v, %v) = %v, want %v", ctx, key, err, nil)
		}
	})
	t.Run("Count", func(t *testing.T) {
		value, err := requestGateway.Count(ctx, key)
		if err != nil || value != 1 {
			t.Errorf("Count(%v, %v) = (%v, %v), want (1, %v)", ctx, key, value, err, nil)
		}
	})
}
