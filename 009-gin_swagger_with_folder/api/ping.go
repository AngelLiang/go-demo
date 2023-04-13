package api

import (
	"net/http"
	"github.com/gin-gonic/gin" // 导入 gin 框架
)

type JSONResult struct {
	// 业务码
    Code    int          `json:"code" `
	// 响应信息
    Message string       `json:"message"`
	// 响应数据
    Data    interface{}  `json:"data"`
}


// @Summary ping
// @Schemes http
// @Description Ping 测试接口
// @Tags 通用
// @Router /ping [get]
// @Accept json
// @Produce json
// @Success 200 {object} JSONResult
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message":  "请求成功",
		"data": gin.H{},
	})
}
