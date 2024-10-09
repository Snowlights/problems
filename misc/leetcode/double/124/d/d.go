package main

import (
	"sort"
)

func maxSelectedElements(nums []int) (ans int) {
	sort.Ints(nums)
	f := map[int]int{}
	for _, x := range nums {
		f[x+1] = f[x] + 1
		f[x] = f[x-1] + 1
	}
	for _, res := range f {
		ans = max(ans, res)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
