package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func writeDemo1() {
	fileObj,err := os.OpenFile("./xx.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v,err")
		return
	}
	// write
	fileObj.Write([]byte("测试1 路飞 \n"))
	//wirtestring
	fileObj.WriteString("波雅汉库克 \n")
}

func writeDemo2() {
	fileObj,err := os.OpenFile("./oo.txt",os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0644)
		if err != nil {
		fmt.Printf("open file failed, err:%v", err)
		return
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello沙河\n") // 写到缓存中
	wr.Flush()                  // 将缓存中的内容写入文件
}

func writeDemo3() {
	str := "罗罗诺亚.索隆"
	err := ioutil.WriteFile("./dd.txt",[]byte(str),0644)
	if err != nil {
		fmt.Println("输入错误,err:",err)
		return
	}
}

func main() {
	// writeDemo1()
	// writeDemo2()
	writeDemo3()
}