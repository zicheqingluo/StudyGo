/*
 * @Author: yangxiaokang
 * @Date: 2019-08-29 15:47:02
 * @Last Modified by: yangxiaokang
 * @Last Modified time: 2019-08-29 17:04:51
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func getfunc() {
	apiUrl := "http://127.0.0.1:9090/"
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("解析url错误,err:%v \n", err)

	}

	u.RawQuery = data.Encode() //url encode
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("发送请求失败,err:%v \n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取返回出错,err:%v \n", err)
	}
	fmt.Println(string(b))
}

func postfunc() {
	apiUrl := "http://127.0.0.1:9090/"
	//data := url.Values{}
	//data.Set("name", "小王子")
	//data.Set("age", "18")
	//u, err := url.ParseRequestURI(apiUrl)
	// if err != nil {
	// 	fmt.Printf("解析url错误,err:%v \n", err)

	// }

	// u.RawQuery = data.Encode() //url encode
	// fmt.Println(u.String())
	// resp, err := http.Get(u.String())
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("发送请求失败,err:%v \n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取返回出错,err:%v \n", err)
	}
	fmt.Println(string(b))
}

func main() {
	postfunc()
}
