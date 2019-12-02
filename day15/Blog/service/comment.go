package service


import (
	"fmt"
	"studygo/day15/Blog/dao/db"
	"studygo/day15/Blog/model"
	"time"
)

//GetCommentList 根据文章id获取评论列表
func GetCommentList(articleId int64) (commentList []*model.CommentInfo,err error) {
	commentList,err = db.GetCommentListByArticleID(articleId)
	if err != nil {
		return 
	}
	return
}

func InsertComment(userName string,content *string,articleId int64)(err error){
	exist,err := db.IsArticleExit(articleId)
	if err != nil {
		fmt.Printf("query database failed, err:%v\n", err)
		return 
	}

	if exist == false {

		err = fmt.Errorf("article id:%d not found", articleId)
		return 
	}
    var article int64 
	article,err = db.InsertComment(userName,content,1,articleId,time.Now())
	fmt.Println("article:",article)
	fmt.Println("err:",err)
	if err != nil || article == 0{
		fmt.Println("插入评论错误")
		return 
	}
	return nil


}