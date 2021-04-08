package main

import (
	"fmt"
)

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if k >= len(arr) {
		return arr
	}

	split := partition(arr, 0, len(arr)-1)

	if split+1 == k {
		return arr[0:k]
	}
	if split+1 > k {
		return getLeastNumbers(arr[0:split], k)
	}
	return append(arr[0:split+1], getLeastNumbers(arr[split+1:], k-split-1)...)
}

func partition(arr []int, left, right int) int {
	povit, index := left, left+1
	for i := index; i <= right; i++ {
		if arr[i] < arr[povit] {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[povit], arr[index-1] = arr[index-1], arr[povit]
	return index - 1
}

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}
