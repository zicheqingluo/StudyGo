package main

//如何优雅的通知groutine退出，全局变量的方式
import (
	"sync"
	"fmt"
	"time"
)
	var wg sync.WaitGroup
	var notify bool

func f() {
	defer wg.Done()
	for {

	fmt.Println("suolong")
	time.Sleep(time.Millisecond * 500)
	if notify {
		break
	}
	}
}

func main() {
	
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}