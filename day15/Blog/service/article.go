package service

import (
	"studygo/day15/Blog/dao/db"
	"studygo/day15/Blog/model"
)

//GetArticleRecordList 获取文章和对应的分类
func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	//.获取文章列表
	articleInfoList, err := db.GetAricleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	//2.获取文章对应的分类
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	//返回页面，做聚合
	//遍历所有文章
	for _, article := range articleInfoList {
		//根据当前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//文章取出分类id
		categoryID := article.CategoryId
		for _, category := range categoryList {
			if categoryID == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)

	}
	return
}

//根据多个文章的id，获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo) (ids []int64) {
	//遍历文章，得到每个文章
LABEL:
	for _, article := range articleInfoList {
		categoryID := article.CategoryId
		//去重
		for _, id := range ids {
			if id == categoryID {
				continue LABEL
			}
		}
		ids = append(ids, categoryID)

	}
	return

}

//GetArticleRecordListBYID  按分类id获取文章
func GetArticleRecordListBYID(categoryID, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {

	//.获取文章列表
	articleInfoList, err := db.GetAricleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0 {
		return
	}
	//2.获取文章对应的分类
	categoryIds := getCategoryIds(articleInfoList)
	categoryList, err := db.GetCategoryList(categoryIds)
	if err != nil {
		return
	}
	//返回页面，做聚合
	//遍历所有文章
	for _, article := range articleInfoList {
		//根据当前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		//文章取出分类id
		categoryID := article.CategoryId
		for _, category := range categoryList {
			if categoryID == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)

	}
	return
}

//getArticleById 根据文章id查询文章详情
func GetArticleById(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	articleDetail, err = db.GetArticleDetail(articleId)
	if err != nil {
		return
	}
	return
}

//getArticleById 根据文章id查询上一篇文章
func GetPrevArticle(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	prevId := articleId - int64(1)
	if prevId >= 0 {
		articleDetail, err = db.GetArticleDetail(prevId)
		if err != nil {
			return
		}
		return
	}
	return
}


//getArticleById 根据文章id查询上一篇文章
func GetNextArticle(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	nextId := articleId + int64(1)
	if nextId >= 0 {
		articleDetail, err = db.GetArticleDetail(nextId)
		if err != nil {
			return
		}

		return
	}
	return
}

// 根据文章id获取对应的分类id
func GetCategoryId(articleId int64) (int64,error){

	articleDetail,err := db.GetArticleCategoryId(articleId)
	if err != nil{
		return 0,err
	}
	categoryId := articleDetail.ArticleInfo.CategoryId
	return categoryId,err
}
