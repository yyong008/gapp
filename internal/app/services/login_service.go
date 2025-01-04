package services

import (
	"gapp1/internal/app/model"
	"gapp1/pkg/db"

	"github.com/gin-gonic/gin"
)

type LoginService struct{}

func (u *LoginService) Login(c *gin.Context) (string, error) {
	// 保存用户到数据库
	username := c.PostForm("username")

	if err := db.Database.Find(&model.User{Username: username}).Error; err != nil {
		return "", err
	}
	token := ""
	return token, nil
}
