package main

import (
	"fmt"
	. "github.com/BaymaxRice/leetcode"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	var retHead = &ListNode{}
	retHead.Next = head

	pre := retHead
	var remain, cur *ListNode

	cur = head
	for {
		tmpK := k
		tmpHead := cur
		// 凑齐k个节点
		for tmpK > 0 && tmpHead != nil {
			// 说明找到了要反转链表的最后一个节点了，这时候需要把链表截断，不然反转链表的时候会多反转
			if tmpK == 1 {
				remain = tmpHead.Next
				tmpHead.Next = nil
				pre.Next = reverse(cur)
				pre = cur
				cur = remain
				tmpK--
				break
			} else {
				tmpHead = tmpHead.Next
				tmpK--
			}
		}
		if tmpK != 0 {
			break
		}
	}
	pre.Next = remain

	return retHead.Next
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp, pre := head.Next, head.Next
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
	fmt.Println(List2Arr(reverseKGroup(GenerateList([]int{1, 2, 3, 4, 5}), 2)))
}
