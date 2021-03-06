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
//需要收集的日志的配置
type LogEntry struct {
	Path string `json:"path"` //日志存放路径
	Topic string `json:"topic"` //日志要发往kafka中的哪个topic
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

// etcd watch
func WatchConf(key	string,newConfCh chan<- []*LogEntry){
	ch := cli.Watch(context.Background(),key)
	for wresp := range ch {
		for _,evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v \n",evt.Type,string(evt.Kv.Key),string(evt.Kv.Value))
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				err := json.Unmarshal(evt.Kv.Value,&newConf)
				if err != nil {
					fmt.Printf("unmarshal failed,err:%v \n",err)
					continue
				}
			}
			fmt.Printf("get new conf:%v \n",newConf)
			newConfCh <- newConf
		}
	}
}