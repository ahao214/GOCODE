package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin路由和路由组
func main() {
	r := gin.Default()
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{
				"method": "GET",
			})
		case "POST":
			c.JSON(http.StatusOK, gin.H{
				"method": "POST",
			})
		case "PUT":
			c.JSON(http.StatusOK, gin.H{
				"method": "PUT",
			})
		case "DELETE":
			c.JSON(http.StatusOK, gin.H{
				"method": "DELETE",
			})
		}
	})

	//没有路由的页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "没有找到页面",
		})
	})

	//路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "userGroup.GET",
			})
		})
		userGroup.POST("/loign", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "userGroup.POST",
			})
		})
	}

	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "shopGroup.GET",
			})
		})
		shopGroup.DELETE("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "shopGroup.DELETE",
			})
		})
	}

	r.Run(":9090")
}
