package main

import "fmt"

func main() {

}

func rotate(nums []int, k int) {

	if k == 0 || len(nums) == 1 {
		return
	}
	arrLen := len(nums)
	k = k % arrLen
	// 翻转整个字符串
	for i := 0; i < arrLen/2; i++ {
		t := nums[i]
		nums[i] = nums[arrLen-i-1]
		nums[arrLen-i-1] = t
	}
	fmt.Println(nums)

	// 翻转前半部分
	for i := 0; i < k/2; i++ {
		t := nums[i]
		nums[i] = nums[k-i-1]
		nums[k-i-1] = t
	}

	// 翻转后半部分
	fmt.Println(k)
	for i := k; i < (arrLen+k)/2; i++ {
		t := nums[i]
		nums[i] = nums[arrLen-i+k-1]
		nums[arrLen-i+k-1] = t
	}
}
