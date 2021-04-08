package main

import (
	"fmt"
	. "github.com/BaymaxRice/leetcode"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return match(root.Left, root.Right)
}

func match(left *TreeNode, right *TreeNode) bool {
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left == nil && right == nil {
		return true
	}
	if left.Val != right.Val {
		return false
	}
	return match(left.Left, right.Right) && match(left.Right, right.Left)
}

func main() {
	fmt.Println(isSymmetric(GenerateTree([]int{1, 2, 2, 3, 4, 4, 3})))
}
