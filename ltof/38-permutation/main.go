package main

import "fmt"

func permutation(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}
	length := 1
	for i:= 1; i <= len(s); i++ {
		length *= i
	}
	ret := make([]string, 0, length)
	used := make(map[int]bool, len(s))
	already := make([]byte, 0, len(s))
	dfs(s, &ret, used, already)
	return ret
}

func dfs(s string, ret *[]string, used map[int]bool, already []byte) {
	if len(already) >= len(s) {
		tmpRet := make([]byte, len(s))
		copy(tmpRet, already)
		*ret = append(*ret, string(tmpRet))
		return
	}

	tmpUsed := make(map[byte]bool, len(s) - len(already))
	for k, v := range []byte(s) {
		if used[k] {
			continue
		}
		if tmpUsed[v] {
			continue
		}
		tmpUsed[v] = true
		used[k] = true
		already = append(already, v)
		dfs(s, ret, used, already)
		used[k] = false
		already = already[:len(already)-1]
	}
}

func main() {
	fmt.Println(permutation("ab"))
}
