package main

import "fmt"

// 如果有多个defer，程序提前return，则后面的defer不执行

func Hello(t int) {
	fmt.Printf("Hello: %d \n", t)
}

func main() {
	defer Hello(1)
	a := 1
	if a > 0 {
		return
	}
	defer Hello(2)
}
