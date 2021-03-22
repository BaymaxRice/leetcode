package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateList(input []int) *ListNode {
	head := &ListNode{
		Val:  input[0],
		Next: nil,
	}
	tmp := head

	for i := 1; i < len(input); i++ {
		tmpL := &ListNode{
			Val: input[i],
		}
		tmp.Next = tmpL
		tmp = tmpL
	}
	return head
}
