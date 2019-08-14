package main

import "fmt"

type animal struct {
	name string
}

func (a animal)move() {
	fmt.Printf("%s会动! \n",a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog)wang () {
	fmt.Printf("%s会叫 \n",d.name)
}

func main() {
	d1 := dog{
		animal : animal{name:"乔巴"},
		feet: 4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}