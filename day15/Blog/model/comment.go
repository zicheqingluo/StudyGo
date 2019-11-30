package model

import "time"

//定义文章评论结构体
type CommentInfo struct{
	Id	int64 `db:"id"`
	Username string	`db:"username"`
	CreateTime time.Time 	`db:"create_time"`
	Content string	`db:"content"`
	Status	int64	`db:"status"`
	ArticleId int64 `db:"article_id"`
	
}