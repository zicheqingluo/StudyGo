package main

import "fmt"

func sum(x int ,y int) (ret int) {
	return x + y
}

func f5() (int,string) {
	return 1,"沙河"
}

func f6(x, y, z int,m,n string) int{
	return x + y
}

func f7(x string,y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	_,n := f5()
	fmt.Println(n)

	a := f6(1,2,3,"杨","xiao")
	fmt.Println(a)
	f7("下雨")
	f7("下雨了",1,2,3)
}