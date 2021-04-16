package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//gin上传文件
func main() {
	r := gin.Default()
	//单个文件
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		file, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("c:/tmp/%s", file.Filename)
		//上传文件到指定的目录
		_ = c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' upload", file.Filename),
		})
	})

	r.Run(":9090")
}
