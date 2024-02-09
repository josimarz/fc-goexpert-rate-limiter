package gateway

import (
	"context"
	"testing"
	"time"
)

var (
	lockGateway *LockRedisGateway
)

func TestNewLockRedisGateway(t *testing.T) {
	lockGateway = NewLockRedisGateway(redisClient)
	if lockGateway == nil {
		t.Errorf("NewLockRedisGateway(%v) = %v, should be defined", redisClient, nil)
	}
}

func TestLockRedisGateway(t *testing.T) {
	ctx := context.Background()
	key := "127.0.0.1"
	ttl := time.Second
	t.Run("Lock", func(t *testing.T) {
		if got := lockGateway.Lock(ctx, key, ttl); got != nil {
			t.Errorf("Lock(%v, %v, %v) = %v, want %v", ctx, key, ttl, got, nil)
		}
	})
	t.Run("IsLocked", func(t *testing.T) {
		if locked, err := lockGateway.IsLocked(ctx, key); !locked || err != nil {
			t.Errorf("IsLocked(%v, %v) = (%v, %v), want (%v, %v)", ctx, key, locked, err, true, nil)
		}
		time.Sleep(ttl)
		if locked, err := lockGateway.IsLocked(ctx, key); locked || err != nil {
			t.Errorf("IsLocked(%v, %v) = (%v, %v), want (%v, %v)", ctx, key, locked, err, false, nil)
		}
	})
}
