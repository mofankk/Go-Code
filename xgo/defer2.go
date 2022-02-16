package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("a")
		fmt.Println("b")
	}()

	fmt.Println("Hello")
}
