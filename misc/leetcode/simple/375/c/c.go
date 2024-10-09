package main

import (
	_ "math/bits"
)

func countSubarrays(nums []int, k int) int64 {
	mx := nums[0]
	for _, v := range nums {
		mx = max(mx, v)
	}
	l, ans, h := 0, 0, make(map[int]int)
	for _, v := range nums {
		h[v]++
		for h[mx] >= k {
			h[nums[l]]--
			l++
		}
		ans += l
	}
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
