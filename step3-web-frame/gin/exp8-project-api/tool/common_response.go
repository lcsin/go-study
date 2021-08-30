package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS int = iota
	FAILED
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": SUCCESS,
		"msg":  "请求成功!",
		"data": data,
	})
}

func Failed(ctx *gin.Context, msg interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": FAILED,
		"msg":  msg,
	})
}
