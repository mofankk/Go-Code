package main

import "fmt"

func main() {

	go func(i int) {
		i++
		fmt.Println(i)
	}(0)

	for {
	}
}
