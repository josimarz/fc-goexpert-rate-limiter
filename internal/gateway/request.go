package gateway

import (
	"context"
)

type RequestGateway interface {
	Save(context.Context, string) error
	Count(context.Context, string) (int, error)
}
