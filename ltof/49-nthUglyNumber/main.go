package main

import (
	"fmt"
	"math"
)

// 超时了，计算了无效的数，但结果应该是对的
func nthUglyNumber(n int) int {
	if n <= 5 {
		return n
	}
	ugly := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true}

	i := 5
	for len(ugly) < n {
		i++
		if i%2 == 0 && ugly[i/2] {
			ugly[i] = true
			continue
		}
		if i%3 == 0 && ugly[i/3] {
			ugly[i] = true
			continue
		}
		if i%5 == 0 && ugly[i/5] {
			ugly[i] = true
		}
	}
	return i
}

func main() {
	fmt.Println(nthUglyNumber1(10))
}

func nthUglyNumber1(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i< n; i++ {
		tmp1, tmp2, tmp3 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[i] = int(math.Min(math.Min(float64(tmp1), float64(tmp2)), float64(tmp3)))
		if dp[i] == tmp1 {
			a++
		}
		if dp[i] == tmp2 {
			b++
		}
		if dp[i] == tmp3 {
			c++
		}
	}
	return dp[n-1]
}
