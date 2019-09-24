package etcd

import (
	"encoding/json"
	"time"
	"go.etcd.io/etcd/clientv3"
	"fmt"
	"context"
)

var (
	cli 	*clientv3.Client
)

type LogEntry struct {
	Path string `json:"path"`
	Topic string `json:"topic"`
}

//Init 初始化ETCD的函数
func Init(addr string,timeout time.Duration) (err error) {
	cli,err = clientv3.New(clientv3.Config{
		Endpoints: []string{addr},
		DialTimeout: 5* time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed,err:%v \n",err)
		return
	}
	return
}

//从etcd中根据key获取配置项

func GetConf(key string) (logEntryConf []*LogEntry ,err error){
	//get
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	resp,err := cli.Get(ctx,key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed,err:%v \n",err)
		return
	}
	for _,ev := range resp.Kvs{
		//fmt.Printf("%s:%s \n",ev.Key,ev.Value)
		err = json.Unmarshal(ev.Value,&logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal failed,%v \n",err)
			return
		}

	}
	return
	}
