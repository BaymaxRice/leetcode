package main

import (
	"fmt"
	"math"
	"sort"
)

type sortArr [][]int

func (a sortArr) Len() int {
	return len(a)
}

func (a sortArr) Less(i, j int) bool {
	return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] < a[j][1])
}

func (a sortArr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	ret := make([][]int, 0, len(intervals))
	sort.Sort(sortArr(intervals))
	ret = append(ret, intervals[0])

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > ret[len(ret)-1][1] {
			ret = append(ret, intervals[i])
		} else {
			ret[len(ret)-1][1] = int(math.Max(float64(ret[len(ret)-1][1]), float64(intervals[i][1])))
		}
	}
	return ret
}

func main() {
	fmt.Println(merge([][]int{{1, 4}, {2, 3}}))
}
