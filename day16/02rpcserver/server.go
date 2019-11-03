package main

import (
	"log"
	"net/http"
	"net/rpc"
)
//计算加减法

type Params struct {
	Width, Height int
}

type Rect struct{}

//乘法
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Height * p.Width
	return nil
}

//除法 
func (r *Rect) Perimeter(p Params, ret,ret1 *int) error {
	*ret = p.Height/p.Width
	*ret1 = p.Height % p.Width
	return nil
}

//主函数
func main() {
	//1.注册服务
	rect := new(Rect)
	rpc.Register(rect)
	//2。服务
	rpc.HandleHTTP()
	//3.监听服务
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
