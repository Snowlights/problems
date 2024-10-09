package main

import "sort"

const mod int = 1e9 + 7

func countWays(ranges [][]int) (ans int) {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	ans, maxR := 2, ranges[0][1]
	for _, p := range ranges[1:] {
		if p[0] > maxR { // 产生了一个新的集合
			ans = ans * 2 % mod
		}
		maxR = max(maxR, p[1])
	}
	ans = (ans%mod + mod) % mod
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
