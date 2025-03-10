package user

import (
	"github.com/yxSakana/gdev_demo/internal/model/do"
	"log"

	"github.com/gin-gonic/gin"

	v1 "github.com/yxSakana/gdev_demo/api/user/v1"
	"github.com/yxSakana/gdev_demo/internal/consts"
	account "github.com/yxSakana/gdev_demo/internal/logic/user"
	"github.com/yxSakana/gdev_demo/internal/model/entity"
)

func Register(c *gin.Context) {
	var req v1.RegisterReq
	var res v1.RegisterRes

	err := c.ShouldBind(&req)
	if err != nil {
		consts.ParamError(c)
		log.Printf("register参数错误: %v", err)
		return
	}

	uEntity := entity.User{
		Username: req.Username,
		Nickname: req.Nickname,
		Password: req.Password,
		Email:    &req.Email,
	}
	err = account.Register(c, &do.User{User: &uEntity})
	if err != nil {
		consts.ServerError(c)
		return
	}

	consts.Success(c, res)
}
