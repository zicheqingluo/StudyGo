package main

import "fmt"

func sum(x int,y int) (ret int) {
	return x + y
}

func f1(x int,y int){
	fmt.Println(x + y)
}

func f2(){
	fmt.Println("f2")
}

func f3() int {
	ret :=3
	return ret
}

func f4(x int,y int) (ret int){
	ret = x + y
	return
}

func f5() (int,string) {
	return 1,"沙河"
}

func f6(x,y,z int,m,n string,i,j bool) int {
	return x + y
}

func f7(x string,y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main(){
	r :=sum(1,2)
	fmt.Println(r)

	_,n := f5()
	fmt.Println(n)
	a := f6(1,2,3,"中","国",true,false)
	fmt.Println(a)

	f7("下雨")
	f7("下雨",1,2,3,4,5,6)
}