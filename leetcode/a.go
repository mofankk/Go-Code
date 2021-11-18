package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 5}
	fmt.Println(searchInsert(nums, 4))
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for {
		if low == high {
			fmt.Println(low)
			if nums[low] < target {
				return low + 1
			} else if nums[low] > target {
				if low-1 < 0 {
					return 0
				} else {
					if nums[low-1] < target {
						return low
					} else {
						return low - 1
					}
				}
			} else {
				return low
			}
		}
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}
