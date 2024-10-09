package main

import "sort"

func minimumBoxes(apple []int, capacity []int) int {
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] > capacity[j]
	})
	sum, ans := 0, 0
	for _, v := range apple {
		sum += v
	}
	for _, v := range capacity {
		if sum >= v {
			ans++
			sum -= v
		} else {
			break
		}
	}
	if sum > 0 {
		ans++
	}
	return ans
}
