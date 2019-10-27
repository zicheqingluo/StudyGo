package model

import (
	"time"
)

//ArticleInfo 定义文章结构体,首页
type ArticleInfo struct {
	ID         int64 `db:"id"`
	CategoryID int64 `db:"category_id"`
	Summary	string `db:"summary"` //文章摘要
	Title string 	`db:"title"`
	ViewCount uint32	`db:"view_count"`
	CreateTime	time.Time	`db:"create_time"`
	CommentCount uint32	`db:"comment_count"`
	Username	string	`db:"username"`
}

//ArticleDetail 用于定义文章详情页实体
type ArticleDetail struct{
	ArticleInfo
	Content	string 	`db:"content"`
	Category
}

//ArticleRecord 用于文章上下页,上一篇、下一篇文章
type ArticleRecord struct{
	ArticleInfo
	Category

}