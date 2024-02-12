package ratelimiter

import (
	"context"
	"net/http"
	"testing"

	"github.com/josimarz/fc-goexpert-rate-limiter/internal/infra/db/inmemory/gateway"
)

var (
	rateLimiter *DefaultRateLimiter
)

func TestNewDefaultRateLimiter(t *testing.T) {
	settings := NewSettings(10, 60, true)
	lockGateway := gateway.NewLockInMemoryGateway()
	requestGateway := gateway.NewRequestInMemoryGateway()
	tokenGateway := gateway.NewTokenInMemoryGateway()
	rateLimiter = NewDefaultRateLimiter(settings, lockGateway, requestGateway, tokenGateway)
	if rateLimiter == nil {
		t.Errorf("NewDefaultRateLimiter(%v, %v, %v, %v) = %v, should be defined", settings, lockGateway, requestGateway, tokenGateway, nil)
	}
}

func TestDefaultRateLimiter(t *testing.T) {
	ctx := context.Background()
	t.Run("CanGo", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.RemoteAddr = "127.0.0.1:0000"
		for i := 0; i < 10; i++ {
			if proceed, err := rateLimiter.CanGo(ctx, req); !proceed || err != nil {
				t.Errorf("CanGo(%v, %v) = (%v, %v), want (%v, %v)", ctx, req, proceed, err, true, nil)
			}
		}
		if proceed, err := rateLimiter.CanGo(ctx, req); proceed || err != nil {
			t.Errorf("CanGo(%v, %v) = (%v, %v), want (%v, %v)", ctx, req, proceed, err, false, nil)
		}
		req.Header.Add("x-api-key", "p7eWgd0PvJcqB3ea45pw3k5thpWaqpI12RGYU3MiP91Kgao5MCXtlFtL2rwISxYL")
		for i := 0; i < 100; i++ {
			if proceed, err := rateLimiter.CanGo(ctx, req); !proceed || err != nil {
				t.Errorf("CanGo(%v, %v) = (%v, %v), want (%v, %v)", ctx, req, proceed, err, true, nil)
			}
		}
		if proceed, err := rateLimiter.CanGo(ctx, req); proceed || err != nil {
			t.Errorf("CanGo(%v, %v) = (%v, %v), want (%v, %v)", ctx, req, proceed, err, false, nil)
		}
	})
}
