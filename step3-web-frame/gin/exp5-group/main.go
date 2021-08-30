package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// r.Group(): 路由组
	userGroup := r.Group("/user")
	{
		// 路由组内的请求，path将会组合为：/user/registry
		userGroup.POST("/registry", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "注册成功!",
			})
		})
		userGroup.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "登录成功!",
			})
		})
		userGroup.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "用户信息",
			})
		})
		userGroup.DELETE("/10086", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"msg":  "删除成功!",
			})
		})
	}

	r.Run()
}
