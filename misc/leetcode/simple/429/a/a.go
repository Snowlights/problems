package main

import "slices"

func minimumOperations(nums []int) int {
	seen := map[int]bool{}
	for i, x := range slices.Backward(nums) {
		if seen[x] {
			return i/3 + 1
		}
		seen[x] = true
	}
	return 0
}
