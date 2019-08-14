package main

import (
	"bufio"
	"fmt"
	"os"
)

//1、定义日志级别，每个级别对应的日志输出位置
//2、定义可选的日志输出位置



//定义日志输出的位置的选型
type logFormat struct {
	level	string
	dir		string
	logName	string
}

//构造函数
func newLogFormat(level,dir,logName string) *logFormat {
	return &logFormat{
		level: level,
		dir: dir,
		logName: logName,
	}
}

//前台打印日志

func (l logFormat) stdoutPrint(logInfo string) {
	fmt.Printf("%s:  %s \n ",l.level,logInfo)

}

//持久化到文件
func (l logFormat) saveFile(logInfo string) {
	logFile := fmt.Sprintf("%v/%v",l.dir,l.logName)
	logInfo = fmt.Sprintf("【%v】 %v \n",l.level,logInfo)
	fileObj,err := os.OpenFile(logFile,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)	
	if err != nil {
		fmt.Println("%s 写入日志文件失败",l.logName)
	}
	defer fileObj.Close()
	wr := bufio.NewWriter(fileObj)
	wr.WriteString(logInfo)
	wr.Flush( )
}

//上传到kafka
func (l logFormat) upKafka() {
	fmt.Println("Kafka: 没定义kafka的接口，换其他方式")

}

// 定义日志级别
// type debug struct {
	
// 	dirName	string
// 	logName	string
// }

// func (d debug) log(loginfo string) {
// 	l := newLogFormat("Debug",d.dirName,d.logName)
// 	l.stdoutPrint(loginfo)

// }



//定义初始化,定义每种日志级别对应的日志输出位置
type initLog struct {
	dirName string
	logName string
}

//构造函数
func newInitLog(dirName,logName string) *initLog {
	return &initLog{
		dirName:dirName,
		logName:logName,
	}
}

//Debug,仅输入日志到屏幕
func (i initLog) Debug(loginfo string) {
	l := newLogFormat("Debug",i.dirName,i.logName)
	l.stdoutPrint(loginfo)
}

//INFO,输入日志到屏幕、kafka,并保存到日志
func (i initLog) Info(loginfo string) {
	l := newLogFormat("INFO",i.dirName,i.logName)
	l.stdoutPrint(loginfo)
	l.saveFile(loginfo)
	l.upKafka()
}
//ERROR,输入日志到屏幕,并保存到日志
func (i initLog) Error(loginfo string) {
	l := newLogFormat("ERROR",i.dirName,i.logName)
	l.stdoutPrint(loginfo)
	l.saveFile(loginfo)
}


func main() {
	logger := newInitLog(".","pay")
	logger.Debug("xxxxxxxx")
	logger.Info("这是一条info11111111111日志")
	logger.Error("返回参数错误")

}