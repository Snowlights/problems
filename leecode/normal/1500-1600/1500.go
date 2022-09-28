package _500_1600

import (
	"sort"
)

// 1508
func rangeSum(nums []int, n int, left int, right int) int {
	ans := make([]int64, 0, (n+1)*n/2)
	mod := int64(1e9 + 7)
	prefix := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			ans = append(ans, prefix[j]-prefix[i])
		}
	}
	sort.Slice(ans, func(i int, j int) bool {
		return ans[i] < ans[j]
	})
	res := int64(0)
	for i := left - 1; i < right; i++ {
		res += ans[i]
		res %= mod
	}
	return int(res)
}
