package gateway

import (
	"context"
)

type TokenGateway interface {
	GetLimit(context.Context, string) (int, error)
}
