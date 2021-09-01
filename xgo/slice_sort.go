package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"app_id", "method", "charset", "sign_type", "timestamp", "biz_content"}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
}
