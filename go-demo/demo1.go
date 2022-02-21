package main

import (
	"fmt"
	"sync"
)

//1. 交替打印数字和字母
func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()

	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		i := 'A'
		for {
			select {
			case <-letter:
				if i > 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				if i > 'Z' {
					wait.Done()
					return
				}
				fmt.Print(string(i))
				i++
				number <- true
			}
		}
	}(&wait)

	number <- true
	wait.Wait()
	fmt.Println()
}
