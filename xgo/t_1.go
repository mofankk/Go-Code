package main

import "fmt"

func main() {

	a := make(chan int, 1)
	b := make(chan int, 1)
	c := make(chan int, 2)

	a <- 1
	b <- 2
	go merge(a, b, c)

	for t := range c {
		fmt.Println(t)
	}

}

func merge(a, b chan int, c chan int) {
	for {
		select {
		case _a := <-a:
			c <- _a
		case _b := <-b:
			c <- _b
		}
	}
}
