package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// r.Use(): 全局使用中间件
	r.Use(RequestInfos())

	// 局部使用中间件
	r.GET("/ping", RequestInfos(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "pong!",
		})
	})

	r.Run()
}

// RequestInfos 自定义中间件
func RequestInfos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		method := ctx.Request.Method
		status := ctx.Writer.Status()
		fmt.Println("请求路径=>", path)
		fmt.Println("请求方法=>", method)

		// ctx.Next(): 中断当前中间件代码的执行，转而执行业务处理程序，当业务处理程序执行完毕后，返回继续执行中间件的代码
		ctx.Next()

		fmt.Println("状态码=>", status)
	}
}
