package main

func numberOfStableArrays(zero, one, limit int) int {
	const mod int = 1e9 + 7
	dp := make([][][]int, zero+1)
	for i := range dp {
		dp[i] = make([][]int, one+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(z, o, pre int) int
	dfs = func(z, o, pre int) int {
		if z == 0 {
			if pre == 1 && o <= limit {
				return 1
			}
			return 0
		}
		if o == 0 {
			if pre == 0 && z <= limit {
				return 1
			}
			return 0
		}

		res := 0
		dv := &dp[z][o][pre]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
		if pre == 0 {
			res = (dfs(z-1, o, 0) + dfs(z-1, o, 1) + mod) % mod
			if z > limit {
				res = (res - dfs(z-limit-1, o, 1) + mod) % mod
			}
		} else {
			res = (dfs(z, o-1, 0) + dfs(z, o-1, 1) + mod) % mod
			if o > limit {
				res = (res - dfs(z, o-limit-1, 0) + mod) % mod
			}
		}
		return res
	}

	return (dfs(zero, one, 1) + dfs(zero, one, 0) + mod) % mod
}
