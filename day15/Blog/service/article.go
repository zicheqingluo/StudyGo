package service

import (
	"studygo/day15/Blog/dao/db"
	"studygo/day15/Blog/model"
)

//GetArticleRecordList 获取文章和对应的分类
func GetArticleRecordList(pageNum,pageSize int) (articleRecordList []model.ArticleRecord){
	//1.获取文章列表
	articleInfoList,err := db.GetAricleList(pageNum,pageSize)
	if err != nil {
		return
	}
	if len(articleInfoList) <= 0{
		return
	}
	//2.获取文章对应的分类
	categoryIds := getCategoryIds(articleInfoList)
	categoryList,err := db.GetCategoryList(categoryIds)
	if err != nil{
		return
	}
	//返回页面，做聚合
	//遍历所有文章
	for _,article := range articleInfoList{
		//根据当前文章，生成结构体
		articleRecord := &model.ArticleRecord{
			ArticleInfo:*article,
		}
		//文章取出分类id
		categoryID := article.CategoryID
		for _,category := range categoryList{
			if categoryID == category.CategoryID{
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList)

	}

}

//根据多个文章的id，获取多个分类id的集合
func getCategoryIds(articleInfoList []*model.ArticleInfo)(ids []int64) {
	//遍历文章，得到每个文章
	for _,article := range articleInfoList{
		categoryID:=article.CategoryID
		//去重
		for _,id := range ids{
			if id != categoryID{
				ids = append(ids,categoryID)
			}
		}
	}


}