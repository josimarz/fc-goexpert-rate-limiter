package gateway

import "context"

type Request struct {
	key   string
	count int
}

type RequestInMemoryGateway struct {
	requests []*Request
}

func NewRequestInMemoryGateway() *RequestInMemoryGateway {
	return &RequestInMemoryGateway{
		requests: []*Request{},
	}
}

func (g *RequestInMemoryGateway) Save(ctx context.Context, key string) error {
	for _, r := range g.requests {
		if r.key == key {
			r.count++
			return nil
		}
	}
	g.requests = append(g.requests, &Request{key: key, count: 1})
	return nil
}

func (g *RequestInMemoryGateway) Count(ctx context.Context, key string) (int, error) {
	for _, r := range g.requests {
		if r.key == key {
			return r.count, nil
		}
	}
	return 0, nil
}
