package main

import (
	"fmt"
   "github.com/gin-gonic/gin"
   docs "gin_apikey_demo/docs"  // NOTE: docs导入的包要和go.mod的module一样
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
   "net/http"
)

// authMiddleware中间件会检查每个请求的身份验证信息是否正确
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apikey := c.Request.Header.Get("Authorization")
		fmt.Printf("apikey:%s\r\n", apikey)
		if (apikey != "123456") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 继续处理请求
		c.Next()
	}
}


// @Summary 检查 ApiKey
// @Schemes
// @Description
// @Tags 登录认证
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /checkApiKey [get]
// @Security ApiKeyAuth
func checkApiKey(g *gin.Context)  {
	apikey := g.Request.Header.Get("Authorization")
	data := map[string]interface{}{
		"code": 1,
		"msg":  "ok",
		"data": apikey,
	}
   g.JSON(http.StatusOK, data)
}

// @securityDefinitions.apikey ApiKeyAuth
// @name Authorization
// @in header
func main()  {
   r := gin.Default()
   docs.SwaggerInfo.BasePath = ""
   r.GET("/checkApiKey", AuthMiddleware(), checkApiKey)
   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
   r.Run(":8080")
}
