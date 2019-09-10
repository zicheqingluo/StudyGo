package main

import (
	"fmt"
	"studygo/day11/logagent/kafka"
	"studygo/day11/logagent/taillog"
	"time"
)

func run() {
	//1.读取日志

	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
	//2.发送到kafka
}

func main() {
	//1.初始化kafka链接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v \n", err)
		return
	}
	fmt.Println("init kafka secuss...")
	//2.打开日志文件准备收集日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v \n", err)
	}
	fmt.Println("init taillog sucess...")
	run()

}
