package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type PostParams struct {
	Name string `json:"name" uri:"name" form:"name" binding:"required"`
	Age  int    `json:"age" uri:"age" form:"age" binding:"required"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/testbind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(200, gin.H{
				"msg":  "报错了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "成功了",
				"data": p,
			})
		}
	})

	r.Run(":8000")
}
