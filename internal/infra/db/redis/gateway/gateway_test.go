package gateway

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	redisClient *redis.Client
	clients     []struct {
		token string
		limit int
	} = []struct {
		token string
		limit int
	}{
		{"O7kGWyOjgv6GzbHZRcW2yg41rfESRfTIL8kVO0IGiFyiVjJiPEKI6G0PPYRdCMOU", 10},
		{"mybEY5EGi3ham5jpuVLqlklsrB7rc0fQOEbZBThE60XH2YIt3djtgZTd9It88JPP", 50},
		{"qr917mOtltQv1wBRcoxKuLL2D1GHvyrZD1EBCHsoJ7oFA0XyTg3mBoPW5wjBuY6c", 100},
		{"FUjwIqlitgOyym5MOxdZOYUHMIe6mwuQT5mteJVf6EtJEkWYLcSz1vZwLaqMXdEk", 150},
		{"WBLUM3qEu6d1elE2QGQi35e1DYsaQ7VBgWGquhtA9TVllLRFIaje1RRTD0iyT9j7", 200},
	}
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}
	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	endpoint, err := redisC.Endpoint(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr: endpoint,
	})
	for _, client := range clients {
		key := fmt.Sprintf("limit:%s", client.token)
		redisClient.Set(ctx, key, client.limit, 0)
	}
	os.Exit(m.Run())
	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}
