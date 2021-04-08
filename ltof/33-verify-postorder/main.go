package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}
	// 后续遍历，根据数组最后一个数，能将数组二分为小于自身，和大于自身两部分
	idx := 0
	for idx < (len(postorder)-1) && postorder[idx] < postorder[len(postorder)-1] {
		idx++
	}
	for i := idx; i < len(postorder)-1; i++ {
		if postorder[i] < postorder[len(postorder)-1] {
			return false
		}
	}

	// 如果这一层二分没问题，则递归二分当中的是否是合法后续遍历
	return verifyPostorder(postorder[0:idx]) && verifyPostorder(postorder[idx:len(postorder)-1])
}

func main() {
	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
}
