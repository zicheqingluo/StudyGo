package main

import (
	"github.com/gin-gonic/gin"
	//"net/http"
	"fmt"
	"studygo/day14/book/db"
)



func list(c *gin.Context) {
	
	bookList := db.QueryMore(0)
	//c.String(http.StatusOK, "hello World!")
	fmt.Println("list后",bookList)
	c.HTML(200,"book_list.html",gin.H{
		"data": bookList,})
 }

 func addBookHtml(c *gin.Context){
	c.HTML(200,"book_new.html",gin.H{})
 }

 func addBook(c *gin.Context){
	title := c.PostForm("title")
	price := c.PostForm("price")
	db.Insert(title,price)
	c.Redirect(302,"http://127.0.0.1:8000/book/list")
 }

 func delBook(c *gin.Context){
	 id := c.DefaultQuery("id","0")
	 db.DeleteRow(id)
	 c.Redirect(302,"http://127.0.0.1:8000/book/list") 
 }
func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Println("链接错误",err)
	}
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	//3.渲染模板
	r.LoadHTMLGlob("templates/*")
	// gin.Context，封装了request和response
	r.GET("/book/list",list)
	r.GET("/book/new",addBookHtml )
	r.POST("/book/new",addBook )
	r.GET("/book/delete",delBook)
	// 4.监听端口，默认在8080
	r.Run(":8000")

	
 }
 