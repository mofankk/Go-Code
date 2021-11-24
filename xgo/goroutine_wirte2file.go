package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	path := []string{
		"H1",
		"H2",
		"H3",
		"H4",
		"H5",
		"H6",
	}
	af, _ := os.Create("aa.log")
	af.Close()

	for i := range path {
		fmt.Println(path[i])
		go func(p string) {
			f, err := os.OpenFile("aa.log", os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Printf("%v", err)
			}
			defer f.Close()
			f.WriteString(time.Now().Format("2006-01-02 15:04:05 url: ") + p + "\n")

		}(path[i])
		time.Sleep(1 * time.Second)
	}
}
