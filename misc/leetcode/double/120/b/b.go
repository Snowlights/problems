package main

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	s := 0
	for _, v := range nums {
		s += v
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if s > 2*nums[i] {
			return int64(s)
		}
		s -= nums[i]
	}
	return -1
}
