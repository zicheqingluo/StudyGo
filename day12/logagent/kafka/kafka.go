package kafka

//专门往kafka写日志

import "github.com/Shopify/sarama"
import "fmt"

type logData struct{
	topic	string
	data	string
}

var (
	client sarama.SyncProducer
	logDataChan	chan *logData
)

//初始化client
func Init(addrs []string,maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 发送完数据需要 leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一一个 partition
	config.Producer.Return.Successes = true
	// 成功交付的消息将在 success channel返回
	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	logDataChan =make(chan *logData,maxSize)
	//开启后台的goroutine从通道中取数据发往kafka
	go sendToKafka()
	return
}

func SendToChan(topic,data string){
	msg := &logData{
		topic:topic,
		data:data,

	}
	logDataChan <- msg
}

func sendToKafka() {

	// 构造一一个消息
	for {
		select{
		case ld := <- logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			//发送到kafka
			pid, offset, err := client.SendMessage(msg)
			fmt.Println("xxx")
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:

		}
	}

}
