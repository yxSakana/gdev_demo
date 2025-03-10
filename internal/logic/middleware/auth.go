package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/internal/consts"
	"github.com/yxSakana/gdev_demo/internal/logic/user"
	"net/http"
)

func Auth(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	_, err := user.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, consts.ApiResponse{
			Code: -1,
			Msg:  "Unauthorized",
			Data: nil,
		})
		c.Abort()
		return
	}

	c.Next()
}
