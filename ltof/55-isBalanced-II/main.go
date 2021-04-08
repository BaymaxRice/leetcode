package main

import (
	"fmt"
	. "github.com/BaymaxRice/leetcode"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	_, isBalanced := dfs(root)
	return isBalanced
}

func dfs(root *TreeNode) (depth int, isBalanced bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, leftIsBalanced := dfs(root.Left)
	if !leftIsBalanced {
		return 0, false
	}
	rightDepth, rightIsBalanced := dfs(root.Right)
	if !rightIsBalanced {
		return 0, false
	}
	return 1 + int(math.Max(float64(leftDepth), float64(rightDepth))), int(math.Abs(float64(leftDepth)-float64(rightDepth))) <= 1
}

func main() {
	tree := GenerateTree([]int{1, 2, 2, 3, 3, -1, -1, 4, 4})
	fmt.Println(isBalanced(tree))
}
