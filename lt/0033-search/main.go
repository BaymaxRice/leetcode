package main

import "fmt"

func search(nums []int, target int) int {
	return calc(nums, target, 0, len(nums)-1)
}

func calc(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] < nums[mid] && target >= nums[left] && target < nums[mid] {
			right = mid - 1
		}
		if nums[mid] < nums[right] && target > nums[mid] && target <= nums[right] {
			left = mid + 1
		}
		leftFind := calc(nums, target, left, mid-1)
		if leftFind != -1 {
			return leftFind
		} else {
			return calc(nums, target, mid+1, right)
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
}
