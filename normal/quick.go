package normal


func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr))
}

// 递归
func quickSort(arr []int, left, right int) []int {
	if left < right {
		split := partition(arr, left, right)
		quickSort(arr, left, split-1)
		quickSort(arr, split+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	povit, index := left, left+1
	for i := index; i <= right; i++ {
		if arr[povit] > arr[i] {
			arr[index], arr[i] = arr[i], arr[index]
			index++
		}
	}
	arr[povit], arr[index-1] = arr[index-1], arr[povit]
	return index - 1
}