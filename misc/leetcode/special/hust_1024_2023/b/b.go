package main

import "sort"

func specialChopsticks(chopsticks []int) int {
	sort.Ints(chopsticks)
	n := len(chopsticks)
	for i := 0; i <= n; i++ {
		idx := sort.SearchInts(chopsticks, i)
		if n-idx == i {
			return i
		}
	}
	return -1
}
