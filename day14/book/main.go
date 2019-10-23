package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"StudyGo/day14/book/db"
)


func list(c *gin.Context) {

	bookList := db.queryMore(0)
	//c.String(http.StatusOK, "hello World!")
	c.HTML(200,bookList.tmpl,gin.H{
		"data": "bookList",)


 }
func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	//3.渲染模板
	r.LoadHTMLGlob("templates/*")
	// gin.Context，封装了request和response
	r.GET("/book/list",list )
	r.POST("/book/new",list )
	// 4.监听端口，默认在8080
	r.Run(":8000")
	err := db.InitDB()
	if err != nil {
		fmt.Println("链接错误",err)
	}
	
 }
 