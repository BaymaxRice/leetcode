package main

import (
	"fmt"
)

// func minArray(numbers []int) int {
//
// 	if len(numbers) == 1 {
// 		return numbers[0]
// 	}
// 	left, right := 0, len(numbers)-1
// 	mid := left + (right-left)/2
// 	min := numbers[0]
// 	if numbers[mid] > min {
// 		return minArray(numbers[mid+1:])
// 	} else if numbers[mid] < min {
// 		return minArray(numbers[0:mid])
// 	} else {
// 		minLeft := minArray(numbers[0:mid-1])
// 		minRight := minArray(numbers[mid:])
// 		if minLeft < minRight {
// 			return minLeft
// 		} else {
// 			return minRight
// 		}
// 	}
// }

func main() {
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
}


// eg: [10,1,10,10,10]
// 迭代不能解决，当中间和最小值相等时，判断最小值在左还是右
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	left, right := 0, len(numbers)-1

	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] > numbers[right] {
			left = mid+1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}
