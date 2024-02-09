package gateway

import (
	"context"
	"testing"
)

var (
	tokenGateway *TokenRedisGateway
)

func TestNewTokenRedisGateway(t *testing.T) {
	tokenGateway = NewTokenRedisGateway(redisClient)
	if tokenGateway == nil {
		t.Errorf("NewTokenRedisGateway(%v) = %v, want to be defined", redisClient, tokenGateway)
	}
}

func TestTokenRedisGateway(t *testing.T) {
	ctx := context.Background()
	t.Run("GetLimit", func(t *testing.T) {
		token := clients[0].token
		limit := clients[0].limit
		if got, err := tokenGateway.GetLimit(ctx, token); got != limit || err != nil {
			t.Errorf("GetLimit(%v, %v) = (%v, %v), want (%v, %v)", ctx, token, got, err, 10, nil)
		}
		token = "Jhdw06vCSW0cKoi8hJe0wmc8j4A4cSqCBbT2spWQj7gsYooXOd1rmfLlguHyASq3"
		if got, err := tokenGateway.GetLimit(ctx, token); got != 0 || err != nil {
			t.Errorf("GetLimit(%v, %v) = (%v, %v), want (%v, %v)", ctx, token, got, err, 0, nil)
		}
	})
}
