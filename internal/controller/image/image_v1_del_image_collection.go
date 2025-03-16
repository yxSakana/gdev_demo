package image

import (
	"context"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/yxSakana/gdev_demo/api/image/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/dao"
	imgDao "github.com/yxSakana/gdev_demo/internal/dao/image"
)

func DelImageCollection(c *gin.Context) {
	var req v1.DelImageCollectionReq
	var res v1.DelImageCollectionRes

	icid, err := strconv.ParseUint(c.Param("collection_id"), 10, 64)
	if err != nil {
		log.Printf("parse uint error: %v", err)
		consts.ParamError(c)
		return
	}

	req.Id = icid
	imgDo := do.ImageCollection{ID: icid}
	if err := imgDo.DelCache(context.Background(), req.Id); err != nil {
		consts.ComError(c, err.Error())
		return
	}
	if err := imgDao.DelCollection(dao.Ctx(c), req.Id); err != nil {
		consts.ComError(c, err.Error())
		return
	}

	consts.Success(c, res)
}
