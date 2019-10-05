package es

import (
	"time"
	"fmt"
	"strings"
	"github.com/olivere/elastic/v7"
	"context"

)

type LogData	struct{
	Topic string `json:"topic"`
	Data	string `json:"data"`
}

var (
	client	*elastic.Client
	ch  chan *LogData
	
)

//初始化ES，准备接收kafka发来的数据

func Init(address	string,chanSize int)(err error){
	if !strings.HasPrefix(address,"http://"){
		address = "http://"+address
	}
	fmt.Println("es地址:",address)
	client,err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogData,chanSize)
	go SendToES()
	return
}
//异步发送到channel
func SendToESChan(msg *LogData){
	ch <- msg
}

// 从chanel发送数据到ES
func SendToES() {
	for {
		select{
		case msg:= <- ch:
			put1, err := client.Index().Index(msg.Topic).BodyJson(msg).Do(context.Background())
			//注意，BodyJson一定得传结构体，能够转出json，如果传msg.Data会报错，因为他是string
			if err != nil {
				fmt.Println("发送es",err)
			}
			fmt.Printf("indexed student %s to index %s,type %s \n",put1.Id,put1.Index,put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}


}