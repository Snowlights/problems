package main

func maxEnergyBoost(a, b []int) int64 {
	n := len(a)
	c := [2][]int{a, b}
	memo := make([][2]int64, n)
	var dfs func(int, int) int64
	dfs = func(i, j int) int64 {
		if i < 0 {
			return 0
		}
		p := &memo[i][j]
		if *p == 0 { // 首次计算
			*p = max(dfs(i-1, j), dfs(i-2, j^1)) + int64(c[j][i])
		}
		return *p
	}
	return max(dfs(n-1, 0), dfs(n-1, 1))
}
