package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Width, Height int
}

func main() {
	//1.链接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	//2.调用方法
	ret := 0
	err2 := conn.Call("Rect.Area", Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("面积：", ret)

}
