package main

func maxProfit(prices []int) int {
	var min, max int
	min = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] - min > max {
			max = prices[i] - min
		}

		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
