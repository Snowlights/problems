package main

func longestString(x, y, z int) (ans int) {

	dp := make([][][][3]int, x+1)
	for i := range dp {
		dp[i] = make([][][3]int, y+1)
		for j := range dp[i] {
			dp[i][j] = make([][3]int, z+1)
			for k := range dp[i][j] {
				for l := range dp[i][j][k] {
					dp[i][j][k][l] = -1
				}
			}
		}
	}

	var dfs func(i, j, k, pre int) int
	dfs = func(i, j, k, pre int) int {
		dv := &dp[i][j][k][pre]
		if *dv >= 0 {
			return *dv
		}
		res := 2
		defer func() {
			*dv = res
		}()

		switch pre {
		case 0:
			if j+1 <= y {
				res += dfs(i, j+1, k, 1)
			}
		case 1:
			cur := res
			if i+1 <= x {
				res = max(res, cur+dfs(i+1, j, k, 0))
			}
			if k+1 <= z {
				res = max(res, cur+dfs(i, j, k+1, 2))
			}
		case 2:
			cur := res
			if i+1 <= x {
				res = max(res, cur+dfs(i+1, j, k, 0))
			}
			if k+1 <= z {
				res = max(res, cur+dfs(i, j, k+1, 2))
			}
		}
		return res
	}

	return max(dfs(1, 0, 0, 0), max(dfs(0, 1, 0, 1), dfs(0, 0, 1, 2)))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
