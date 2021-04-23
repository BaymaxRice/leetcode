package main

import "fmt"

func smallestSubsequence(s string) string {
	count := make(map[byte]int, len(s))
	for _, v := range []byte(s) {
		count[v] += 1
	}

	exist := make(map[byte]bool, len(s))
	minQueue := make([]byte, 0, len(s))
	for _, v := range []byte(s) {
		count[v] -= 1
		if exist[v] {
			continue
		}
		for len(minQueue) > 0 && minQueue[len(minQueue)-1] > v && count[minQueue[len(minQueue)-1]] > 0 {
			exist[minQueue[len(minQueue)-1]] = false
			minQueue = minQueue[:len(minQueue)-1]

		}
		minQueue = append(minQueue, v)
		exist[v] = true
	}
	return string(minQueue)
}

func main() {
	fmt.Println(smallestSubsequence("bcabc"))
	fmt.Println(smallestSubsequence("cbacdcbc"))
	fmt.Println(smallestSubsequence("bcbcbcababa"))
}
