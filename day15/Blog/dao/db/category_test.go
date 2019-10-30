package db

import (
	"testing"
)

func init(){
	dns:= "tom:123456@tcp(10.127.63.26:3306)/blogger?parseTime=true"
	err := Init(dns)
	if err !=nil {
		panic(err)
	}
}

func TestGetCategoryByID(t *testing.T) {
	category,err := GetCategoryByID(1)
	if err != nil {
		panic(err)
	}
	t.Logf("catetory:%#v",category)

}

func TestGetCategoryList(t *testing.T){
	var categoryIds []int64
	categoryIds  = append(categoryIds,1,2,3)
	categoryList,err := GetCategoryList(categoryIds)
	if err !=nil {
		panic(err)

	}
	for _,v := range categoryList{
		t.Logf("id:%d category:%#v \n",v.CategoryID,v)
	}
}

func TestGetAllCategoryList(t *testing.T){
	categoryList,err := GetAllCategoryList()
	if err !=nil {
		panic(err)

	}
	for _,v := range categoryList{
		t.Logf("id:%d category:%#v \n",v.CategoryID,v)
	}
}