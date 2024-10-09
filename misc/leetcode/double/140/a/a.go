package main

import "math"

func minElement(nums []int) int {
	ans := math.MaxInt
	for _, x := range nums {
		s := 0
		for x > 0 {
			s += x % 10
			x /= 10
		}
		ans = min(ans, s)
	}
	return ans
}
