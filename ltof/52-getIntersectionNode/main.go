package main

import (
	. "baymax/leetcode"
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lengthA := lengthList(headA)
	lengthB := lengthList(headB)

	tmpA, tmpB := headA, headB
	if lengthA > lengthB {
		lengthA, lengthB = lengthB, lengthA
		tmpA, tmpB = headB, headA
	}

	// 先将指针进度调到一致
	for i := 0; i < int(math.Abs(float64(lengthA)-float64(lengthB))); i++ {
		tmpB = tmpB.Next
	}

	for tmpA != tmpB {
		tmpA = tmpA.Next
		tmpB = tmpB.Next
	}
	return tmpA
}

func lengthList(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return length
}
