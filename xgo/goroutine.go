package main

import (
	"fmt"
	"time"
)

func g() {
	fmt.Println("g start")

	time.Sleep(3 * time.Second)

	fmt.Println("g end")
}

func t1() {
	fmt.Println("t1 start")

	go g()

	fmt.Println("t1 end")
}

func main() {
	fmt.Println("start")
	t1()
	time.Sleep(5 * time.Second)

	fmt.Println("end")
}
