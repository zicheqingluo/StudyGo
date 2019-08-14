package main

import (
	"fmt"
	"strings"


)


var (
	abc string
	def string
	
)

func tes(){
	abc = "你好123jj"
	def = "2345啦啦"
	ddd := strings.Split(def,"5")
	fmt.Println(ddd)
	fmt.Println(strings.HasPrefix(abc,"你"))
	fmt.Println(strings.HasSuffix(def,"啦"))
	fmt.Println(strings.Index(abc,"j"))
	fmt.Println(strings.Join(ddd,"+"))
}

func main(){
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

	tes()
}