package main

import (
	"math"
)

func maxSum(nums []int, k, m int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x
	}

	f := make([]int, n+1)
	for i := 1; i <= k; i++ {
		nf := make([]int, n+1)
		for j := range nf {
			nf[j] = math.MinInt / 2
		}
		mx := math.MinInt
		// 左右两边留出足够空间给其他子数组
		for j := i * m; j <= n-(k-i)*m; j++ {
			// mx 表示最大的 f[L]-s[L]，其中 L 在区间 [(i-1)*m, j-m] 中
			mx = max(mx, f[j-m]-s[j-m])
			nf[j] = max(nf[j-1], mx+s[j]) // 不选 vs 选
		}
		f = nf
	}
	return f[n]
}
