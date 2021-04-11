package normal

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	return merge(MergeSort(arr[0:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

// 开辟空间
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	leftIdx, rightIdx := 0, 0
	for leftIdx < len(left) || rightIdx < len(right) {
		if leftIdx < len(left) && rightIdx < len(right) {
			if left[leftIdx] < right[rightIdx] {
				result = append(result, left[leftIdx])
				leftIdx++
			} else {
				result = append(result, right[rightIdx])
				rightIdx++
			}
		} else {
			if leftIdx < len(left) {
				result = append(result, left[leftIdx])
				leftIdx++
			}
			if rightIdx < len(right) {
				result = append(result, right[rightIdx])
				rightIdx++
			}
		}
	}
	return result
}
