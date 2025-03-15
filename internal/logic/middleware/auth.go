package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"net/http"
)

func Auth(c *gin.Context) {
	abort := func() {
		c.JSON(http.StatusUnauthorized, consts.ApiResponse{
			Code: -1,
			Msg:  "Unauthorized",
			Data: nil,
		})
		c.Abort()
	}

	uid, err := user.GetUserID(c)
	if err != nil {
		abort()
		return
	}
	exist, err := user.IsExist(c, uid)
	if err != nil || !exist {
		abort()
		return
	}

	c.Next()
}
