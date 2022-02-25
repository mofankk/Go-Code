package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			// 这里会造成10个groutine都处于阻塞状态
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout, exit")
	}
	time.Sleep(time.Second * 2)
}

func WaitTimeout(wg *sync.WaitGroup, t time.Duration) bool {
	ch := make(chan bool)
	go func() {
		time.AfterFunc(t, func() {
			ch <- true
		})
	}()

	go func() {
		wg.Wait()
		ch <- false
	}()

	return <-ch
}
