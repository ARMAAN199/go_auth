package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/ARMAAN199/Go_EcomApi/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, cfg *config.RedisConfig) (*redis.Client, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       0,
		Protocol: 2,
	})

	res, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
		return nil, err
	}

	fmt.Println("Connected to Redis:", res)

	return redisClient, nil
}
