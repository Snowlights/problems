package main

import "sort"

func findNonMinOrMax(a []int) (ans int) {
	if len(a) <= 2 {
		return -1
	}
	sort.Ints(a)
	return a[1]
}
