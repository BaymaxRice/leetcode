package main

type MaxQueue struct {
	data    []int
	maxData []int
}

func Constructor() MaxQueue {
	m := MaxQueue{
		data:    make([]int, 0, 10),
		maxData: make([]int, 0, 10),
	}
	return m
}

func (this *MaxQueue) Max_value() int {
	if len(this.maxData) > 0 {
		return this.maxData[0]
	}
	return -1
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
	for len(this.maxData) > 0 && this.maxData[len(this.maxData)-1] < value {
		this.maxData = this.maxData[:len(this.maxData)-1]
	}
	this.maxData = append(this.maxData, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.data) > 0 {
		if this.maxData[0] == this.data[0] {
			this.maxData = this.maxData[1:]
		}
		tmp := this.data[0]
		this.data = this.data[1:]
		return tmp
	}
	return -1
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
