package main

func maxScore(a []int, x int) int64 {
	ans := 0
	f := [2]int{-1e9, -1e9}
	f[a[0]%2] = a[0]
	ans = a[0]
	for _, v := range a[1:] {
		p := v % 2
		t := max(f[p]+v, f[p^1]+v-x)
		f[p] = max(f[p], t)
		ans = max(ans, t)
	}

	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
