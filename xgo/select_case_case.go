package main

import (
	"fmt"
	"time"
)

/*
select 中一个case触法后，处理的过程中其他case就算触发也不会执行
select 中多个case同时触发，随机执行其中一个case
*/

func main() {

	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)

	go h2(ch2)
	go h1(ch1)

	select {
	case <-ch1:
		fmt.Println("1")
		time.Sleep(4 * time.Second)
		fmt.Println("2")
	case <-ch2:
		fmt.Println("3")
		time.Sleep(6 * time.Second)
		fmt.Println("4")
	}

}

func h1(t chan bool) {
	t <- true
	//	time.Sleep(1 * time.Second)
}

func h2(t chan bool) {
	t <- true
	time.Sleep(4 * time.Second)
}
