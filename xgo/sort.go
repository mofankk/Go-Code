package main

import "sort"

func main() {
	nums := []int{1, 5, 4, 2}
	sort.Sort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return 0

}
