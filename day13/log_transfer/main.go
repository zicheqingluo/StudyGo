package main

import(
	"studygo/day13/log_transfer/conf"
	"studygo/day13/log_transfer/kafka"
	"studygo/day13/log_transfer/es"
	"gopkg.in/ini.v1"
	"fmt"
)
// log transfer

func main() {
	//0.加载配置文件
	var cfg = new(conf.LogTransferCfg)  ///cfg 是指针
	err := ini.MapTo(cfg,"./conf/cfg.ini") //修改变量一定得传指针
	if err != nil {
		fmt.Printf("init config failed,err:%v \n",err)
	}
	fmt.Printf("cfg:%v \n",cfg)
	
	//1.初始化ES
	err = es.Init(cfg.ESCfg.Address,cfg.ESCfg.ChanSize)
	if err != nil {
		fmt.Printf("init ES clent failed,err:%v\n",err)
		return
	}
	//2.初始化kafka
	err = kafka.Init([]string{cfg.KafkaCfg.Address},cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init kafka failed:%v\n",err)
		return
	}
	fmt.Println("必须调到这")
	select {}
}