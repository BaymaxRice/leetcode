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
