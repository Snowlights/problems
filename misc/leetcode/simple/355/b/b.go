package main

func maxArrayValue(a []int) int64 {
	s := []int{}
	for i := len(a) - 1; i >= 0; i-- {
		if len(s) == 0 || s[len(s)-1] < a[i] {
			s = append(s, a[i])
		} else {
			s[len(s)-1] += a[i]
		}
	}
	ans := 0
	for _, v := range s {
		ans = max(ans, v)
	}
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
