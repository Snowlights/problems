package main

import "strings"

func minimumOperations(num string) int {
	ans := len(num)
	if strings.Contains(num, "0") {
		ans-- // 可以删除 len(num)-1 次得到 "0"
	}
	f := func(tail string) {
		i := strings.LastIndexByte(num, tail[1])
		if i < 0 {
			return
		}
		i = strings.LastIndexByte(num[:i], tail[0])
		if i < 0 {
			return
		}
		ans = min(ans, len(num)-i-2)
	}
	f("00")
	f("25")
	f("50")
	f("75")
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
