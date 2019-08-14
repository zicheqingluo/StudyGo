package mylogger

import (
	"fmt"
	"time"
)

//往终端写日志



type Logger struct {
	Level LogLevel

}
//logger 日志结构体
//newlog构造函数

func NewLog(levelStr string) Logger {
	level,err := parseLogLevel(levelStr)
	if err != nil{
		panic(err)
	}
	return Logger{
		Level:level,
	}
}

func (l Logger) enable(LogLevel LogLevel)bool {
	return LogLevel >= l.Level 
}

func getLogString(lv LogLevel) string{
	switch lv{
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	}
	return "DEBUG"
}

func (l Logger) log(lv LogLevel,format string,a ...interface{}) {
	if l.enable(lv) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	now1 := now.Format("2006-01-02 15:04:05")
	funcName,fileName,lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n",now1,getLogString(lv),funcName,fileName,lineNo,msg)
	}

}

func (l Logger) Debug(format string,a ...interface{}) {
		l.log(DEBUG,format,a...)

}
func (l Logger) Info(format string,a ...interface{}) {

		l.log(INFO,format,a...)

}