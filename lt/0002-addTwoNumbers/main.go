package main

import . "baymax/leetcode"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := l1
	pre := 0
	for l1 != nil || l2 != nil {
		// 保证l1一定比l2长
		if l1 != nil && l2 != nil && l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			l2.Next = nil
		}
		tmpAdd := pre
		tmpAdd += l1.Val
		if l2 != nil {
			tmpAdd += l2.Val
		}
		pre = tmpAdd / 10
		l1.Val = tmpAdd % 10
		if l1.Next == nil && pre != 0 {
			l1.Next = &ListNode{
				Val: pre,
			}
			break
		} else {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head
}
