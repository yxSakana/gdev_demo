package novel

//
//import (
//	"context"
//	"errors"
//	"github.com/redis/go-redis/v9"
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
//func GetFromCache(id uint64) (*do.Novel, error) {
//	ctx := context.Background()
//	key := rediscon.GetCacheKey(id, consts.NovelCt)
//
//	ret, err := rediscon.Rdb.HGet(ctx, key, "empty").Result()
//	if errors.Is(err, redis.Nil) || (err == nil && ret == "1") {
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
//	uv, _, err := rediscon.GetUvAndPv(id, consts.NovelCt)
//	if err != nil {
//		return nil, err
//	}
//
//	e := do.Novel{
//		ID:            utility.MustUint64(cacheRet["id"]),
//		UserID:        utility.MustUint64(cacheRet["uid"]),
//		Uploader:      cacheRet["uploader"],
//		Title:         cacheRet["title"],
//		Description:   desc,
//		CoverUrl:      cacheRet["cu"],
//		Tags:          strings.Split(cacheRet["tags"], ","),
//		Status:        utility.MustUint8(cacheRet["status"]),
//		ChapterNumber: utility.MustUint(cacheRet["chapter_number"]),
//		WordCount:     utility.MustUint(cacheRet["wc"]),
//		View:          uint(uv), // TODO: 是 mysql.view + cache.uv 吗
//		Like:          utility.MustUint(cacheRet["like"]),
//		CreatedAt:     utility.MustTime(cacheRet["created_at"]),
//		UpdatedAt:     utility.MustTime(cacheRet["updated_at"]),
//	}
//
//	return &e, nil
//}
//
//func SetCache(id uint64, n *do.Novel) error {
//	ctx := context.Background()
//	key := rediscon.GetCacheKey(id, consts.NovelCt)
//
//	if n == nil {
//		expire := 60 * time.Second
//		rediscon.Rdb.HSet(ctx, key, "empty", 1)
//		rediscon.Rdb.Expire(ctx, key, expire)
//		return nil
//	}
//
//	if err := rediscon.Rdb.Set(ctx, key+":desc", n.Description, 0).Err(); err != nil {
//		return err
//	}
//
//	return rediscon.Rdb.HMSet(ctx, key, map[string]interface{}{
//		"id":             n.ID,
//		"uid":            n.UserID,
//		"uploader":       n.Uploader,
//		"title":          n.Title,
//		"cu":             n.CoverUrl,
//		"tags":           strings.Join(n.Tags, ","),
//		"status":         n.Status,
//		"chapter_number": n.ChapterNumber,
//		"wc":             n.WordCount,
//		"view":           n.View,
//		"like":           n.Like,
//		"created_at":     n.CreatedAt,
//		"updated_at":     n.UpdatedAt,
//		"empty":          0,
//	}).Err()
//}
//
//func DelCache(id uint64) error {
//	ctx := context.Background()
//	key := rediscon.GetCacheKey(id, consts.NovelCt)
//
//	if err := rediscon.Rdb.Del(ctx, key+":desc").Err(); err != nil {
//		return err
//	}
//	if err := rediscon.Rdb.Del(ctx, key).Err(); err != nil {
//		return err
//	}
//	return nil
//}
