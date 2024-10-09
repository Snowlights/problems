package main

import "math"

func numberOfSubarrays(nums []int) int64 {
	ans := len(nums)
	type pair struct{ x, cnt int }
	st := []pair{{math.MaxInt, 0}}
	for _, x := range nums {
		for x > st[len(st)-1].x {
			st = st[:len(st)-1]
		}
		if x == st[len(st)-1].x {
			ans += st[len(st)-1].cnt
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{x, 1})
		}
	}
	return int64(ans)
}
