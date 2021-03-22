package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	ret := 1.0
	for n > 0 {
		if n % 2 == 0 {
			x *= x
			n /= 2
		} else {
			ret *= x
			n -= 1
		}
	}
	return ret
}

func main() {
	fmt.Println(myPow(2, 10))
}
