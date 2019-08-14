package main

import "fmt"

func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	// var b map[string]int
	// b["沙河娜扎"] = 100
	// fmt.Println(b)
}