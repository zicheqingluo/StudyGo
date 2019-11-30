package model

import (
	"time"
)

//ArticleInfo 定义文章结构体,首页
type ArticleInfo struct {
	Id         int64 `db:"id"`
	CategoryId int64 `db:"category_id"` //文章分类id
	Summary	string `db:"summary"` //文章摘要
	Title string 	`db:"title"` //文章标题
	ViewCount uint32	`db:"view_count"` //阅读次数
	CreateTime	time.Time	`db:"create_time"` //创建时间
	CommentCount uint32	`db:"comment_count"` //评论次数
	Username	string	`db:"username"` //用户名
}

//ArticleDetail 用于定义文章详情页实体
type ArticleDetail struct{
	ArticleInfo
	Content	string 	`db:"content"` //文章内容
	Category //分类结构体
}

//ArticleRecord 用于文章上下页,上一篇、下一篇文章
type ArticleRecord struct{
	ArticleInfo //首页结构体
	Category //文章分类结构体

}