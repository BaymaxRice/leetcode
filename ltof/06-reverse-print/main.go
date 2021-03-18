package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	pre, tmp := head.Next, head
	num := 1
	head.Next = nil

	for pre != nil {
		head = pre
		pre = pre.Next
		head.Next = tmp
		tmp = head
		num++
	}

	ret := make([]int, num)
	idx := 0

	for head != nil {
		ret[idx] = head.Val
		head = head.Next
		idx++
	}
	return ret
}

func main() {
	input := []int{1}

	head := &ListNode{
		Val: input[0],
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


	fmt.Println(reversePrint(head))
}
