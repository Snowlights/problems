package main

func minimumCost(s string) int64 {
	ans, n := 0, len(s)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += min(i, n-i)
		}
	}
	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
