package image

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yxSakana/gdev_demo/api/image/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	server "github.com/yxSakana/gdev_demo/internal/logic/image"
	"github.com/yxSakana/gdev_demo/internal/model"
	"strings"
)

func Create(c *gin.Context) {
	var req v1.CreateReq
	var res v1.CreateRes

	if err := c.ShouldBind(&req); err != nil {
		consts.ParamError(c)
		return
	}

	for i, t := range req.Tags {
		req.Tags[i] = strings.TrimSpace(t)
	}

	if err := server.Create(c, model.CreateImageCollectionInput{
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Tags:        req.Tags,
	}); err != nil {
		consts.ComError(c, err.Error())
		return
	}

	consts.Success(c, res)
	//file, err := c.FormFile("cover")
	//if err != nil {
	//	return
	//}
}
