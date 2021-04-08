package main

import (
	"fmt"
	. "github.com/BaymaxRice/leetcode"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	if (A.Val == B.Val && match(A, B)) ||
		isSubStructure(A.Left, B) ||
		isSubStructure(A.Right, B) {
		return true
	}
	return false
}

func match(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val == B.Val && match(A.Left, B.Left) && match(A.Right, B.Right) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSubStructure(GenerateTree([]int{3, 4, 5, 1, 2}), GenerateTree([]int{3, 4})))
}
