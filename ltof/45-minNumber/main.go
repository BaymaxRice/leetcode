package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type sortString []string

func (s sortString) Len() int {
	return len(s)
}

func (s sortString) Less(i, j int) bool {
	return s[i] + s[j] < s[j] + s[i]
}

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func minNumber(nums []int) string {
	tmp := make([]string, len(nums))
	for k, v := range nums {
		tmp[k] = strconv.Itoa(v)
	}
	sortValue := sortString(tmp)
	sort.Sort(sortValue)
	return strings.Join(sortValue, "")
}

func main() {
	// fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
	// fmt.Println(minNumber([]int{0, 10, 0}))
	// fmt.Println(minNumber([]int{20, 1}))
	fmt.Println(minNumber([]int{12, 121}))
}
