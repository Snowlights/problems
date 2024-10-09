package main

import "sort"

func maxScore(nums []int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	sum := 0
	for _, x := range nums {
		sum += x
		if sum <= 0 {
			break
		}
		ans++
	}
	return
}
