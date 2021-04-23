package normal

func HeapSort(arr []int) []int {
	heapSort(arr, len(arr))
	return arr
}

func heapSort(arr []int, len int) {
	buildMaxHeap(arr, len)

	for i := len - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}

func buildMaxHeap(arr []int, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		heapify(arr, i, len)
	}
}

func heapify(arr []int, idx, len int) {
	left, right, largest := 2*idx+1, 2*idx+2, idx

	if left < len && arr[left] > arr[largest] {
		largest = left
	}
	if right < len && arr[right] > arr[largest] {
		largest = right
	}
	if largest != idx {
		arr[idx], arr[largest] = arr[largest], arr[idx]
		heapify(arr, largest, len)
	}
}
