package image

//
//import (
//	"context"
//	"gorm.io/gorm"
//	"strings"
//	"time"
//
//	"github.com/yxSakana/gdev_demo/internal/consts"
//	"github.com/yxSakana/gdev_demo/internal/model/do"
//	"github.com/yxSakana/gdev_demo/internal/rediscon"
//	"github.com/yxSakana/gdev_demo/utility"
//)
//
//func GetFromCache(id uint64) (*do.ImageCollection, error) {
//	ctx := context.Background()
//	key := rediscon.GetCacheKey(id, consts.ImageCt)
//
//	ret, err := rediscon.Rdb.HGet(ctx, key, "empty").Result()
//	if err == nil && ret == "1" {
//		return nil, gorm.ErrRecordNotFound
//	}
//
//	desc, err := rediscon.Rdb.Get(ctx, key+":desc").Result()
//	if err != nil {
//		return nil, err
//	}
//	cacheRet, err := rediscon.Rdb.HGetAll(ctx, key).Result()
//	if err != nil {
//		return nil, err
//	}
//	uv, _, err := rediscon.GetUvAndPv(id, consts.ImageCt)
//	if err != nil {
//		return nil, err
//	}
//
//	e := do.ImageCollection{
//		ID:          utility.MustUint64(cacheRet["id"]),
//		UserID:      utility.MustUint64(cacheRet["uid"]),
//		Uploader:    cacheRet["uploader"],
//		Title:       cacheRet["title"],
//		Description: desc,
//		CoverUrl:    cacheRet["cu"],
//		Tags:        strings.Split(cacheRet["tags"], ","),
//		Number:      utility.MustInt(cacheRet["number"]),
//		View:        int(uv), // TODO: 是 mysql.view + cache.uv 吗
//		Like:        utility.MustInt(cacheRet["like"]),
//		CreatedAt:   utility.MustTime(cacheRet["created_at"]),
//		UpdatedAt:   utility.MustTime(cacheRet["updated_at"]),
//	}
//
//	return &e, nil
//}
//
//func SetCache(id uint64, ic *do.ImageCollection) error {
//	ctx := context.Background()
//	key := rediscon.GetCacheKey(id, consts.ImageCt)
//
//	if ic == nil {
//		expire := 60 * time.Second
//		rediscon.Rdb.HSet(ctx, key, "empty", 1)
//		//rediscon.Rdb.Set(ctx, key+":desc", "", expire)
//		rediscon.Rdb.Expire(ctx, key, expire)
//		return nil
//	}
//
//	if err := rediscon.Rdb.Set(ctx, key+":desc", ic.Description, 0).Err(); err != nil {
//		return err
//	}
//
//	return rediscon.Rdb.HMSet(ctx, key, map[string]interface{}{
//		"id":         ic.ID,
//		"uid":        ic.UserID,
//		"uploader":   ic.Uploader,
//		"title":      ic.Title,
//		"cu":         ic.CoverUrl,
//		"tags":       strings.Join(ic.Tags, ","),
//		"number":     ic.Number,
//		"view":       ic.View,
//		"like":       ic.Like,
//		"created_at": ic.CreatedAt,
//		"updated_at": ic.UpdatedAt,
//		"empty":      0,
//	}).Err()
//}
//
//func DelCache(id uint64) error {
//	ctx := context.Background()
//	key := rediscon.GetCacheKey(id, consts.ImageCt)
//
//	if err := rediscon.Rdb.Del(ctx, key+":desc").Err(); err != nil {
//		return err
//	}
//	if err := rediscon.Rdb.Del(ctx, key).Err(); err != nil {
//		return err
//	}
//	return nil
//}
