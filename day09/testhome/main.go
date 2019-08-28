package main

//import "studygo/day09/hwhuiwen"
import "fmt"

func main() {
ss := "bbabb"
bd := ";b"
for _,b := range ss {
	//fmt.Printf("%v %T \n",a,b)
	for _,c := range bd {
		if b != c {
			fmt.Println("不相等")
		}else{
			fmt.Println("相等")
		}
	}
}

//hwhuiwen.BackToText(ss)
}