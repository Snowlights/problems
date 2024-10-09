package main

const mod int = 1e9 + 7

func specialPerm(a []int) (ans int) {
	n := len(a)
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) (res int) {
		if i == 0 {
			return 1
		}
		dv := &dp[i][j]
		if *dv >= 0 {
			return *dv
		}
		defer func() {
			*dv = res
		}()
		for k, x := range a {
			if i>>k&1 > 0 && (a[j]%x == 0 || x%a[j] == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		return
	}

	for i := range a {
		ans = (ans + dfs(((1<<n)-1)^(1<<i), i)%mod) % mod
	}

	ans = (ans%mod + mod) % mod
	return
}
