package main

import (
	"sort"
)

func minOperationsToMakeMedianK(nums []int, k int) (ans int64) {
	sort.Ints(nums)
	m := len(nums) / 2
	if nums[m] > k {
		for i := m; i >= 0 && nums[i] > k; i-- {
			ans += int64(nums[i] - k)
		}
	} else {
		for i := m; i < len(nums) && nums[i] < k; i++ {
			ans += int64(k - nums[i])
		}
	}
	return
}
