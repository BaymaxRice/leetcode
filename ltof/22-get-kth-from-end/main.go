package main

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

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 1 {
		fast = fast.Next
		k--
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
