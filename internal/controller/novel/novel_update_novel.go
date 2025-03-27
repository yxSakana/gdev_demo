package novel

import (
	"github.com/yxSakana/gdev_demo/internal/model"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/novel"
)

// UpdateNovel is a test doc
func UpdateNovel(c *gin.Context) {
	var req v1.UpdateNovelReq
	var res v1.UpdateNovelRes

	nid, err := strconv.ParseUint(c.Param("novel_id"), 10, 64)
	if err != nil {
		log.Printf("parse uint error: %v", err)
		consts.ParamError(c)
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("ShouldBind error: %v", err)
		consts.ParamError(c)
		return
	}

	if err := novel.UpdateNovel(c, nid, model.UpdateNovelInput{
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Status:      req.Status,
		Tags:        req.Tags,
	}); err != nil {
		log.Printf("UpdateNovel err: %v", err)
		consts.ServerError(c)
		return
	}

	consts.Success(c, res)
}
