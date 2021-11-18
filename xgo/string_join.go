package main

import (
	"fmt"
	"strings"
)

func main() {

	strs := []string{"hello", "world", "cui"}

	t1 := strings.Join(strs, ",")

	var strs2 []string

	t2 := strings.Join(strs2, ",")

	fmt.Println(t1)
	fmt.Println(t2)

}
