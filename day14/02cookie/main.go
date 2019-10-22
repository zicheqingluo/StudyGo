package main


import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
 )
 
 // 定义中间
 func cook() gin.HandlerFunc {
	return func(c *gin.Context) {


	r.GET("cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			c.JSON(200, gin.H{"request": "未登录"})


		}


	   c.Next()

	}
 }
}
 
 func main() {
	r := gin.Default()
	// 注册中间件
	r.Use(cook())
	// {}为了代码规范
	{
	   r.GET("/login", func(c *gin.Context) {
		  // 取值
		  //req, _ := c.Get("request")
		  //fmt.Println("request:", req)
		  //time.Sleep(3 * time.Second)
		  c.SetCookie("key_cookie", "value_cookie", 60, "/",
		  "localhost", false, true)
		  // 页面接收
		  c.JSON(200, gin.H{"request": "ok"})
	   })
	   r.GET("/home",func(c *gin.Context){
		c.JSON(200, gin.H{"request": "已登录"})
 
	   })
 
	}
	r.Run(":8000")
 }