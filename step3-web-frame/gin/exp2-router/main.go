package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// r.Handle: 通用路由请求处理器
	r.Handle("GET", "/ping", func(ctx *gin.Context) {
		// ctx.FullPath: 获取完整的请求路径
		path := ctx.FullPath()
		fmt.Println("fullPath:", path)

		// ctx.DefaultQuery: 获取query参数
		name := ctx.DefaultQuery("name", "缺省值")
		fmt.Println("query param:", name)

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "pong",
		})
	})

	r.Handle("POST", "/registry", func(ctx *gin.Context) {
		// ctx.PostForm: 获取表单参数
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		fmt.Println("form param: ", username, " ", password)

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "registry success",
		})
	})

	r.DELETE("/user/:id", func(ctx *gin.Context) {
		// ctx.Param: 获取path参数
		id := ctx.Param("id")
		fmt.Println("user id:", id)

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "delete success",
		})
	})

	r.Run()
}
