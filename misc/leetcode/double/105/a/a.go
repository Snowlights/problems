package main

import "sort"

func buyChoco(a []int, money int) (ans int) {
	sort.Ints(a)
	if a[0]+a[1] <= money {
		return money - a[0] - a[1]
	}
	return money
}
