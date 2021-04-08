package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	max := 0
	exist := make(map[byte]int, len(s))
	left, right := 0, 0
	for right < len(s) {
		value, ok := exist[s[right]]
		if ok && value >= left {
			left = value + 1
		}
		if max < right-left+1 {
			max = right - left + 1
		}
		exist[s[right]] = right
		right++
	}
	if max < right-left {
		max = right - left
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring("bbbbbbb"))  // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // 3
	fmt.Println(lengthOfLongestSubstring("au"))       // 2
}
