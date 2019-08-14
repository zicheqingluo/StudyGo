package main

import "fmt"

func assign(a interface{}) {
	fmt.Printf("%T\n",a)
	str,ok := a.(string)
	if !ok {
		fmt.Println("猜错了")

	} else{
		fmt.Println("传进来的是一个字符串:",str)

	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n",a)
	switch t := a.(type){
	case string:
		fmt.Println("字符串:",t)
	case int:
		fmt.Println("是Int:",t)
	case int64:
		fmt.Println("是int64:",t)
	case bool:
		fmt.Println("是bool:",t)
	}
}

	func main() {
		assign2(true)
		assign2("哈哈")
		assign(8)
	}