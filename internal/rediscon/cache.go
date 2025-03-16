package rediscon

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/yxSakana/gdev_demo/internal/consts"
)

type Cacheable interface {
	GetCacheKey(id uint64) string
	GetFromCache(ctx context.Context, id uint64) error
	SaveToCache(ctx context.Context, id uint64) error
	DelCache(ctx context.Context, id uint64) error
}

type HashCacheable interface {
	Cacheable
	ToCacheMap() map[string]interface{}
	RefreshFromCacheMap(cacheRet map[string]string)
}

func AddUvAndPv(ctx context.Context, uid, id uint64, obj Cacheable) error {
	key := obj.GetCacheKey(id)

	_, err := Rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.PFAdd(ctx, key+":uv", uid).Err(); err != nil {
			return err
		}
		return p.Incr(ctx, key+":pv").Err()
	})
	return err
}

func GetUvAndPv(ctx context.Context, id uint64, obj Cacheable) (uv int64, pv int64, err error) {
	key := obj.GetCacheKey(id)

	cmds, err := Rdb.Pipelined(ctx, func(p redis.Pipeliner) error {
		if err := p.PFCount(ctx, key+":uv").Err(); err != nil {
			return err
		}
		return p.Get(ctx, key+":pv").Err()
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
	pv, err = cmds[1].(*redis.StringCmd).Int64()
	if err != nil {
		return
	}

	return uv, pv, nil
}

func AddUV(uid, id uint64, obj Cacheable) error {
	key := obj.GetCacheKey(id) + ":uv"
	return Rdb.PFAdd(context.Background(), key, uid).Err()
}

func GetUV(id uint64, obj Cacheable) (int64, error) {
	key := obj.GetCacheKey(id) + ":uv"
	return Rdb.PFCount(context.Background(), key).Result()
}

func AddPV(id uint64, obj Cacheable) error {
	key := obj.GetCacheKey(id) + ":pv"
	return Rdb.Incr(context.Background(), key).Err()
}

func GetPV(id uint64, obj Cacheable) (int64, error) {
	key := obj.GetCacheKey(id) + ":pv"
	return Rdb.Get(context.Background(), key).Int64()
}

func NilCache(ctx context.Context, id uint64, obj Cacheable) error {
	key := obj.GetCacheKey(id)
	Rdb.HSet(ctx, key, "empty", 1)
	Rdb.Expire(ctx, key, consts.NilCacheExpire)
	return nil
}
