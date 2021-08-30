package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/go-study/step3/gin/exp8-project-api/controller"
	"github.com/lcsin/go-study/step3/gin/exp8-project-api/tool"
	"log"
)

func main() {
	config, err := tool.ParseConfig("exp8-project-api/config/app.json")
	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(config)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	app := gin.Default()
	registryRouter(app)

	app.Run(config.AppHost + ":" + config.AppPort)
}

// 设置路由(封装路由)
func registryRouter(router *gin.Engine) {
	// 注册路由
	new(controller.HellController).Router(router)
}
