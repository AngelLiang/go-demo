package main

import (
   "github.com/gin-gonic/gin"
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
   "net/http"

   docs "gin_httptest/docs"  // NOTE: docs导入的包要和go.mod的module一样
)

// PingExample godoc
// @Summary helloworld example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/v1/helloworld [get]
func Helloworld(g *gin.Context)  {
   g.JSON(http.StatusOK, "helloworld")
}

func InitRouter(r *gin.Engine) {
	r.GET("/api/v1/helloworld", Helloworld)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func main()  {
   r := gin.Default()
   docs.SwaggerInfo.Title = "接口文档"
   docs.SwaggerInfo.Description = "接口文档"
   docs.SwaggerInfo.Version = "1.0"
//    docs.SwaggerInfo.BasePath = "/api/v1"
   InitRouter(r)
   r.Run("0.0.0.0:8080")
}
