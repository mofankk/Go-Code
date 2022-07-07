package main

import "fmt"

func main() {

	// 如果 len(s) - 7 < 0，会 panic: runtime error: slice bounds out of range [-2:]
	s := "hello"
	fmt.Println(s[len(s)-7:])

}
