package homework1

import (
	"time"
	"fmt"
	"path"
	"os"
	"strconv"
	
)

//如用户选择输出至文件，初始化文件
func (c *consoleLogger)FileInit(isFile bool,filePath,fileName string)(error){
	c.isFile = isFile
	c.filePath = filePath
	c.fileName = fileName
	//fmt.Println(c.isFile,isFile,filePath) //为毛不打印
	logPath := path.Join(filePath,fileName)
	fileObj,err := os.OpenFile(logPath,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil{
		fmt.Printf("open logpath failed,err \n",err)
		return err
	}
	c.fileObj = fileObj
	c.ch = make(chan Ch1,10)
	for j := 1; j<10;j++{

		go c.SaveFile()
	}
	return nil
}

//检查当前时间是否为每分钟的00秒
func (c *consoleLogger)checkTime()bool{
	timef := time.Now().Format("200601021504") + "00"
	checkTime,err := strconv.Atoi(timef)
	nowTime, err := strconv.Atoi(time.Now().Format("20060102150405"))
	if err != nil {
		fmt.Printf("转换错误：%s",err)
	}
	//fmt.Printf("时间 %v %T,标志时间:%v %T \n",nowTime,nowTime,checkTime,checkTime)	

	return checkTime == nowTime

}

//切割日志文件
func (c *consoleLogger)MoveFile(){
	c.fileObj.Close()
	nowStr := time.Now().Format("20060102150405")
	logPath := path.Join(c.filePath,c.fileName)
	newLogPath := logPath + nowStr
	os.Rename(logPath,newLogPath)

	fileObj,err := os.OpenFile(logPath,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("打开新的日志文件错误,%s",err)
		return 
	}
	c.fileObj = fileObj




}