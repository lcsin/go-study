package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 获取gin默认路由,gin.Default创建的路由默认使用了Logger和Recovery中间件
	r := gin.Default()

	// 处理一个get请求
	r.GET("/ping", func(c *gin.Context) {
		// c.JSON: 返回Json数据
		// gin.H: gin框架提供的 map[string]interface{} 类型
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "pong",
		})
	})

	// 默认端口为8080
	r.Run()
}
