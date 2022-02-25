package main

import "fmt"

func main() {
	arr1 := new([]int)
	arr2 := make([]int, 5)
	*arr1 = append(*arr1, 1, 2, 3)
	arr2 = append(arr2, 4, 5, 6)
	fmt.Printf("arr1: %v, len: %d, cap: %d \n", arr1, len(*arr1), cap(*arr1))
	fmt.Printf("arr2: %v, len: %d, cap: %d \n", arr2, len(arr2), cap(arr2))
}

/*
arr1: &[1 2 3], len: 3, cap: 3
arr2: [0 0 0 0 0 4 5 6], len: 8, cap: 10
*/
