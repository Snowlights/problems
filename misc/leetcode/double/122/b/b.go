package main

import (
	"math/bits"
	"sort"
)

func canSortArray(nums []int) bool {
	i, n := 0, len(nums)

	for i < n {
		s := i
		for i+1 < n && bits.OnesCount(uint(nums[i])) ==
			bits.OnesCount(uint(nums[i+1])) {
			i++
		}
		i++
		sort.Ints(nums[s:i])
	}

	return sort.IntsAreSorted(nums)
}
