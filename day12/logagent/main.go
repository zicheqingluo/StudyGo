package main

import (
	"studygo/day12/logagent/conf"
	"fmt"
	"studygo/day12/logagent/kafka"
	//"studygo/day12/logagent/taillog"
	"studygo/day12/logagent/etcd"
	"time"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	//1.读取日志

// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			//kafka.SendToKafka("web_log", line.Text)
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}
// 	//2.发送到kafka
// }

func main() {
	//0.加载配置文件

	err := ini.MapTo(cfg,"./conf/config.ini")
	if err != nil {
		fmt.Println("init kafka failed",err)
		return
	}

	//1.初始化kafka链接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafka failed,err:%v \n", err)
		return
	}

	fmt.Println("init kafka secuss...")
	//2.初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address,time.Duration(cfg.EtcdConf.Timeout) * time.Second)
	if err != nil {
		fmt.Println("init etcd failed",err)
		return
	}

	fmt.Println("init etcd success...")
	//2.1从etcd中获取日志收集项的配置信息
	logEntryConf,err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("从etcd中获取配置失败GetConf %v \n",err)
		return
	}
	fmt.Println("get conf from etcd success,%v \n",logEntryConf)
	for _,v := range logEntryConf{
		fmt.Println(v)
	}
	//2.2派一个哨兵监视日志收集项的变化，有变化及时通知logagent实现热加载配置

	//2.打开日志文件准备收集日志
	//err = taillog.Init("./my.log")
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("Init taillog failed,err:%v \n", err)
	// }
	// fmt.Println("init taillog sucess...")
	// run()

}
