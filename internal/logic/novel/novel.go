package novel

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/dao"
	novelDao "github.com/yxSakana/gdev_demo/internal/dao/novel"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"github.com/yxSakana/gdev_demo/internal/model"
	"github.com/yxSakana/gdev_demo/internal/model/conv"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	"github.com/yxSakana/gdev_demo/internal/rediscon"
	novelService "github.com/yxSakana/gdev_demo/internal/service/novel"
	"github.com/yxSakana/gdev_demo/utility"
)

func CreateNovel(c *gin.Context, in *v1.CreateNovelReq) (uint64, error) {
	userInfo, err := user.GetUserinfo(c)
	if err != nil {
		return 0, err
	}

	filePath, err := utility.SaveFile(c, in.Cover, utility.CoverFt)

	novelDo := do.Novel{
		UserID:      userInfo.ID,
		Uploader:    userInfo.Username,
		Title:       in.Title,
		Description: in.Description,
		CoverUrl:    filePath,
		Status:      0,
	}

	db := dao.Ctx(c)
	if err := db.Transaction(func(tx *gorm.DB) error {
		err = novelDao.CreateNovel(tx, &novelDo)
		if err != nil {
			return err
		}
		return novelService.LinkNovelAndTags(tx, novelDo.ID, in.Tags)
	}); err != nil {
		return 0, err
	}

	return novelDo.ID, nil
}

func UploadChapter(c *gin.Context, req *v1.UploadChapterReq) (uint64, error) {
	db := dao.Ctx(c)

	ncDo := do.NovelChapter{
		NovelId:   req.NovelId,
		Title:     req.Title,
		Number:    req.Number,
		Content:   req.Content,
		WordCount: uint(len(req.Content)),
	}

	err := novelDao.CreateChapter(db, &ncDo)
	if err != nil {
		return 0, err
	}

	return ncDo.ID, nil
}

func QueryNovel(c *gin.Context, query model.NovelQueryInput) ([]do.Novel, error) {
	db := dao.Ctx(c)

	queryRes := db.Model(&entity.Novel{})
	if query.Id != nil {
		queryRes = queryRes.Where("id=?", *query.Id)
	}
	if query.Author != nil {
		queryRes = queryRes.Where("author LIKE ?", fmt.Sprintf("%%%s%%", *query.Author))
	}
	if query.Title != nil {
		queryRes = queryRes.Where("title LIKE ?", fmt.Sprintf("%%%s%%", *query.Title))
	}
	if query.WordCount != nil {
		queryRes = queryRes.Where("word_count >= ?", *query.WordCount)
	}
	if query.View != nil {
		queryRes = queryRes.Where("view >= ?", *query.View)
	}
	if query.Like != nil {
		queryRes = queryRes.Where("like >= ?", *query.Like)
	}
	if query.Tag != nil {
		queryRes = queryRes.
			Joins("LEFT JOIN novel_tag_rel ON novels.id = novel_tag_rel.novel_id").
			Joins("LEFT JOIN novel_tags ON novel_tag_rel.novel_tag_id = novel_tags.id").
			Where("novel_tags.name = ?", *query.Tag)
	}

	var novelIds []uint64
	err := queryRes.Pluck("id", &novelIds).Error
	if err != nil {
		return nil, err
	}

	var outs []do.Novel
	for _, i := range novelIds {
		n, err := DetailNovelByID(c, i)
		if err != nil {
			log.Printf("DetailNovelByID error: %v", err)
			continue
		}
		outs = append(outs, *n)
	}
	return outs, nil
}

func DetailNovelByID(c *gin.Context, nid uint64) (nDo *do.Novel, err error) {
	db := dao.Ctx(c)
	uid, err := user.GetUserID(c)
	if err != nil {
		return nil, err
	}
	nDo = new(do.Novel)
	defer func() {
		if err == nil && nDo != nil {
			if err := rediscon.AddUvAndPv(context.Background(), uid, nid, nDo); err != nil {
				log.Printf("add Uv&Pv err: %v", err)
			}
			if err := novelDao.UpdateNovel(db, nid, map[string]any{"view": nDo.View}); err != nil {
				log.Printf("from cache update novel: %#v", nDo)
			}
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := rediscon.NilCache(context.Background(), nid, nDo); err != nil {
				log.Printf("save to cache err: %v", err)
			}
		}
	}()

	if err := nDo.GetFromCache(context.Background(), nid); err == nil || errors.Is(err, consts.ErrCacheIsNil) {
		log.Printf("Get from cache novel: %#v", nDo)
		return nDo, err
	}

	nEntity, err := novelDao.GetNovelByID(db, nid)
	if err != nil {
		return nil, err
	}

	tags, err := GetNovelTags(db, nEntity.ID)
	if err != nil {
		return nil, err
	}

	*nDo = conv.NovelToDo(*nEntity, tags)
	if err := nDo.SaveToCache(context.Background(), nid); err != nil {
		log.Printf("save to cache err: %v", err)
	}
	return nDo, nil
}

func GetNovelTags(db *gorm.DB, nid uint64) ([]string, error) {
	var tags []string
	err := db.Model(&entity.NovelTag{}).
		Joins("LEFT JOIN novel_tag_rel ON novel_tags.id = novel_tag_rel.novel_tag_id").
		Joins("LEFT JOIN novels ON novel_tag_rel.novel_id = novels.id").
		Where("novels.id = ?", nid).
		Pluck("name", &tags).Error
	return tags, err
}

func GetNovelChapterIds(c *gin.Context, nid uint64) ([]uint64, error) {
	var ids []uint64
	err := dao.Ctx(c).Model(&entity.NovelChapter{}).
		Where("novel_id = ?", nid).
		Pluck("id", &ids).Error
	return ids, err
}

func UpdateNovel(c *gin.Context, nid uint64, in model.UpdateNovelInput) error {
	updateMap := make(map[string]interface{})

	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	for i := 0; i < t.NumField(); i++ {
		fieldVal := v.Field(i)
		if fieldVal.Kind() == reflect.Ptr && fieldVal.IsNil() {
			continue
		}
		field := t.Field(i)
		tag := field.Tag.Get("gorm")
		if tag == "-" {
			continue
		}

		updateMap[tag] = v.Field(i).Interface()
	}

	filePath, err := utility.SaveFile(c, in.Cover, utility.CoverFt)
	if err != nil && !errors.Is(err, utility.ErrFileHeaderIsNil) {
		log.Printf("save file err: %v", err)
		return err
	}
	updateMap["cover_url"] = filePath

	db := dao.Ctx(c)
	err = db.Transaction(func(tx *gorm.DB) error {
		if in.Tags != nil {
			if err := novelDao.DelNrt(tx, nid); err != nil {
				return err
			}

			if err := novelService.LinkNovelAndTags(tx, nid, *in.Tags); err != nil {
				return err
			}
		}

		return novelDao.UpdateNovel(tx, nid, updateMap)
	})
	if err != nil {
		return err
	}

	nDo := new(do.Novel)
	if err := nDo.DelCache(context.Background(), nid); err != nil {
		log.Printf("del cache err: %v", err)
	}
	return nil
}
