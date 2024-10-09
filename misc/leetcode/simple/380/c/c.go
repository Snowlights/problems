package main

import (
	"sort"
	"strconv"
)

func findMaximumNumber(k int64, x int) int64 {
	return int64(sort.Search(int(1e18), func(i int) bool {
		i++
		return countDigitOne(i, x) > int(k)
	}))
}

func countDigitOne(n, x int) int {
	s := strconv.FormatInt(int64(n), 2)
	n = len(s)

	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	btoi := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}
	var dfs func(i, pre int, limit bool) int
	dfs = func(i, pre int, limit bool) (res int) {
		if i == n {
			return pre
		}

		if !limit {
			dv := &memo[i][pre]
			if *dv != -1 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}

		low, high := 0, 1
		if limit {
			high = int(s[i] - '0')
		}

		for ; low <= high; low++ {
			res += dfs(i+1, pre+btoi(low == 1 && (n-i)%x == 0), limit && low == high)
		}
		return res
	}
	return dfs(0, 0, true)
}
