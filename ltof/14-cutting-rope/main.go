package main

import (
	"fmt"
	"math"
)

func cuttingRope(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}

	num3 := n / 3
	l := n % 3
	if l == 1 {
		return int(math.Pow(3, float64(num3-1))) * 4
	} else if l == 0{
		return int(math.Pow(3, float64(num3)))
	} else {
		return int(math.Pow(3, float64(num3))) * 2
	}
}

func main() {
	fmt.Println(cuttingRope(10))
}
