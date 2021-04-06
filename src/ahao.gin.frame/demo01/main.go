package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		//user为空，获取默认值aho
		user := c.DefaultQuery("user", "aho")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "aho")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})
	r.PUT("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "aho")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})
	r.Run(":8000")
}
