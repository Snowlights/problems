package main

import "sort"

const mod int = 1e9 + 7

func sumDistance(a []int, s string, d int) (ans int) {
	for i, v := range a {
		switch s[i] {
		case 'L':
			a[i] = v - d
		case 'R':
			a[i] = v + d
		}
	}
	sort.Ints(a)

	sum := 0
	for i, v := range a {
		ans = (ans + i*v%mod) - sum
		sum += v
	}

	ans = (ans%mod + mod) % mod
	return
}
