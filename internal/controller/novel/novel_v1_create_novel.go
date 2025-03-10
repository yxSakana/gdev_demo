package novel

import (
	"log"

	"github.com/gin-gonic/gin"

	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/novel"
)

func CreateNovel(c *gin.Context) {
	var req v1.CreateNovelReq
	var res v1.CreateNovelRes

	if err := c.ShouldBind(&req); err != nil {
		consts.ParamError(c)
		log.Printf("upload novel 参数错误: %v", err)
		return
	}

	if nid, err := novel.CreateNovel(c, &req); err == nil {
		res.NovelID = nid
		consts.Success(c, res)
		return
	} else {
		log.Printf("upload novel err: %v", err)
	}
}

func UploadChapter(c *gin.Context) {
	var req v1.UploadChapterReq
	var res v1.UploadChapterRes

	if err := c.ShouldBind(&req); err != nil {
		consts.ParamError(c)
		log.Printf("upload chapter 参数错误: %v", err)
		return
	}

	if _, err := novel.UploadChapter(c, &req); err == nil {

		consts.Success(c, res)
		return
	} else {
		log.Printf("upload novel err: %v", err)
	}
}
