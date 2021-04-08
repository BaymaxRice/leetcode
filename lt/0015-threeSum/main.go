package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0, 8)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		left, right := i+1, len(nums)-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right] > -nums[i] {
				right--
			} else if nums[left]+nums[right] < -nums[i] {
				left++
			} else {
				tmp := []int{nums[i], nums[left], nums[right]}
				ret = append(ret, tmp)
				left++
				right--
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))
	fmt.Println(threeSum([]int{0}))
}
