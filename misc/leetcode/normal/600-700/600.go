package _00_700

import (
	"sort"
	"strconv"
)

// 600
func findIntegers(n int) int {
	s := strconv.FormatInt(int64(n), 2)
	m := len(s)

	dp := make([][2]int64, m)
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(i, pre int, limit bool) int64
	f = func(i, pre int, limit bool) int64 {
		if i == m {
			return 1
		}

		res := int64(0)
		if !limit {
			dv := &dp[i][pre]
			if *dv >= 0 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}

		up := int64(1)
		if limit {
			up = int64(s[i] & 1)
		}
		res += f(i+1, 0, limit && up == 0)
		if pre == 0 && up == 1 {
			res += f(i+1, 1, limit)
		}
		return res
	}

	return int(f(0, 0, true))
}

// 611
func triangleNumber(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return
}

// 621
func leastInterval(tasks []byte, n int) int {
	cnt := make([]int, 26)
	for _, v := range tasks {
		cnt[v-'A']++
	}
	sort.Ints(cnt)
	minTime := (n+1)*(cnt[25]-1) + 1
	for i := 0; i < 25; i++ {
		if cnt[i] == cnt[25] {
			minTime++
		}
	}
	return max(minTime, len(tasks))
}
