package main

import (
	"fmt"
)

/**
此题使用了used数组来表示已经用的格子，实际上为了减少空间，还可以破坏已经用的元素，比如将已使用的元素标记为'\'
这样下层递归的时候，也就不可能再使用这个元素了
 */
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	used := make([][]bool, len(board))
	for k := range board {
		used[k] = make([]bool, len(board[k]))
	}

	// 先找到第一个字母
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if board[row][col] == word[0] {
				used[row][col] = true
				has := dfs(board, used, row, col, word[1:])
				if has {
					return true
				}
				used[row][col] = false
			}
		}
	}
	return false
}

func dfs(board [][]byte, used [][]bool, row, col int, word string) bool {
	if len(word) == 0 {
		return true
	}

	if row > 0 && board[row-1][col] == word[0] && !used[row-1][col] {
		used[row-1][col] = true
		has := dfs(board, used, row-1, col, word[1:])
		if has {
			return true
		}
		used[row-1][col] = false
	}

	if row+1 < len(board) && board[row+1][col] == word[0] && !used[row+1][col] {
		used[row+1][col] = true
		has := dfs(board, used, row+1, col, word[1:])
		if has {
			return true
		}
		used[row+1][col] = false
	}

	if col > 0 && board[row][col-1] == word[0] && !used[row][col-1] {
		used[row][col-1] = true
		has := dfs(board, used, row, col-1, word[1:])
		if has {
			return true
		}
		used[row][col-1] = false
	}

	if col+1 < len(board[0]) && board[row][col+1] == word[0] && !used[row][col+1] {
		used[row][col+1] = true
		has := dfs(board, used, row, col+1, word[1:])
		if has {
			return true
		}
		used[row][col+1] = false
	}


	return false
}

func main() {
	// input := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	// path := "ABCCED"

	input := [][]byte{{'a', 'b'}, {'c', 'd'}}
	path := "abcd"
	fmt.Println(exist(input, path))
}
