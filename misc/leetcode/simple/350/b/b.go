package main

import (
	"math"
	"sort"
)

func findValueOfPartition(a []int) (ans int) {
	sort.Ints(a)
	ans = math.MaxInt32
	for i := 1; i < len(a); i++ {
		ans = min(ans, a[i]-a[i-1])
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
