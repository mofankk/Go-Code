package main

import (
	"fmt"
	"math/bits"
)

func main() {
	quo, rem := bits.Div64(10, 4, 20)
	fmt.Println(quo, rem)
}
