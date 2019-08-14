package main

import "fmt"



func main(){
	// 十进制
	var a int = 10
	fmt.Println("")
	fmt.Printf("十进制数字10：%d \n",a)
	fmt.Printf("二进制数字10：%b \n",a)
	fmt.Printf("八进制数字10：%o \n",a)
	fmt.Printf("十六进制数字10：%x \n",a)
	i2 :=077 //八进制
	i3 :=0x1234567 //十六进制
	fmt.Printf("%T \n",i2)
	fmt.Printf("%T \n",i3)
	//建立一个int8类型的变量
	i4 := int8(3)
	fmt.Printf("%T \n",i4)
}