package gateway

import (
	"context"
	"testing"
)

var (
	tokenGateway *TokenInMemoryGateway
)

func TestNewTokenInMemoryGateway(t *testing.T) {
	tokenGateway = NewTokenInMemoryGateway()
	if tokenGateway == nil {
		t.Errorf("NewTokenInMemoryGateway() = %v, should be defined", nil)
	}
}

func TestTokenInMemoryGateway(t *testing.T) {
	ctx := context.Background()
	token := "p7eWgd0PvJcqB3ea45pw3k5thpWaqpI12RGYU3MiP91Kgao5MCXtlFtL2rwISxYL"
	t.Run("GetLimit", func(t *testing.T) {
		if limit, err := tokenGateway.GetLimit(ctx, token); limit != 100 || err != nil {
			t.Errorf("GetLimit(%v, %v) = (%v, %v), want (%v, %v)", ctx, token, limit, err, 100, nil)
		}
		token = "7Wk8BMQZjsfYluEhUy1KS0pbrMv1YEheLi93kXmGGlIy0XEsgpZpH128AYzTUEEr"
		if limit, err := tokenGateway.GetLimit(ctx, token); limit != 0 || err != nil {
			t.Errorf("GetLimit(%v, %v) = (%v, %v), want (%v, %v)", ctx, token, limit, err, 0, nil)
		}
	})
}
