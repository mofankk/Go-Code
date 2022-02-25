package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	// 数组不能append, append针对slice而言
	// arr = append(arr, 3)
	fmt.Println(arr)
}
