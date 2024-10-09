package main

import "sort"

func relocateMarbles(a []int, moveFrom []int, moveTo []int) (ans []int) {

	h := make(map[int]bool)
	for _, v := range a {
		h[v] = true
	}

	for i, v := range moveFrom {
		delete(h, v)
		h[moveTo[i]] = true
	}
	a = []int{}
	for i := range h {
		a = append(a, i)
	}
	sort.Ints(a)
	return a
}
