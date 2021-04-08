package main

import (
	"fmt"
	"math"
)

func findNthDigit(n int) int {
	if n <= 0 {
		return 0
	}

	digit := 1
	base := 1
	for n > 9*base*digit {
		n = n - 9*base*digit
		digit++
		base *= 10
	}

	num := (n - 1) / digit
	num += base
	index := (n - 1) % digit

	return num / int(math.Pow(10, float64(digit-index-1))) % 10
}

func main() {
	fmt.Println(findNthDigit(25))
}
