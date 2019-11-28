package controller

import (
	"strconv"
	"net/http"
	"studygo/day15/Blog/service"
	"github.com/gin-gonic/gin"
	"fmt"
)

//访问主页的控制器
func IndexHandle(c *gin.Context){
	//从service取数据
	articleRecordList,err := service.GetArticleRecordList(0,15)
	if err !=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}

	//2.加载分类数据
	categoryList,err := service.GetALLCategoryList()
	if err !=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	c.HTML(http.StatusOK,"views/index.html",gin.H{
		"article_list": articleRecordList,
		"category_list":categoryList,
	})

}

//点击分类云进行分类
func CategoryList(c *gin.Context){
	categoryIDStr := c.Query("category_id")
	//转int
	categoryID,err := strconv.ParseInt(categoryIDStr,110,64)
	if err !=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//根据分类id获取文章列表
	articleRecordList,err := service.GetArticleRecordListBYID(int(categoryID),0,15)
	if err !=nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	//再次加载所有分类数据，用于分类云显示
	categoryList,err:= service.GetALLCategoryList()
	c.HTML(http.StatusOK,"views/index.html",gin.H{
		"article_list": articleRecordList,
		"category_list":categoryList,
	})

}

//ArticleDetail 文章评论
func ArticleDetail(c *gin.Context){
	articleIdStr := c.Query("article_id")
	articleId,err := strconv.ParseInt(articleIdStr,10,64)
	if err != nil {
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}

	commentList,err := service.GetCommentList(articleId)
	if err != nil {
		fmt.Printf("get comment list failed,err:%v\n",err)
	}
}