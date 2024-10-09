package main

func addMinimum(s string) int {
	t := 1
	for i := 1; i < len(s); i++ {
		if s[i-1] >= s[i] { // 必须生成一个新的 abc
			t++
		}
	}
	return t*3 - len(s)
}
