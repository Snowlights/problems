package main

import (
	"slices"
)

func minOperations(nums []int, k int) int {
	mn := slices.Min(nums)
	if k > mn {
		return -1
	}
	set := map[int]struct{}{}
	for _, x := range nums {
		set[x] = struct{}{}
	}
	if k == mn {
		return len(set) - 1
	}
	return len(set)
}
