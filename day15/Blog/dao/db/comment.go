package db

import (
	"studygo/day15/Blog/model"

	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)

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
