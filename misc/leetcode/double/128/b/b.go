package main

import "sort"

func minRectanglesToCoverPoints(points [][]int, w int) int {
	x, ans := make([]int, 0, len(points)), 1
	for _, v := range points {
		x = append(x, v[0])
	}
	sort.Ints(x)
	base := x[0]
	for _, v := range x {
		if base+w < v {
			ans++
			base = v
		}
	}

	return ans
}
