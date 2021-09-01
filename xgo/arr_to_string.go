package main

import (
	"fmt"
	"strings"
)

func main() {
	value := [][]int64{{1, 2}, {3, 4}, {5, 6}}

	str := fmt.Sprint(value)
	str = strings.Replace(str, " ", ",", -1)

	str = strings.ReplaceAll(str, "[", "(")
	str = strings.ReplaceAll(str, "]", ")")
	fmt.Println(str)

	//fmt.Println(value)

}
