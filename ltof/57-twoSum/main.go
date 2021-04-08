package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else {
			break
		}
	}
	return []int{nums[left], nums[right]}
}

func main() {

	fmt.Println(twoSum([]int{2, 7, 10, 13, 15}, 9))
}
