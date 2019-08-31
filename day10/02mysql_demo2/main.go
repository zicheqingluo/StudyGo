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

func queryOne(){
	var u1 user
	//1.单条语句查询
	sqlStr := `select id,name,age from user where id=?;`
	//2.执行语句
	// for i:=0;i<13;i++{

	// 	db.QueryRow(sqlStr,2)
	// 	fmt.Printf("建立第%v个链接\n",i)

	// }
		db.QueryRow(sqlStr,2).Scan(&u1.id,&u1.name,&u1.age)
	//3.拿结果

	//rowObj.Scan(&u1.id,&u1.name,&u1.age)  //不调用scan会一直建立链接
	//4.打印结果
	fmt.Printf("u1:%#v \n",u1)
}

func queryMore(n int) {

	sqlStr :=`select id,name,age from user where id>?;` 
	rows,err := db.Query(sqlStr,n)
	if err != nil {
		fmt.Printf("exec %s query fail",err)
	}
	defer rows.Close()
	for rows.Next() {
		var u2 user
		err := rows.Scan(&u2.id,&u2.name,&u2.age)
		if err != nil {
			fmt.Printf("exec %s query fail",err)	
		}
		fmt.Printf("u1:%#v \n",u2)
	}

}

func insert() {
	//1.写sql语句
	sqlStr := `insert into user(name,age) values("wusuopu",34)`
	//2.exec
	ret,err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("错误",err)
		return 
	}
	//如果插入数据能拿到插入数据的id
	fmt.Println(ret.LastInsertId())
}

func updateRow(na,id int) {
	sqlStr := `update user set age=? where id = ?`
	ret,err := db.Exec(sqlStr,na,id)
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

func deleteRow(id int) {
	sqlStr := `delete from user where id = ?`
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

//预处理方式插入
func prepareInsert(){
	sqlStr := `insert into user(name,age) values(?,?) `
	stmt,err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预处理执行失败",err)
		return
	}
	defer stmt.Close()
	//后续只需要拿stmt执行操作
	var m = map[string]int{
		"鸣人":10,
		"佐助":20,
		"小樱":30,
	}
	for k,v := range m {

		stmt.Exec(k,v)
	}

}


func main() {
	
	err := initDB()
	if err != nil {
		fmt.Println("链接错误",err)
	}

	//queryMore(0)
	//insert()
	//updateRow(100,1)
	//deleteRow(1)
	prepareInsert()




}