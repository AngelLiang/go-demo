package main

import "github.com/gin-gonic/gin"
import "net/http"
import swaggerfiles "github.com/swaggo/files"
import ginSwagger "github.com/swaggo/gin-swagger"
import docs "gin_path_parameter_demo/docs"


// @Summary Hello
// @Schemes
// @Description 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
// @Tags default
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /user/{name} [get]
// @Param name path string true "姓名"
func Hello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

// @Summary Action
// @Schemes
// @Description 此 handler 将匹配 /user/john/ 和 /user/john/send。如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
// @Tags default
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /user/{name}/{action} [get]
// @Param name path string true "姓名"
// @Param action path string false "动作"
func Action(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func main() {
	router := gin.Default()

	// 此 handler 将匹配 /user/john 但不会匹配 /user/ 或者 /user
	router.GET("/user/:name", Hello)

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	router.GET("/user/:name/*action", Action)

	docs.SwaggerInfo.Title = "接口文档"
	docs.SwaggerInfo.Description = "接口文档"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
