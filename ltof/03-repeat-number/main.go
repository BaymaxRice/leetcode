package main

import "fmt"

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber(nums []int) int {
	ret := -1
	handle := 0
	for {
		if nums[handle] != handle {
			if nums[nums[handle]] == nums[handle] {
				ret = nums[handle]
				break
			}

			tmpHandle := nums[nums[handle]]
			nums[nums[handle]] = nums[handle]
			handle = tmpHandle
		} else {
			handle++
		}
	}
	return ret
}
