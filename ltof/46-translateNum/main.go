package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	if len(numStr) == 1 {
		return 1
	}
	first := 1
	second := 1
	if numStr[0:2] < "26" && numStr[0:2] >= "10" {
		second = 2
	}

	for i := 2; i < len(numStr); i++ {
		if numStr[i-1:i+1] < "26" && numStr[i-1:i+1] >= "10" {
			first, second = second, first + second
		} else {
			first, second = second, second
		}
	}
	return second
}

func main() {
	fmt.Println(translateNum(12258))
}
