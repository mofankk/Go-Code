package main

import "fmt"

func main() {
	tt(1)
}

func tt(t int) {

	fmt.Println("before defer")
	if t == 1 {
		return
	}

	defer fmt.Println("defer run")

	fmt.Println("after defer")
}
