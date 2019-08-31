package main

import (
	"fmt"
	"github.com/sqlx"
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
