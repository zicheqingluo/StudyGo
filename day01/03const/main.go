package main

import "fmt"

const pi = 3.1415926
const e = 2.8192

// const (
// 	a = 1
// 	b = 'c'
// 	c 
// 	d
// )

const (
	n1 = iota
	n2 
	_
	n4
)

const (
	a,b = iota +1 ,iota +2
	c,d
)

func foo2() {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	fmt.Printf("%d %d \n" ,KB,MB)
}
func main(){
	// fmt.Printf("a:%d b:%d c:%d d:%d \n",a,b,c,d)
	// fmt.Printf("a:%d b:%c c:%c d:%d \n",a,b,c,d)
	fmt.Printf("n1:%d n2:%d n4:%d \n",n1,n2,n4)
	fmt.Printf("a:%d b:%d c:%d d:%d \n ",a,b,c,d)
	foo2()
}