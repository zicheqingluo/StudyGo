package db

import (
	"github.com/jmoiron/sqlx"
	"studygo/day15/Blog/model"
)
//InsertCategory 添加分类
func InsertCategory(category *model.Category)(categoryID int64,err error){
	sqlstr:="insert into category(category_name,category_no) value(?,?)"
	result,err:=DB.Exec(sqlstr,category.CategoryName,category.CategoryNo)
	if err !=nil {
		return
	}
	categoryID ,err = result.LastInsertId()
	return
}


//GetCategoryByID 获取单个文章分类
func GetCategoryByID(id int64)(category *model.Category,err error) {
	category = &model.Category{}
	sqlstr := "select id,category_name,category_no from category where id = ?"
	err = DB.Get(category,sqlstr,id)
	if err !=nil{
		return nil,err
	} 
	return

}
//GetCategoryList 获取多个分类
func GetCategoryList(categoryIds []int64)(categoryList []*model.Category,err error){
	sqlstr,args,err := sqlx.In("select id,category_name,category_no from category where id in (?)",categoryIds)
	if err != nil{
		return
	}
	err = DB.Select(&categoryList,sqlstr,args...)
	return
}
//GetAllCategoryList 获取所有分类
func GetAllCategoryList() (categoryList []*model.Category,err error){
	sqlstr := "select id,category_name,category_no from category order by category_no asc"
	err = DB.Select(&categoryList,sqlstr)
	return
}