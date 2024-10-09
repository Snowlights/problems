package main

import (
	"sort"
)

func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	s, n := 0, len(nums)
	id := make([]int, n)
	for i, x := range nums {
		s += x
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return id[i] > id[j]
	})

	ans := make([]int64, len(queries))
	j := 0
	for qi, p := range queries {
		i, k := p[0], p[1]
		s -= nums[i]
		nums[i] = 0 // 标记
		for ; j < n && k > 0; j++ {
			i := id[j]
			if nums[i] > 0 { // 没有标记
				s -= nums[i]
				nums[i] = 0
				k--
			}
		}
		ans[qi] = int64(s)
	}
	return ans
}
