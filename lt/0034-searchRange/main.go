package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1

	// 二分查找左边界
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	retLeft := left
	right = len(nums) - 1
	// 二分查找右边界
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{retLeft, right}
}

func main() {
	fmt.Println(searchRange([]int{1, 1, 2, 2, 3, 3, 3}, 3))
}
