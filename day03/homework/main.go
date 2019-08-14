package main

import (
	"fmt"
	"strings"
)


var (
	coins = 5000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	E int
	I int
	O int
	U int
	sum int 
	sum1 int
)


func dispatchCoin() (left int) {
	for _,str := range users{
		// fmt.Println(str)
		str = strings.ToUpper(str)
		// fmt.Println("修改后",str)
		E,I,O,U = 0,0,0,0
		for _,st := range str{

			switch st {
			case 69:
				E += 1
			case 73:
				O += 2
			case 79:
				I += 3
			case 85:
				U += 4
			default:
				
			}
		sum = E + I + U + O
		distribution[str] = sum
	}
	
}
// fmt.Println(distribution)
	
	for _,v := range distribution{
		
		sum1 += v
	}
	
	
return coins - sum1
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

