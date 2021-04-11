package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin获取form参数
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html")
	})

	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的数据
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")

		username = c.DefaultPostForm("username", "somebody")
		pwd = c.DefaultPostForm("pwd", "***")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": username,
			"pwd":  pwd,
		})
	})
	r.Run(":9090")
}
