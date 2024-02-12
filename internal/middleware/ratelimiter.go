package middleware

import (
	"context"
	"net/http"

	"github.com/josimarz/fc-goexpert-rate-limiter/internal/ratelimiter"
)

type RateLimiterMiddleware struct {
	rateLimiter ratelimiter.RateLimiter
}

func NewRateLimiterMiddleware(rateLimiter ratelimiter.RateLimiter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{rateLimiter}
}

func (m *RateLimiterMiddleware) Execute(ctx context.Context, next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		proceed, err := m.rateLimiter.CanGo(ctx, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !proceed {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
