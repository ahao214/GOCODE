package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type msg struct {
	Name    string
	Message string
	Age     int
}

//gin框架返回json
func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map
		data := map[string]interface{}{
			"name": "aho",
			"msg":  "hello msg",
			"age":  90,
		}
		//使用gin.H
		data = gin.H{
			"name": "aho",
			"msg":  "hello msg",
			"age":  90,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/struct_josn", func(c *gin.Context) {
		//方法二：使用结构体
		data := msg{
			Name:    "Jack",
			Message: "hello jack",
			Age:     34,
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run()
}
