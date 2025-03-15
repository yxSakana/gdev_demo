package do

import (
	"fmt"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/utility"
	"strings"
	"time"
)

type Novel struct {
	ID            uint64    `json:"id"`
	UserID        uint64    `json:"user_id"`
	Uploader      string    `json:"Uploader"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverUrl      string    `json:"cover_url"`
	Tags          []string  `json:"tags"`
	Status        uint8     `json:"status"`
	ChapterNumber uint      `json:"chapter_number"`
	WordCount     uint      `json:"word_count"`
	View          uint      `json:"view"`
	Like          uint      `json:"like"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type NovelChapter struct {
	ID        uint64    `json:"id"`
	NovelId   uint64    `json:"novel_id"`
	Title     string    `json:"title"`
	Number    int       `json:"num"`
	Content   string    `json:"content"`
	WordCount uint      `json:"word_count"`
	View      uint      `json:"view"`
	Like      uint      `json:"like"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *Novel) GetCacheKey(id uint64) string {
	return fmt.Sprintf("%s:%d", consts.CacheNovel, id)
}

func (n *Novel) ToCacheMap() map[string]interface{} {
	return map[string]interface{}{
		"id":             n.ID,
		"uid":            n.UserID,
		"uploader":       n.Uploader,
		"title":          n.Title,
		"cu":             n.CoverUrl,
		"tags":           strings.Join(n.Tags, ","),
		"status":         n.Status,
		"chapter_number": n.ChapterNumber,
		"wc":             n.WordCount,
		"view":           n.View,
		"like":           n.Like,
		"created_at":     n.CreatedAt,
		"updated_at":     n.UpdatedAt,
		"empty":          0,
	}
}

func (n *Novel) FromCacheMap(cacheRet map[string]string) {
	if n == nil {
		return
	}

	n.ID = utility.MustUint64(cacheRet["id"])
	n.UserID = utility.MustUint64(cacheRet["uid"])
	n.Uploader = cacheRet["uploader"]
	n.Title = cacheRet["title"]
	n.Description = cacheRet["desc"]
	n.CoverUrl = cacheRet["cu"]
	n.Tags = strings.Split(cacheRet["tags"], ",")
	n.Status = utility.MustUint8(cacheRet["status"])
	n.ChapterNumber = utility.MustUint(cacheRet["chapter_number"])
	n.WordCount = utility.MustUint(cacheRet["wc"])
	n.View = utility.MustUint(cacheRet["view"]) // TODO: 是 mysql.view + cache.uv 吗
	n.Like = utility.MustUint(cacheRet["like"])
	n.CreatedAt = utility.MustTime(cacheRet["created_at"])
	n.UpdatedAt = utility.MustTime(cacheRet["updated_at"])
}
