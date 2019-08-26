package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request) {
	b,err := ioutil.ReadFile("./xx.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v",err)))
	}
	w.Write(b)
}

func main(){
	http.HandleFunc("/heelo/",f1)
	http.ListenAndServe("0.0.0.0:9090",nil)
}