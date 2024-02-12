package ratelimiter

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/josimarz/fc-goexpert-rate-limiter/internal/gateway"
)

type DefaultRateLimiter struct {
	settings       *Settings
	lockGateway    gateway.LockGateway
	requestGateway gateway.RequestGateway
	tokenGateway   gateway.TokenGateway
}

func NewDefaultRateLimiter(settings *Settings, lockGateway gateway.LockGateway, requestGatewat gateway.RequestGateway, tokenGateway gateway.TokenGateway) *DefaultRateLimiter {
	return &DefaultRateLimiter{settings, lockGateway, requestGatewat, tokenGateway}
}

func (rt *DefaultRateLimiter) CanGo(ctx context.Context, r *http.Request) (bool, error) {
	key := r.Header.Get("x-api-key")
	if key == "" || !rt.settings.LimitByToken {
		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			return false, err
		}
		key = host
	}
	locked, err := rt.lockGateway.IsLocked(ctx, key)
	if locked {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	limit, err := rt.tokenGateway.GetLimit(ctx, key)
	if err != nil {
		return false, err
	}
	if limit == 0 {
		limit = rt.settings.Ratelimit
	}
	total, err := rt.requestGateway.Count(ctx, key)
	if err != nil {
		return false, err
	}
	err = rt.requestGateway.Save(ctx, key)
	if err != nil {
		return false, err
	}
	if total >= limit {
		if err := rt.lockGateway.Lock(ctx, key, time.Second*time.Duration(rt.settings.ExpirationTime)); err != nil {
			return false, err
		}
		return false, nil
	}
	return total <= limit, nil
}
