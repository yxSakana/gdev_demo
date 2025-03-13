package rediscon

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/yxSakana/gdev_demo/settings"
)

var Rdb *redis.Client

func init() {
	cfg := settings.Settings.Database.Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
