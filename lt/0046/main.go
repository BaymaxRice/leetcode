package main

import "fmt"

func main() {
	input := []int{1,2,3}
	fmt.Printf("%+v", permute(input))
}

func permute(nums []int) [][]int {
	length := 1
	for k := range nums {
		length *= k+1
	}
	ret := make([][]int, 0, length)
	for i:= 0; i < len(nums); i++ {
		if i < len(nums)-1 {
			already := []int{nums[i]}
			input := make([]int, 0, len(nums)-1)
			copy(input, nums[0:i])
			input = append(input, nums[i+1:]...)
			calcRet(ret, already, input)
		} else {
			calcRet(ret, []int{nums[i]}, nums[0:i])
		}
	}
	return ret
}

func calcRet(ret [][]int, already []int, input []int) {
	if len(input) == 1 {
		tmp := make([]int, len(already) + 1)
		tmp = append(already, input[0])
		ret = append(ret, tmp)
		return
	}

	for i:= 0; i < len(input); i++ {
		if i < len(input)-1 {
			calcRet(ret, append(already, input[i]), append(input[0:i], input[i+1:]...))
		} else {
			calcRet(ret, append(already, input[i]), input[0:i])
		}
	}
}
