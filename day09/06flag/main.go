package main

import (
	"time"
	"fmt"
	"flag"
)

func main()  {
	// name := flag.String("name","lucy","请输入名字")
	// age := flag.Int("age",20,"请输入年龄")
	// married := flag.Bool("married",false,"请输入年龄")
	// cTime := flag.Duration("ct",time.Second,"结婚多久")
	// flag.Parse()
	// fmt.Println(*name)
	// fmt.Println(*age)
	// fmt.Println(*married)
	// fmt.Println(*cTime)dir

	var name string
	var age int
	var married bool
	var delay time.Duration
	
	flag.StringVar(&name,"name","lucy","姓名")
	flag.IntVar(&age,"age",20,"年龄")
	flag.BoolVar(&married,"married",false,"婚否")
	flag.DurationVar(&delay,"d",0,"时间间隔")
	flag.Parse()
	fmt.Println(name,age,married,delay)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}