package main

const mod int = 1e9 + 7

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func count(x string, y string, min_sum int, max_sum int) (ans int) {

	var f func(s string) int
	f = func(s string) int {
		n := len(s)
		dp := make([][]int, n)
		for i := range dp {
			dp[i] = make([]int, min(9*n, max_sum)+1)
			for j := range dp[i] {
				dp[i][j] = -1
			}
		}

		var dfs func(i, sum int, limit bool) int
		dfs = func(i, sum int, limit bool) int {
			if sum > max_sum {
				return 0
			}
			if i == n {
				if sum >= min_sum {
					return 1
				}
				return 0
			}
			res := 0
			if !limit {
				dv := &dp[i][sum]
				if *dv >= 0 {
					return *dv
				}
				defer func() {
					*dv = res
				}()
			}

			down, up := 0, 9
			if limit {
				up = int(s[i] - '0')
			}
			for ; down <= up; down++ {
				res += dfs(i+1, sum+down, limit && down == up) % mod
			}
			return res
		}
		return dfs(0, 0, true)
	}

	ans = f(y) - f(x) + mod
	sum := 0
	for _, v := range x {
		sum += int(v - '0')
	}
	if min_sum <= sum && sum <= max_sum {
		ans++
	}

	ans = (ans%mod + mod) % mod
	return
}
