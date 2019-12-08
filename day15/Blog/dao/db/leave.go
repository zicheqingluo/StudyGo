package db

import (
	"fmt"
	"studygo/day15/Blog/model"
	"time"

	_ "github.com/go-sql-driver/mysql"

	//"database/sql"
)

//GetLeaveList 查询留言
func GetLeaveList() (leaveList []*model.Leave, err error) {
	sqlStr := `select 
	username,email,create_time,content 
	from leavelist order by create_time desc limit 0,3`
	err = DB.Select(&leaveList, sqlStr)
	if err != nil {
		fmt.Println("db get leave list ",err)
		return
	}
	return
}

//addLeave 插入留言
func AddLeave(leaveInfo *model.Leave) (articleId int64,err error) {
	sqlStr := `insert into 
	leavelist(username,email,content,update_time) 
	values(?,?,?,?)`
	result, err := DB.Exec(sqlStr, leaveInfo.Username, leaveInfo.Email,leaveInfo.Content, time.Now())
	if err != nil{
		return
	}
	articleId,err = result.LastInsertId()
	if err != nil{
		return
	}
	return
}
