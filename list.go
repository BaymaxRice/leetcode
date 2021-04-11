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

func List2Arr(head *ListNode) []int {
	retArr := make([]int, 0, 8)
	tmpHead := head
	for tmpHead != nil {
		retArr = append(retArr, tmpHead.Val)
		tmpHead = tmpHead.Next
	}
	return retArr
}
