package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin" // 导入 gin 框架

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "gin_logger_demo/docs"
)

func main() {
	// 创建 gin 实例
	r := gin.Default()

	// 使用 Logger 中间件，记录每个请求的日志信息
	r.Use(gin.Logger())

	// 使用 Recovery 中间件，处理 panic 异常
	r.Use(gin.Recovery())

	// 自定义中间件，记录请求处理时间
	r.Use(customMiddleware)

	r.GET("/", index)

	// 添加处理 GET /panic 请求的路由，模拟 panic 异常
	r.GET("/panic", panicHandle)

	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 启动 Web 服务，监听 8080 端口
	r.Run(":8080")
}

// @Summary index
// @Router / [get]
func index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
}

// @Summary 处理 panic
// @Schemes
// @Description 
// @Tags default
// @Accept json
// @Produce json
// @Success 200 {string} ok
// @Router /panic [get]
func panicHandle(c *gin.Context) {
	panic("发生了一个错误")
}

// 自定义中间件
func customMiddleware(c *gin.Context) {
	fmt.Println("执行自定义中间件前...")
	start := time.Now()
	c.Next() // 执行下一个中间件或者路由处理函数
	fmt.Println("执行自定义中间件后...")
	end := time.Now()
	latency := end.Sub(start) // 计算请求处理时间
	fmt.Printf("请求处理时间为：%v\n", latency)
}
