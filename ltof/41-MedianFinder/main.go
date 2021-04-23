package main

import "fmt"

// 维护一个最大堆和一个最小堆，最小堆的最小值大于等于最大堆的最大值，并且保证两个堆中的数据量相差在1以内，取中位数的时候直接读取两个值就可以了
type MedianFinder struct {
	minHeap []int
	maxHeap []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: make([]int, 0, 8),
		maxHeap: make([]int, 0, 8),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if len(this.minHeap) < len(this.maxHeap) { // 要往最小堆中添加数据
		if num < this.maxHeap[0] { // 如果小于最大堆的最大值，则和最大堆的最大值交换
			num, this.maxHeap[0] = this.maxHeap[0], num
			shiftDown(this.maxHeap, 0, true)

		}
		this.minHeap = append(this.minHeap, num)
		shiftUp(this.minHeap, len(this.minHeap)-1, false)
	} else {
		// 要往最大堆中添加数据
		if len(this.minHeap) > 0 && num > this.minHeap[0] {
			num, this.minHeap[0] = this.minHeap[0], num
			shiftDown(this.minHeap, 0, false)
		}
		this.maxHeap = append(this.maxHeap, num)
		shiftUp(this.maxHeap, len(this.maxHeap)-1, true)
	}
}

func shiftDown(arr []int, idx int, isMax bool) {
	left, right, tmp := 2*idx+1, 2*idx+2, idx
	if isMax {
		if left < len(arr) && arr[left] > arr[tmp] {
			tmp = left
		}
		if right < len(arr) && arr[right] > arr[tmp] {
			tmp = right
		}
	} else {
		if left < len(arr) && arr[left] < arr[tmp] {
			tmp = left
		}
		if right < len(arr) && arr[right] < arr[tmp] {
			tmp = right
		}
	}
	if tmp != idx {
		arr[tmp], arr[idx] = arr[idx], arr[tmp]
		shiftDown(arr, tmp, isMax)
	}
}

func shiftUp(arr []int, idx int, isMax bool) {
	if idx <= 0 {
		return
	}
	parent := (idx - 1) / 2
	if (isMax && arr[parent] < arr[idx]) || (!isMax && arr[parent] > arr[idx]) {
		arr[parent], arr[idx] = arr[idx], arr[parent]
		shiftUp(arr, parent, isMax)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (len(this.minHeap)+len(this.maxHeap))%2 == 1 {
		if len(this.minHeap) > len(this.maxHeap) {
			return float64(this.minHeap[0])
		} else {
			return float64(this.maxHeap[0])
		}
	} else {
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

func main() {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Println(obj.FindMedian())
	obj.AddNum(3)
	fmt.Println(obj.FindMedian())
}
