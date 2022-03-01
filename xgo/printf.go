package main

import "fmt"

func main() {
	i := -5
	j := +5
	z := 5

	fmt.Printf("- : %-d, %-d, %-d \n", i, j, z)
	fmt.Printf("+ : %+d, %+d, %+d \n", i, j, z)
	fmt.Printf("++ : %++d, %++d, %++d \n", i, j, z)
	fmt.Printf("-- : %--d, %--d, %--d \n", i, j, z)
}
