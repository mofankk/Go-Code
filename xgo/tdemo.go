package main

import "fmt"

func main() {
	ints := make([]int, 1)
	ints = append(ints, 2)
	fmt.Println(ints[0])
}