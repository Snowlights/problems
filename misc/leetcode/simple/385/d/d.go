package main

func calcZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}
	z[0] = n
	return z
}

func countPrefixSuffixPairs(words []string) (ans int64) {
	type node struct {
		son [26]*node
		cnt int
	}
	root := &node{}
	for _, t := range words {
		z := calcZ(t)
		cur := root
		for i, c := range t {
			c -= 'a'
			if cur.son[c] == nil {
				cur.son[c] = &node{}
			}
			cur = cur.son[c]
			if z[len(t)-1-i] == i+1 { // t[:i+1] == t[len(t)-1-i:]
				ans += int64(cur.cnt)
			}
		}
		cur.cnt++
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
