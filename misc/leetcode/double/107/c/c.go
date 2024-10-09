package main

func minimizeConcatenatedLength(a []string) (ans int) {

	n := len(a)
	dp := make([][26][26]int, n+1)
	for i := range dp {
		for j := range dp[i] {
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var dfs func(i, j, k int) int
	dfs = func(i, j, k int) int {
		if i == n {
			return 0
		}

		dv := &dp[i][j][k]
		if *dv >= 0 {
			return *dv
		}
		res := 0
		defer func() {
			*dv = res
		}()

		cur := a[i]
		if int(cur[len(cur)-1]-'a') == j {
			res = dfs(i+1, int(cur[0]-'a'), k) + len(cur) - 1
		} else {
			res = dfs(i+1, int(cur[0]-'a'), k) + len(cur)
		}
		if int(cur[0]-'a') == k {
			res = min(res, dfs(i+1, j, int(cur[len(cur)-1]-'a'))+len(cur)-1)
		} else {
			res = min(res, dfs(i+1, j, int(cur[len(cur)-1]-'a'))+len(cur))
		}

		return res
	}

	return dfs(1, int(a[0][0]-'a'), int(a[0][len(a[0])-1]-'a')) + len(a[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
