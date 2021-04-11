package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin获取querystring参数
func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取哦浏览器那边发送请求携带的query string参数
		//name := c.Query("query")
		//获取不到数据，就用知道的默认值
		//name := c.DefaultQuery("query", "somebody")
		name, ok := c.GetQuery("query")
		if !ok {
			name = "somebody"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	r.Run()
}
