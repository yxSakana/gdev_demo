package novel

import (
	"context"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"github.com/yxSakana/gdev_demo/internal/rediscon"
	"github.com/yxSakana/gdev_demo/utility"
)

func GetFromCache(key string) *do.Novel {
	ctx := context.Background()

	if exists, err := rediscon.Rdb.Exists(ctx, key).Result(); err != nil || exists == 0 {
		return nil
	}

	cacheRet, err := rediscon.Rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil
	}

	e := entity.Novel{
		ID:          utility.MustUint64(cacheRet["id"]),
		UserID:      utility.MustUint64(cacheRet["uid"]),
		Uploader:    cacheRet["uploader"],
		Title:       cacheRet["title"],
		Description: cacheRet["desc"],
		CoverUrl:    cacheRet["cu"],
		Status:      utility.MustUint8(cacheRet["status"]),
		WordCount:   utility.MustUint(cacheRet["wc"]),
		View:        utility.MustUint(cacheRet["view"]),
		Like:        utility.MustUint(cacheRet["like"]),
		CreatedAt:   utility.MustTime(cacheRet["created_at"]),
		UpdatedAt:   utility.MustTime(cacheRet["updated_at"]),
	}
	return &do.Novel{Novel: &e}

	//var e entity.Novel
	//if err := mapstructure.Decode(cacheRet, &e); err != nil {
	//	log.Println(err)
	//	return nil
	//}
	//return &do.Novel{Novel: &e}
}

func SetCache(key string, n *do.Novel) error {
	ctx := context.Background()

	return rediscon.Rdb.HMSet(ctx, key, map[string]interface{}{
		"id":         n.ID,
		"uid":        n.UserID,
		"uploader":   n.Uploader,
		"title":      n.Title,
		"desc":       n.Description,
		"cu":         n.CoverUrl,
		"status":     n.Status,
		"wc":         n.WordCount,
		"view":       n.View,
		"like":       n.Like,
		"created_at": n.CreatedAt,
		"updated_at": n.UpdatedAt,
	}).Err()

	//var cac map[string]interface{}
	//if err := mapstructure.Decode(*n.Novel, &cac); err != nil {
	//	log.Printf("set cache err: %v", err)
	//	return err
	//}
	//log.Printf("set cache key: %#v", cac)
	//return rediscon.Rdb.HMSet(ctx, key, cac).Err()
}
