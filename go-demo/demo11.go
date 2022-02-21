package main

import (
	"fmt"
	"time"
)

// 11. 请找出下面代码的问题所在。
func main() {
	ch := make(chan int, 20)
	go func() {
		for i := 0; i < 10; i++ {
			// 可能会造成panic
			ch <- i
		}
	}()

	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("closed!")
				return
			}
			fmt.Println(a)
		}
	}()

	close(ch)
	time.Sleep(time.Second * 20)
}

/*
解析:
在 golang 中 goroutine 的调度时间是不确定的，在题目中，第一个写 channel 的 goroutine 可能还未调 用，或已调用但没有写完时直接 close 管道，可能导致写失败，既然出现 panic 错误。
*/
