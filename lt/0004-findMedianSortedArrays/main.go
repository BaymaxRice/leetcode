package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// 如果有个数组为空，则直接返回结果
	if len(nums1) == 0 {
		return float64(nums2[len(nums2)/2]+nums2[(len(nums2)-1)/2]) / 2
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
	fmt.Println(findMedianSortedArrays1([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	// 如果有个数组为空，则直接返回结果
	if len(nums1) == 0 {
		return float64(nums2[len(nums2)/2]+nums2[(len(nums2)-1)/2]) / 2
	}

	iMin, iMax := 0, len(nums1)
	for {
		mid1 := (iMin + iMax) / 2
		mid2 := (len(nums1)+len(nums2)+1)/2 - mid1
		if mid2 != 0 && mid1 != len(nums1) && nums2[mid2-1] > nums1[mid1] {
			iMin = mid1 + 1
		} else if mid1 != 0 && mid2 != len(nums2) && nums1[mid1-1] > nums2[mid2] {
			iMax = mid1 - 1
		} else { // 要么是有个数组到了边界，要么是找到了中位数
			var maxLeft, minRight int
			if mid1 == 0 {
				maxLeft = nums2[mid2-1]
			} else if mid2 == 0 {
				maxLeft = nums1[mid1-1]
			} else {
				maxLeft = int(math.Max(float64(nums1[mid1-1]), float64(nums2[mid2-1])))
			}

			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(maxLeft)
			}

			if mid1 == len(nums1) {
				minRight = nums2[mid2]
			} else if mid2 == len(nums2) {
				minRight = nums1[mid1]
			} else {
				minRight = int(math.Min(float64(nums2[mid2]), float64(nums1[mid1])))
			}
			return float64(maxLeft+minRight) / 2
		}
	}
}
