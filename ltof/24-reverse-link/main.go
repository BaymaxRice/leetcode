package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, tmp := head.Next, head.Next
	head.Next = nil

	for pre != nil {
		pre = pre.Next
		tmp.Next = head
		head = tmp
		tmp = pre
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
	fmt.Println(reverseList(head))
}
