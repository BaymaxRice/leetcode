package main

import "fmt"

func dicesProbability(n int) []float64 {
	dp := make([][]float64, 0, n)
	dp = append(dp, make([]float64, 7))
	for i := 1; i <= 6; i++ {
		dp[0][i] = 1.0 / 6
	}
	for i := 1; i < n; i++ {
		dp = append(dp, make([]float64, i*6+7))
		for j := i + 1; j <= 6*(i+1); j++ {
			for x := 1; x <= 6; x++ {
				if j-x >= 1 && j-x <= 6*i {
					dp[i][j] += 1.0 / 6 * dp[i-1][j-x]
				}
			}
		}
	}
	return dp[n-1][n:]
}

func main() {
	fmt.Println(dicesProbability(2))
}
