package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启128个goroutine
	for w := 1; w <= 128; w++ {
		go worker(w, jobs, results)
	}
	// go worker(1, jobs, results)
	// 500个任务
	for j := 1; j <= 500; j++ {
		jobs <- j
	}
	fmt.Println(len(jobs))
	close(jobs)
	
	// 输出结果

	for x := range results {
		fmt.Println(x)
	}
}
