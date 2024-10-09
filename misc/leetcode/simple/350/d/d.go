package main

import "math"

func paintWalls(cost, time []int) int {
	n := len(cost)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n*2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	// 把 j 定义为【付费时间和】减去【免费时间和】，这样递归到终点时，如果 j >= 0
	// 说明这种方案是合法的，否则不合法。这样一来，状态个数就大大减少了。
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if j-n > i { // 注意 j 加上了偏移量 n
			return 0 // 剩余的墙都可以免费刷
		}
		if i < 0 {
			return math.MaxInt32
		}
		p := &memo[i][j]
		if *p != -1 {
			return *p
		}
		*p = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
		return *p
	}
	return dfs(n-1, n) // 加上偏移量 n，防止负数
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
