package main

func numberOfStableArrays(zero, one, limit int) int {
	const mod int = 1e9 + 7
	dp := make([][][]int, 205)
	for i := range dp {
		dp[i] = make([][]int, 205)
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
			if o > limit {
				return 0
			}
			return 1
		}
		if o == 0 {
			if z > limit {
				return 0
			}
			return 1
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
			for i := 1; i <= min(limit, o); i++ {
				res = (res + dfs(z, o-i, 1)) % mod
			}
		} else {
			for i := 1; i <= min(limit, z); i++ {
				res = (res + dfs(z-i, o, 0)) % mod
			}
		}
		return res
	}

	return (dfs(zero, one, 1) + dfs(zero, one, 0)) % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
