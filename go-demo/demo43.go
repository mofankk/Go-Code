package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Done()
		// panic: sync: WaitGroup is reused before previous Wait has returned
		// WaitGroup 在调用 Wait 之后是不能再调用 Add 方法的。
		wg.Add(1)
	}()
	wg.Wait()
}
