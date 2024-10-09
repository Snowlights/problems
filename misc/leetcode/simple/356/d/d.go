package main

const mod int = 1e9 + 7

func countSteppingNumbers(low string, high string) (ans int) {

	var f func(s string) int
	f = func(s string) int {
		dp := make([][10]int, len(s))
		for i := range dp {
			for j := range dp[i] {
				dp[i][j] = -1 // -1 表示没有计算过
			}
		}
		var dfs func(i, pre int, limit, num bool) int
		dfs = func(i, pre int, limit, num bool) int {
			if i == len(s) {
				if num {
					return 1
				}
				return 0
			}
			res := 0
			if !limit && num {
				dv := &dp[i][pre]
				if *dv >= 0 {
					return *dv
				}
				defer func() {
					if !limit && num {
						*dv = res
					}
				}()
			}
			if !num {
				res += dfs(i+1, pre, false, false)
			}

			down, up := 0, 9
			if !num {
				down = 1
			}
			if limit {
				up = int(s[i] - '0')
			}

			for ; down <= up; down++ {
				if !num || abs(down-pre) == 1 {
					res += dfs(i+1, down, limit && down == up, true) % mod
				}
			}
			return res % mod
		}
		return dfs(0, 0, true, false)
	}

	valid := func(s string) int {
		for i := 1; i < len(s); i++ {
			if abs(int(s[i-1])-int(s[i])) != 1 {
				return 0
			}
		}
		return 1
	}

	ans = f(high) - f(low) + valid(low)
	ans = (ans%mod + mod) % mod
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
