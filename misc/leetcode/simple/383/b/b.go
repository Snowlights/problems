package main

func minimumTimeToInitialState(s string, k int) int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
		if i%k == 0 && z[i] >= n-i {
			return i / k
		}
	}
	return (n-1)/k + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
