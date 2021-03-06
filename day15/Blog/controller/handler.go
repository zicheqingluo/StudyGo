package controller

import (
	"path/filepath"
	"studygo/day15/Blog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"studygo/day15/Blog/model"
	"studygo/day15/Blog/service"
	"github.com/satori/go.uuid"
	"path"
	"log"
	//"path/filepath"
)

//访问主页的控制器
func IndexHandle(c *gin.Context) {
	//从service取数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//2.加载分类数据
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})

}

//点击分类云进行分类
func CategoryList(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	//转int
	categoryID, err := strconv.ParseInt(categoryIDStr, 110, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//根据分类id获取文章列表
	articleRecordList, err := service.GetArticleRecordListBYID(int(categoryID), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//再次加载所有分类数据，用于分类云显示
	categoryList, err := service.GetALLCategoryList()
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})

}

//ArticleDetail 文章评论
func ArticleDetail(c *gin.Context) {
	articleIdStr := c.Query("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//获取评论详情
	commentList, err := service.GetCommentList(articleId)
	if err != nil {
		fmt.Printf("get comment list failed,err:%v\n", err)
	}

	//获取文章详情
	articleDetail, err := service.GetArticleById(articleId)
	if err != nil {
		fmt.Printf("get article detail failed,err:%v\n", err)
	}

	//获取上一篇文章
	prev, err := service.GetPrevArticle(articleId)
	if err != nil {
		fmt.Printf("get prev detail failed,err:%v\n", err)
	}

	//获取下一篇文章
	next, err := service.GetNextArticle(articleId)
	if err != nil {
		fmt.Printf("get next detail failed,err:%v\n", err)
	}

	//获取相关文章
	categoryId, err := service.GetCategoryId(articleId)
	if err != nil {
		fmt.Printf("get  categoryid failed,err:%v\n", err)
	}
	fmt.Println("分类id:", categoryId)

	relativeArticle, err := service.GetArticleRecordListBYID(int(categoryId), 0, 100)
	if err != nil {
		fmt.Printf("get  detail failed,err:%v\n", err)
	}

	//获取分类数据
	categoryList, err := service.GetALLCategoryList()

	if err != nil {
		fmt.Printf("get  categoryList failed,err:%v\n", err)
	}

	c.HTML(http.StatusOK, "views/detail.html", gin.H{
		"comment_list":     commentList,
		"detail":           articleDetail,
		"prev":             prev,
		"next":             next,
		"relative_article": relativeArticle,
		"category":         categoryList,
		"article_id":       articleId,
	})
}

//增加评论
func CommentSubmit(c *gin.Context) {
	comment := c.PostForm("comment")
	author := c.PostForm("author")
	articleIdStr := c.PostForm("article_id")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	err = service.InsertComment(author, &comment, articleId)
	if err != nil {
		fmt.Println("service :", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	url := fmt.Sprintf("/article/detail/?article_id=%d", articleId)
	c.Redirect(http.StatusMovedPermanently, url)

}

//投稿
func NewArticle(c *gin.Context) {
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		fmt.Println("newarticle :", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	c.HTML(http.StatusOK, "views/post_article.html", categoryList)
}

//提交投稿
func ArticleSubmit(c *gin.Context) {
	author := c.PostForm("author")
	title := c.PostForm("title")
	categoryId := c.PostForm("category_id")
	content := c.PostForm("content")

	_, err := service.CreateArticle(author, content, title, categoryId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	c.Redirect(http.StatusMovedPermanently, "/")

}

//留言
func NewLeave(c *gin.Context) {
	leaveList, err := service.GetLeaveList()
	if err != nil {
		fmt.Println("newarticle :", err)
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	c.HTML(http.StatusOK, "views/gbook.html", leaveList)
}

//提交留言
func LeaveSubmit(c *gin.Context) {
	userName := c.PostForm("author")
	email := c.PostForm("email")
	content := c.PostForm("content")
	leaveInfo := &model.Leave{
		Username: userName,
		Email:    email,
		Content:  content,
	}
	err := service.AddLeave(leaveInfo)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	c.Redirect(http.StatusMovedPermanently, "/leave/new/")

}

func AboutMe(c *gin.Context){
	c.HTML(http.StatusOK,"views/about.html","")
}

// func UploadFile(c *gin.Context){
// 	file,err := c.FormFile("upload")
// 	if err != nil {
// 		fmt.Println("upload",err)
// 		c.JSON(http.StatusInternalServerError,gin.H{
// 			"message":err.Error(),
// 		})
// 		return
// 	}


// 	rootPath:= utils.GetRootDir()
// 	exc := path.Ext(file.Filename)
// 	aaa := uuid.NewV4()
// 	url := fmt.Sprintf("/static/upload/%s%s",aaa,exc)
// 	pat := filepath.Join(rootPath,url)
// 	_ = c.SaveUploadedFile(file,pat)
// 	c.JSON(http.StatusOK,gin.H{
// 		"uploaded":"true",
// 		"url":url,
// 	})

// }

func UploadFile(c *gin.Context) {
	// single file
	file, err := c.FormFile("upload")
	if err != nil {
		fmt.Println("load file",err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println(file.Filename)
	rootPath := utils.GetRootDir()
	u2 := uuid.NewV4()
	ext := path.Ext(file.Filename)
	url := fmt.Sprintf("/static/upload/%s%s", u2, ext)
	fmt.Println("file name:",url)
	dst := filepath.Join(rootPath, url)
	fmt.Println("路径",dst)
	// Upload the file to specific dst.
	_ = c.SaveUploadedFile(file, dst)
	c.JSON(http.StatusOK, gin.H{
		"uploaded": true,
		"url":      url,
	})
}
