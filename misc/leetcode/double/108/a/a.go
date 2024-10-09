package main

func alternatingSubarray(a []int) (ans int) {
	ans = -1
	for i := range a {
		j, base := i+1, 1
		for j < len(a) {
			if a[j]-a[j-1] != base {
				break
			}
			base *= -1
			j++
		}
		if j-i > 1 {
			ans = max(ans, j-i)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
