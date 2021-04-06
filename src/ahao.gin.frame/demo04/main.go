package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//中间件
func middle() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前，我是方法一")
		c.Next()
		fmt.Println("我在方法后，我是方法一")
	}
}

func middletwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法前，我是方法二")
		c.Next()
		fmt.Println("我在方法后，我是方法二")
	}
}

func main() {
	r := gin.Default()
	//第一种调用
	v1 := r.Group("v1").Use(middle(), middletwo())
	//第二种调用
	//v1=r.Group("v1").Use(middle()).Use(middletwo())
	v1.GET("test", func(c *gin.Context) {
		fmt.Println("我在分组方法内部")
		c.JSON(200, gin.H{
			"sucdess": true,
		})
	})

	r.Run(":8080")
}
