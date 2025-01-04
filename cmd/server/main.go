package main

import "gapp1/internal/app"

// @title Swagger Example API
// @version 1.0
// @description This is a sample API for demonstrating Swagger with Gin.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 创建应用实例并启动
	a := app.NewApp()
	a.Run()
}
