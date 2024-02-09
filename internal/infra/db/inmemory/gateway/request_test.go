package gateway

import (
	"context"
	"testing"
)

var (
	requestGateway *RequestInMemoryGateway
)

func TestNewRequestInMemoryGateway(t *testing.T) {
	requestGateway = NewRequestInMemoryGateway()
	if requestGateway == nil {
		t.Errorf("NewRequestInMemoryGateway() = %v, should be defined", nil)
	}
}

func TestRequestInMemoryGateway(t *testing.T) {
	ctx := context.Background()
	key := "31ce7OfgTPZ38JMwFKKVZSdqeeMEKb7fYF1wZ8EtufoP0eGbI0PHZy6gbGrZqyCb"
	t.Run("Save", func(t *testing.T) {
		if err := requestGateway.Save(ctx, key); err != nil {
			t.Errorf("Save(%v, %v) = %v, want %v", ctx, key, err, nil)
		}
		if err := requestGateway.Save(ctx, key); err != nil {
			t.Errorf("Save(%v, %v) = %v, want %v", ctx, key, err, nil)
		}
	})
	t.Run("Count", func(t *testing.T) {
		if total, err := requestGateway.Count(ctx, key); total != 2 || err != nil {
			t.Errorf("Count(%v, %v) = (%v, %v), want (%v, %v)", ctx, key, total, err, 2, nil)
		}
		key = "Wz5T1xugbVrrlZ6bFuPiHAEJPhx6YX1gFuhD6gvCvaDWVLzaLLOdAXCOgi9rb52M"
		if total, err := requestGateway.Count(ctx, key); total != 0 || err != nil {
			t.Errorf("Count(%v, %v) = (%v, %v), want (%v, %v)", ctx, key, total, err, 0, nil)
		}
	})
}
