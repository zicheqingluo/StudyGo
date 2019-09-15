package main

//如何优雅的通知groutine退出，全局变量的方式
import (
	"sync"
	"fmt"
	"time"
)
	var wg sync.WaitGroup
	var exitChan = make(chan bool,1)

func f() {
	defer wg.Done()
LOOP:
	for {

	fmt.Println("suolong")
	time.Sleep(time.Millisecond * 500)
	select {
	case <-exitChan:
		break LOOP
	default:

	}
	}
}

func main() {
	
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	exitChan <- true
	wg.Wait()
}