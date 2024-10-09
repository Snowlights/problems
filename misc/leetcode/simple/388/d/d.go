package main

import (
	"math"
)

func maximumStrength(nums []int, k int) int64 {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}
	f := make([][]int, k+1)
	f[0] = make([]int, n+1)
	for i := 1; i <= k; i++ {
		f[i] = make([]int, n+1)
		f[i][i-1] = math.MinInt
		mx := math.MinInt
		w := (k - i + 1) * (i%2*2 - 1)
		// j 不能太小也不能太大，要给前面留 i-1 个数，后面留 k-i 个数
		for j := i; j <= n-k+i; j++ {
			mx = max(mx, f[i-1][j-1]-s[j-1]*w)
			f[i][j] = max(f[i][j-1], s[j]*w+mx)
		}
	}
	return int64(f[k][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
