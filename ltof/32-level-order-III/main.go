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
	queue = append(queue, root)
	right := true
	for len(queue) > 0 {
		tmp := make([]int, 0, len(queue))
		length := len(queue)
		idx := 0
		if !right {
			idx = len(queue) - 1
		}
		// 顺序添加
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[idx].Val)
			if !right {
				idx--
			} else {
				idx++
			}
		}
		queue = queue[length:]
		right = !right
		ret = append(ret, tmp)
	}

	return ret
}

func main() {
	tree := GenerateTree([]int{1, 2, 3, 4, -1, -1, 5})
	fmt.Println(levelOrder(tree))
}
