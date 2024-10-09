package main

import "strconv"

var pow5 []string

func init() {
	for p5 := 1; p5 < 1<<15; p5 *= 5 {
		pow5 = append(pow5, strconv.FormatUint(uint64(p5), 2))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumBeautifulSubstrings(s string) (ans int) {

	n := len(s)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		f[i] = n + 1
		if s[i] == '0' {
			continue
		}
		for _, t := range pow5 {
			if i+len(t) > n {
				break
			}
			if s[i:i+len(t)] == t {
				f[i] = min(f[i], f[i+len(t)]+1)
			}
		}
	}
	if f[0] > n {
		return -1
	}
	return f[0]
}
