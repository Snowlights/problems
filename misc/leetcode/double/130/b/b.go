package main

import (
	"math/bits"
	"sort"
)

func maxPointsInsideSquare(points [][]int, s string) int {
	ans := 0
	sort.Search(1e18, func(size int) bool {
		mask := 0
		for i, v := range points {
			c := int(s[i] - 'a')
			if abs(v[0]) <= size && abs(v[1]) <= size {
				if mask>>c&1 == 1 {
					return true
				}
				mask |= 1 << c
			}
		}
		ans = bits.OnesCount(uint(mask))
		return false
	})

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
