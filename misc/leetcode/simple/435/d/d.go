package main

import (
	"math"
)

func maxDifference(s string, k int) int {
	const inf = math.MaxInt / 2
	ans, n := -inf, 5
	for x := range n {
		for y := range n {
			if y == x {
				continue
			}
			curS := [5]int{}
			preS := [5]int{}
			minS := [2][2]int{{inf, inf}, {inf, inf}}
			left := 0
			for i, b := range s {
				curS[b-'0']++
				r := i + 1
				for r-left >= k && curS[x] > preS[x] && curS[y] > preS[y] {
					p := &minS[preS[x]&1][preS[y]&1]
					*p = min(*p, preS[x]-preS[y])
					preS[s[left]-'0']++
					left++
				}
				ans = max(ans, curS[x]-curS[y]-minS[curS[x]&1^1][curS[y]&1])
			}
		}
	}
	return ans
}
