package main

import "fmt"

func main() {

	type Content struct {
		Name string
		Age  int
	}

	c := &Content{
		Name: "hello",
		Age:  12,
	}

	d := c

	d.Name = "Ping"

	fmt.Println(c.Name)

}
