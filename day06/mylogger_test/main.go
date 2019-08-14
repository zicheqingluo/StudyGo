package main
import (
	"fmt"
	"time"
	"studygo/day06/homework1"
)


func main() {
	log := homework1.NewConsoleLogger("info")
	err := log.FileInit(true,"./","detail.log")
	if err != nil {
		fmt.Printf("日志文件初始化失败,%s",err)
		return
	}
	for {

	log.Debug("这是一条debug日志")
	log.Info("这是一条info日志")
	//time.Sleep(time.Millisecond * 300)
	time.Sleep(time.Second)
	}
}