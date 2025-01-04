package v1

import (
	"gapp1/internal/app/handlers"
	"gapp1/internal/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup) {
	users := &handlers.UserHandler{}
	userGroup := r.Group("/users")
	userGroup.Use(middleware.Auth())

	{
		userGroup.GET("/", users.GetUserListHandler)
		userGroup.POST("/", users.CreateUserHandler)
		userGroup.GET("/:id", users.GetUserByIDHandler)
		userGroup.PUT("/:id", users.UpdateUserByIDHandler)
		userGroup.DELETE("/", users.DeleteUserByIDsHandler)
	}

}
