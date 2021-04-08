package main

import "fmt"

func reverseWords(s string) string {
	words := make([]string, 0, 5)
	index := 0
	for index < len(s) {
		for index < len(s) && s[index] == ' ' {
			index++
		}

		if index < len(s) {
			tmpIndex := index
			for index < len(s) && s[index] != ' ' {
				index++
			}
			words = append(words, s[tmpIndex:index])
		}
	}
	var ret string
	for k := range words {
		if k != 0 {
			ret += " "
		}
		ret += words[len(words)-1-k]
	}
	return ret
}

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}
