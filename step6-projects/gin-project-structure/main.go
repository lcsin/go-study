package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/go-study/step6/gin-prj-structrue/controller"
	"github.com/lcsin/go-study/step6/gin-prj-structrue/dao"
)

func main() {
	// 初始化DB
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	// 初始化gin路由
	r := gin.Default()
	registryRouter(r)
	r.Run()
}

// 路由封装
func registryRouter(router *gin.Engine) {
	// 注册路由
	new(controller.HelloController).Router(router)
	new(controller.UserController).Router(router)
}
