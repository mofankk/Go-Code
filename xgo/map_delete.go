package main

import "fmt"

func main() {

	m := make(map[string]string)
	m["1"] = "1"
	m["2"] = "2"
	m["3"] = "3"
	m["4"] = "4"
	m["5"] = "5"

	for k, _ := range m {

		if k == "2" || k == "5" {
			delete(m, k)
		}
	}

	for k, v := range m {
		fmt.Println(k, ": ", v)
	}
}
