package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()

	// 绑定 URI Query 参数
	r.GET("/hello", func(ctx *gin.Context) {
		var u User
		// ctx.ShouldBindQuery: 实体绑定Query参数,可以通过实体类的tag指定绑定名称
		err := ctx.ShouldBindQuery(&u)
		if err != nil {
			log.Fatalln("should exp3-bind failed,err:", err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  u.Name,
		})
	})

	// 绑定 Form 表单参数
	r.POST("/registry", func(ctx *gin.Context) {
		var u User
		// ctx.ShouldBind: 实体绑定from表单参数
		err := ctx.ShouldBind(&u)
		if err != nil {
			log.Fatalln("param exp3-bind failed,err:", err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  u.Name + "=>" + u.Password,
		})
	})

	// 绑定 JSON 参数
	r.GET("/login", func(ctx *gin.Context) {
		var u User
		// ctx.ShouldBindJSON: 绑定json参数
		err := ctx.ShouldBindJSON(&u)
		if err != nil {
			log.Fatalln("json exp3-bind failed,err:", err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "login success",
		})
	})

	r.Run()
}

type User struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Age      int    `form:"age"`
}
