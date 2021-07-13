package main

import (
	"fmt"
	"strconv"
)

func main() {
	f := 5.21
	fmt.Println(strconv.FormatFloat(f, 'g', -1, 32))
}
