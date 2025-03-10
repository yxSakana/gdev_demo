package consts

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

func ParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, ApiResponse{
		Code: http.StatusBadRequest,
		Msg:  "参数错误",
		Data: nil,
	})
}

func ServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, ApiResponse{
		Code: http.StatusInternalServerError,
		Msg:  "Server Internal Error",
		Data: nil,
	})
}

func ComError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, ApiResponse{
		Code: -1,
		Msg:  msg,
		Data: nil,
	})
}
