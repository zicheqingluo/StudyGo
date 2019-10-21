package main


import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
 )
 
 // 定义中间
 func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
	   t := time.Now()
	   fmt.Println("中间件开始执行了")
	   // 设置变量到Context的key中，可以通过Get()取
	   //c.Set("request", "中间件")
	   // 执行函数
	   c.Next()
	   // 中间件执行完后续的一些事情
	   status := c.Writer.Status()
	   fmt.Println("中间件执行完毕", status)
	   t2 := time.Now().Sub(t)

	   fmt.Println("执行时间:", t2)
	}
 }
 
 func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare())
	// {}为了代码规范
	{
	   r.GET("/middleware", func(c *gin.Context) {
		  // 取值
		  //req, _ := c.Get("request")
		  //fmt.Println("request:", req)
		  time.Sleep(3 * time.Second)
		  // 页面接收
		  c.JSON(200, gin.H{"request": "ok"})
	   })
 
	}
	r.Run(":8000")
 }