package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return []int{}
	}
	// 存放结果
	ret := make([]int, len(nums)-k+1)
	// 存放最大值
	queue := make([]int, 0, k)
	queue = append(queue, nums[0])
	for i := 1; i < k; i++ {
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	ret[0] = queue[0]
	for i := k; i < len(nums); i++ {
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		ret[i-k+1] = queue[0]
	}
	return ret
}

func main() {
	// fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
	// fmt.Println(maxSlidingWindow([]int{7, 2, 4}, 2))
	fmt.Println(maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
