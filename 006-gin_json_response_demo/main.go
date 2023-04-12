package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gin_json_response_demo/docs"
)

type LoginUser struct {
	// 用户名
	Username string `bind:"required" json:"username"`
	
	// 密码
	Password string `bind:"required" json:"password"`
}


// @Summary 用户登录
// @Schemes
// @Description 
// @Tags 登录认证
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /login [post]
// @Param string body LoginUser false "用户名"
func Login(c *gin.Context) {
	var login LoginUser

	// 使用ShouldBind函数从请求中获取登录表单数据，并将其自动绑定到login变量中
	if c.ShouldBind(&login) == nil && login.Password != "" && login.Username != "" {
		// 如果登录表单数据获取成功，并且用户名和密码都不为空，返回成功信息
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"msg":   "登录成功",
			"data": gin.H{
				"token": "token",
			},
		})
	} else {
		// 否则，返回失败信息
		c.JSON(http.StatusOK, gin.H{
			"code":  1,
			"msg":   "登录失败",
			"data": gin.H{},
		})
	}

	// 打印登录表单数据
	fmt.Println(login)
}



func main() {
    router := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	router.POST("/login", Login)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
