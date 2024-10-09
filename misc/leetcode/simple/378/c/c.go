package main

import (
	_ "math/bits"
	"sort"
)

func maximumLength(s string) int {
	h, ans := make(map[int][]int), 0
	i, n, cnt := 0, len(s), 0
	for ; i < n; i++ {
		cnt++
		if i == n-1 || s[i] != s[i+1] {
			h[int(s[i]-'a')] = append(h[int(s[i]-'a')], cnt)
			cnt = 0
		}
	}

	for _, v := range h {
		v = append(v, 0, 0)
		sort.Ints(v)
		n = len(v)
		ans = max(ans, v[n-1]-2, min(v[n-1]-1, v[n-2]), v[n-3])
	}

	if ans == 0 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
