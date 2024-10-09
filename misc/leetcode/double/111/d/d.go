package main

import (
	"strconv"
)

func numberOfBeautifulIntegers(low int, high int, k int) int {

	cal := func(s string) int {
		n := len(s)
		var dp = [10][10][10][21]int{}
		for i := range dp {
			for j := range dp[i] {
				for k := range dp[i][j] {
					for l := range dp[i][j][k] {
						dp[i][j][k][l] = -1
					}
				}
			}
		}
		var dfs func(i, pre, odd, even, mod int, limit, num bool) int
		dfs = func(i, pre, odd, even, mod int, limit, num bool) int {
			if i == n {
				if odd == even && mod == 0 {
					return 1
				}
				return 0
			}

			res := 0
			if num && !limit {
				dv := &dp[i][odd][even][mod]
				if *dv >= 0 {
					return *dv
				}
				defer func() {
					*dv = res
				}()
			}

			if !num {
				res += dfs(i+1, 0, 0, 0, 0, false, false)
			}

			low, up := 0, 9
			if !num {
				low = 1
			}
			if limit {
				up = int(s[i] - '0')
			}

			for ; low <= up; low++ {
				res += dfs(i+1, pre*10+low, odd+low%2, even+(low+1)%2, (pre*10+low)%k, limit && low == up, true)
			}
			return res
		}
		return dfs(0, 0, 0, 0, 0, true, false)
	}
	return cal(strconv.Itoa(high)) - cal(strconv.Itoa(low-1))
}
