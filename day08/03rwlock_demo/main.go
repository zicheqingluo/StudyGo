package main

import (
	"sync"
	"time"
	"fmt"
)


//rwlock
var (
	x = 0
	lock sync.Mutex
	wg sync.WaitGroup
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()

	lock.Lock()
	//rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	lock.Unlock()
	//rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	//rwlock.Lock()
	x = x +1
	time.Sleep(time.Millisecond * 10)
	//rwlock.Unlock()
	lock.Unlock()
}


func main() {
	start := time.Now()

	for i :=0;i<10;i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i:=0;i<1000;i++{
		go read()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}