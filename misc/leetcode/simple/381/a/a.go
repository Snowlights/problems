package main

import "sort"

func minimumPushes(word string) (ans int) {
	cnt := [26]int{}
	for _, b := range word {
		cnt[b-'a']++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cnt[:])))

	for i, c := range cnt {
		ans += c * (i/8 + 1)
	}
	return
}
