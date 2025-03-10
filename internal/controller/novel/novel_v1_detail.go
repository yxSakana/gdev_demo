package novel

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/novel"
)

func DetailNovel(c *gin.Context) {
	var req v1.DetailNovelReq
	var res v1.DetailNovelRes

	nid, err := strconv.Atoi(c.Param("novel_id"))
	if err != nil {
		log.Printf("detail novel 参数错误: %v", err)
		consts.ParamError(c)
		return
	}
	req.NovelID = uint64(nid)

	novelOutput, err := novel.DetailNovelByID(c, req.NovelID)
	if err != nil {
		log.Printf("detail novel err: %v", err)
		consts.ServerError(c)
		return
	}

	res = v1.DetailNovelRes{
		NovelDetail: v1.NovelDetail{
			NovelID:      novelOutput.NovelID,
			UploaderID:   novelOutput.UserID,
			Author:       novelOutput.Author,
			Title:        novelOutput.Title,
			Tags:         novelOutput.Tags,
			Description:  novelOutput.Description,
			CoverUrl:     novelOutput.CoverUrl,
			Status:       novelOutput.Status,
			WordCount:    novelOutput.WordCount,
			View:         novelOutput.View,
			Like:         novelOutput.Like,
			ChapterCount: novelOutput.ChapterCount,
			ChapterIds:   novelOutput.ChapterIds,
		},
	}
	consts.Success(c, res)
}
