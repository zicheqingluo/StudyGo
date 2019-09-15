package main

import (
	"studygo/day11/logagent/conf"
	"fmt"
	"studygo/day11/logagent/kafka"
	"studygo/day11/logagent/taillog"
	"time"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func run() {
	//1.读取日志

	for {
		select {
		case line := <-taillog.ReadChan():
			//kafka.SendToKafka("web_log", line.Text)
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
	//2.发送到kafka
}

func main() {
	//0.加载配置文件
	// cfg,err := ini.Load("./conf/config.ini")
	// addr := cfg.Section("kafka").Key("address").String()
	// topic := cfg.Section("topic").Key("topic").String()
	// path := cfg.Section("taillog").Key("path").String()
	// fmt.Println(addr)
	err := ini.MapTo(cfg,"./conf/config.ini")
	if err != nil {
		fmt.Println("init kafka failed",err)
		return
	}

	//1.初始化kafka链接
	//err = kafka.Init([]string{"127.0.0.1:9092"})
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	fmt.Println("地址",cfg.KafkaConf.Address,cfg.KafkaConf.Topic,cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("init kafka failed,err:%v \n", err)
		return
	}
	fmt.Println("init kafka secuss...")
	//2.打开日志文件准备收集日志
	//err = taillog.Init("./my.log")
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed,err:%v \n", err)
	}
	fmt.Println("init taillog sucess...")
	run()

}
