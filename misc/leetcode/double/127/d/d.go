package main

import "sort"

func sumOfPowers(nums []int, k int) int {
	const (
		MOD int = 1e9 + 7
		INF int = 1e18
	)

	n := len(nums)
	sort.Ints(nums)
	type pair struct {
		i, rem, lst, mn int
	}

	dp := make(map[pair]int, n+10)
	var dfs func(i, rem, lst, mn int) int
	dfs = func(i, rem, lst, mn int) int {
		if n-i < rem {
			return 0
		}
		if rem == 0 {
			return mn
		}
		p := pair{
			i:   i,
			rem: rem,
			lst: lst,
			mn:  mn,
		}
		if res, ok := dp[p]; ok {
			return res
		}
		res := dfs(i+1, rem-1, nums[i], min(mn, nums[i]-lst)) % MOD
		res = (res + dfs(i+1, rem, lst, mn)) % MOD
		dp[p] = res
		return res
	}

	return dfs(0, k, -INF, INF) % MOD
}
