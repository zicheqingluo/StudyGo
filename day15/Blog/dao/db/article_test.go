package db

import (
	"time"
	"studygo/day15/Blog/model"
	"testing"
)

func init(){
	dns:= "tom:123456@tcp(10.127.63.26:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err !=nil {
panic(err)
	}
}

//插入文章
func TestInsertArticle(t *testing.T){
	//构建对象
article := &model.ArticleDetail{}
article.ArticleInfo.CategoryID = 1
article.ArticleInfo.CommentCount = 0
article.ArticleInfo.CreateTime = time.Now()
article.ArticleInfo.Title = "5qi go"
article.ArticleInfo.Username = "yang"
article.ArticleInfo.Summary = "abc fd"
article.ArticleInfo.ViewCount = 1
article.Content = "联想小新pro13 锐龙"
articleID,err := InsertArticle(article)
if err != nil {
	return
}
t.Logf("articleID: %d \n",articleID)


}