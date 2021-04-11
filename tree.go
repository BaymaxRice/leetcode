package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTree(input []int) *TreeNode {
	head := &TreeNode{
		Val:   input[0],
		Left:  nil,
		Right: nil,
	}

	tmp := make([]*TreeNode, len(input))
	tmp[0] = head

	for i := 1; i < len(input); i++ {
		if input[i] == -1 {
			continue
		}
		tmpTree := &TreeNode{
			Val:   input[i],
			Left:  nil,
			Right: nil,
		}
		if i&1 == 0 {
			tmp[(i-1)/2].Right = tmpTree
		} else {
			tmp[(i-1)/2].Left = tmpTree
		}
		tmp[i] = tmpTree
	}
	return head
}

func Tree2Arr(root *TreeNode) []int {
	retArr := make([]int, 0, 8)
	queue := make([]*TreeNode, 0, 8)
	queue = append(queue, root)
	count := 1
	for {
		loop := 0
		has := false
		for loop < count {
			if queue[loop] != nil {
				retArr = append(retArr, queue[loop].Val)
				if queue[loop].Left != nil {
					queue = append(queue, queue[loop].Left)
					has = true
				} else {
					queue = append(queue, nil)
				}
				if queue[loop].Right != nil {
					queue = append(queue, queue[loop].Right)
					has = true
				} else {
					queue = append(queue, nil)
				}
			} else {
				retArr = append(retArr, -1)
				queue = append(queue, nil, nil)
			}

			loop++
		}
		queue = queue[count:]
		count *= 2
		if !has {
			break
		}
	}
	return retArr
}
