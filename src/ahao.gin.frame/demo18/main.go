package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//定义一个中间件m1:统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//计时
	start := time.Now()
	c.Next()  //调用后面的处理函数
	c.Abort() //组织调用后续的处理函数

	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 int...")
	c.Next()
	c.Abort() //组织调用后续的处理函数
	fmt.Println("m2 out...")

}

//中间件
//中间件适合处理一些公共的业务逻辑
func main() {
	r := gin.Default()
	//全局注册中间件函数m1和m2
	r.Use(m1, m2)

	//使用中间件
	r.GET("/index", m1, indexHandle)
	r.Run(":9090")
}
