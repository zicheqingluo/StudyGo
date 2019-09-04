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

	_ "github.com/go-sql-driver/mysql"
)

func static(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println(r.URL)
		staticUrl := "." + fmt.Sprint(r.URL)
		b, err := ioutil.ReadFile(staticUrl)

		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
		w.Write(b)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// 对于GET请求，参数都放在URL上（query param）, 请求体中是没有数据的。
	if r.Method == "GET" {
		queryParam := r.URL.Query() // 自动帮我们识别URL中的query param
		name := queryParam.Get("name")
		age := queryParam.Get("age")
		fmt.Println(name, age)
		b, err := ioutil.ReadFile("./login.html")
		if err != nil {
			w.Write([]byte(fmt.Sprintf("%v", err)))
		}
		w.Write(b)
	} else if r.Method == "POST" {
		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		r.Body.Close()
		type User struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		r1 := &User{}

		err = json.Unmarshal([]byte(result), r1)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("res:", string(r1.Password), r1.Username)
		w.Write([]byte("ok"))
	}

}

func main() {

	http.HandleFunc("/login/", login)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("0.0.0.0:9090", nil)
}
