package main

import "sort"

func maxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	a := make([]int, 0, m*n)
	for _, row := range values {
		a = append(a, row...)
	}
	sort.Ints(a)
	ans := 0
	for i, x := range a {
		ans += x * (i + 1)
	}
	return int64(ans)
}
