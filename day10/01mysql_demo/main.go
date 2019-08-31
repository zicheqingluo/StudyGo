package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)




func main() {
	//数据库信息
	dsn := "root:@tcp(192.168.16.35:3306)/godata"
	//链接
	db,err := sql.Open("mysql",dsn)
	if err != nil {//dsn格式不正确会报错
		fmt.Println("dsn invalid",err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("链接不成功") //数据库链接失败报错
		return
	}
	fmt.Println("链接成功")
	



}