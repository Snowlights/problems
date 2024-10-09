package main

import (
	"strconv"
	"strings"
)

func numberOfPowerfulInt(start, finish int64, limit int, s string) int64 {
	low, high := strconv.FormatInt(start, 10), strconv.FormatInt(finish, 10)
	low = strings.Repeat("0", len(high)-len(low)) + low
	memo, n := make([]int, len(high)), len(high)
	diff := n - len(s)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int
	dfs = func(i int, lowLimit, highLimit bool) int {
		if i == n {
			return 1
		}

		var res int
		if !lowLimit && !highLimit {
			if memo[i] != -1 {
				return memo[i]
			}
			defer func() {
				memo[i] = res
			}()
		}

		l, h := 0, 9
		if lowLimit {
			l = int(low[i] - '0')
		}
		if highLimit {
			h = int(high[i] - '0')
		}

		if i < diff {
			for d := l; d <= min(h, limit); d++ {
				res += dfs(i+1, lowLimit && d == l, highLimit && d == h)
			}
		} else {
			x := int(s[i-diff] - '0')
			if l <= x && x <= min(h, limit) {
				res += dfs(i+1, lowLimit && x == l, highLimit && x == h)
			}
		}
		return res
	}

	return int64(dfs(0, true, true))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
