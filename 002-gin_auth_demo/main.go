package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	docs "gin_auth_demo/docs"  // NOTE: docs导入的包要和go.mod的module一样
)


// @title MyAPI API
// @version 1.0
// @description This is a sample API
// @BasePath /
// @securityDefinitions.basic BasicAuth
func main() {
	// 创建gin引擎实例
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	// 添加gin-swagger中间件处理文档接口
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 定义路由和中间件，并添加swagger注释
	router.GET("/api/secure", secureHandler)

	router.GET("/api/unsecure", unsecureHandler)

	router.GET("/api/protected", authMiddleware(), protectedHandler)

	// 启动HTTP服务器
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// secureHandler和unsecureHandler是不需要认证的API
// @Summary This is a secure API.
// @Tags Secure
// @Produce text/plain
// @Success 200 {string} string "OK"
// @Router /api/secure [get]
func secureHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a secure API.")
}

// @Summary This is an unsecure API.
// @Tags Unsecure
// @Produce text/plain
// @Success 200 {string} string "OK"
// @Router /api/unsecure [get]
func unsecureHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is an unsecure API.")
}

// protectedHandler需要进行身份验证和授权才能访问
// @Summary This is a protected API.
// @Tags Protected
// @Security BasicAuth
// @Produce text/plain
// @Success 200 {string} string "OK"
// @Router /api/protected [get]
func protectedHandler(c *gin.Context) {
	c.String(http.StatusOK, "This is a protected API.")
}

// authMiddleware中间件会检查每个请求的身份验证信息是否正确
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查认证信息是否正确
		username, password, ok := c.Request.BasicAuth()
		if !ok || !isValidCredentials(username, password) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 继续处理请求
		c.Next()
	}
}

// isValidCredentials函数验证用户名和密码是否有效
func isValidCredentials(username, password string) bool {
	// 在这里实现验证逻辑
	return username == "admin" && password == "password"
}
