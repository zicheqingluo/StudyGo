package main

import "fmt"
import "strings"

var a string = "abcdABC"
var (
	A int
	B int
	C int
)

func main(){
	a = strings.ToUpper(a)
	// b := make(map[string]int,26)
	A = 0
	B = 0
	C = 0
	for _,c := range a{

		switch c {
		case 65:
			A += 1
		case 66:
			B += 1
		case 67:
			C += 1
		default:
			fmt.Println(" ")

			
		}
	}
	fmt.Println(A,B,C)

	

}