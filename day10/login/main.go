/*
 * @Author: yangxiaokang
 * @Date: 2019-09-03 14:43:31
 * @Last Modified by: yangxiaokang
 * @Last Modified time: 2019-09-03 19:11:26
 */

package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	

	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
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
		//fmt.Println(ioutil.ReadAll(r.Body)) // 我在服务端打印客户端发来的请求的body
		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("result:", string(result))
		//data := fmt.Sprintf("`%s`",result)
		

		r.Body.Close()
		type Class struct {
			username string 
			password string 
		}
		var res Class
		s1 := `{"username":"111","password":"2222"}`
		err = json.Unmarshal([]byte(s1), &res)
		if err != nil {
			fmt.Println("err:", err)
		}
		fmt.Println("res:", res)
		w.Write([]byte("ok"))
	}

	//fmt.Println(r.Method)

}

func main() {

	http.HandleFunc("/login/", login)
	//http.HandleFunc("/static/", static)
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("0.0.0.0:9090", nil)
}
