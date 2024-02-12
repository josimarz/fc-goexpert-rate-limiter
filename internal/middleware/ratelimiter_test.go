package middleware

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/josimarz/fc-goexpert-rate-limiter/internal/handler"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/infra/db/inmemory/gateway"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/ratelimiter"
)

var (
	mid *RateLimiterMiddleware
	rt  *MockRateLimiter
)

type MockRateLimiter struct {
	settings       *ratelimiter.Settings
	lockGateway    *gateway.LockInMemoryGateway
	requestGateway *gateway.RequestInMemoryGateway
	tokenGateway   *gateway.TokenInMemoryGateway
}

func (rt *MockRateLimiter) CanGo(ctx context.Context, r *http.Request) (bool, error) {
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

func TestMain(m *testing.M) {
	settings := ratelimiter.NewSettings(10, 5, true)
	lockGateway := gateway.NewLockInMemoryGateway()
	requestGateway := gateway.NewRequestInMemoryGateway()
	tokenGateway := gateway.NewTokenInMemoryGateway()
	rt = &MockRateLimiter{settings, lockGateway, requestGateway, tokenGateway}
	os.Exit(m.Run())
}

func TestNewRateLimiterMiddleware(t *testing.T) {
	mid = NewRateLimiterMiddleware(rt)
	if mid == nil {
		t.Errorf("NewRateLimiterMiddleware(%v) = %v, should be defined", rt, nil)
	}
}

func TestRateLimiterMiddleware(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.RemoteAddr = "127.0.0.1:8080"
	t.Run("Execute", func(t *testing.T) {
		h := mid.Execute(ctx, &handler.DefaultHandler{})
		for i := 0; i < 10; i++ {
			rr := httptest.NewRecorder()
			h.ServeHTTP(rr, req)
			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
			expected := "Welcome"
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
			}
		}
	})
}
