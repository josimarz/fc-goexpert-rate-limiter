package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/josimarz/fc-goexpert-rate-limiter/configs"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/handler"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/infra/db/redis/gateway"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/middleware"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/ratelimiter"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	settings := ratelimiter.NewSettings(config.RateLimit, config.ExpirationTime, config.LimitByToken)
	lockGateway := gateway.NewLockRedisGateway(redisClient)
	requestGateway := gateway.NewRequestRedisGateway(redisClient)
	tokenGateway := gateway.NewTokenRedisGateway(redisClient)
	rateLimiter := ratelimiter.NewDefaultRateLimiter(settings, lockGateway, requestGateway, tokenGateway)
	middleware := middleware.NewRateLimiterMiddleware(rateLimiter)
	mux := http.NewServeMux()
	mux.Handle("/", middleware.Execute(ctx, &handler.DefaultHandler{}))
	http.ListenAndServe(":8080", mux)
}
