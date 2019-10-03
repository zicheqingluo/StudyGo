package taillog

import (
	"fmt"
	"studygo/day12/logagent/etcd"
	"time"

)

var tskMgr *tailLogMgr

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	tskMap map[string]*TailTask
	newConfChan	chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:logEntryConf,
		tskMap: make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}
	for _,logEntry := range logEntryConf {
		tailObj:= NewTailTask(logEntry.Path,logEntry.Topic)
		mk := fmt.Sprintf("%s_%s",logEntry.Path,logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
}
go tskMgr.run()

}

// 监听自己的newConfChan ,有了新的配置过来之后就做对应的处理
// 1.配置新增
// 2.配置删除
// 3.配置变更

func (t *tailLogMgr)run(){
	for {
		select {
		case newConf := <- t.newConfChan:
			fmt.Printf("新配置来了:%v",newConf)
			for _,conf := range newConf{
				mk := fmt.Sprintf("%s_%s",conf.Path,conf.Topic)

				_,ok := t.tskMap[mk]
				if ok{
						continue
				}else{
					//新增
					tailObj := NewTailTask(conf.Path,conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			for _,c1 := range t.logEntry{
				isDelete := true
				for _,c2:= range newConf{
					if c2.Path == c1.Path && c2.Topic == c1.Topic{
						isDelete = false
						continue
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s",c1.Path,c1.Topic)
					t.tskMap[mk].cancelFunc()


				}
			}
			// 找到原来有现在没有的，停掉
		default:
			time.Sleep(time.Second)
		}
	}
}

// 向外暴露一个函数，用于将更新后的logEntry放入chan
func NewConfChan() chan<- []*etcd.LogEntry  {
	return tskMgr.newConfChan
}