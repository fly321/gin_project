package service

import (
	"flyblog/db"
	"flyblog/model"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (u UserService) FindById(ctx *gin.Context) {
	user := model.User{}
	form := ctx.Query("id")
	db.DB.Debug().Where("id = ?", form).First(&user)
	ctx.JSON(200, user)
}
