package main

import "fmt"

func main() {
	arr := [2]int{0, 1}
	q(arr)
	fmt.Println(arr)

	x(&arr)
	fmt.Println(arr)

	t(arr[:])
	fmt.Println(arr)

}

// the arrary is converted to slice, passed using the default reference the slice
func t(arr []int) {
	arr[1] *= 10
}

// value
func q(arr [2]int) {
	arr[1] *= 20
}

// use a pointer
func x(arr *[2]int) {
	arr[1] *= 30
}
