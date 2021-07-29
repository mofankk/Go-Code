package main

import "fmt"

func main() {

	type Wo struct {
		A int
	}

	var wo Wo
	wo.A = 1

	he := getIn()
	if he == nil {
		fmt.Println("he is nil")
		return
	}

	h := he.(Wo)

	fmt.Println(h)
}

func getIn() (in interface{}) {
	return
}
