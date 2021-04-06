package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//保存文件
		c.SaveUploadedFile(file, "./"+file.Filename)

		/* 手写保存文件
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)
		*/

		c.JSON(200, gin.H{
			"msg": file,
		})
	})

	r.Run(":8000")
}
