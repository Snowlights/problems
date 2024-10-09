package main

func minimumOperations(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	cnt := make([][10]int, m)
	for _, v := range grid {
		for j, vv := range v {
			cnt[j][vv]++
		}
	}

	memo := make([][11]int, m)
	for i := range memo {
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, pre int) int
	dfs = func(i, pre int) int {
		if i == m {
			return 0
		}
		dv := &memo[i][pre]
		if *dv >= 0 {
			return *dv
		}
		res := 0
		defer func() {
			*dv = res
		}()

		for j := 0; j < 10; j++ {
			if j == pre {
				continue
			}
			res = max(res, dfs(i+1, j)+cnt[i][j])
		}

		return res
	}

	return n*m - dfs(0, 10)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
