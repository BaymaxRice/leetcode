package main

import "fmt"

func xx(input [][]int) [][]int {
	if len(input) == 0 {
		return [][]int{}
	}
	ret := make([][]int, 0, 8)

	col := 0
	for {
		has := false
		tmp := make([]int, 0, len(input))
		for i := 0; i < len(input); i++ {
			if col < len(input[i]) {
				tmp = append(tmp, input[i][col])
				has = true
			}
		}

		if !has {
			break
		}
		ret = append(ret, tmp)
		col++
	}
	return ret
}

func main() {
	fmt.Println(xxx([][]int{{1, 2}, {1}, {1, 2, 3}}))
}

func printArr1(arr [][]int) [][]int {
	if len(arr) == 0 {
		return [][]int{}
	}
	res := [][]int{}

	i := 0
	j := 0

	item := []int{}
	// for i := 0; i < len(arr); i++ {
	for {
		fmt.Println("outer", i, j)
		if i == len(arr)-1 {
			j++
			i = 0
		}

		// if i > len(arr)-1 {
		//      break
		//      fmt.Println(i, j)
		// }

		for {
			if i == len(arr)-1 {
				break
			}

			if j > len(arr[i])-1 {
				break
			}

			b := []int{}
			copy(b, item)
			item = append(b, arr[i][j])
			res = append(res, item)

			i++

		}

	}
	return res
}

func xxx(input [][]int) [][]int {
	if len(input) == 0 {
		return [][]int{}
	}
	ret := make([][]int, 0, 8)
	dfs(input, &ret)
	return ret
}

func dfs(input [][]int, ret *[][]int) {
	tmp := make([]int, 0, len(input))
	for k, v := range input {
		if len(v) > 0 {
			tmp = append(tmp, v[0])
			input[k] = input[k][1:]
		}
	}
	if len(tmp) == 0 {
		return
	}
	*ret = append(*ret, tmp)
	dfs(input, ret)
}
