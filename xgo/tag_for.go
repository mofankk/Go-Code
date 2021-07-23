package main

import "fmt"

func main() {

	fmt.Print("start")
	i := 0

EXITLOOP:
	for i = 0; i < 20; i++ {
		if i == 5 {
			break EXITLOOP
		}

	}
	fmt.Print(i)

}
