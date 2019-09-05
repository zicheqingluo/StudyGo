/*
 * @Author: yangxiaokang
 * @Date: 2019-09-03 14:43:31
 * @Last Modified by: yangxiaokang
 * @Last Modified time: 2019-09-04 18:31:23
 */

package main

import (
	//"bytes"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


var db *sql.DB
//定义用户注册信息结构体
type user struct{
	id int
	name string
}
//initDB数据库配置
func initDB() (err error) {
	//数据库信息
	dsn := "root:@tcp(127.0.0.1:3306)/godata"
	//链接
	db,err = sql.Open("mysql",dsn)
	if err != nil {//dsn格式不正确会报错
		fmt.Println("数据库连接信息格式不正确：",err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库链接失败",err)
		return
	}
	db.SetMaxOpenConns(10)
	return
	
}
//queryOne查询
func queryOne(username string) error{
	var u1 user
	//1.单条语句查询
	sqlStr := `select id,name from user where name=?;`
	err := db.QueryRow(sqlStr,username).Scan(&u1.id,&u1.name)
	if err != nil{

		return err
	}
	return nil
}
//insert 写入数据库
func insert(username ,password string,age string)(error){
	sqlStr := `insert into user(name,password,age) values(?,?,?)`
	//2.exec
	_,err := db.Exec(sqlStr,username,password,age)
	if err != nil {
		fmt.Println("错误",err)
		return err
	}
	return nil
}


//register用户注册
func register(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，返回注册html。
	if r.Method == "GET" {
		queryParam := r.URL.Query() // 自动帮我们识别URL中的query param
		name := queryParam.Get("name")
		age := queryParam.Get("age")
		fmt.Println(name, age)
		b, err := ioutil.ReadFile("./register.html")
		if err != nil {
			fmt.Println("register.html返回失败")
			return
		}
		w.Write(b)
	} else if r.Method == "POST" {
		//POST请求，提交用户注册信息，写入数据库
		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		r.Body.Close()
		//接收用户注册信息
		type User struct {
			Username string
			Password string 
			Age string
		}
		//返回注册结果
		type Result struct {
			Status bool `json:"status"`
			Info	string `json:"info"`
		}

		r1 := &User{}

		err = json.Unmarshal([]byte(result), r1)
		if err != nil {
			fmt.Println("err:", err)
		}

		for {

		
		if r1.Password == ""{
			result := &Result{
				Status: false,
				Info: "用户名不得为空",
			}
			resData,err := json.Marshal(result)
			if err != nil {
				fmt.Println("返回值序列化失败",err)
			}

			w.Write([]byte(resData))
			fmt.Println("用户名为空")
			break

		}
		err = queryOne(r1.Username)
		if err == nil {
			result := &Result{
				Status: false,
				Info: "用户名重复，请更改用户名",
			}
			resData,err := json.Marshal(result)
			if err != nil {
				fmt.Println("返回值序列化失败",err)
			}

			w.Write([]byte(resData))
			fmt.Println("用户名重复")
			break
		}

		//fmt.Println("res:", string(r1.Password), r1.Username)
		err = insert(r1.Username,r1.Password,r1.Age)

		if err != nil{
			fmt.Println("注册失败",err)
			result := &Result{
				Status: false,
				Info: "注册失败",
			}
			resData,err := json.Marshal(result)
			if err != nil {
				fmt.Println("返回值序列化失败",err)
			}

			w.Write([]byte(resData))
			break
			

		}else {


			result := &Result{
				Status: true,
				Info: "恭喜您注册成功",
			}

			resData,err := json.Marshal(result)
			if err != nil {
				fmt.Println("返回值序列化失败",err)
			}

			w.Write([]byte(resData))
			fmt.Println("注册成功")
			break
		}

		}	
	}

}

func main() {
	//初始化数据库链接
	err := initDB()
	if err != nil {
		fmt.Println("链接错误",err)
	}


	http.HandleFunc("/register/", register) //注册
	http.Handle("/static/", http.FileServer(http.Dir("."))) //静态文件
	http.ListenAndServe("0.0.0.0:9090", nil) //监听9090端口
}
