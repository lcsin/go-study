package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lcsin/go-study/step6/gin-prj-structrue/common"
	"github.com/lcsin/go-study/step6/gin-prj-structrue/models"
)

type UserController struct {
}

func (u *UserController) Router(engine *gin.Engine) {
	engine.GET("/user", u.List)
}

func (u UserController) List(ctx *gin.Context) {
	users, err := models.ListUser()
	if err != nil {
		common.Failed(ctx, err.Error())
	}
	common.Success(ctx, users)
}
