package main

import (
   "github.com/gin-gonic/gin"
   docs "gin_swagger_demo/docs"  // NOTE: docs导入的包要和go.mod的module一样
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
   "net/http"
)
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context)  {
   g.JSON(http.StatusOK, "helloworld")
}

func main()  {
   r := gin.Default()

   docs.SwaggerInfo.Title = "接口文档"
   docs.SwaggerInfo.Description = "接口文档"
   docs.SwaggerInfo.Version = "1.0"
   docs.SwaggerInfo.BasePath = "/api/v1"
   v1 := r.Group("/api/v1")
   {
      eg := v1.Group("/example")
      {
         eg.GET("/helloworld",Helloworld)
      }
   }
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   r.Run(":8080")

}
