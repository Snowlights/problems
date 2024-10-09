package main

import (
	"sort"
)

func minimumDeletions(word string, k int) int {
	cnt := make([]int, 26)
	for _, b := range word {
		cnt[b-'a']++
	}
	sort.Ints(cnt)

	maxSave := 0
	for i, base := range cnt {
		sum := 0
		for _, c := range cnt[i:] {
			sum += min(c, base+k)
		}
		maxSave = max(maxSave, sum)
	}
	return len(word) - maxSave
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
