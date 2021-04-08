package main

import "fmt"

var letter = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	ret := make([]string, 0, 8)
	cur := make([]byte, len(digits))
	dfs([]byte(digits), &ret, &cur, 0)
	return ret
}

func dfs(digit []byte, ret *[]string, cur *[]byte, k int) {
	if k >= len(digit) {
		tmp := make([]byte, len(digit))
		copy(tmp, *cur)
		*ret = append(*ret, string(tmp))
		return
	}
	for _, v := range letter[digit[k]] {
		(*cur)[k] = v
		dfs(digit, ret, cur, k+1)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations(""))
}
