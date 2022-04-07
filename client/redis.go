package client

import (
	"context"
	"log"

	"github.com/go-redis/redis"
)

// RedisProxy redis的proxy
type RedisProxy struct {
	Address  string
	Password string
	DB       int
}

// NewRedisClient 新建一个redis的client
func NewRedisClient(ctx context.Context, proxy RedisProxy) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     proxy.Address,
		Password: proxy.Password,
		DB:       proxy.DB,
	})
	_, err := cli.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("New Redis Client Error %+v", err)
		return nil, err
	}
	return cli, nil
}
