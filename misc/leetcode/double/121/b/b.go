package main

import (
	"math/bits"
)

func minOperations(nums []int, k int) int {
	for _, v := range nums {
		k ^= v
	}
	return bits.OnesCount(uint(k))
}
