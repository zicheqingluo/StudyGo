package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
	"studygo/day12/logagent/kafka"
	"context"
)

// 专门从日志文件收集日志的模块

var (
	tailObj *tail.Tail
	LogChan chan string
	
)

type TailTask struct {
	path string
	topic string
	instance *tail.Tail
	//为了能实现退出t.run()
	ctx context.Context
	cancelFunc	context.CancelFunc
}

func NewTailTask(path,topic string) (tailObj *TailTask){
	ctx,cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path : path,
		topic: topic,
		ctx:ctx,
		cancelFunc:cancel,
	}
	tailObj.init() //根据路径打开对应的日志
	return
}

func (t *TailTask)init() () {
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}
	// 当goroutine执行的函数退出的时候，goroutine就结束了
	go t.run() //采集日志
}

func (t *TailTask)run(){
	for {

		select {
			
		case line:= <- t.instance.Lines:
				kafka.SendToChan(t.topic,line.Text)
		case <- t.ctx.Done():
			fmt.Printf("tail task:%v 退出 \n",t.path+t.topic)
			return 
			
		}
	}
}


