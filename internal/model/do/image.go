package do

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/rediscon"
	"github.com/yxSakana/gdev_demo/utility"
)

type ImageCollection struct {
	ID          uint64    `json:"id" redis:"id"`
	UserID      uint64    `json:"user_id" redis:"user_id"`
	Uploader    string    `json:"uploader" redis:"uploader"`
	Title       string    `json:"title" redis:"title"`
	Description string    `json:"description" redis:"description"`
	CoverUrl    string    `json:"cover_url" redis:"cover_url"`
	Tags        []string  `json:"tags" redis:"tags"`
	Number      int       `json:"number" redis:"number"`
	View        int       `json:"view" redis:"view"`
	Like        int       `json:"like" redis:"like"`
	CreatedAt   time.Time `json:"created_at" redis:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" redis:"updated_at"`
}

type Image struct {
	ID           uint64 `json:"id"`
	CollectionID uint64 `json:"collection_id"`
	ImageUrl     string `json:"image_url"`
}

func (ic *ImageCollection) GetCacheKey(id uint64) string {
	return fmt.Sprintf("%s:%d", consts.CacheImage, id)
}

func (ic *ImageCollection) ToCacheMap() map[string]interface{} {
	empty := "0"
	if ic == nil {
		empty = "1"
	}
	return map[string]interface{}{
		"id":         ic.ID,
		"uid":        ic.UserID,
		"uploader":   ic.Uploader,
		"title":      ic.Title,
		"cu":         ic.CoverUrl,
		"tags":       strings.Join(ic.Tags, ","),
		"number":     ic.Number,
		"view":       ic.View,
		"like":       ic.Like,
		"created_at": ic.CreatedAt,
		"updated_at": ic.UpdatedAt,
		"empty":      empty,
	}
}

func (ic *ImageCollection) RefreshFromCacheMap(cacheRet map[string]string) {
	if ic == nil {
		return
	}

	ic.ID = utility.MustUint64(cacheRet["id"])
	ic.UserID = utility.MustUint64(cacheRet["uid"])
	ic.Uploader = cacheRet["uploader"]
	ic.Title = cacheRet["title"]
	ic.Description = cacheRet["desc"]
	ic.CoverUrl = cacheRet["cu"]
	ic.Tags = strings.Split(cacheRet["tags"], ",")
	ic.Number = utility.MustInt(cacheRet["number"])
	ic.View = utility.MustInt(cacheRet["view"]) // TODO: 是否 mysql.view + cache.uv？
	ic.Like = utility.MustInt(cacheRet["like"])
	ic.CreatedAt = utility.MustTime(cacheRet["created_at"])
	ic.UpdatedAt = utility.MustTime(cacheRet["updated_at"])
}

func (ic *ImageCollection) GetFromCache(ctx context.Context, id uint64) error {
	key := ic.GetCacheKey(id)

	ret, err := rediscon.Rdb.HGet(ctx, key, "empty").Result()
	if err != nil {
		return err
	}
	if ret == "1" {
		return consts.ErrCacheIsNil
	}

	desc, err := rediscon.Rdb.Get(ctx, key+":desc").Result()
	if err != nil {
		return err
	}
	cacheRet, err := rediscon.Rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return err
	}
	uv, _, err := rediscon.GetUvAndPv(ctx, id, ic)
	if err != nil {
		return err
	}

	cacheRet["desc"] = desc
	cacheRet["view"] = fmt.Sprintf("%d", uv)

	ic.RefreshFromCacheMap(cacheRet)
	return nil
}

func (ic *ImageCollection) SaveToCache(ctx context.Context, id uint64) error {
	key := ic.GetCacheKey(id)

	if err := rediscon.Rdb.Set(ctx, key+":desc", ic.Description, 0).Err(); err != nil {
		return err
	}
	return rediscon.Rdb.HSet(ctx, key, ic.ToCacheMap()).Err()
}

func (ic *ImageCollection) DelCache(ctx context.Context, id uint64) error {
	key := ic.GetCacheKey(id)

	if err := rediscon.Rdb.Del(ctx, key+":desc").Err(); err != nil {
		return err
	}
	if err := rediscon.Rdb.Del(ctx, key).Err(); err != nil {
		return err
	}
	return nil
}
