package main

import "math"

func minimumSubstringsInPartition(s string) int {
	n := len(s)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		dv := &memo[i]
		if *dv >= 0 {
			return *dv
		}
		res := math.MaxInt
		cnt := make([]int, 26)
		defer func() {
			*dv = res
		}()

	loop:
		for j := i; j >= 0; j-- {
			c := int(s[j] - 'a')
			cnt[c]++
			for _, v := range cnt {
				if v > 0 && cnt[c] != v {
					continue loop
				}
			}
			res = min(res, dfs(j-1)+1)
		}
		return res
	}
	return dfs(n - 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
