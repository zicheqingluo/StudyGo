package homework1

import (
	"os"
	"strings"
	"fmt"
	"time"
)
//定义结构体
type Ch1 struct {
	msg *string
	fo *os.File
}

//var ch chan Ch1

//定义结构体
type consoleLogger struct {
	level	logLevel	
	isFile	bool
	filePath	string
	fileName	string
	fileObj		*os.File
	ch		chan Ch1  //定义Ch1类型的通道
}
//构造函数
func NewConsoleLogger(level string) *consoleLogger{
	level = strings.ToUpper(level)
	Level,err := parseLogLevel(level)
	if err != nil {
		fmt.Printf("解析级别出现错误，err",err)
	}
	
	//ch = make(chan string,10)
	return &consoleLogger{
		level:Level,
		isFile:false,
	}
}

//日志开关，确定当前日志级别是否可以打印，返回true则可以打印
func (c consoleLogger)isTrue(cLevel logLevel) bool {
	return cLevel >= c.level
}

//定义日志格式，打印日志到文件或屏幕
func (c *consoleLogger)messageFormat(lv logLevel,message *string){
		now := time.Now()
		funcName,fileName,lineNo := getInfo(3)
		//定义日志格式
		msg := fmt.Sprintf("[%s] [%v] [%s:%s:%d] %s\n",now.Format("2006-01-02 15:04:05"),getLogString(lv),fileName, funcName, lineNo,*message)
	if c.isTrue(lv) {
		fmt.Printf(msg) //输出至屏幕
		if c.isFile{
			if c.checkTime() {
				
				c.MoveFile()
				}

        //将格式化好的日志内容和要写入文件的句柄放入通道，异步写入磁盘
		ch2 := Ch1{&msg,c.fileObj,}
		c.ch <- ch2
		}

			//fmt.Fprintf(c.fileObj,msg)

		

}
}

	//写文件线程，从通道中读取格式化好的日志内容和文件句柄。
	func (c consoleLogger)SaveFile(){
		for ch3 := range c.ch {
			fo := ch3.fo
			msg := ch3.msg
			fmt.Fprintf(fo,*msg)
		}
		// ch3,ok := <- c.ch
		// if !ok {
			// fmt.Printf("errrr: %s",ok)
		// }
			// fo := ch3.fo
			// msg := ch3.msg
			// fmt.Fprintf(fo,*msg)


		
	}
	

func (c *consoleLogger)Debug(message string){
	c.messageFormat(DEBUG,&message)
}

func (c *consoleLogger)Info(message string){
	c.messageFormat(INFO,&message)
}

