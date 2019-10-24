package db


import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

type Book struct{
	Id int
	Title string
	Price int
}

func InitDB() (err error) {
	//数据库信息
	dsn := "tom:123456@tcp(10.127.63.26:3306)/cloud"
	//链接
	db,err = sql.Open("mysql",dsn)
	if err != nil {//dsn格式不正确会报错
		
		return
	}
	err = db.Ping()
	if err != nil {
		
		return
	}
	//db.SetMaxOpenConns(10)
	return
	
}

func Insert(title ,price string) {
	//1.写sql语句
	sqlStr := `insert into book(title,price) values(?,?)`
	//2.exec
	ret,err := db.Exec(sqlStr,title,price)
	if err != nil {
		fmt.Println("错误",err)
		return 
	}
	//如果插入数据能拿到插入数据的id
	fmt.Println(ret.LastInsertId())
}

func QueryMore(n int)[]Book {
	//BookList:= make([]Book,100)
	var BookList []Book

	sqlStr :=`select id,title,price from book where id>?;` 
	
	rows,err := db.Query(sqlStr,n)
	
	if err != nil {
		fmt.Printf("exec %s query fail",err)
	}
	defer rows.Close()
	for rows.Next() {
		var u2 Book
		err := rows.Scan(&u2.Id,&u2.Title,&u2.Price)
		if err != nil {
			fmt.Printf("exec %s query fail",err)	
		}
		BookList = append(BookList,u2)
		fmt.Printf("u1:%v \n",u2)
	}
	return BookList

}

func DeleteRow(id string) {
	sqlStr := `delete from book where id = ?`
	ret,err := db.Exec(sqlStr,id)
	if err != nil {
		fmt.Printf("update failed %s",err)
		return
	}
	n,err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("update failed %s",err)
		return
	}
	fmt.Println("影响行数为",n)
}