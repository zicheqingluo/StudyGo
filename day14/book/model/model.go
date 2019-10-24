package model

import (
	"github.com/gin-gonic/gin"
	"studygo/day14/book/db"
	"fmt"
)

func List(c *gin.Context) {
	
	bookList := db.QueryMore(0)
	//c.String(http.StatusOK, "hello World!")
	fmt.Println("list后",bookList)
	c.HTML(200,"book_list.html",gin.H{
		"data": bookList,})
 }

 func AddBookHtml(c *gin.Context){
	c.HTML(200,"book_new.html",gin.H{})
 }

 func AddBook(c *gin.Context){
	title := c.PostForm("title")
	price := c.PostForm("price")
	db.Insert(title,price)
	c.Redirect(302,"http://127.0.0.1:8000/book/list")
 }

 func DelBook(c *gin.Context){
	 id := c.DefaultQuery("id","0")
	 db.DeleteRow(id)
	 c.Redirect(302,"http://127.0.0.1:8000/book/list") 
 }