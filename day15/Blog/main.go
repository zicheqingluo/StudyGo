package main

import (
	"studygo/day15/Blog/controller"
	"studygo/day15/Blog/dao/db"

	"github.com/gin-gonic/gin"
)



func main(){
	router := gin.Default()
	dns := "tom:123456@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil{
		panic(err)
	}
	//加载静态文件
	router.Static("/static/","./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/",controller.IndexHandle)
	router.GET("/category/",controller.CategoryList)
	router.GET("/article/detail/",controller.ArticleDetail)
	router.POST("/comment/submit/",controller.CommentSubmit)
	router.GET("/article/new/",controller.NewArticle)
	router.POST("/article/submit/",controller.ArticleSubmit)
	router.GET("/leave/new/",controller.NewLeave)
	router.POST("/leave/submit/",controller.LeaveSubmit)
	router.GET("/about/me/",controller.AboutMe)
	router.POST("/upload/file/",controller.UploadFile)
	_ = router.Run(":8080")
}