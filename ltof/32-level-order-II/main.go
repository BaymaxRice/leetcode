package main

import (
	"fmt"
	. "github.com/BaymaxRice/leetcode"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode, 0, 10)
	ret := make([][]int, 0, 10)
	queue = append(queue, root, nil)
	for len(queue) > 0 {
		tmp := make([]int, 0, 5)
		for queue[0] != nil {
			tmp = append(tmp, queue[0].Val)
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
		// 去除本层nil，表示本层打印完了
		queue = queue[1:]
		ret = append(ret, tmp)
		if len(queue) > 0 {
			queue = append(queue, nil)
		}
	}

	return ret
}

func main() {
	tree := GenerateTree([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(levelOrder(tree))
}
