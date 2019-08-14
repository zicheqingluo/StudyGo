package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//cat实现mover借口
func (c *cat) move() {
	fmt.Println("走猫步...")
}

//cat 实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s...\n",food)
}

func main(){
	var a1 animal

	c1 := cat{
		name : "小黄",
		feet: 4,
	}

	a1 = &c1
	a1.move()

	var m1 mover
	m1 = &c1
	m1.move()

	var e1 eater
	e1 = &c1
	e1.eat("老鼠")
}