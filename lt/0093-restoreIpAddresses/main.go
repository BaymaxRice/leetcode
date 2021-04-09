package main

import (
	"fmt"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	ip := make([]string, 0, 4)
	ret := make([]string, 0, 8)
	dfs(s, 0, &ip, &ret)

	return ret
}

func dfs(s string, idx int, ip, ret *[]string) {
	// 剩余的数字组成的ip一定不合法
	if (len(*ip) == 4 && idx < len(s)-1) || len(s)-idx < 4-len(*ip) || len(s)-idx > (4-len(*ip))*3 {
		return
	}

	// 表示这个ip已经是合法的了
	if idx >= len(s)-1 && len(*ip) == 4 {
		tmpIp := strings.Join(*ip, ".")
		*ret = append(*ret, tmpIp)
		return
	}

	for i := 0; i < 3 && (i == 0 || s[idx] != '0') && idx+i+1 <= len(s); i++ {
		if i == 2 && s[idx:idx+i+1] > "255" {
			return
		}
		*ip = append(*ip, s[idx:idx+i+1])
		dfs(s, idx+i+1, ip, ret)
		*ip = (*ip)[:len(*ip)-1]
	}
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
}
