package main

import "fmt"

func maxArea(height []int) int {
	ret := 0
	len := len(height)

	for k, v := range height {
		tmpMax := 0
		for j := len-1; j > k; j-- {
			if v > height[j] && height[j] * (j-k) > tmpMax {
				tmpMax = height[j] * (j-k)
			}
			if v <= height[j] {
				if v * (j - k) > tmpMax {
					tmpMax = v * (j-k)
				}
				break
			}
		}
		if tmpMax > ret {
			ret = tmpMax
		}
	}
	return ret
}

func maxArea1(height []int) int {
	ret := 0
	pre := 0
	fro := len(height)-1
	for pre < fro {
		tmpArea := 0
		if height[pre] < height[fro] {
			tmpArea = (fro - pre) * height[pre]
			pre++
		} else {
			tmpArea = (fro - pre) * height[fro]
			fro--
		}
		if tmpArea > ret {
			ret = tmpArea
		}
	}
	return ret
}

func main() {
	fmt.Println(maxArea1([]int{1,1}))
	fmt.Println(maxArea1([]int{1,8,6,2,5,4,8,3,7}))
}