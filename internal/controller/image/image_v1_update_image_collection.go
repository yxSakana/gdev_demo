package image

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

	v1 "github.com/yxSakana/gdev_demo/api/image/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/image"
	"github.com/yxSakana/gdev_demo/internal/model"
)

func UpdateImageCollection(c *gin.Context) {
	var req v1.UpdateImageCollectionReq
	var res v1.UpdateImageCollectionRes

	icid, err := strconv.ParseUint(c.Param("collection_id"), 10, 64)
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

	if err := image.UpdateImageCollection(c, icid, model.UpdateImageCollectionInput{
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Tags:        req.Tags,
	}); err != nil {
		log.Printf("UpdateNovel err: %v", err)
		consts.ServerError(c)
		return
	}

	consts.Success(c, res)
}
