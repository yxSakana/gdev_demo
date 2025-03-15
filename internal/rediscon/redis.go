package rediscon

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yxSakana/gdev_demo/internal/consts"
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

func AddUvAndPv(uid, id uint64, ct consts.ContentType) error {
	key := GetCacheKey(id, ct)
	_, err := Rdb.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		Rdb.PFAdd(context.Background(), key, uid)
		Rdb.Incr(context.Background(), key)
		return nil
	})
	return err
}

func GetUvAndPv(id uint64, ct consts.ContentType) (uv int64, pv int64, err error) {
	key := GetCacheKey(id, ct)
	cmds, err := Rdb.Pipelined(context.Background(), func(rdb redis.Pipeliner) error {
		Rdb.PFCount(context.Background(), key)
		Rdb.Get(context.Background(), key)
		return nil
	})
	if err != nil {
		return
	}

	if len(cmds) != 2 {
		return 0, 0, fmt.Errorf("unexpected pipeline response length")
	}

	uv, err = cmds[0].(*redis.IntCmd).Result()
	if err != nil {
		return
	}
	pv, err = cmds[1].(*redis.IntCmd).Result()
	if err != nil {
		return
	}

	return uv, pv, nil
}

func AddUV(uid, id uint64, ct consts.ContentType) error {
	key := GetCacheKey(id, ct) + ":uv"
	return Rdb.PFAdd(context.Background(), key, uid).Err()
}

func GetUV(id uint64, ct consts.ContentType) (int64, error) {
	key := GetCacheKey(id, ct) + ":uv"
	return Rdb.PFCount(context.Background(), key).Result()
}

func AddPV(id uint64, ct consts.ContentType) error {
	key := GetCacheKey(id, ct) + ":pv"
	return Rdb.Incr(context.Background(), key).Err()
}

func GetPV(id uint64, ct consts.ContentType) (int64, error) {
	key := GetCacheKey(id, ct) + ":pv"
	return Rdb.Get(context.Background(), key).Int64()
}
