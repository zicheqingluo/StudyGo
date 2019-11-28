package main

import (
	"studygo/day15/Blog/controller"
	"studygo/day15/Blog/dao/db"

	"github.com/gin-gonic/gin"
)



func main(){
	router := gin.Default()
	dns := "tom:123456@tcp(10.127.63.21:3306)/blogger?parseTime=true"
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
	_ = router.Run(":8000")
}