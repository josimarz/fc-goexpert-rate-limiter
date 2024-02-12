package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/josimarz/fc-goexpert-rate-limiter/configs"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/gateway"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/handler"
	redis_gateway "github.com/josimarz/fc-goexpert-rate-limiter/internal/infra/db/redis/gateway"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/middleware"
	"github.com/josimarz/fc-goexpert-rate-limiter/internal/ratelimiter"
	"github.com/redis/go-redis/v9"
)

type Token struct {
	Token string `json:"token"`
	Limit int    `json:"limit"`
}

var (
	config         *configs.Config
	redisClient    *redis.Client
	settings       *ratelimiter.Settings
	lockGateway    gateway.LockGateway
	requestGateway gateway.RequestGateway
	tokenGateway   gateway.TokenGateway
	rateLimiter    ratelimiter.RateLimiter
	mid            middleware.Middleware
)

func main() {
	loadConfig()
	connectToRedis()
	createTokens()
	initSettings()
	initGateways()
	initRateLimiter()
	initMiddleware()
	startServer()
}

func loadConfig() {
	var err error
	config, err = configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}
}

func connectToRedis() {
	addr := fmt.Sprintf("%v:%v", config.Host, config.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.Password,
		DB:       config.DB,
	})
}

func createTokens() {
	r, err := os.ReadFile("./assets/tokens.json")
	if err != nil {
		log.Fatalf("unable to load tokens: %v", err)
	}
	tokens := []Token{}
	err = json.Unmarshal([]byte(r), &tokens)
	if err != nil {
		log.Fatalf("unable to parse tokens: %v", err)
	}
	ctx := context.Background()
	for _, token := range tokens {
		redisClient.Set(ctx, token.Token, token.Limit, 0)
	}
}

func initSettings() {
	settings = ratelimiter.NewSettings(config.RateLimit, config.ExpirationTime, config.LimitByToken)
}

func initGateways() {
	lockGateway = redis_gateway.NewLockRedisGateway(redisClient)
	requestGateway = redis_gateway.NewRequestRedisGateway(redisClient)
	tokenGateway = redis_gateway.NewTokenRedisGateway(redisClient)
}

func initRateLimiter() {
	rateLimiter = ratelimiter.NewDefaultRateLimiter(settings, lockGateway, requestGateway, tokenGateway)
}

func initMiddleware() {
	mid = middleware.NewRateLimiterMiddleware(rateLimiter)
}

func startServer() {
	ctx := context.Background()
	mux := http.NewServeMux()
	mux.Handle("/", mid.Execute(ctx, &handler.DefaultHandler{}))
	http.ListenAndServe(":8080", mux)
}
