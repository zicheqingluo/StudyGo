package mylogger

//往文件写日志
import (
	"path"
	"os"

	"fmt"
	"time"
)

type FileLogger struct {
	level LogLevel
	filePath string
	fileName string
	fileObj *os.File
	errFileObj *os.File
	maxFileSize int64
}

func NewFileLogger(levelStr,fp,fn string, maxFileSize int64) (*FileLogger) {
	logLevel,err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	f1:= &FileLogger{
		level :logLevel,
		filePath: fp,
		fileName :fn,
		maxFileSize:maxFileSize,
	}
	err = f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *FileLogger) initFile()(error){
	fullFileName := path.Join(f.filePath,f.fileName)
	fileObj,err := os.OpenFile(fullFileName,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open log file failed,%v \n",err)
		return err
	}
	
	errFullFileName := path.Join(f.filePath,f.fileName)
	errFileObj,err := os.OpenFile(errFullFileName+".err",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open log file failed,%v \n",err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger)Close(){
	f.fileObj.Close()
	f.errFileObj.Close()
}

func (f *FileLogger) enable(LogLevel LogLevel)bool {
	return LogLevel >= f.level 
}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo ,err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err%v \n",err)
		return false
	}

	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File) (*os.File,error) {
	//需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo,err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v \n",err)
		return nil,err
	}
	logName := path.Join(f.filePath,fileInfo.Name())
	newLogName := fmt.Sprintf("%s.%s",logName,nowStr)
	// 1.关闭当前的日志文件
	file.Close()
	//2.备份文件
	os.Rename(logName,newLogName)
	//3.打开新的日志文件
	fileObj,err := os.OpenFile(logName,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open new file failed,err:%v \n",err)
		return nil ,err
	}
	//4.将新打开的日志文件赋值给f.fileobj
	return fileObj,nil
}


func (f *FileLogger) log(lv LogLevel,format string,a ...interface{}) {
	if f.enable(lv) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	now1 := now.Format("2006-01-02 15:04:05")
	funcName,fileName,lineNo := getInfo(3)

	if f.checkSize(f.fileObj) {
		newFile,err := f.splitFile(f.fileObj)
		if err != nil {
			return
		}
		f.fileObj = newFile


	}

	fmt.Fprintf(f.fileObj,"[%s] [%s] [%s:%s:%d] %s \n",now1,getLogString(lv),funcName,fileName,lineNo,msg)
	if lv > DEBUG{
		//如果记录日志大于等于ERROR日志，还要在error日志记录一份
		if f.checkSize(f.errFileObj) {
			newFile,err := f.splitFile(f.errFileObj)
			if err != nil {
				return
			}
			f.errFileObj = newFile
		}
	fmt.Fprintf(f.errFileObj,"[%s] [%s] [%s:%s:%d] %s \n",now1,getLogString(lv),funcName,fileName,lineNo,msg)
	}
	}

}
func log(lv LogLevel,format string,a ...interface{}) {
	msg := fmt.Sprintf(format,a...)
	now := time.Now()
	now1 := now.Format("2006-01-02 15:04:05")
	funcName,fileName,lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s \n",now1,getLogString(lv),funcName,fileName,lineNo,msg)

}

func (f *FileLogger) Debug(format string,a ...interface{}) {
		f.log(DEBUG,format,a...)

}
func (f *FileLogger) Info(format string,a ...interface{}) {

		f.log(INFO,format,a...)

}