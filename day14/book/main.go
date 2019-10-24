package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"fmt"
	"studygo/day14/book/db"
	"studygo/day14/book/model"
)

func initDB(){
	err := db.InitDB()
	if err != nil {
		fmt.Println("链接错误",err)
	}
}

func initModule(){
		// 1.创建路由
		r := gin.Default()
		// 2.绑定路由规则，执行的函数
		//3.渲染模板
		r.LoadHTMLGlob("templates/*")
		// gin.Context，封装了request和response
		r.GET("/book/list",model.List)
		r.GET("/book/new",model.AddBookHtml )
		r.POST("/book/new",model.AddBook )
		r.GET("/book/delete",model.DelBook)
		// 4.监听端口，默认在8080
		r.Run(":8000")
}


func main() {

	initDB()
	initModule()

	
 }
 