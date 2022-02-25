package main

import "fmt"

func main() {

	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		// panic: close of nil channel
		// ch 没有初始化，关闭时会panic
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}
