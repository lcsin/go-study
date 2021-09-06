package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/go-study/step6/gin-prj-structrue/common"
)

type HelloController struct {
}

// Router 注册路由
func (h *HelloController) Router(engine *gin.Engine) {
	engine.GET("/ping", h.Ping)
	engine.GET("/error", h.Error)
}

// Ping ping
func (h *HelloController) Ping(ctx *gin.Context) {
	common.Success(ctx, "Pong!")
}

// Error error
func (h *HelloController) Error(ctx *gin.Context) {
	common.Failed(ctx, "error!")
}
