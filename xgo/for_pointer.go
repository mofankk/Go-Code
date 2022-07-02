package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2
		c = 3
	)

	var arr = []*int{&a, &b, &c}
	//res := make([]*int, len(arr))
	var res []*int
	var v *int
	for _, v = range arr {
		//res[i] = v
		res = append(res, v)
	}

	for i := range res {
		fmt.Println(res[i])
	}

}
