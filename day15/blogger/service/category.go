package service

import (
	"studygo/day15/blogger/dao/db"
	"studygo/day15/blogger/model"
	"fmt"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}
