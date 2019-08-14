package main

import "fmt"

func sum1(){
	var a = [...] int {1,3,5,7,8}
	s := 0
	for _,v := range a{
		s += v
	}
	fmt.Println(s)
}

func sum2(){
	var a = [...] int {1,3,5,7,8}
	s := 0
	l := fmt.Sprintf(len(a))
	for ;s<l;s++;{
		a :=0
		a += s
		for ;a<l;
	}
	fmt.Println(s)
}

func main(){
	sum1()
}


