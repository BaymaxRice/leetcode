package main

import "fmt"

func main() {
	fmt.Println(replaceSpace("we are family"))
}


func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	numSpace := 0
	for _, v := range s {
		if string(v) == " " {
			numSpace++
		}
	}

	ret := make([]byte, len(s) + numSpace*2)
	idx := len(ret)-1

	for i := len(s)-1; i >= 0; i-- {
		if string(s[i]) != " " {
			ret[idx] = s[i]
			idx--
		} else {
			ret[idx] = '0'
			ret[idx-1] = '2'
			ret[idx-2] = '%'
			idx -= 3
		}
	}

	return string(ret)
}