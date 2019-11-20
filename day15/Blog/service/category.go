package service

import (
	"studygo/day15/Blog/dao/db"
	"studygo/day15/Blog/model"
)

//GetALLCategoryList 获取所有分类
func GetALLCategoryList()(categoryList []*model.Category,err error){
	categoryList,err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}