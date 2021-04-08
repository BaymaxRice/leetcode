package main

import (
	. "baymax/leetcode"
	"fmt"
)

func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	res := -1
	dfs(root, &res, &k)
	return res
}

func dfs(root *TreeNode, res, k *int) {
	if root == nil {
		return
	}
	dfs(root.Right, res, k)
	if *k == 1 {
		*res = root.Val
	}
	*k--
	dfs(root.Left, res, k)
}

func main() {
	fmt.Println(kthLargest(GenerateTree([]int{5, 3, 6, 2, 4, -1, -1, 1}), 3))
}
