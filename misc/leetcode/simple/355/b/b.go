package main

func maxArrayValue(a []int) int64 {
	var s []int
	var ans int
	for i := len(a) - 1; i >= 0; i-- {
		if len(s) == 0 || s[len(s)-1] < a[i] {
			s = append(s, a[i])
		} else {
			s[len(s)-1] += a[i]
			ans = max(ans, s[len(s)-1])
		}
	}
	return int64(max(ans, s[len(s)-1]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
