package main

import (
	
	"fmt"
	"strings"
)

func main() {
	a := "how do you do"
	m := make(map[string]int,10)

	for _,v := range strings.Split(a," ") {
		m[v] +=1
		
	}
	for k,v := range m {
		fmt.Printf("单词 %v 出现了 %v 次 \n",k,v)
	}
	

}