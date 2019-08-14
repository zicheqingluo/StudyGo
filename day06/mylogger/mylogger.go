package mylogger

import (
	"path"
	"fmt"
	"runtime"
	"strings"
	"errors"
)
type LogLevel uint16



const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR 
	FATAL 
)

func parseLogLevel(s string) (LogLevel,error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG,nil
	case "trace":
		return TRACE,nil
	case "info":
		return INFO,nil

	case "warning":
		return WARNING,nil
	case "fatal":
		return FATAL,nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN,err
	}

}

func getInfo(skip int) (funcName,fileName string,lineNo int) {
	pc,file,lineNo,ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed \n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return 
}