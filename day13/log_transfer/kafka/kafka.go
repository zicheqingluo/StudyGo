package kafka
import (
	"fmt"
	"github.com/Shopify/sarama"
	"studygo/day13/log_transfer/es"

)
// Init 初始化client
func Init(addrs []string,topic string)(err error) {
	consumer,err := sarama.NewConsumer(addrs,nil)
	if err != nil {
		fmt.Printf("fail to start consumer,err:%v \n",err)
		return err
	}
	partitionList,err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition,err:%v \n",err)
		return err 
	}
	fmt.Println("分区列表：",partitionList)
	for partition := range partitionList{
		pc,err := consumer.ConsumePartition(topic,int32(partition),sarama.OffsetNewest)
		if err != nil {
		fmt.Printf("fail to get list of partition,err:%v \n",err)
	}
	//defer pc.AsyncClose()
	// 异步从每个分区中消费信息
		go func(sarama.PartitionConsumer){	
			for msg:= range pc.Messages(){
				fmt.Printf("partition:%d offset:%d key:%v value:%v",msg.Partition,msg.Offset,msg.Key, string(msg.Value))
				// 发送到ES

				ld := es.LogData{
					Topic:topic,
					Data:string(msg.Value),
				}
				es.SendToESChan(&ld)
				fmt.Println("sendToES",ld)
				
			}
		
		}(pc)
}
fmt.Println("即将结束")
return err
	
}