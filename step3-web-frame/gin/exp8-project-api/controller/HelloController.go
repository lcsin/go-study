package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HellController struct {
}

func (h *HellController) Router(engine *gin.Engine) {
	engine.GET("/ping", h.Pong)
}

func (h *HellController) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Pong!",
	})
}
