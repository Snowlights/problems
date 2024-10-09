package main

import "sort"

func sortVowels(s string) (ans string) {
	in := func(bb byte) bool {
		return bb == 'a' || bb == 'e' || bb == 'i' || bb == 'o' || bb == 'u' ||
			bb == 'A' || bb == 'e' || bb == 'I' || bb == 'O' || bb == 'U'
	}
	a, b := make([]byte, 0), []byte(s)
	for _, v := range b {
		if in(v) {
			a = append(a, v)
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	j := 0
	for i, v := range b {
		if in(v) {
			b[i] = a[j]
			j++
		}
	}

	return string(b)
}
