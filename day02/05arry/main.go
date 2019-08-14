package main

import "fmt"


func main(){

	a1 := [3][2]string{

	{"北京","上海"},

	{"广州","重庆"},

	{"成都","深圳"},

	}

//遍历

	for _,v1 := range a1{

		for _,v2 := range v1 {

			fmt.Println(v2)

	}

	}

	b1 := [3]int{1,2,3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1,b2)

}

