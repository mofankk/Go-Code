package main

import "fmt"

func main() {
	a, b, c := 1, 2, 3

	switch {
	case a == 1:
		fmt.Println("a", a)
	case b == 2:
		fmt.Println("a", a)
	case c == 3:
		fmt.Println("a", a)
	default:
		fmt.Println("default")
	}
}

/*
执行结果:
a 1

*/
