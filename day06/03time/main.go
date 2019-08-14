package main

import (
	"strings"
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Printf("time:%v\n",now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%d-%02d %d:%d:%d \n",year,month,day,hour,minute,second)

}

func main() {
	timeDemo()

}