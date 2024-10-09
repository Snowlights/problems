package main

import "sort"

func maxFrequencyScore(nums []int, k int64) int {

	sort.Ints(nums)
	pre := make([]int, len(nums)+1)
	for i, v := range nums {
		pre[i+1] = pre[i] + v
	}

	cost := func(l, m, r int) int {
		left := (m-l)*nums[m] - (pre[m] - pre[l])
		right := pre[r+1] - pre[m+1] - nums[m]*(r-m)
		return left + right
	}
	l, ans := 0, 1
	for r := range nums {
		for int64(cost(l, (l+r)>>1, r)) > k {
			l++
		}
		ans = max(ans, r-l+1)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
