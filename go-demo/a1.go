package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	defer func() {
		fmt.Println("first defer")
		printCode(printCode(1))
	}()

	defer func() {
		fmt.Println("second defer")
		printCode(printCode(2))
	}()
}

func printCode(code int) int {
	fmt.Println("the code is ", code)
	return code + 10
}
