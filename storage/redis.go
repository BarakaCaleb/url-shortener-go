package storage

import (
	"context"
	"errors"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient interface {
	Set(code string, longURL string) error
	Get(code string) (string, error)
}

type redisStorage struct {
	client *redis.Client
}

func NewRedisClient(addr string) RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &redisStorage{client: rdb}
}

func (r *redisStorage) Set(code string, longURL string) error {
	return r.client.Set(ctx, code, longURL, 0).Err()
}

func (r *redisStorage) Get(code string) (string, error) {
	val, err := r.client.Get(ctx, code).Result()
	if err == redis.Nil {
		return "", errors.New("code not found")
	}
	return val, err
}
