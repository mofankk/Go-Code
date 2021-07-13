package main

import "fmt"

func main() {

	fmt.Println(add(3,1))
	
	sub(3,1)
}


func add(a, b int) (c int) {
	c = a + b
	return 
}

func sub(a, b int) {
	c := a - b
	fmt.Println(c)
}


