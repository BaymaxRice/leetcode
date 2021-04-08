package main

import (
	"fmt"
	"math"
)

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for k := range s {
		tmp := make([]int, len(s))
		dp[k] = tmp
		dp[k][k] = 1
		for i := k - 1; i >= 0; i-- {
			if s[i] != s[k] {
				dp[i][k] = int(math.Max(float64(dp[i][k-1]), float64(dp[i+1][k])))
			} else {
				if i+1 < len(s) && k-1 > 0 {
					dp[i][k] = dp[i+1][k-1] + 2
				} else {
					dp[i][k] = 2
				}
			}
		}
	}
	return dp[0][len(s)-1]
}

func main() {
	fmt.Println(longestPalindromeSubseq("bbbb"))
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}
