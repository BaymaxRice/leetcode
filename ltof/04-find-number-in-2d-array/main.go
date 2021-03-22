package main

import (
	"fmt"
)

func main() {
	fmt.Println(findNumberIn2DArray([][]int{
		// {1, 4, 7, 11, 15},
		// {2, 5, 8, 12, 19},
		// {3, 6, 9, 16, 22},
		// {10, 13, 14, 17, 24},
		// {18, 21, 23, 26, 3},
		{-1, 3},
	}, 3))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	for row, col := 0, len(matrix[0])-1; row < len(matrix) && col >= 0; {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return true
		}
	}
	return false
}

// 用递归来写，超时
/*func findNumberIn2DArray(matrix [][]int, target int) bool {
	not := make([][]bool, len(matrix))
	for k := range matrix {
		tmp := make([]bool, len(matrix[0]))
		not[k] = tmp
	}
	return calc(matrix, not, 0, 0, target)
}

func calc(matrix [][]int, not [][]bool, row, col, target int) bool {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) {
		return false
	}
	if matrix[row][col] > target {
		not[row][col] = true
		return false
	}
	if matrix[row][col] == target {
		return true
	}

	if row < len(matrix)-1 && !not[row+1][col] {
		downRet := calc(matrix, not, row+1, col, target)
		if downRet {
			return true
		}
	}

	if col < len(matrix[0])-1 && !not[row][col] {
		rightRet := calc(matrix, not, row, col+1, target)
		return rightRet
	}
	not[row][col] = true
	return false
}*/
