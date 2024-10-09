package main

import (
	"math"
	"slices"
)

func minimumAverage(nums []int) float64 {
	slices.Sort(nums)
	ans := math.MaxInt
	for i, n := 0, len(nums); i < n/2; i++ {
		ans = min(ans, nums[i]+nums[n-1-i])
	}
	return float64(ans) / 2
}
