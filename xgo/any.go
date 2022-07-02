package main

import "fmt"

func main() {
	var (
		a int     = 1
		b string  = "hello"
		c float64 = 1.11
	)

	f(c)
	f(a)
	f(b)
}

func f(t any) {
	switch t.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("flat64")
	default:
		fmt.Println("none")
	}
}
