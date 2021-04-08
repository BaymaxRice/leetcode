package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// 两种方式，一个用map存储子串已出现的字母，或者从left像right递归，判断新字符出现过没有
	// 空间换时间
	left, right, max := 0, 0, 0
	has := make(map[byte]bool, len(s))
	for right < len(s) {
		for right < len(s) && !has[s[right]] {
			has[s[right]] = true
			right++
		}
		if right-left > max {
			max = right - left
		}
		// 加入右边新元素，剔除掉左边旧元素
		if right < len(s) {
			for s[left] != s[right] {
				has[s[left]] = false
				left++
			}
			left++
			right++
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
