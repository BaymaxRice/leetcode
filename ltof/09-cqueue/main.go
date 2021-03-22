package main

type CQueue struct {
	data []int
}

// 题目要求的是双栈实现队列，所以一个栈入数据，一个栈出数据，出数据的栈没数据之后，就把入栈的数据全部导过来，就ok了


func Constructor() CQueue {
	return CQueue{data:make([]int, 0, 10)}
}


func (this *CQueue) AppendTail(value int)  {
	this.data = append(this.data, value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.data) == 0 {
		return -1
	}
	ret := this.data[0]
	this.data = this.data[1:]
	return ret
}
