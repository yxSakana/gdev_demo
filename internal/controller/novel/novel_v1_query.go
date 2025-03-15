package novel

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/novel"
	"github.com/yxSakana/gdev_demo/internal/model"
	"log"
)

func Query(c *gin.Context) {
	var req v1.QueryReq
	var res v1.QueryRes

	if err := c.ShouldBindQuery(&req); err != nil {
		log.Printf("Query novel err: %v", err)
		consts.ParamError(c)
		return
	}

	novelItems, err := novel.QueryNovel(c, model.NovelQueryInput{
		Id:        req.Id,
		Tag:       req.Tag,
		Author:    req.Author,
		Title:     req.Title,
		WordCount: req.WordCount,
		View:      req.View,
		Like:      req.Like,
	})
	if err != nil {
		log.Printf("Query novel err: %v", err)
		consts.ServerError(c)
		return
	}

	for _, n := range novelItems {
		res.List = append(res.List, v1.NovelDetail{
			NovelID:       n.ID,
			UploaderID:    n.UserID,
			Author:        n.Uploader,
			Title:         n.Title,
			Tags:          n.Tags,
			Description:   n.Description,
			CoverUrl:      n.CoverUrl,
			Status:        n.Status,
			WordCount:     n.WordCount,
			View:          n.View,
			Like:          n.Like,
			ChapterNumber: n.ChapterNumber,
		})
	}
	res.Total = uint(len(res.List))
	consts.Success(c, res)

	//queries := c.Request.URL.Query()
	//v := reflect.ValueOf(req)
	//t := reflect.TypeOf(req)
	//for i := 0; i < v.NumField(); i++ {
	//	formTag := t.Field(i).Tag.Get("form")
	//	if q, exists := queries[formTag]; exists {
	//
	//	}
	//}
}
