package main

import "math"

func minCost(a []int, x int) int64 {
	ans, n := math.MaxInt64, len(a)
	sum := make([]int, n)
	for i := range sum {
		sum[i] = i * x
	}

	for i, v := range a {
		for j := i; j < n+i; j++ {
			v = min(v, a[j%n])
			sum[j-i] += v
		}
	}
	for _, v := range sum {
		ans = min(ans, v)
	}

	return int64(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
