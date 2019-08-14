package main

import "fmt"

//1.判断回文
func huiWen(x string){
	var breakFlag bool
	for i :=0;i< len(x)/2 -1;i++{
		if x[i] != x[len(x)-1-i]{
			breakFlag = true

			break
		} 
	}
	if breakFlag {
		fmt.Println("不是回文")
	} else {
		fmt.Println("是回文")
	}
		
	
}

//2. 字符统计
func stringCount(x string){
	//数字48-57
	//英文字符65-90 97-122
	//中文大于127
	//剩余是其他字符
	for _,v := range x {
		switch  {

		case v > 64 && v < 91:
			fmt.Printf("%c 是大写字母 \n",v)
		case v > 96 && v < 123:
			fmt.Printf("%c 是小写字母 \n",v)
		case v > 47 && v < 58:
			fmt.Printf("%c 是数字 \n",v)
		case v > 127:
			fmt.Printf("%c 是中文 \n",v)
		default:
			fmt.Printf("%c 是其他字符 \n",v)
		}
	}
}



func main(){

	//回文作业
	b := "abvba"
	huiWen(b)

	//字符统计作业
	// a := "1a,中国人G"
	// stringCount(a)





}

