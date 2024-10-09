package main

import (
	"math"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	f0, f1 := 0, math.MinInt32
	for _, v := range nums {
		f0, f1 = max(v^k+f1, f0+v), max(v^k+f0, f1+v)
	}

	return int64(f0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
