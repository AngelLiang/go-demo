package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gin_json_parser_demo/docs"
)

// 定义了一个名为LoginUser的结构体
type LoginUser struct {
	// Username字段表示登录用户名
	// bind:"required"表示这个字段是必需的，需要在绑定时提供值
	// json:"username"表示在JSON序列化/反序列化时使用的键名
	Username string `binding:"required" json:"username"`
	
	// Password字段表示登录密码
	// bind:"required"表示这个字段是必需的，需要在绑定时提供值
	// json:"password"表示在JSON序列化/反序列化时使用的键名
	Password string `binding:"required" json:"password"`
}



// @Summary 用户登录
// @Schemes
// @Description 
// @Tags 登录认证
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /loginWithJSON [post]
// @Param string body LoginUser false "用户名"
func Login(c *gin.Context) {
	// 定义一个LoginUser类型的变量login
	var login LoginUser

	// 使用ShouldBind函数从请求中获取登录表单数据，并将其自动绑定到login变量中
	if c.ShouldBind(&login) == nil && login.Password != "" && login.Username != "" {
		// 如果登录表单数据获取成功，并且用户名和密码都不为空，返回成功信息
		c.String(http.StatusOK, "login successfully !")
	} else {
		// 否则，返回失败信息
		c.String(http.StatusBadRequest, "login failed !")
	}

	// 打印登录表单数据
	fmt.Println(login)
}



func main() {
    router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.POST("/loginWithJSON", Login)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
