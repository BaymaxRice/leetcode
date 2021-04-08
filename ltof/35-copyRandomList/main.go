package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	copy := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	oriHead := head
	copyHead := copy
	visited[head] = copy
	for head.Next != nil {
		tmpCopy := &Node{
			Val:    head.Next.Val,
			Next:   nil,
			Random: nil,
		}
		visited[head.Next] = tmpCopy
		copy.Next = tmpCopy
		copy = copy.Next
		head = head.Next
	}

	copy = copyHead
	for copy != nil {
		if oriHead.Random != nil {
			copy.Random = visited[oriHead.Random]
		}
		copy = copy.Next
		oriHead = oriHead.Next
	}

	return copyHead
}