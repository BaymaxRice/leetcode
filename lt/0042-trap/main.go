package main

import "fmt"

func trap(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	max := height[0]
	maxIndex := 0
	for k, v := range height {
		if v > max {
			maxIndex, max = k, v
		}
	}
	tmpMax := height[0]
	ret := 0
	for i := 1; i < maxIndex; i++ {
		if height[i] < tmpMax {
			ret += tmpMax - height[i]
		} else {
			tmpMax = height[i]
		}
	}
	tmpMax = height[len(height)-1]
	for i := len(height) - 1; i > maxIndex; i-- {
		if height[i] < tmpMax {
			ret += tmpMax - height[i]
		} else {
			tmpMax = height[i]
		}
	}
	return ret
}

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}
