package main

import "github.com/gin-gonic/gin"
import "net/http"
import swaggerfiles "github.com/swaggo/files"
import ginSwagger "github.com/swaggo/gin-swagger"
import docs "gin_querystring_demo/docs"



func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.GET("/welcome", welcome)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}


// @Summary 查询字符串参数
// @Schemes
// @Description 示例 URL：/welcome?firstname=Jane&lastname=Doe
// @Tags default
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /welcome [get]
// @Param firstname query string false "名称"
// @Param lastname query string false "姓氏"
func welcome(c *gin.Context){
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // c.Request.URL.Query().Get("lastname") 的一种快捷方式
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
