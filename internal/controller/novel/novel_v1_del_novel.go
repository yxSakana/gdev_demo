package novel

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/dao"
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"log"
	"strconv"

	v1 "github.com/yxSakana/gdev_demo/api/novel/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	novelDao "github.com/yxSakana/gdev_demo/internal/dao/novel"
)

func DelNovel(c *gin.Context) {
	var req v1.DelNovelReq
	var res v1.DelNovelRes

	nid, err := strconv.ParseUint(c.Param("novel_id"), 10, 64)
	if err != nil {
		log.Printf("parse uint error: %v", err)
		consts.ParamError(c)
		return
	}

	req.Id = nid
	nDo := do.Novel{ID: nid}
	if err := nDo.DelCache(context.Background(), req.Id); err != nil {
		consts.ComError(c, err.Error())
		return
	}
	if err := novelDao.DelNovel(dao.Ctx(c), req.Id); err != nil {
		consts.ComError(c, err.Error())
		return
	}

	consts.Success(c, res)
}
