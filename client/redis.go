package client

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

// RedisConfig redis的配置文件
type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

// RedisProxy redis的proxy
type RedisProxy struct {
	redisCli redis.Client
}

// NewRedisClient 新建一个redis的client
func NewRedisClient(proxy RedisConfig) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     proxy.Address,
		Password: proxy.Password,
		DB:       proxy.DB,
	})
	_, err := cli.Ping().Result()
	if err != nil {
		log.Fatalf("New Redis Client Error %+v", err)
		return nil, err
	}
	return cli, nil
}

// RedisBatchMSet 批量加入redis
func (cli *RedisProxy) RedisBatchMSet(ctx context.Context, keys, values []string, count int) error {
	if len(keys) != len(values) {
		return fmt.Errorf("key value not match")
	}
	params := make([]interface{}, 0)
	for i := 0; i < len(keys); i++ {
		params = append(params, keys[i], values[i])
		if len(params)%count == 0 {
			_, err := cli.redisCli.MSet(ctx, params).Result()
			if err != nil {
				log.Fatalf("Mset in redis err: [%+v],the datas is [%+v]", err, params)
				continue
			}
		}
		params = make([]interface{}, 0)
	}

	if len(params) != 0 {
		_, err := cli.redisCli.MSet(ctx, params).Result()
		if err != nil {
			log.Fatalf("Mset in redis err: [%+v],the datas is [%+v]", err, params)
			return err
		}
	}

	return nil
}
