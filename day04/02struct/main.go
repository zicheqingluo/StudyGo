package main

import "fmt"

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main() {
	var p person
	p.name = "鸣人"
	p.age = 100
	p.gender = "男"
	p.hobby = []string{"足疗","大保健","按摩"}
	fmt.Println(p)
	fmt.Printf("type:%T value:%v \n",p,p)
}