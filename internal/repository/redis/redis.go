package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/madxiii/hackatone/configs"
)

func InitRDB(ctx context.Context, cfg configs.Redis) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Address,
	})

	return client, client.Ping(ctx).Err()
}
