package main

import "sort"

func minProcessingTime(processorTime []int, tasks []int) (ans int) {
	sort.Ints(processorTime)
	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	for i, v := range processorTime {
		ans = max(ans, v+tasks[i*4])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
