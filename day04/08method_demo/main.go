package main

import "fmt"

type dog struct {
	name string
}

type person struct {
	name string
	age int
}

func newDog(name string) dog {
	return dog{
		name:name,
	}
}

func newPerson(name string,age int) person {
	return person{
		name:name,
		age:age,
	}
}

func (d dog) wang(){
	fmt.Printf("%s:www \n",d.name)
}

func (p person) guonian() {
	p.age++
}

func (p *person) guonian1() {
	p.age++
}

func main(){
	d1 := newDog("路飞")
	d1.wang()
	p1 := newPerson("索隆",20)
	p1.guonian()
	fmt.Println(p1.age)
	p1.guonian1()
	fmt.Println(p1.age)

}