package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.RWMutex
var count int

/*
1. 在A方案下，由于main中要加写锁，当有一个协程要加写锁时，其他协程不能获取到锁，但在A方案中， A、B、C整个调用链占用读锁5s，而主协程2s后就要加写锁，因此panic。
2. 在B方案下，A、B、C只占用读锁1s，而主协程要在2s后加写锁，因此程序正常运行。
*/
func main() {
	go A()
	//A: time.Sleep(2 * time.Second)
	//B:
	time.Sleep(3 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}
func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}
func B() {
	//A: time.Sleep(5 * time.Second)
	//B:
	time.Sleep(1 * time.Second)
	C()
}
func C() {
	mu.RLock()
	defer mu.RUnlock()
}
