package main

import (
	"fmt"
	"strings"
)

func main() {
	var buffer strings.Builder

	str := "We are happy."
	for _, v := range str {
		if v == ' ' {
			buffer.WriteString("%20")
			continue
		}
		buffer.WriteString(string(v))
	}
	fmt.Printf("buffer: %v\n", buffer.String())
}
