package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启3个goroutine
	for id := 0; id < 3; id++ {
		go works(id, jobs, results)
	}

	// 生成5个job
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	// 输出结果
	for i := 0; i < 5; i++ {
		<-results
	}
}

func works(id int, jobs, result chan int) {
	for j := range jobs {
		time.Sleep(1 * time.Second)
		result <- j * 2
		fmt.Printf("JobID: {%d}, Job is: %d \n", id, j)
	}
}
