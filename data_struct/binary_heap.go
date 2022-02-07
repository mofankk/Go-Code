package main

// 二叉堆这种数据结构的实现

import "fmt"

func main() {
	fmt.Println("vim-go")
}

// 父节点索引
func parent(k int) int {
	return k / 2
}

// 左子节点索引
func left(k int) int {
	return k * 2
}

// 右子节点索引
func right(k int) int {
	return k * 2 + 1
}

func maxHeapSwim(arr []int, k int) {
	for k >= 1 {
		if arr[parent(k)] < arr[k] {
			arr[parent(k)], arr[k] = arr[k], arr[parent(k)]
			k = parent(k)
			continue
		}
		break
	}
}

func maxHeapSink(arr []int, k int) {
	length := len(arr)
	for k < length {
		

}



