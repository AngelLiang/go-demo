package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gin_swagger_with_folder/docs"
	api "gin_swagger_with_folder/api"
)

// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header
func main() {
	r := gin.Default()
	r.GET("/ping", api.Ping)

	docs.SwaggerInfo.Title = "接口文档"
	docs.SwaggerInfo.Description = "接口文档"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
