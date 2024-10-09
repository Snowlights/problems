package main

import "math"

func maximumJumps(nums []int, target int) int {
	n := len(nums)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示没有计算过
	}
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前算过了
			return *p
		}
		res := math.MinInt
		for j, x := range nums[:i] {
			if -target <= nums[i]-x && nums[i]-x <= target {
				res = max(res, dfs(j)+1)
			}
		}
		*p = res // 记忆化
		return res
	}
	ans := dfs(n - 1)
	if ans < 0 {
		return -1
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
