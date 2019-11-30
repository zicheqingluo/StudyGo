package db

import (
	"studygo/day15/Blog/model"

	_ "github.com/go-sql-driver/mysql"
)

//InsertArticle 插入文章
func InsertArticle(article *model.ArticleDetail) (articleID int64, err error) {
	//验证article是否为空
	if article == nil {
		return
	}
	sqlstr := `insert into
					article(content,summary,title,username,category_id,view_count,
						comment_count) values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlstr, article.Content, article.Summary, article.Title, article.Username,
		article.ArticleInfo.CategoryId, article.ArticleInfo.ViewCount, article.ArticleInfo.CommentCount)
	if err != nil {
		return
	}
	articleID, err = result.LastInsertId()
	return

}

//GetAricleList 获取文章列表，做分页
func GetAricleList(pageNum, pageSize int) (aricleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	//时间降序排列
	sqlstr := `select
	id,summary,title,view_count,create_time,comment_count,username,category_id
from
article
where
status = 1
order by create_time desc
limit ?,?`
	err = DB.Select(&aricleList, sqlstr, pageNum, pageSize)
	return
}

//GetArticleDetail 根据文章id，查询单个文章
func GetArticleDetail(articleID int64) (articleDetail *model.ArticleDetail, err error) {
	if articleID < 0 {
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select
	id,summary,title,view_count,content,create_time,comment_count,username,category_id
from
article
where
id = ?
and
status = 1
`
	err = DB.Get(articleDetail, sqlstr, articleID)
	return
}

//GetArticleListByCategoryID 根据分类id，查询这一类的文章
func GetArticleListByCategoryID(category_id, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlstr := `select
	id,summary,title,view_count,create_time,comment_count,username,category_id
from
article
where
status = 1
and
category_id = ?
order by create_time desc
limit ?,?`
	err = DB.Select(&articleList, sqlstr, category_id, pageNum, pageSize)
	return
}

//GetArticleDetail 根据文章id，查询对应的分类id
func GetArticleCategoryId(articleID int64) (articleDetail *model.ArticleDetail, err error) {
	if articleID < 0 {
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlstr := `select
	category_id
from
article
where
id = ?
and
status = 1
`
	err = DB.Get(articleDetail, sqlstr, articleID)
	return
}
