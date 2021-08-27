package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/byte", func(ctx *gin.Context) {
		// ctx.Writer.Write: 返回byte切片数据
		ctx.Writer.Write([]byte("pong"))
	})

	r.GET("/string", func(ctx *gin.Context) {
		// ctx.Writer.WriteString: 返回string数据
		ctx.Writer.WriteString("hello,world!")
	})

	r.GET("/json", func(ctx *gin.Context) {
		// ctx.JSON: 返回json数据
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "response json data",
		})
	})

	r.GET("/html", func(ctx *gin.Context) {
		path := ctx.FullPath()

		// r.LoadHTMLGlob: 加载html文件目录
		r.LoadHTMLGlob("exp4-response/*")

		// ctx.HTML: 返回html页面,其中gin.H中的数据为传递给html页面的数据
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
			"path":  path,
		})
	})

	// r.Static: 设置静态文件目录 relativePath:映射路径 root:静态文件目录路径
	r.Static("/img", "exp4-response/img")
	r.GET("static", func(ctx *gin.Context) {
		// r.LoadHTMLFiles: 加载html文件
		r.LoadHTMLFiles("exp4-response/index.html")

		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run()
}
