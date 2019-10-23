package db


import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)

var db *sql.DB

type Book struct{
	id int
	title string
	price int
}
type BookList []Book

func InitDB() (err error) {
	//数据库信息
	dsn := "tom:@tcp(10.127.63.26:3306)/cloud"
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

// func insert() {
// 	//1.写sql语句
// 	sqlStr := `insert into user(name,age) values("wusuopu",34)`
// 	//2.exec
// 	ret,err := db.Exec(sqlStr)
// 	if err != nil {
// 		fmt.Println("错误",err)
// 		return 
// 	}
// 	//如果插入数据能拿到插入数据的id
// 	fmt.Println(ret.LastInsertId())
// }

func queryMore(n int)BookList {
	BookList:= make([]Book,100,100)

	sqlStr :=`select id,title,price from book where id>?;` 
	rows,err := db.Query(sqlStr,n)
	if err != nil {
		fmt.Printf("exec %s query fail",err)
	}
	defer rows.Close()
	for rows.Next() {
		var u2 Book
		err := rows.Scan(&u2.id,&u2.title,&u2.price)
		if err != nil {
			fmt.Printf("exec %s query fail",err)	
		}
		BookList = append(BookList,Book)
		fmt.Printf("u1:%#v \n",u2)
	}
	return BookList

}