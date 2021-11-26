package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case i := <-ch:
		fmt.Println(i)

	default:
		fmt.Println("default")
	}
}
