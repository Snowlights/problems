package main

import (
	"math"
)

func maxSum(nums []int) (ans int) {
	set := map[int]struct{}{}
	mx := math.MinInt
	for _, x := range nums { // 一次遍历
		if x < 0 {
			mx = max(mx, x)
		} else if _, ok := set[x]; !ok {
			set[x] = struct{}{}
			ans += x
		}
	}
	if len(set) == 0 {
		return mx
	}
	return
}
