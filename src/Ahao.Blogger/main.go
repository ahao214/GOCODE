package main

import (
	"Ahao.Blogger/controller"
	"Ahao.Blogger/dao/db"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:admin@tcp(localhost:3306)/Ahao.Blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.GET("/article/new/", controller.NewArticle)
	router.POST("/article/submit/", controller.ArticleSubmit)
	router.GET("/article/detail/", controller.ArticleDetail)
	router.POST("/upload/file/", controller.UploadFile)
	router.GET("/leave/new/", controller.LeaveNew)
	router.GET("/about/me/", controller.AboutMe)
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.POST("/leave/submit/", controller.LeaveSubmit)
	_ = router.Run(":8000")
}
