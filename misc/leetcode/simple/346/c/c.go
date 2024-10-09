package main

import "strconv"

func punishmentNumber(n int) (ans int) {

	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i * i)
		n := len(s)
		var dfs func(int, int) bool
		dfs = func(p, sum int) bool {
			if p == n {
				return sum == i
			}
			x := 0
			for j := p; j < n; j++ {
				x = x*10 + int(s[j]-'0')
				if dfs(j+1, sum+x) {
					return true
				}
			}
			return false
		}
		if dfs(0, 0) {
			ans += i * i
		}
	}

	return
}
