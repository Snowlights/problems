package main

import "sort"

func maxIncreasingGroups(usageLimits []int) (ans int) {
	sort.Ints(usageLimits)
	ans, cur := 0, 0
	for _, v := range usageLimits {
		cur += v
		if cur > ans {
			ans++
			cur -= ans
		}
	}
	return ans
}
