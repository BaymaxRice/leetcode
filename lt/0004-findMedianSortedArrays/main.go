package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// 如果有个数组为空，则直接返回结果
	if len(nums1) == 0 {
		return float64(nums2[len(nums2)/2]+nums2[(len(nums2)+1)/2]) / 2
	}

	return float64(kthNumber(nums1, nums2, (len(nums1)+len(nums2))/2)+kthNumber(nums1, nums2, (len(nums1)+len(nums2)+1)/2)) / 2
}

func kthNumber(arr1, arr2 []int, k int) int {
	idx1, idx2 := 0, 0

	for {
		if idx1 == len(arr1) {
			return arr2[idx2+k-1]
		}
		if idx2 == len(arr2) {
			return arr1[idx2+k-1]
		}

		if k == 1 {
			return int(math.Min(float64(arr1[idx1]), float64(arr2[idx2])))
		}

		half := k / 2
		newIdx1 := int(math.Min(float64(idx1+half), float64(len(arr1)))) - 1
		newIdx2 := int(math.Min(float64(idx2+half), float64(len(arr2)))) - 1
		if arr1[newIdx1] > arr2[newIdx2] {
			k -= newIdx2 - idx2 + 1
			idx2 = newIdx2 + 1
		} else {
			k -= newIdx1 - idx1 + 1
			idx1 = newIdx1 + 1
		}
	}
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
