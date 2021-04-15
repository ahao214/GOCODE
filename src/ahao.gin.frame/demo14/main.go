package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"username"`
}

//gin参数绑定
func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		u := UserInfo{
			Username: username,
			Password: password,
		}
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u UserInfo
		err := c.ShouldBindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":9090")
}
