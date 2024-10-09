package main

import "slices"

func clearStars(s string) string {
	st := [26][]int{}
	for i, c := range s {
		if c != '*' {
			st[c-'a'] = append(st[c-'a'], i)
			continue
		}
		for j, p := range st {
			if len(p) > 0 {
				st[j] = p[:len(p)-1]
				break
			}
		}
	}

	idx := []int{}
	for _, p := range st {
		idx = append(idx, p...)
	}
	slices.Sort(idx)

	t := make([]byte, len(idx))
	for i, j := range idx {
		t[i] = s[j]
	}
	return string(t)
}
