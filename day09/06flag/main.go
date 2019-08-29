package main

import (
	"time"
	"fmt"
	"flag"
)

func main()  {
	name := flag.String("name","lucy","请输入名字")
	age := flag.Int("age",20,"请输入年龄")
	married := flag.Bool("married",false,"请输入年龄")
	cTime := flag.Duration("ct",time.Second,"结婚多久")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
}