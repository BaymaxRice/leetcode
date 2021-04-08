package main

import "fmt"

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if row == 0 && col == 0 {
				continue
			}
			if row == 0 {
				grid[row][col] += grid[row][col-1]
				continue
			}
			if col == 0 {
				grid[row][col] += grid[row-1][col]
				continue
			}
			if grid[row-1][col]+grid[row][col] > grid[row][col-1]+grid[row][col] {
				grid[row][col] += grid[row-1][col]
			} else {
				grid[row][col] += grid[row][col-1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	fmt.Println(maxValue([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
