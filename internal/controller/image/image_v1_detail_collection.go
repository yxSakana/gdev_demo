package image

import (
	"errors"
	"github.com/gin-gonic/gin"
	v1 "github.com/yxSakana/gdev_demo/api/image/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	logic "github.com/yxSakana/gdev_demo/internal/logic/image"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func DetailImgCollection(c *gin.Context) {
	var req v1.DetailImgCollectionReq
	var res v1.DetailImgCollectionRes

	nid, err := strconv.Atoi(c.Param("collection_id"))
	if err != nil {
		log.Printf("DetailImgCollection 参数错误: %v", err)
		consts.ParamError(c)
		return
	}
	req.CollectionID = uint64(nid)

	imgCollection, err := logic.DetailImgCollectionByID(c, req.CollectionID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		consts.ComError(c, "作品不存在")
		return
	}
	if err != nil {
		consts.ComError(c, err.Error())
		return
	}

	res = v1.DetailImgCollectionRes{
		ImageCollection: v1.ImageCollection{
			ID:          imgCollection.ID,
			UserID:      imgCollection.UserID,
			Uploader:    imgCollection.Uploader,
			Title:       imgCollection.Title,
			Description: imgCollection.Description,
			CoverUrl:    imgCollection.CoverUrl,
			Number:      imgCollection.Number,
			Tags:        imgCollection.Tags,
			CreatedAt:   imgCollection.CreatedAt,
			UpdatedAt:   imgCollection.UpdatedAt,
		},
	}

	consts.Success(c, res)
}
