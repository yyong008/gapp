package v1

import (
	"gapp1/internal/app/handlers"

	docs "gapp1/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Static("/static", "./static") // 静态文件路由

	v1Group := r.Group("/api/v1")
	InitAuthRoute(v1Group)
	InitUserRoutes(v1Group)
	{
		ping := &handlers.Ping{}
		v1Group.GET("/ping", ping.PingHandler)
	}
}
