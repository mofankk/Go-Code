package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello, 世界"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(len([]rune(str)))
}
