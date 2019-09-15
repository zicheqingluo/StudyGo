package main

import (
	"fmt"
	"time"
	"sync"
	"context"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	LOOP:
		for {
			fmt.Println("db connecting...")
			time.Sleep(time.Millisecond * 10) //假设链接数据库10毫秒
			select {
			case <- ctx.Done():
				break LOOP
			default:

			}
		}
			fmt.Println("worker done1")
			wg.Done()
}


func main() {
	//设置50毫秒超时
	ctx,_ := context.WithTimeout(context.Background(),time.Millisecond * 50)
	// d := time.Now().Add(50 * time.Millisecond)
	// ctx, _ := context.WithDeadline(context.Background(), d)	
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	//cancel()	
	wg.Wait()
	fmt.Println("over")
}