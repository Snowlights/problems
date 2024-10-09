package main

import "math"

func continuousSubarrays(a []int) int64 {
	left, ans := 0, 0
	cnt := make(map[int]int)
	max := func() int {
		m := 0
		for k := range cnt {
			if k > m {
				m = k
			}
		}
		return m
	}
	min := func() int {
		m := math.MaxInt32
		for k := range cnt {
			if k < m {
				m = k
			}
		}
		return m
	}

	for r, v := range a {
		cnt[v]++
		for max()-min() > 2 {
			y := a[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans += r - left + 1
	}
	return int64(ans)
}
