package main

import "fmt"

//作业一见03const


//作业二
func foo1(){
	// math.MaxFloat32
	f1 := 1.2345
	fmt.Printf("%T %v \n",f1,f1)

	f2 := float32(1.2344)
	fmt.Printf("%T %v \n",f2,f2)

	b1 := true
	var b2 bool
	fmt.Printf("%T ,%v \n",b1,b1)
	fmt.Printf("%T ,%v \n",b2,b2)

	i1 := 10
	i2 := 'b'
	fmt.Printf("%T %v \n",i1,i1)
	fmt.Printf("%T %v \n ",i2,i2)
}
//作业三
func foo3(){
	a := "hello沙河小王子"
	count := 0
	for _,i := range(a){
		if i > 127{  // byte最大127
			count ++
		}
	}

	fmt.Printf("中文个数是：%d \n",count)
	
	// for i:=0;i<len(a);i++ {
	// 	fmt.Printf("%v(%c) \n",a[i],a[i])
	// for _,i := range a{
	// 	fmt.Printf("%v(%c) \n",i,i)

	// }
	


	// b := "abc"
	// byteb := []byte(b)
	// byteb[0] = 'x'
	// fmt.Println(string(byteb))

	// c := "你好"
	// runec := []rune(c)
	// runec[0] = '我'
	// fmt.Println(string(runec))


}
//作业四
func foo4(){
	i :=1
	for i<10{
		l := 1
		for l<=i{
			fmt.Printf("%d*%d=%d ",l,i,i*l)
			l++
		}
		fmt.Println()
		i++
	}
}

func main(){
	foo1()
	// foo2()
	// foo3()
	// foo4()
}