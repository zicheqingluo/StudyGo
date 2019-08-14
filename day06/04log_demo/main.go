package main

import (
	"time"
	"log"
)

func main() {
	for {
		log.Println("测试日志")
		time.Sleep(time.Second * 3)
	}
}