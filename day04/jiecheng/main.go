package main

import "fmt"

func f(a int64)(int64){
	if a<=1{
		return 1
	}
			return a * f(a-1)
}

func main(){
	b := f(7)
				fmt.Println(b)
}