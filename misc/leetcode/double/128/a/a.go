package main

func scoreOfString(s string) int {
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += abs(int(s[i]) - int(s[i-1]))
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
