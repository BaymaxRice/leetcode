package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	ret := &TreeNode{Val: preorder[0]}
	idx := -1
	for k, v := range inorder {
		if v == preorder[0] {
			idx = k
			break
		}
	}
	ret.Left = buildTree(preorder[1:1+idx], inorder[0:idx])
	if idx < len(inorder)-1 {
		ret.Right = buildTree(preorder[1+idx:], inorder[idx+1:])
	}
	return ret
}

func main() {
	ret := buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	fmt.Println(ret)
}
