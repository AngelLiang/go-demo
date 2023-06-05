package main

import (
	"net/http"
	"path/filepath"

	docs "gin_upload_file/docs"  // NOTE: docs导入的包要和go.mod的module一样
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @Summary 上传文件接口
// @Schemes
// @Description 上传文件接口
// @Tags example
// @Accept json
// @Produce json
// @Param file formData file true "file"
// @Param string formData string false "name"
// @Success 200 {string} ok
// @Router /upload [post]
func UploadFileApi(c *gin.Context) {
	// name := c.PostForm("name")
	// email := c.PostForm("email")

	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	// 保存文件
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	// c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "上传成功"})
}

func main() {
	router := gin.Default()
	docs.SwaggerInfo.Title = "接口文档"
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// router.Static("/", "./public")
	router.POST("/upload", UploadFileApi)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
