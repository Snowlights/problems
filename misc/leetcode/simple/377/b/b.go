package main

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
	h, v := make(map[int]bool), make(map[int]bool)
	cal := func(arr []int, m map[int]bool) {
		for i, v := range arr {
			for _, vv := range arr[i+1:] {
				m[abs(vv-v)] = true
			}
		}
	}
	hFences = append(hFences, 1, m)
	vFences = append(vFences, 1, n)
	cal(hFences, h)
	cal(vFences, v)
	ans, mod := -1, 1_000_000_007
	for i := range h {
		if v[i] {
			ans = max(ans, i)
		}
	}
	if ans == -1 {
		return ans
	}
	return ans * ans % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
