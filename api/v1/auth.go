package v1

import (
	"gapp1/internal/app/handlers"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(r *gin.RouterGroup) {
	auth := &handlers.LoginHandler{}
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/register", auth.Register)
		authGroup.POST("/logout", auth.Logout)
		authGroup.POST("/refresh_token", auth.RefreshToken)
	}

}
