package main

import (
	"net/http"
	// "reflect"
	"time"

	docs "gin_custom_validation/docs"  // NOTE: docs导入的包要和go.mod的module一样
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)


var isAfterToday validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

type Booking struct {
	// 判断输入的日期是否是今天之后的日期 格式 2006-01-02
	AfterToday  time.Time `form:"afterToday" binding:"required,isAfterToday" time_format:"2006-01-02"`
}


// @Summary getBookable
// @Schemes
// @Description bookable
// @Tags example
// @Accept json
// @Produce json
// @Param string query Booking false "入参"
// @Success 200 {string} ok
// @Router /bookable [get]
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

// 初始化校验器
func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isAfterToday", isAfterToday)
	}
}

func main() {
	route := gin.Default()
	InitValidator()
	docs.SwaggerInfo.BasePath = ""

	route.GET("/bookable", getBookable)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	route.Run(":8080")
}
