package main

func minimumLength(s string) (ans int) {
	cnt := [26]int{}
	for _, b := range s {
		cnt[b-'a']++
	}
	for _, c := range cnt {
		ans += (c-1)%2 + 1
	}
	return
}
