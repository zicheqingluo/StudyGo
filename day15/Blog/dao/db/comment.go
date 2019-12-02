package db

import (
	"studygo/day15/Blog/model"

	_ "github.com/go-sql-driver/mysql"
		
	"fmt"
	"time"
)

//GetCommentListByArticleID 根据文章id获取文章评论
func GetCommentListByArticleID(articleId int64) (commentList []*model.CommentInfo, err error) {
	sqlstr := `select
	id,article_id,username,create_time,content,status
from
comment
where
status = 1
and
article_id = ?`
	err = DB.Select(&commentList, sqlstr, articleId)
	return
}

//InsertComment 插入文章评论
func InsertComment(userName string,content *string,status,articleId int64,createTime time.Time)(int64,error){
	sqlstr := `insert into 
	comment(content,username,create_time,status,article_id)
	values(?,?,?,?,?)
	`
	result,err := DB.Exec(sqlstr,content,userName,createTime,status,articleId)
	if err != nil{
		fmt.Println("db err:",err)
		return 0,err
	}
	articleId,err = result.LastInsertId()
	return articleId,err
}