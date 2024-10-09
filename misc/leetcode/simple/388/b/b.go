package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	ans := 0
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	for i, v := range happiness[:k] {
		if v-i > 0 {
			ans += v - i
		}
	}
	return int64(ans)
}
