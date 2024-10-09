package main

import "math"

func minOperations(k int) int {
	rt := max(int(math.Sqrt(float64(k-1))), 1)
	return min(rt-1+(k-1)/rt, rt+(k-1)/(rt+1))
}

func minOperations2(k int) int {
	ans := math.MaxInt
	for m := 1; m <= k; m++ {
		ans = min(ans, m-1+(k-1)/m)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
