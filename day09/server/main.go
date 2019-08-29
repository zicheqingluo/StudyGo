package main

import (
	"fmt"
	"net/http"
)

/*
 * @Author: yangxiaokang
 * @Date: 2019-08-29 15:46:56
 * @Last Modified by: yangxiaokang
 * @Last Modified time: 2019-08-29 16:47:03
 */
// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello 沙河！")
// }
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
func main() {
	http.HandleFunc("/", getHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
