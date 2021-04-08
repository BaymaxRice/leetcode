package main


func majorityElement(nums []int) int {
	index, num := 0, 0
	for k, v := range nums {
		if num == 0 {
			index = k
			num++
		}else if nums[index] == v {
			num++
		} else {
			num--
		}
	}
	return nums[index]
}
