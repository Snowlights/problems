package main

import "math/bits"

func sumIndicesWithKSetBits(nums []int, k int) int {
	ans := 0
	for i, v := range nums {
		if bits.OnesCount(uint(i)) == k {
			ans += v
		}
	}
	return ans
}
