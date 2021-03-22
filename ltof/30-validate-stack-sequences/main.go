package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stash := make([]int, 0, len(pushed))
	store := false
	for len(pushed) > 0 || store {
		for len(popped) > 0 && len(stash) > 0 && stash[len(stash)-1] == popped[0] {
			stash = stash[:len(stash)-1]
			popped = popped[1:]
		}
		store = false

		if len(pushed) > 0 {
			stash = append(stash, pushed[0])
			pushed = pushed[1:]
			store = true
		}
	}
	if len(pushed) == 0 && len(popped) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(validateStackSequences1([]int{8, 9, 2, 3, 7, 0, 5, 4, 6, 1}, []int{6, 8, 2, 1, 3, 9, 0, 7, 4, 5}))
}

func validateStackSequences1(pushed []int, popped []int) bool {
	stash := make([]int, 0, len(pushed))
	popIndex := 0
	for _, v := range pushed {
		stash = append(stash, v)
		for len(stash) > 0 && stash[len(stash)-1] == popped[popIndex] {
			stash = stash[0:len(stash)-1]
			popIndex++
		}
	}
	return len(stash) == 0
}
