package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//模板解析
	//r.LoadHTMLFiles("template/index.tmpl")
	r.LoadHTMLGlob("template/**/*")
	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //模板渲染
			"title": "aho",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "aho.users.index",
		})
	})
	r.Run(":9000") //启动server
}
