package main

import (
	"fmt"
)

func movingCount(m int, n int, k int) int {
	count := 1
	can := make([][]bool, m)
	for i := 0; i < m; i++ {
		can[i] = make([]bool, n)
	}
	can[0][0] = true

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if ((col > 0 && can[row][col-1]) || (row > 0 && can[row-1][col])) &&
				(getValue(row)+getValue(col)) <= k {
				can[row][col] = true
				count++
			}
		}
	}
	return count
}

func getValue(input int) int {
	ret := 0
	for input > 0 {
		ret += input % 10
		input /= 10
	}
	return ret
}

func main() {
	fmt.Println(movingCount(3, 1, 0))
}
