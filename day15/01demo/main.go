package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
	
	
)

func main() {
	// 初始化数据库
	//err := initDB()
	//if err != nil {
	//	panic(err)
	//}
	r := gin.Default()
	// 加载页面
	//r.LoadHTMLGlob("./book/templates/*")
	// 查询所有图书
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "ok"})
	})
	err := r.Run(":8000")
	if err != nil{
		panic(err)
	}
}
