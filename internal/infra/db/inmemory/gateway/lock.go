package gateway

import (
	"context"
	"time"
)

type Lock struct {
	key       string
	expiresAt time.Time
}

type LockInMemoryGateway struct {
	locks []*Lock
}

func NewLockInMemoryGateway() *LockInMemoryGateway {
	return &LockInMemoryGateway{
		locks: []*Lock{},
	}
}

func (g *LockInMemoryGateway) Lock(ctx context.Context, key string, ttl time.Duration) error {
	for _, lock := range g.locks {
		if lock.key == key {
			return nil
		}
	}
	g.locks = append(g.locks, &Lock{key: key, expiresAt: time.Now().Local().Add(ttl)})
	return nil
}

func (g *LockInMemoryGateway) IsLocked(ctx context.Context, key string) (bool, error) {
	for _, lock := range g.locks {
		if lock.key == key {
			return lock.expiresAt.After(time.Now()), nil
		}
	}
	return false, nil
}
