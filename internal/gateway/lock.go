package gateway

import (
	"context"
	"time"
)

type LockGateway interface {
	Lock(context.Context, string, time.Duration) error
	IsLocked(context.Context, string) (bool, error)
}
