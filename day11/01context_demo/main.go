package main

import (
	"sync"
	"fmt"
	"time"
)
	var wg sync.WaitGroup

func f() {
	defer wg.Done()
	for {

	fmt.Println("suolong")
	time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	
	wg.Add(1)
	go f()
	wg.Wait()
}