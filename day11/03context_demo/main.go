package main

import (
	"time"
	"fmt"
	"context"
	"sync"
)

var wg sync.WaitGroup

func f(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("suolong")
		time.Sleep(time.Millisecond * 500)
		select {
		case <- ctx.Done():
			break FORLOOP

		default:	
		}
	}

}

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}