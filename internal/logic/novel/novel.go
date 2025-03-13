package novel

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/utility"
	"gorm.io/gorm"
	"log"

	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/dao"
	novelDao "github.com/yxSakana/gdev_demo/internal/dao/novel"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"github.com/yxSakana/gdev_demo/internal/model"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
	novelService "github.com/yxSakana/gdev_demo/internal/service/novel"
)

func CreateNovel(c *gin.Context, in *v1.CreateNovelReq) (uint64, error) {
	userInfo, err := user.GetUserinfo(c)
	if err != nil {
		return 0, err
	}

	filePath, err := utility.GenerateFilePath(in.Cover)
	if err != nil {
		return 0, err
	}

	if err := utility.CheckCoverFile(in.Cover, filePath); err != nil {
		return 0, err
	}

	if err := c.SaveUploadedFile(in.Cover, filePath); err != nil {
		return 0, err
	}

	novelEntity := entity.Novel{
		UserID:      userInfo.ID,
		Uploader:    userInfo.Username,
		Title:       in.Title,
		Description: in.Description,
		CoverUrl:    filePath,
		Status:      0,
	}
	novelDo := do.Novel{Novel: &novelEntity}

	db := dao.Ctx(c)
	if err := db.Transaction(func(tx *gorm.DB) error {
		err = novelDao.CreateNovel(db, &novelDo)
		if err != nil {
			return err
		}
		return novelService.LinkNovelAndTags(db, novelDo.ID, in.Tags)
	}); err != nil {
		return 0, err
	}

	return novelEntity.ID, nil
}

func UploadChapter(c *gin.Context, req *v1.UploadChapterReq) (uint64, error) {
	db := dao.Ctx(c)

	cEntity := entity.NovelChapter{
		NovelId:   req.NovelId,
		Title:     req.Title,
		Number:    req.Number,
		Content:   req.Content,
		WordCount: uint(len(req.Content)),
	}
	ncDo := do.NovelChapter{NovelChapter: &cEntity}

	err := novelDao.CreateChapter(db, &ncDo)
	if err != nil {
		return 0, err
	}

	return cEntity.ID, nil
}

func QueryNovel(c *gin.Context, query model.NovelQueryInput) ([]model.NovelOutput, error) {
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

	var outs []model.NovelOutput
	for _, i := range novelIds {
		n, err := DetailNovelByID(c, i)
		if err != nil {
			return nil, err
		}
		outs = append(outs, *n)
	}
	return outs, nil
}

func DetailNovelByID(c *gin.Context, nid uint64) (*model.NovelOutput, error) {
	cacheKey := fmt.Sprintf("%s:%d", consts.CacheNovel, nid)
	d := novelService.GetFromCache(cacheKey)
	if d != nil {
		log.Printf("get from cache: %#v", *d.Novel)
	}

	nEntity, err := novelDao.GetNovelByID(dao.Ctx(c), nid)
	if err != nil {
		return nil, err
	}

	tags, err := GetNovelTags(c, nEntity.ID)
	if err != nil {
		return nil, err
	}

	if err := novelService.SetCache(cacheKey, &do.Novel{Novel: nEntity}); err != nil {
		log.Printf("set cache err:%v", err)
	}

	ret := &model.NovelOutput{
		NovelID:     nid,
		UserID:      nEntity.UserID,
		Author:      nEntity.Uploader,
		Title:       nEntity.Title,
		Tags:        tags,
		Description: nEntity.Description,
		CoverUrl:    nEntity.CoverUrl,
		Status:      nEntity.Status,
		WordCount:   nEntity.WordCount,
		View:        nEntity.View,
		Like:        nEntity.Like,
		//ChapterCount: uint(len(chapterIds)), TODO: novel chapter count
	}

	return ret, nil
}

func GetNovelTags(c *gin.Context, nid uint64) ([]string, error) {
	var tags []string
	err := dao.Ctx(c).Model(&entity.NovelTag{}).
		Joins("LEFT JOIN novel_tag_rel ON novel_tags.id = novel_tag_rel.novel_tag_id").
		Joins("LEFT JOIN novels ON novel_tag_rel.novel_id = novels.id").
		Where("novels.id = ?", nid). // TODO: 是否正确
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
