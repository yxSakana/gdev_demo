package image

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yxSakana/gdev_demo/api/image/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	logic "github.com/yxSakana/gdev_demo/internal/logic/image"
	"log"
)

func UploadImage(c *gin.Context) {
	var req v1.UploadImageReq
	var res v1.UploadImageRes
	if err := c.ShouldBind(&req); err != nil {
		consts.ParamError(c)
		return
	}

	if err := logic.UploadImage(c, req.CollectionID, req.Image); err != nil {
		consts.ComError(c, err.Error())
		return
	}
	consts.Success(c, res)
}

func UploadImages(c *gin.Context) {
	var req v1.UploadImagesReq
	var res v1.UploadImagesRes
	if err := c.ShouldBind(&req); err != nil {
		log.Printf("%v", err)
		consts.ParamError(c)
		return
	}

	for _, image := range req.Images {
		if err := logic.UploadImage(c, req.CollectionID, image); err != nil {
			consts.ComError(c, err.Error())
			return
		}
	}

	consts.Success(c, res)
}
