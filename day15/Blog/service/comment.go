package service


import (
	"studygo/day15/Blog/dao/db"
	"studygo/day15/Blog/model"
)

//GetCommentList 根据文章id获取评论列表
func GetCommentList(articleId int64) (commentList []*model.CommentInfo,err error) {
	commentList,err = db.GetCommentListByArticleID(articleId)
	if err != nil {
		return 
	}
	return
}