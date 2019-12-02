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
	//获取评论详情
	commentList,err := service.GetCommentList(articleId)
	if err != nil {
		fmt.Printf("get comment list failed,err:%v\n",err)
	}

	//获取文章详情
	articleDetail,err := service.GetArticleById(articleId)
	if err != nil {
		fmt.Printf("get article detail failed,err:%v\n",err)
	}

	//获取上一篇文章
	prev,err := service.GetPrevArticle(articleId)
	if err != nil {
		fmt.Printf("get prev detail failed,err:%v\n",err)
	}

	//获取下一篇文章
	next,err := service.GetNextArticle(articleId)
	if err != nil {
		fmt.Printf("get next detail failed,err:%v\n",err)
	}

	//获取相关文章
	categoryId,err := service.GetCategoryId(articleId)
	if err != nil {
		fmt.Printf("get  categoryid failed,err:%v\n",err)
	}

	relativeArticle,err := service.GetArticleRecordListBYID(int(categoryId),0,100)
	if err != nil {
		fmt.Printf("get  detail failed,err:%v\n",err)
	}

	//获取分类数据
	categoryList,err := service.GetALLCategoryList()

	if err != nil {
		fmt.Printf("get  categoryList failed,err:%v\n",err)
	}
	

	c.HTML(http.StatusOK,"views/detail.html",gin.H{
		"comment_list":commentList,
		"detail":articleDetail,
		"prev":prev,
		"next":next,
		"relative_article":relativeArticle,
		"category":categoryList,
		"article_id": articleId,
	})
}

//增加评论
func CommentSubmit(c *gin.Context)(){
	fmt.Println("postform:",c.PostForm(""))
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	articleIdStr := c.PostForm("article_id")
	articleId,err := strconv.ParseInt(articleIdStr,10,64)
	if err != nil {
		fmt.Println("strconv :",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	err = service.InsertComment(author,&comment,articleId)
	if err != nil {
		fmt.Println("service :",err)
		c.HTML(http.StatusInternalServerError,"views/500.html",nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d",articleId)
	c.Redirect(http.StatusMovedPermanently,url)

}