package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	//数据库信息
	dsn := "root:@tcp(127.0.0.1:3306)/godata"
	//链接
	db,err = sql.Open("mysql",dsn)
	if err != nil {//dsn格式不正确会报错
		
		return
	}
	err = db.Ping()
	if err != nil {
		
		return
	}
	db.SetMaxOpenConns(10)
	return
	
}


type user struct{
	id int
	name string
	age int
}


func transactionDemo(){
	//1.开启事务
	tx,err := db.Begin()
	if err != nil {
		fmt.Println("begin error",err)
	}
	//执行多个sql、操作
	sqlStr1 := `update user set age = age-2 where id = 2`
	sqlStr2 := `update user set age = age+2 where id = 3`
	_,err = tx.Exec(sqlStr1)
	if err != nil {
		//回滚
		err := tx.Rollback()
		if err != nil {
			fmt.Println("error:",err)
			return
		}
	}
	_,err = tx.Exec(sqlStr2)
	if err != nil {
		//回滚
		err := tx.Rollback()
		if err != nil {
			fmt.Println("error:",err)
			return
		}
	}
	//如无报错提交事务
	err = tx.Commit()
	if err != nil {
		//要回滚
		tx.Rollback()
		fmt.Println("提交出错，要回顾年")
		return

	}
	fmt.Println("执行成功")

	



}

func main(){
	err := initDB()
	if err != nil {
		fmt.Println("链接错误",err)
	}
	transactionDemo()
}