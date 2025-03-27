package user

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/yxSakana/gdev_demo/api/user/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"log"
)

func Login(c *gin.Context) {
	var req v1.LoginReq
	var res v1.LoginRes

	if err := c.ShouldBind(&req); err != nil {
		consts.ParamError(c)
		log.Printf("login参数错误: %v", err)
		return
	}

	token, err := user.Login(c, req.Username, req.Password)
	if err != nil {
		consts.ServerError(c)
		log.Println(err)
		return
	}
	res.Token = token

	consts.Success(c, res)
}
