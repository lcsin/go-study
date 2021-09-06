package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": RequestSuccess,
		"msg":  "请求成功!",
		"data": data,
	})
}

func Failed(ctx *gin.Context, msg interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": RequestFailed,
		"msg":  msg,
	})
}
