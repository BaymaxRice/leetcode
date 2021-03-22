package main

import "fmt"

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

func spiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])
	startRow, startCol := 0, 0

	ret := make([]int, 0, len(matrix)*len(matrix[0]))

	for {
		if col > 0 {
			for i := 0; i < col; i++ {
				ret = append(ret, matrix[startRow][startCol])
				startCol++
			}
			startCol--
			startRow++
			row--
		} else {
			break
		}

		if row > 0 {
			for i := 0; i < row; i++ {
				ret = append(ret, matrix[startRow][startCol])
				startRow++
			}
			startRow--
			startCol--
			col--
		} else {
			break
		}

		if col > 0 {
			for i := 0; i < col; i++ {
				ret = append(ret, matrix[startRow][startCol])
				startCol--
			}
			startCol++
			startRow--
			row--
		} else {
			break
		}

		if row > 0 {
			for i := 0; i < row; i++ {
				ret = append(ret, matrix[startRow][startCol])
				startRow--
			}
			startRow++
			startCol++
			col--
		} else {
			break
		}
	}
	return ret
}
