package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	ret := make([][]int, 0, target/3)
	left, right := 1, 2
	stop := (target + 1) / 2
	target = - target + 3
	for left < right && right <= stop {
		if target < 0 {
			right++
			target += right
		} else if target > 0 {
			target -= left
			left++
		} else {
			tmp := make([]int, 0, right-left+1)
			for i := left; i <= right; i++ {
				tmp = append(tmp, i)
			}
			ret = append(ret, tmp)
			right++
			target += right
		}
	}
	return ret
}

func main() {
	fmt.Println(findContinuousSequence(15))
}
