package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func one() {
	once.Do(func() {
		fmt.Println("hello")
	})

	fmt.Println("world")
}

func main() {
	for i := 0; i < 15; i++ {
		one()

	}

}
