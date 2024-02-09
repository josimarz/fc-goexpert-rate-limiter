package gateway

import (
	"context"
	"testing"
	"time"
)

var (
	lockGateway *LockInMemoryGateway
)

func TestNewLockInMemoryGateway(t *testing.T) {
	lockGateway = NewLockInMemoryGateway()
	if lockGateway == nil {
		t.Errorf("NewLockInMemoryGateway() = %v, should be defined", lockGateway)
	}
}

func TestLockInMemoryGateway(t *testing.T) {
	ctx := context.Background()
	key := "vcQtRsHMepRoQdROl0ziBk7zcfyPza2AyloYC1XVtCbP6hxdloOYHb7JZ1wcOpCT"
	ttl := time.Second
	t.Run("Lock", func(t *testing.T) {
		if err := lockGateway.Lock(ctx, key, ttl); err != nil {
			t.Errorf("Lock(%v, %v, %v) = %v, want %v", ctx, key, ttl, err, nil)
		}
		if err := lockGateway.Lock(ctx, key, ttl); err != nil {
			t.Errorf("Lock(%v, %v, %v) = %v, want %v", ctx, key, ttl, err, nil)
		}
	})
	t.Run("IsLocked", func(t *testing.T) {
		if locked, err := lockGateway.IsLocked(ctx, key); !locked || err != nil {
			t.Errorf("IsLocked(%v, %v) = (%v, %v), want (%v, %v)", ctx, key, locked, err, true, nil)
		}
		key := "XCb6H3c6zPG4M5yRG0vy1p1pFY5KBqiy3fqIghDXh3AZSo78cRHGM2eqmIJc5QgF"
		if locked, err := lockGateway.IsLocked(ctx, key); locked || err != nil {
			t.Errorf("IsLocked(%v, %v) = (%v, %v), want (%v, %v)", ctx, key, locked, err, false, nil)
		}
	})
}
