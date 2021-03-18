package main

import "fmt"

func exchange(nums []int) []int {
	// 双向移动指针
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left] & 1 != 0 {
			left++
			continue
		}
		if nums[right] & 1 == 0 {
			right--
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{2,16,3,5,13,1,16,1,12,18,11,8,11,11,5,1}))
}