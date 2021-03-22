package main

import "fmt"
import . "baymax/leetcode"


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val > l2.Val {
		head = l2
		l2 = l2.Next
	} else {
		head = l1
		l1 = l1.Next
	}
	tmp := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}

	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return head
}

func main() {
	input := []int{1, 2, 3, 4, 5}

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

	input2 := []int{1, 2, 3, 4, 5}

	head2 := &ListNode{
		Val:  input2[0],
		Next: nil,
	}
	tmp2 := head2

	for i := 1; i < len(input2); i++ {
		tmpL := &ListNode{
			Val: input2[i],
		}
		tmp2.Next = tmpL
		tmp2 = tmpL
	}

	fmt.Println(mergeTwoLists(head, head2))
}

